package main

// the type used for frequency calculations: wasn't sure what to use,
// (still don't), so change this as needed.
type tf float32

// a frequency
type freq tf

// length of a syllable
type length uint

// tone, expressed in semitones above/below the tonic
type semi int

// tone and length in a convenient package
type note struct {
	semi
	length
}

type key []freq

// keyOf generates a key for given tonic and scale
func keyOf(tonic freq, s scale) key {
	k := make(key, len(s))
	for i := range k {
		k[i] = tonic.shift(s[i])
	}
	return k
}

// scale is a list of intervals abv (canonically, first is 0)
type scale []semi

// Major / minor scales
var major = scale{0, 2, 4, 5, 7, 9, 11, 12}
var minor = scale{0, 1, 3, 5, 6, 8, 10, 12} // really? probably sounds awful
