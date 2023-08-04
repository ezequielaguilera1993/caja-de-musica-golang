package tocadiscos

import (
	"fmt"
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"math"
	"time"
)

type Tempo struct {
	BPM         int // beats per minutes
	PulseFigure FiguraMusical
}

type NoteToPlay struct {
	Tone   Nota
	Figure FiguraMusical
}

type Song struct {
	NotesToPlay []NoteToPlay
	SampleRate  beep.SampleRate
	Tempo
}
type Nota float32

type FiguraMusical float32

const (
	R  FiguraMusical = 1
	B  FiguraMusical = 2
	N  FiguraMusical = 4
	C  FiguraMusical = 8
	SC FiguraMusical = 16
	F  FiguraMusical = 32
	SF FiguraMusical = 64
	G  FiguraMusical = 128
	SG FiguraMusical = 256
)

func (f FiguraMusical) AgregarPuntillo() FiguraMusical {
	return f * (2.0 / 3.0)
}

func PlaySong(songConfig Song) {
	fmt.Println("")
	fmt.Println("///////////// Song config /////////////")
	fmt.Println("BPM: ", songConfig.BPM)
	fmt.Println("PulseFigure: ", songConfig.PulseFigure)
	fmt.Println("SampleRate: ", songConfig.SampleRate)
	fmt.Println("///////////////////////////////////////")

	beatDuration := float64(songConfig.PulseFigure) * float64(60000) / float64(songConfig.BPM)
	for _, toneConfig := range songConfig.NotesToPlay {
		reproduceTone(toneConfig, beatDuration, songConfig.SampleRate, songConfig.PulseFigure)
	}

	speaker.Close()
}

func reproduceTone(config NoteToPlay, beatDuration float64, sampleRate beep.SampleRate, pulseFigure FiguraMusical) {
	figureDuration := time.Duration(beatDuration / float64(config.Figure))
	toneDuration := time.Millisecond * figureDuration

	// Initialize speaker
	sampleDuration := time.Second / time.Duration(config.Tone)
	speaker.Init(sampleRate, sampleRate.N(sampleDuration))
	fmt.Println("------------------------")
	fmt.Println("sampleDuration: ", sampleDuration)
	fmt.Println("beatDuration: ", beatDuration)
	fmt.Println("toneDuration: ", toneDuration)
	fmt.Println("Figure: ", config.Figure)
	fmt.Println("PulseFigure: ", pulseFigure)
	fmt.Println("figureDuration: ", figureDuration)

	// Create a streamer that generates a sine wave for the given duration
	s := beep.StreamerFunc(func(samples [][2]float64) (n int, ok bool) {
		for i := range samples {
			t := float64(i) / float64(sampleRate)
			if t >= toneDuration.Seconds() {
				return i, false // Stop generating samples after some time
			}
			samples[i][0] = math.Sin(2*math.Pi*float64(config.Tone)*t) * 10 // Left channel
			samples[i][1] = samples[i][0]                                   // Right channel (same as left)
		}
		return len(samples), true
	})

	// Play the sound for the given duration
	speaker.Play(s)
	time.Sleep(toneDuration)
}
