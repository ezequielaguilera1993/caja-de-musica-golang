package main

import (
	"fmt"
	"github.com/faiface/beep/speaker"
	"math"
	"time"

	"github.com/faiface/beep"
)

var looserSong = song{
	tempo: tempo{
		100,
		N,
	},
	sampleRate: beep.SampleRate(44100),

	notesToPlay: []noteToPlay{
		{F1, N},
		{E1, N},
		{D1_SOSTENIDO, N},
		{D1, B},
	},
}

var odeHim = song{
	tempo: tempo{
		150,
		N,
	},
	sampleRate: beep.SampleRate(44100),

	notesToPlay: []noteToPlay{
		{B1 / 2, N},
		{B1 / 2, N},
		{C2 / 2, N},
		{D2 / 2, N},
		{D2 / 2, N},
		{C2 / 2, N},
		{B1 / 2, N},
		{A1 / 2, N},
		{G1 / 2, N},
		{G1 / 2, N},
		{A1 / 2, N},
		{B1 / 2, N},
		{B1 / 2, N.agregarPuntillo()},
		{A1 / 2, C},
		{A1 / 2, N.agregarPuntillo()},
	},
}

func main() {
	playSong(odeHim)
	playSong(looserSong)
}

type tempo struct {
	bpm         int //beats per minutes
	pulseFigure figuraMusical
}

type noteToPlay struct {
	tone   nota
	figure figuraMusical
}

type song struct {
	notesToPlay []noteToPlay
	sampleRate  beep.SampleRate
	tempo
}

func playSong(songConfig song) {
	fmt.Println("")
	fmt.Println("///////////// Song config /////////////")
	fmt.Println("bpm: ", songConfig.tempo.bpm)
	fmt.Println("pulseFigure: ", songConfig.tempo.pulseFigure)
	fmt.Println("sampleRate: ", songConfig.sampleRate)
	fmt.Println("///////////////////////////////////////")

	for _, toneConfig := range songConfig.notesToPlay {
		reproduceTone(toneConfig, songConfig.bpm, songConfig.sampleRate, songConfig.pulseFigure)
	}
}

func reproduceTone(config noteToPlay, tempo int, sampleRate beep.SampleRate, pulseFigure figuraMusical) {
	beatDuration := float64(pulseFigure) * float64(60000) / float64(tempo)
	figureDuration := time.Duration(beatDuration / float64(config.figure))
	toneDuration := time.Millisecond * figureDuration

	// Initialize speaker
	sampleDuration := time.Second / time.Duration(config.tone)
	speaker.Init(sampleRate, sampleRate.N(sampleDuration))
	fmt.Println("------------------------")
	fmt.Println("sampleDuration: ", sampleDuration)
	fmt.Println("beatDuration: ", beatDuration)
	fmt.Println("toneDuration: ", toneDuration)
	fmt.Println("figure: ", config.figure)
	fmt.Println("pulseFigure: ", pulseFigure)
	fmt.Println("figureDuration: ", figureDuration)

	// Create a streamer that generates a sine wave for 10 seconds
	s := beep.StreamerFunc(func(samples [][2]float64) (n int, ok bool) {
		for i := range samples {
			t := float64(i) / float64(sampleRate)
			if t >= toneDuration.Seconds() {
				return i, false // Stop generating samples after some time
			}
			samples[i][0] = math.Sin(2*math.Pi*float64(config.tone)*t) * 10 // Left channel
			samples[i][1] = samples[i][0]                                   // Right channel (same as left)
		}
		return len(samples), true
	})
	// Play the sound for 10 seconds
	speaker.Play(s)
	time.Sleep(toneDuration)
}

type figuraMusical float32

const (
	R  figuraMusical = 1
	B  figuraMusical = 2
	N  figuraMusical = 4
	C  figuraMusical = 8
	SC figuraMusical = 16
	F  figuraMusical = 32
	SF figuraMusical = 64
	G  figuraMusical = 128
	SG figuraMusical = 256
)

func (f figuraMusical) agregarPuntillo() figuraMusical {
	return f * (2.0 / 3.0)
}

type nota float32

const (
	C8           nota = 4186.01
	B7           nota = 3951.07
	A7_SOSTENIDO nota = 3729.31
	A7           nota = 3520.00
	G7_SOSTENIDO nota = 3322.44
	G7           nota = 3135.96
	F7_SOSTENIDO nota = 2959.96
	F7           nota = 2793.83
	E7           nota = 2637.02
	D7_SOSTENIDO nota = 2489.02
	D7           nota = 2349.32
	C7_SOSTENIDO nota = 2217.46
	C7           nota = 2093.00
	B6           nota = 1975.53
	A6_SOSTENIDO nota = 1864.66
	A6           nota = 1760.00
	G6_SOSTENIDO nota = 1661.22
	G6           nota = 1567.98
	F6_SOSTENIDO nota = 1479.98
	F6           nota = 1396.91
	E6           nota = 1318.51
	D6_SOSTENIDO nota = 1244.51
	D6           nota = 1174.66
	C6_SOSTENIDO nota = 1108.73
	C6           nota = 1046.50
	B5           nota = 987.767
	A5_SOSTENIDO nota = 932.328
	A5           nota = 880.000
	G5_SOSTENIDO nota = 830.609
	G5           nota = 783.991
	F5_SOSTENIDO nota = 739.989
	F5           nota = 698.456
	E5           nota = 659.255
	D5_SOSTENIDO nota = 622.254
	D5           nota = 587.330
	C5_SOSTENIDO nota = 554.365
	C5           nota = 523.251
	B4           nota = 493.883
	A4_SOSTENIDO nota = 466.164
	A4           nota = 440.000
	G4_SOSTENIDO nota = 415.305
	G4           nota = 391.995
	F4_SOSTENIDO nota = 369.994
	F4           nota = 349.228
	E4           nota = 329.628
	D4_SOSTENIDO nota = 311.127
	D4           nota = 293.665
	C4_SOSTENIDO nota = 277.183
	C4           nota = 261.626
	B3           nota = 246.942
	A3_SOSTENIDO nota = 233.082
	A3           nota = 220.000
	G3_SOSTENIDO nota = 207.652
	G3           nota = 195.998
	F3_SOSTENIDO nota = 184.997
	F3           nota = 174.614
	E3           nota = 164.814
	D3_SOSTENIDO nota = 155.563
	D3           nota = 146.832
	C3_SOSTENIDO nota = 138.591
	C3           nota = 130.813
	B2           nota = 123.471
	A2_SOSTENIDO nota = 116.541
	A2           nota = 110.000
	G2_SOSTENIDO nota = 103.826
	G2           nota = 979.989
	F2_SOSTENIDO nota = 924.986
	F2           nota = 873.071
	E2           nota = 824.069
	D2_SOSTENIDO nota = 777.817
	D2           nota = 734.162
	C2_SOSTENIDO nota = 692.957
	C2           nota = 654.064
	B1           nota = 617.354
	A1_SOSTENIDO nota = 582.705
	A1           nota = 550.000
	G1_SOSTENIDO nota = 519.130
	G1           nota = 489.995
	F1_SOSTENIDO nota = 462.493
	F1           nota = 436.536
	E1           nota = 412.035
	D1_SOSTENIDO nota = 388.909
	D1           nota = 367.081
	C1_SOSTENIDO nota = 346.479
	C1           nota = 327.032
	B0           nota = 308.677
	A0_SOSTENIDO nota = 291.353
	A0           nota = 275.000
)
