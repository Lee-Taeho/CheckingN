package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Jsonify(a interface{}) string {
	output, _ := json.MarshalIndent(a, "", "  ")
	return string(output)
}

func JsonifyHttpResponse(resp http.Response) string {
	contents, _ := ioutil.ReadAll(resp.Body)
	return string(contents)
}

func PrintStruct(a interface{}) {
	s, _ := json.MarshalIndent(a, "", "\t")
	fmt.Println(string(s))
}
