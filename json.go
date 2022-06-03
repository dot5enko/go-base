package gobase

import "encoding/json"

var JsonLastError error

func ToJson(obj interface{}) string {
	var bytes []byte
	bytes, JsonLastError = json.Marshal(obj)
	return string(bytes)
}
