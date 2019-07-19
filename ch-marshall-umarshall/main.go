package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

type input struct {
	EncryptedMessage string `json:"encryptedMessage"`
	Key              int    `json:"key"`
}

type sampleInput struct {
	Input input `json:"input"`
}

type sample struct {
	SampleInput sampleInput `json:"sampleInput"`
}

type output struct {
	Message string `json:"message"`
}

type sampleOutput struct {
	Output output `json:"output"`
}

type response struct {
	SampleOutput sampleOutput `json:"sampleOnput"`
}

func caserCipherdecode(str string, key int) string {
	len := len(str)
	msg := ""
	str = strings.ToUpper(str)
	for i := 0; i < len; i++ {

		k := int(str[i])
		if k < 65 {
			msg = msg + string(k)
		} else {
			j := (int(str[i]) - key)
			if j < 65 {
				j = j + 26
			}
			msg = msg + string(j)
		}
	}

	return msg
}

func main() {

	file, _ := ioutil.ReadFile("input-nested.json")

	var data sample
	err := json.Unmarshal([]byte(file), &data)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(data)

	msg := caserCipherdecode(data.SampleInput.Input.EncryptedMessage, data.SampleInput.Input.Key)

	var resp = response{
		SampleOutput: sampleOutput{
			Output: output{
				Message: msg,
			},
		},
	}

	outputJSON, _ := json.Marshal(resp)
	fmt.Println(string(outputJSON))

}
