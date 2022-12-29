package models

import (
	"bytes"
	"fmt"
	"os"
	"time"

	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto/v2"
)

func PlayframeCore(Mp3FlieSrc string) {
	fileBytes, err := os.ReadFile(AudioSourceFile + Mp3FlieSrc + ".mp3")

	if err != nil {
		panic("reading my-file.mp3 failed: " + err.Error())
	}
	fmt.Println("音乐准备完备！！！")
	fileBytesReader := bytes.NewReader(fileBytes)
	// fileBytesReader, err := mp3.NewDecoder(fileBytes)
	// if err != nil {
	// 	panic("mp3.NewDecoder failed: " + err.Error())
	// }
	// Decode file
	decodedMp3, err := mp3.NewDecoder(fileBytesReader)
	if err != nil {
		panic("mp3.NewDecoder failed: " + err.Error())
	}

	samplingRate := 44100

	numOfChannels := 2 //播放速度

	audioBitDepth := 2
	otoCtx, readyChan, err := oto.NewContext(samplingRate, numOfChannels, audioBitDepth)
	if err != nil {
		panic("oto.NewContext failed: " + err.Error())
	}
	<-readyChan
	player := otoCtx.NewPlayer(decodedMp3)
	player.Play()
	for player.IsPlaying() {
		time.Sleep(time.Millisecond)
	}
	err = player.Close()
	if err != nil {
		panic("player.Close failed: " + err.Error())
	}
}
