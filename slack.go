package slack

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"

	"golang.org/x/net/websocket"
)

type API interface {
	Start(context interface{}) error
	Send(v interface{}) error
	Handle(kind string, handlerFunc interface{})
	User() User
}

type Connection struct {
	Token  string
	Debug  bool
	socket *websocket.Conn
	serveMux
	user             User
	errors           chan error
	messageIdCounter int
	lastPong         time.Time
}

type serveMux struct {
	context  reflect.Value
	handlers map[string]handler
}

type handler struct {
	v reflect.Value
	t reflect.Type
}

type User struct {
	Id   string `json: "id"`
	Name string `json: "name"`
}

// https://api.slack.com/methods/rtm.connect
type rtmResponse struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error"`
	URL   string `json:"url"`
	Self  User   `json:"self"`
}

type baseEvent struct {
	Type    string `json:"type"`
	SubType string `json:"subtype"`
}

type errorEvent struct {
	Error struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	} `json:"error"`
}

type pingPongEvent struct {
	Type string    `json:"type"`
	Time time.Time `json:"time"`
}

func (c *Connection) Start(context interface{}) error {
	c.context, c.errors = reflect.ValueOf(context), make(chan error)
	r := rtmResponse{}
	err := get(fmt.Sprintf("https://slack.com/api/rtm.connect?token=%s", c.Token), &r)
	if err != nil {
		return err
	}
	if !r.Ok {
		return errors.New(r.Error)
	}
	s, err := websocket.Dial(r.URL, "", "https://api.slack.com/")
	if err != nil {
		return err
	}
	c.socket, c.user = s, r.Self
	if c.Debug {
		log.Println("Started:", prettyPrintJSON(r.Self))
	}
	go c.pingLoop()
	go c.messageLoop()
	err = <-c.errors
	c.socket.Close()
	c.socket = nil
	return err
}

func (c *Connection) Send(v interface{}) error {
	bytes, err := json.Marshal(v)
	if err != nil {
		return err
	}
	m := map[string]interface{}{}
	err = json.Unmarshal(bytes, &m)
	if err != nil {
		return err
	}
	m["id"] = c.messageIdCounter
	c.messageIdCounter++
	if c.Debug {
		log.Println("Sent:", prettyPrintJSON(m))
	}
	bytes, err = json.Marshal(m)
	if err != nil {
		return err
	}
	if len(bytes) >= 16000 {
		return fmt.Errorf("message must be under 16k bytes long: %#v", m)
	}
	return websocket.Message.Send(c.socket, bytes)
}

func (c *Connection) receive() ([]byte, string, error) {
	bytes, e := []byte{}, baseEvent{}
	err := websocket.Message.Receive(c.socket, &bytes)
	if err != nil {
		return nil, "", err
	}
	if c.Debug {
		log.Println("Received:", prettyPrintMessage(bytes))
	}
	if err := json.Unmarshal(bytes, &e); err != nil {
		return nil, "", err
	}
	kind := e.Type + "/" + e.SubType
	return bytes, kind, nil
}

func (c *Connection) messageLoop() {
	for {
		bytes, kind, err := c.receive()
		switch {
		case err != nil:
		case kind == "error/":
			err = c.handleErrorEvent(bytes)
		case kind == "pong/":
			err = c.handlePongEvent(bytes)
		default:
			err = c.callHandler(kind, bytes, c.Debug)
		}
		if err != nil {
			c.errors <- err
			return
		}
	}
}

func (c *Connection) pingLoop() {
	interval := 10 * time.Second
	for {
		if c.lastPong != (time.Time{}) && time.Since(c.lastPong) > 2*interval {
			c.errors <- fmt.Errorf("did not receive pong for %s", time.Since(c.lastPong))
			return
		}
		err := c.Send(pingPongEvent{"ping", time.Now()})
		if err != nil {
			c.errors <- err
			return
		}
		time.Sleep(interval)
	}
}

func (c *Connection) handleErrorEvent(bytes []byte) error {
	e := errorEvent{}
	err := json.Unmarshal(bytes, &e)
	if err != nil {
		return errors.New(string(bytes))
	}
	return fmt.Errorf("%s (code: %d)", e.Error.Msg, e.Error.Code)
}

func (c *Connection) handlePongEvent(bytes []byte) error {
	e := pingPongEvent{}
	err := json.Unmarshal(bytes, &e)
	if err != nil {
		return fmt.Errorf("bad pong: %s", string(bytes))
	}
	c.lastPong = e.Time
	return nil
}

func (s *serveMux) Handle(kind string, handlerFunc interface{}) {
	v := reflect.ValueOf(handlerFunc)
	if t := v.Type(); t.NumIn() != 2 || t.NumOut() != 1 || t.Out(0) != reflect.TypeOf((*error)(nil)).Elem() {
		panic(fmt.Errorf("handlerFunc must be in the format func(Context, T) error"))
	}
	if _, ok := s.handlers[kind]; ok {
		panic(fmt.Errorf("handler for event kind %s has already been registered", kind))
	}
	if s.handlers == nil {
		s.handlers = map[string]handler{}
	}
	s.handlers[kind] = handler{v, v.Type().In(1)}
}

func (s *serveMux) callHandler(kind string, bytes []byte, debug bool) error {
	handler, handlerKind, ok := s.getHandler(kind)
	if !ok {
		return nil
	}
	v := reflect.New(handler.t)
	if err := json.Unmarshal(bytes, v.Interface()); err != nil {
		return err
	}
	if debug {
		log.Printf("Calling %s (%s): %T", kind, handlerKind, handler.v.Interface())
	}
	if err := handler.v.Call([]reflect.Value{s.context, v.Elem()})[0].Interface(); err != nil {
		return err.(error)
	}
	return nil
}

func (s *serveMux) getHandler(kind string) (handler, string, bool) {
	handler, ok := s.handlers[kind]
	if ok {
		return handler, kind, true
	}
	baseType := strings.Split(kind, "/")[0]
	handler, ok = s.handlers[baseType]
	return handler, baseType, ok
}

func prettyPrintMessage(bytes []byte) string {
	raw := map[string]interface{}{}
	if err := json.Unmarshal(bytes, &raw); err == nil {
		return prettyPrintJSON(raw)
	} else {
		return string(bytes)
	}
}

func prettyPrintJSON(v interface{}) string {
	out := strings.Builder{}
	json := json.NewEncoder(&out)
	json.SetEscapeHTML(false)
	json.SetIndent("", "  ")
	json.Encode(v)
	return out.String()
}

func get(url string, v interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, v)
}