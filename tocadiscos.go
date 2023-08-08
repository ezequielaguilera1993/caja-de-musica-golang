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
}

type NoteToPlay struct {
	Nota          Nota
	FiguraMusical FiguraMusical
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
	fmt.Println("SampleRate: ", songConfig.SampleRate)
	fmt.Println("///////////////////////////////////////")

	beatDuration := time.Millisecond * time.Duration(float64(60000)/float64(songConfig.BPM))

	for _, noteToPlay := range songConfig.NotesToPlay {
		reproduceTone(noteToPlay, beatDuration, songConfig.SampleRate)
	}
	speaker.Close()
}

func reproduceTone(noteToPlay NoteToPlay, beatDurationInMs time.Duration, sampleRate beep.SampleRate) {

	toneDuration := beatDurationInMs * time.Duration(N/noteToPlay.FiguraMusical)
	// Initialize speaker
	sampleDuration := time.Second / time.Duration(noteToPlay.Nota)
	err := speaker.Init(sampleRate, sampleRate.N(sampleDuration))
	if err != nil {
		panic(err)
	}

	fmt.Println("------------------------")
	fmt.Println("Frecuencia de la nota: ", noteToPlay.Nota.Frequency())
	fmt.Println("Figura musical de la nota: ", noteToPlay.FiguraMusical)
	fmt.Println("Duración: ", toneDuration)
	fmt.Println("SampleDuration: ", sampleDuration)

	loopsCounter := 0
	// Create a streamer that generates a sine wave for the given duration
	s := beep.StreamerFunc(func(samples [][2]float64) (n int, ok bool) {
		loopsCounter++

		for i := range samples {
			sampleTimePositionInSeconds := float64(i) / float64(sampleRate)                                //cuando sampleTimePositionInSeconds == 1 es porque paso un segundo
			samples[i][0] = math.Sin(2*math.Pi*float64(noteToPlay.Nota)*sampleTimePositionInSeconds) * 1.1 // Left channel 2*math.Pi, por cada "1". Completa un hz. Pero yo quiero 440hz por ejemplo. El multiplicador de volumen es para saturarlo un poco y sacarle la pureza de la función seno.
			samples[i][1] = samples[i][0]                                                                  // Right channel (same as left)
		}

		return len(samples), true
	})

	// Play the sound for the given duration
	speaker.Play(s)
	time.Sleep(toneDuration)
	fmt.Println("------------------------")
	fmt.Println("loops: ", loopsCounter)

}
