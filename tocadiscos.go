package tocadiscos

import (
	"fmt"
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"math"
	"time"
)

type Tempo struct {
	BPM int // beats per minutes
	FiguraMusical
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

func (f FiguraMusical) AgregarPuntillo() FiguraMusical {
	return f * (2.0 / 3.0)
}

func PlaySong(songConfig Song) {
	fmt.Println()
	fmt.Println("///////////// Song config /////////////")
	fmt.Println("BPM: ", songConfig.BPM)
	fmt.Println("FiguraMusical: ", songConfig.FiguraMusical)
	fmt.Println("SampleRate: ", songConfig.SampleRate)
	fmt.Println("///////////////////////////////////////")

	beatDuration := float64(songConfig.FiguraMusical) * float64(60000) / float64(songConfig.BPM)

	for _, noteToPlay := range songConfig.NotesToPlay {
		reproduceTone(noteToPlay, beatDuration, songConfig.SampleRate, songConfig.FiguraMusical)
	}

	speaker.Close()
}

func reproduceTone(noteToPlay NoteToPlay, beatDuration float64, sampleRate beep.SampleRate, pulseFigure FiguraMusical) {
	figureDuration := time.Duration(beatDuration / float64(noteToPlay.Figure))
	toneDuration := time.Millisecond * figureDuration

	// Initialize speaker
	sampleDuration := time.Second / time.Duration(noteToPlay.Tone)
	speaker.Init(sampleRate, sampleRate.N(sampleDuration))
	fmt.Println("------------------------")
	fmt.Println("sampleDuration: ", sampleDuration)
	fmt.Println("beatDuration: ", beatDuration)
	fmt.Println("toneDuration: ", toneDuration)
	fmt.Println("FiguraMusical: ", noteToPlay.Figure)
	fmt.Println("FiguraMusical: ", pulseFigure)
	fmt.Println("figureDuration: ", figureDuration)

	// Create a streamer that generates a sine wave for the given duration
	s := beep.StreamerFunc(func(samples [][2]float64) (n int, ok bool) {
		for i := range samples {
			t := float64(i) / float64(sampleRate)
			if t >= toneDuration.Seconds() {
				return i, false // Stop generating samples after some time
			}
			samples[i][0] = math.Sin(2*math.Pi*float64(noteToPlay.Tone)*t) * 10 // Left channel
			samples[i][1] = samples[i][0]                                       // Right channel (same as left)
		}
		return len(samples), true
	})

	// Play the sound for the given duration
	speaker.Play(s)
	time.Sleep(toneDuration)
}
