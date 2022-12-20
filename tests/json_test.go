package tests

import (
	"encoding/json"
	"log"
	"testing"
)

type ServerResponse struct {
	ResponseStatus int    `json:"response_status"`
	Message        string `json:"message"`
}

func TestJson(t *testing.T) {
	responseExample := ServerResponse{
		ResponseStatus: 200,
		Message:        "OK",
	}
	responseByteArray, err := json.Marshal(responseExample)
	if err != nil {
		t.Error("json marshal error", err)
		t.Fail()
	}
	responseString := string(responseByteArray)
	log.Println(responseString)

	var unmarshalled ServerResponse
	err = json.Unmarshal(responseByteArray, &unmarshalled)
	if err != nil {
		t.Error("json unmarshal error", err)
		t.Fail()
	}
	log.Println(unmarshalled)
}
