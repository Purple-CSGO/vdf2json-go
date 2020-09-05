package vdf2json

import (
	"bytes"
	"encoding/json"
	"runtime"
	"strings"
)

func ToJson(vdfData string) (string, error) {
	// - 删除最开始的"xxx"
	// - 按行遍历 如果有 \"\t\t\" -> \": \" 同时在结尾加上,
	// - 按行遍历 如果有 { → : { ... 如果有 } → },
	// - trim每行左边的 \t 和每行右边的换行符 \r 和 \n
	// - 遍历 如果有 },} → }}
	// - 删除开头的 ": " 和结尾的 ","
	linebreak, jsonData := "", ""
	switch runtime.GOOS {
	case "darwin":
		linebreak = "\r"
	case "windows":
	case "linux":
		linebreak = "\n"
	default:
		linebreak = "\n"
	}
	startpoint := strings.Index(vdfData, "{")
	vdfData = vdfData[startpoint:]
	lines := strings.Split(vdfData, linebreak)
	for _, line := range lines {
		line = strings.TrimRight(line, "\r")
		if strings.Contains(line, "\"\t\t\"") {
			line = strings.Replace(line, "\"\t\t\"", "\": \"", -1)
			line += ","
		}
		line = strings.Replace(line, "{", ": {", -1)
		line = strings.Replace(line, "\t", "", -1)
		line = strings.Replace(line, "}", "},", -1)
		jsonData += line
	}
	jsonData = strings.Replace(jsonData, ",}", "}", -1)
	jsonData = strings.TrimLeft(jsonData, ": ")
	jsonData = strings.TrimRight(jsonData, ",")
	//	fmt.Println(jsonData)
	var str bytes.Buffer
	err := json.Indent(&str, []byte(jsonData), "", "    ")
	if err != nil {
		return "", err
	}
	return str.String(), nil
}
