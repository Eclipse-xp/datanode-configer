package handler

import (
	"io/ioutil"
	"encoding/json"
	"net/http"
)

func bodyAsMap(r *http.Request) map[string]interface{} {
	reqMap := make(map[string]interface{})
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &reqMap)
	return reqMap
}
