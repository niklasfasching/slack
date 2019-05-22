package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"reflect"
	"regexp"
	"sort"
	"strings"

	"github.com/niklasfasching/soup"
)

var client = &http.Client{}

type Type struct {
	names []string
	value map[string]interface{}
}

type Types []*Type

type Any struct{}

func (ts *Types) addType(name string, v map[string]interface{}) {
	for _, t := range *ts {
		if reflect.DeepEqual(t.value, v) {
			return
		}
	}
	*ts = append(*ts, &Type{[]string{name}, v})
}

func (ts *Types) getStructType(name string, v map[string]interface{}, nested bool) string {
	for _, t := range *ts {
		if t.names[0] != name {
			continue
		}
		valid := true
		for k := range v {
			if _, exists := t.value[k]; !exists {
				valid = false
			}
		}
		if !valid || !nested {
			continue
		}
		return name
	}
	definition, keys := "struct {\n", []string{}
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		definition += fmt.Sprintf("%s %s `json: \"%s\"`\n", k, ts.definition(k, v[k], true), k)
	}
	return definition + "\n}"
}

func (ts *Types) addTypes(name string, v interface{}) {
	v = standardize(v)
	switch v := v.(type) {
	case []interface{}:
		if len(v) == 0 {
			return
		}
		ts.addTypes(name, v[0])
	case map[string]interface{}:
		if len(v) == 0 {
			return
		}
		for k, x := range v {
			ts.addTypes(k, x)
		}
		ts.addType(name, v)
	}
}

func (ts *Types) merge() {
	mergedTs, m := Types{}, map[string][]*Type{}
	for _, t := range *ts {
		if len(t.names) == 1 {
			m[t.names[0]] = append(m[t.names[0]], t)
		} //  else {
		// 	mergedTs = append(mergedTs, t)
		// } more than one name -> remove as we cannot use it
	}
	for _, ts := range m {
		mergedT := ts[0]
		for _, t := range ts[1:] {
			for k, v := range t.value {
				mergedT.value[k] = v
			}
		}
		mergedTs = append(mergedTs, mergedT)
	}
	*ts = mergedTs
}

func (ts *Types) definitions() string {
	definitions := ""
	for _, t := range *ts {
		definitions += fmt.Sprintf("\ntype %s %s\n\n", t.names[0], ts.definition(t.names[0], t.value, false))
	}
	return definitions
}

func (ts *Types) definition(name string, v interface{}, nested bool) string {
	switch v := v.(type) {
	case string, int, float64, bool:
		return fmt.Sprintf("%T", v)
	case []interface{}:
		if len(v) == 0 {
			return "[]interface{}"
		}
		return "[]" + ts.definition(name, v[0], true)
	case map[string]interface{}:
		return ts.getStructType(name, v, nested)
	default:
		return "interface{}"
	}
}

// standardize recursively standardizes v to default values for the respective go type
func standardize(v interface{}) interface{} {
	switch v := v.(type) {
	case string:
		return ""
	case bool:
		return false
	case int, float64:
		return 0
	case []interface{}:
		if len(v) == 0 {
			return []interface{}{Any{}}
		}
		return []interface{}{standardize(v[0])}
	case map[string]interface{}:
		m := map[string]interface{}{}
		for k, x := range v {
			m[camelCase(k)] = standardize(x)
		}
		return m
	default:
		return Any{}
	}
}

func main() {
	m, err := load("api.json")
	if err != nil {
		log.Panic(err)
	}
	types := Types{}
	for url, value := range m {
		name := camelCase(filepath.Base(url))
		types.addTypes(name, value)
	}
	// types.merge()
	ioutil.WriteFile("generated_types.go", []byte(types.definitions()), 0644)
}

func download(path string) error {
	urls := append(listTypeURLs("https://api.slack.com/types"), listTypeURLs("https://api.slack.com/events")...)
	m := map[string]map[string]interface{}{}
	for _, url := range urls {
		codeBlock, err := getCodeBlock(url)
		if err != nil {
			return err
		}
		value := map[string]interface{}{}
		err = json.Unmarshal([]byte(codeBlock), &value)
		m[url] = value
	}
	bs, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, bs, 0644)
}

func load(path string) (m map[string]map[string]interface{}, err error) {
	bs, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return m, json.Unmarshal(bs, &m)
}

func listTypeURLs(url string) (urls []string) {
	for _, a := range soup.MustLoad(client, url).All("table td a") {
		urls = append(urls, "https://api.slack.com"+a.Attribute("href"))
	}
	return urls
}

func generateType(v interface{}) string {
	switch v := v.(type) {
	case string, int, float64, bool:
		return fmt.Sprintf("%T", v)
	case []interface{}:
		if len(v) == 0 {
			return "[]interface{}"
		}
		return "[]" + generateType(v[0])
	case map[string]interface{}:
		if len(v) == 0 {
			return "map[string]interface{}"
		}
		return structType(v)
	default:
		return "interface{}"
	}
}

func structType(m map[string]interface{}) string {
	name, keys := "struct {\n", []string{}
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		name += fmt.Sprintf("%s %s `json: \"%s\"`\n", camelCase(k), generateType(m[k]), k)
	}
	return name + "\n}"
}

func camelCase(snakeCased string) (camelCased string) {
	for _, word := range strings.Split(snakeCased, "_") {
		camelCased += strings.Title(word)
	}
	return strings.ReplaceAll(camelCased, ".", "")
}

func getCodeBlock(url string) (string, error) {
	d, err := soup.Load(client, url)
	if err != nil {
		return "", err
	}
	codeBlock := d.First("pre code").Text()

	// remove ellipsis
	codeBlock = regexp.MustCompile(`{\s*(…|\.\.\.)\s*}`).ReplaceAllString(codeBlock, "{}")
	codeBlock = regexp.MustCompile(`{\s*.*\s*(\.\.\.)\s*.*\s*}`).ReplaceAllString(codeBlock, `{}`)
	codeBlock = regexp.MustCompile(`,?\s*(\.\.\.)\s*\]`).ReplaceAllString(codeBlock, `]`)

	// fix commas
	codeBlock = regexp.MustCompile(`,\s*,`).ReplaceAllString(codeBlock, ",")
	codeBlock = regexp.MustCompile(`,\s*(}|]|$)`).ReplaceAllString(codeBlock, "$1")
	codeBlock = regexp.MustCompile(`("|}|\])\n\s*({|"|\[)`).ReplaceAllString(codeBlock, `$1,$2`)
	return codeBlock, err
}
