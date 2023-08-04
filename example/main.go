package main

import (
	. "github.com/ezequielaguilera1993/tocadiscos"
	"github.com/faiface/beep"
)

func main() {
	PlaySong(LooserSong)
	PlaySong(OdeHim)
}

var LooserSong = Song{
	Tempo: Tempo{
		BPM:         100,
		PulseFigure: N,
	},
	SampleRate: beep.SampleRate(44100),
	NotesToPlay: []NoteToPlay{
		{Tone: F1, Figure: N},
		{Tone: E1, Figure: N},
		{Tone: D1_SOSTENIDO, Figure: N},
		{Tone: D1, Figure: B},
	},
}

var OdeHim = Song{
	Tempo: Tempo{
		BPM:         250,
		PulseFigure: N,
	},
	SampleRate: beep.SampleRate(44100),
	NotesToPlay: []NoteToPlay{
		{Tone: B1 / 2, Figure: N},
		{Tone: B1 / 2, Figure: N},
		{Tone: C2 / 2, Figure: N},
		{Tone: D2 / 2, Figure: N},
		{Tone: D2 / 2, Figure: N},
		{Tone: C2 / 2, Figure: N},
		{Tone: B1 / 2, Figure: N},
		{Tone: A1 / 2, Figure: N},
		{Tone: G1 / 2, Figure: N},
		{Tone: G1 / 2, Figure: N},
		{Tone: A1 / 2, Figure: N},
		{Tone: B1 / 2, Figure: N},
		{Tone: B1 / 2, Figure: N.AgregarPuntillo()},
		{Tone: A1 / 2, Figure: C},
		{Tone: A1 / 2, Figure: N.AgregarPuntillo()},
	},
}
