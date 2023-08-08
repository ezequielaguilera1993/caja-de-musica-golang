package main

import (
	. "github.com/ezequielaguilera1993/tocadiscos"
	"time"
)

func main() {

	newSong := Song{
		NotesToPlay: []NoteToPlay{
			{
				Nota:          A2,
				FiguraMusical: N,
			},
			{
				Nota:          A3,
				FiguraMusical: N,
			},
			{
				Nota:          A4,
				FiguraMusical: N,
			},
		},
		SampleRate: 44100,
		Tempo: Tempo{
			BPM: 100,
		},
	}

	PlaySong(newSong)
	time.Sleep(1 * time.Second)

	PlaySong(LooserSong)
	time.Sleep(1 * time.Second)

	PlaySong(OdeHim)
}
