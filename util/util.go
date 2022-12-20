package util

import (
	"fmt"
	"log"
	"time"
)

type Info struct {
	Message string `json:"message"`
	Tt      string `json:"tt"`
}

func TranslateTime(tt time.Time) string {
	return tt.UTC().String()
}

func TimeNow() {
	now := time.Now()
	info := Info{
		Message: "Hello",
		Tt:      TranslateTime(now),
	}
	log.Println(fmt.Sprintf("%v", info))
}
