package transform

import (
	"encoding/json"

	simplejson "github.com/bitly/go-simplejson"
)

const testJSONInput = `{"rating": {"primary": {"value": 3}, "example": {"value": 3}}}`

func getConfig(spec string, require bool) Config {
	var f map[string]interface{}
	json.Unmarshal([]byte(spec), &f)
	return Config{Spec: &f, Require: require}
}

func getTransformTestWrapper(tform func(spec *Config, data *simplejson.Json) error, cfg Config, input string) (string, error) {
	json, e := simplejson.NewJson([]byte(input))
	if e != nil {
		return "", e
	}
	e = tform(&cfg, json)
	if e != nil {
		return "", e
	}
	tmp, e := json.MarshalJSON()
	if e != nil {
		return "", e
	}
	return string(tmp), nil
}
