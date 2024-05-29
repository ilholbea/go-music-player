package main

import (
	"github.com/faiface/beep"
	"github.com/faiface/beep/flac"
	"github.com/faiface/beep/speaker"
	"log"
	"os"
	"time"
)

func main() {
	file, err := os.Open("test.flac")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	streamer, format, err := flac.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Minute/1))
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
}
