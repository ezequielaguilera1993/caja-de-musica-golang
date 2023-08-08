package tocadiscos

import "github.com/faiface/beep"

var LooserSong = Song{
	Tempo: Tempo{
		BPM: 100,
	},
	SampleRate: beep.SampleRate(44100),
	NotesToPlay: []NoteToPlay{
		{Nota: F4, FiguraMusical: N},
		{Nota: E4, FiguraMusical: N},
		{Nota: D4_SOSTENIDO, FiguraMusical: N},
		{Nota: D4, FiguraMusical: B},
	},
}

var OdeHim = Song{
	Tempo: Tempo{
		BPM: 220,
	},
	SampleRate: beep.SampleRate(44100),
	NotesToPlay: []NoteToPlay{
		{Nota: B3, FiguraMusical: N},
		{Nota: B3, FiguraMusical: N},
		{Nota: C4, FiguraMusical: N},
		{Nota: D4, FiguraMusical: N},
		{Nota: D4, FiguraMusical: N},
		{Nota: C4, FiguraMusical: N},
		{Nota: B3, FiguraMusical: N},
		{Nota: A3, FiguraMusical: N},
		{Nota: G3, FiguraMusical: N},
		{Nota: G3, FiguraMusical: N},
		{Nota: A3, FiguraMusical: N},
		{Nota: B3, FiguraMusical: N},
		{Nota: B3, FiguraMusical: B},
		{Nota: A3, FiguraMusical: C},
		{Nota: A3, FiguraMusical: B.AgregarPuntillo()},
	},
}
