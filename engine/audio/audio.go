package audio

import (
	"github.com/faiface/beep"
	"os"
	"github.com/faiface/beep/wav"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"strings"
	"errors"
	"time"
	"fmt"
)

const SAMPLER_RATE = int(44100 * time.Duration(time.Second/150) / time.Second)

func Decode(file *os.File) (s beep.StreamSeekCloser, format beep.Format, err error) {
	if strings.HasSuffix(file.Name(), ".mp3") {
		return mp3.Decode(file)
	} else if strings.HasSuffix(file.Name(), ".wav") {
		return wav.Decode(file)
	}
	return nil, beep.Format{}, errors.New("Invalid audio type! only mp3 and wav are supported currently")
}

type AudioClip struct {
	Source beep.StreamSeekCloser
	Format *beep.Format
	read uint64
}

func Init(samplerate beep.SampleRate) {
	err := speaker.Init(samplerate, samplerate.N(time.Second/150))
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (ac *AudioClip) Play(endcallback func()) {
	go func() {
		ac.Seek(0)
		if ac.Format.SampleRate != beep.SampleRate(44100) {
			speaker.Play(beep.Seq(beep.Resample(3, ac.Format.SampleRate, beep.SampleRate(44100), ac.Source), beep.Callback(endcallback)))
			return
		}
		speaker.Play(beep.Seq(ac.Source, beep.Callback(endcallback)))
	}()
}

func (ac *AudioClip) Stop() {
	ac.Seek(ac.Source.Len())
}

func (ac *AudioClip) Period(bpm, divisions float64) float64 {
	return (60 / (bpm/divisions))*float64(44100)
}

func (ac *AudioClip) Seek(position int) {
	speaker.Lock()
	ac.read = uint64(ac.Source.Position())
	ac.Source.Seek(position)
	speaker.Unlock()
}

func (ac *AudioClip) Position() int {
	return ac.Source.Position() - int(ac.read)
}