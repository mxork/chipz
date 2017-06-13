package main

// root12 is the ratio between two adjacent semitones
// (ie. the twelfth root of 2).
const root12 = 1.0594630943592953

// Reference octave.
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

// Major / minor scales
var major = scale{0, 2, 4, 5, 7, 9, 11, 12}
var minor = scale{0, 2, 3, 5, 7, 8, 10, 12} // natural minor
