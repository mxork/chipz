package chip

import (
	"encoding/json"
	"math"
	"os"
	"testing"
)

func playIntro(w *wav, b *metronome, s scale) {
	w.playBeat(b, 1, s[0])
	w.playBeat(b, 1, s[4])
	w.playBeat(b, 2, s[7])
}

func playExit(w *wav, b *metronome, s scale) {
	w.playBeat(b, 1, s[7])
	w.playBeat(b, 1, s[4])
	w.playBeat(b, 2, s[0])
}

func playWorkout(w *wav, s scale) {
	b := &metronome{bpm: 80.0, accents: []float64{1.2,1.0,1.1,1.0} }

	playIntro(w, b, s)

	b.beat(4)
	for i := 0; i < 10; i++ {
		w.playBeat(b, 4, s[0])
		b.beat(4)
		w.playBeat(b, 4, s[2])
		b.beat(4)
	}

	playExit(w, b, s)
}

func hotCross(w *wav, s scale) {
	b := &metronome{bpm: 160.0, accents: []float64{1.8,1.0,1.6,1.0} }
	one := func() {
		w.playBeat(b, 2, s[2])
		w.playBeat(b, 2, s[1])
		w.playBeat(b, 2, s[0])
		b.beat(2)
	}
	two := func(p player) {
		for i := 0; i < 4; i++ {
			w.playBeat(b, 1, p)
		}
	}

	one()
	one()
	two(s[0])
	two(s[1])
	one()
}

func happyBirthday(w *wav) {
	s := chromatic(392.0)
	b := &metronome{bpm: 160.0, accents: []float64{1.8,1.0,1.4} }

	b.beat(2) // upbeat

	one := func() {
		w.playBeat(b, 1, s[0])
		w.playBeat(b, 1, s[0])
		w.playBeat(b, 2, s[2])
		w.playBeat(b, 2, s[0])
	}
	one()
	w.playBeat(b, 2, s[5])
	w.playBeat(b, 4, s[4])
	one()
	w.playBeat(b, 2, s[7])
	w.playBeat(b, 4, s[5])

	w.playBeat(b, 1, s[0])
	w.playBeat(b, 1, s[0])
	w.playBeat(b, 2, s[9])
	w.playBeat(b, 2, s[5])
	w.playBeat(b, 2, s[4])
	w.playBeat(b, 2, s[4])
	w.playBeat(b, 4, s[2])


	w.playBeat(b, 1, s[10])
	w.playBeat(b, 1, s[10])
	w.playBeat(b, 2, s[9])
	w.playBeat(b, 2, s[5])
	w.playBeat(b, 2, s[7])
	w.playBeat(b, 4, s[5])
}

func newJfile() []jnote {
	jfile, _ := os.Open("./a.jtune")
	defer jfile.Close()

	buf := make([]byte, 300)
	jtune := []jnote{}
	//amajor := chromatic(392.0)

	n, _ := jfile.Read(buf)

	err := json.Unmarshal(buf[:n], &jtune )
	if err != nil {
		panic(err)
	}

	return jtune
}

func TestMake(t *testing.T) {
	w := newSilence(60.0, 44100)
	//gmajor := majorScale(392.0)
	g:= minorScale(392.0)
	hotCross(w, gmajor)
	f := os.Stdout
	setvol(int16(math.MaxInt16), w.samples)
	w.encodeRiff(f)
}


