package utils

import (
	"encoding/json"
)

func Obj2Json(obj interface{}) string {
	bts, _ := json.Marshal(obj)
	return string(bts)
}
