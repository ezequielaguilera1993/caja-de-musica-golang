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
		BPM: 250,
	},
	SampleRate: beep.SampleRate(44100),
	NotesToPlay: []NoteToPlay{
		{Nota: B1 / 2, FiguraMusical: N},
		{Nota: B1 / 2, FiguraMusical: N},
		{Nota: C2 / 2, FiguraMusical: N},
		{Nota: D2 / 2, FiguraMusical: N},
		{Nota: D2 / 2, FiguraMusical: N},
		{Nota: C2 / 2, FiguraMusical: N},
		{Nota: B1 / 2, FiguraMusical: N},
		{Nota: A1 / 2, FiguraMusical: N},
		{Nota: G1 / 2, FiguraMusical: N},
		{Nota: G1 / 2, FiguraMusical: N},
		{Nota: A1 / 2, FiguraMusical: N},
		{Nota: B1 / 2, FiguraMusical: N},
		{Nota: B1 / 2, FiguraMusical: N.AgregarPuntillo()},
		{Nota: A1 / 2, FiguraMusical: C},
		{Nota: A1 / 2, FiguraMusical: N.AgregarPuntillo()},
	},
}
