package template

import (
	"encoding/json"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

func newFuncMap() map[string]interface{} {
	m := make(map[string]interface{})
	m["base"] = path.Base
	m["split"] = strings.Split
	m["json"] = UnmarshalJsonObject
	m["jsonArray"] = UnmarshalJsonArray
	m["dir"] = path.Dir
	m["getenv"] = os.Getenv
	m["join"] = strings.Join
	m["datetime"] = time.Now
	m["toUpper"] = strings.ToUpper
	m["toLower"] = strings.ToLower
	m["contains"] = strings.Contains
	m["replace"] = strings.Replace
	m["add"] = Add
	m["subtract"] = Subtract
	return m
}

func addFuncs(out, in map[string]interface{}) {
	for name, fn := range in {
		out[name] = fn
	}
}

func UnmarshalJsonObject(data string) (map[string]interface{}, error) {
	var ret map[string]interface{}
	err := json.Unmarshal([]byte(data), &ret)
	return ret, err
}

func UnmarshalJsonArray(data string) ([]interface{}, error) {
	var ret []interface{}
	err := json.Unmarshal([]byte(data), &ret)
	return ret, err
}

func Add(a, b string) (string, error) {
	convA, err := strconv.Atoi(a)
	if err != nil {
		return "", err
	}
	convB, err := strconv.Atoi(b)
	if err != nil {
		return "", err
	}

	return strconv.Itoa(convA + convB), nil
}

func Subtract(a, b string) (string, error) {
	convA, err := strconv.Atoi(a)
	if err != nil {
		return "", err
	}
	convB, err := strconv.Atoi(b)
	if err != nil {
		return "", err
	}

	return strconv.Itoa(convA - convB), nil
}
