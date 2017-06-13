package main

import (
	"math"
)

// player is like io.Reader/Writer, but for the underlying frequency type,
// but there aren't any meaningful errors. See `playableSample`
// for the canonical example
//
// TODO players are assumed responsible, and should do
// something (even zero) the entire buffer they are given.
type player interface {
	play([]tf)
}

type stereoPlayer interface {
	stereoPlay([]tf, []tf)
}

// sample is a discrete chunk of audio data
type sample []tf

// createSample returns a slice of values sampled from
// a sine wave of frequency f, at points s/Hz apart
func sineSample(f, hz freq) sample {
	smp := make(sample, int(hz/f))
	for i := range smp {
		smp[i] = tf(math.Sin(2 * math.Pi * float64(i) * float64(f/hz)))
	}
	return smp
}

// playableSample wraps a sample with some state to allow implementing the
// player interface
type playableSample struct {
	*mint
	sample
}

// implements player
func (s playableSample) play(buf []tf) {
	for i := range buf {
		buf[i] = s.sample[s.val()]
		s.inc()
	}
}
