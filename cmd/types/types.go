package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/niklasfasching/soup"
)

var client = &http.Client{}

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
			return []interface{}{struct{}{}}
		}
		return []interface{}{standardize(v[0])}
	case map[string]interface{}:
		m := map[string]interface{}{}
		for k, x := range v {
			m[camelCase(k)] = standardize(x)
		}
		return m
	default:
		return struct{}{}
	}
}

func main() {
	m, err := load("api.json")
	if err != nil {
		log.Panic(err)
	}
	out := "package slack\n"
	for url, value := range m {
		name := camelCase(filepath.Base(url))
		out += fmt.Sprintf("type %s %s\n\n", name, generateType(value))
	}
	ioutil.WriteFile("generated_types.go", []byte(out), 0644)
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
