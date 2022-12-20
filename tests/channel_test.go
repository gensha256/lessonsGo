package tests

import (
	"log"
	"testing"
	"time"
)

func TestChannels(t *testing.T) {
	stream := make(chan int)

	go WriteStream(stream, 5)
	ReadStream(stream)
}

func WriteStream(stream chan int, num int) {
	for i := 0; i < num; i++ {
		stream <- i
	}
	//close(stream)
}

func ReadStream(intStream chan int) {
	for value := range intStream {
		log.Println("get", value)
	}
}

func TestShitChannel(t *testing.T) {
	stream := make(chan int)

	go WriteStream(stream, 2)
	element := <-stream
	log.Println(element)
	element = <-stream
	log.Println(element)
}

func TestBufferedCh(t *testing.T) {
	stream := make(chan time.Time, 1)

	go func(str chan time.Time) {
		for {
			time.Sleep(1 * time.Second)
			str <- time.Now()
		}
	}(stream)

	for {
		tts := <-stream
		log.Println(tts)
		time.Sleep(1 * time.Second)
	}
}
