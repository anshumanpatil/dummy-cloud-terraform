package helper

import (
	"bytes"
	"encoding/json"
)

type ENVVars struct {
	Host     string
	Username string
	Password string
}

var ENV ENVVars = ENVVars{
	Host:     "",
	Username: "",
	Password: "",
}

func PrettyJson(data interface{}) string {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent("", "\t")

	err := encoder.Encode(data)
	if err != nil {
		return ""
	}
	return buffer.String()
}
