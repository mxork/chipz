package main

import (
	"math"
)

const RATE = 44100

// sample is a discrete sinewave
type sample []tf

// createSample returns a slice of values sampled from
// a sine wave of frequency f, at points s/Hz apart
func newSample(f, hz freq) sample {
	smp := make(sample, int(hz/f))
	for i := range smp {
		smp[i] = tf(math.Sin(2 * math.Pi * float64(i) * float64(f/hz)))
	}
	return smp
}

type playableSample struct {
	*mint
	sample
}

func (s *playableSample) play(buf []tf) (n int) {
	for i := range buf {
		buf[i] = s.sample[s.val()]
		s.inc()
	}
	return len(buf) // will always be satisfied
}

// àla io.Reader/Writer, but for tf, and there
// aren't any meaningful errors. See `playableSample`
// for the canonical example
type player interface {
	play([]tf) int
}
