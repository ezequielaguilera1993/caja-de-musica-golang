# Tocadiscos Package

The tocadiscos package provides functionality to play songs on a virtual "turntable" using the Go programming language. It utilizes the github.com/faiface/beep library to generate audio streams and the github.com/faiface/beep/speaker package to play the audio.

## Installation

To use the tocadiscos package, you must have Go installed. If you don't have Go installed, you can download it from the official website: https://golang.org/

Once you have Go installed, you can install the tocadiscos package by running the following command:

go get github.com/ezequielaguilera1993/tocadiscos

## Usage

Here's an example of create and play two songs:
```go
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

```

## Notes and Tempo

The package defines a set of musical notes and tempo. Here are some of the available notes:

```go
// Note constants (frequency in Hz)
const (
	C8           tocadiscos.nota = 4186.01
	B7           tocadiscos.nota = 3951.07
	A7_SOSTENIDO tocadiscos.nota = 3729.31
	// ... (more notes available)
)
```


## License

This code is provided under the MIT License. Feel free to use, modify, and distribute it as you see fit.

## Acknowledgments

The tocadiscos package uses the github.com/faiface/beep library, which is a fantastic audio library for Go. Special thanks to the contributors and maintainers of the beep project.
