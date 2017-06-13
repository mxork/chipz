package main

import "io"

// intoBytes encodes the floating point values into byte
// values
func intoBytes(floats []tf, bytes []byte) {
	assert(len(floats) == len(bytes))

	for i, f := range floats {
		assert(f <= 1.0) // we could also just clip
		bytes[i] = byte(255 * window(f))
	}
}

// restricts FP value to be within 0.0-1.0
func window(x tf) tf {
	if x > 1.0 {
		x = 1.0
	} else if x < 0.0 {
		x = 0.0
	}

	return x
}

// bytePlayer turns a player into an io.WriterTo
type bytePlayer struct {
	player
	floats []tf
	bytes  []byte
}

func (bp bytePlayer) WriteTo(w io.Writer) (n int64, err error) {
	bp.play(bp.floats)
	intoBytes(bp.floats, bp.bytes)
	n_int, err := w.Write(bp.bytes) // return signature mismatch Writer-WriterTo
	return int64(n_int), err
}
