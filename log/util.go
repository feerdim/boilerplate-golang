package log

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func parseInt(i int, s string) int {
	if s == "" {
		return i
	}

	o, e := strconv.Atoi(s)
	if e != nil || o <= 0 {
		return i
	}

	return o
}

func ParseJSON(data any) string {
	JSON, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err.Error())
	}

	return string(JSON)
}

func ParsePrettyJSON(data any) string {
	JSON, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		fmt.Println(err.Error())
	}

	return string(JSON)
}

func generateMessage(msg string, fields []any) string {
	if len(fields) == 0 {
		return msg
	}

	var json string

	if len(fields) > 1 {
		for i := 1; i < len(fields); i++ {
			json = fmt.Sprintf("%s\n%s : %s", json, fields[i-1], ParseJSON(fields[i]))
			i++
		}
	}

	return fmt.Sprintf("%s%s", msg, json)
}
