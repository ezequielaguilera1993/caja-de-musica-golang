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

```

## Notes and Tempo

The package defines a set of musical notes and tempo. Here are some of the available notes:

```go
// Note constants (frequency in Hz)
const (
	// ... (more notes available)
	C5           Nota = NoteFrequency(3)
	B4           Nota = NoteFrequency(2)
	A4_SOSTENIDO Nota = NoteFrequency(1)
	A4           Nota = NoteFrequency(0)
	G4_SOSTENIDO Nota = NoteFrequency(-1)
	G4           Nota = NoteFrequency(-2)
	F4_SOSTENIDO Nota = NoteFrequency(-3)
	// ... (more notes available)
)
```


## License

This code is provided under the MIT License. Feel free to use, modify, and distribute it as you see fit.

## Acknowledgments

The tocadiscos package uses the github.com/faiface/beep library, which is a fantastic audio library for Go. Special thanks to the contributors and maintainers of the beep project.
