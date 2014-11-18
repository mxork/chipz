package main

const (
	A freq = 440.0
	_      = 466.163761518
	B      = 493.883301256
	C      = 523.251130601
	_      = 554.365261954
	D      = 587.329535835
	_      = 622.253967444
	E      = 659.255113826
	F      = 698.456462866
	_      = 739.988845423
	G      = 783.990871963
	_      = 830.60939516
)

// 12th root of 2
const root12 = 1.0594630943592953

//var root12 tf = tf(math.Pow(2, 1.0/12))

func (f freq) sharp() freq {
	return f * root12
}

func (f freq) flat() freq {
	return f / root12
}

// shift moves a freq up or down by n semitones
//
// Implemenataion uses iterated mult/division, as
// for small integer n, it is probably faster than math.Pow
func (f freq) shift(n semi) freq {
	of := f
	switch {
	case n > 0:
		for n != 0 {
			of *= root12
			n--
		}
	case n < 0:
		for n != 0 {
			of /= root12
			n++
		}
	}
	return of
}

// lut is a lookup table for frequencies of notes over 5 octaves.
// uniform (piano?) tuning
//var lut [60]freq
