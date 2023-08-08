package main

import (
	. "github.com/ezequielaguilera1993/tocadiscos"
)

func main() {

	newSong := Song{
		NotesToPlay: []NoteToPlay{
			{
				Nota:          A4,
				FiguraMusical: N,
			},
			{
				Nota:          A4,
				FiguraMusical: N,
			},
			{
				Nota:          A4,
				FiguraMusical: N,
			},
		},
		SampleRate: 0,
		Tempo: Tempo{
			BPM: 100,
		},
	}
	PlaySong(newSong)

	PlaySong(LooserSong)
	PlaySong(OdeHim)
}
