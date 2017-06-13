package main

import "math"

// the type used for frequency calculations
type tf float32

// freq is pitch, expressed in Hz
type freq tf

// sharp returns a frequency one semitone higher in pitch
func (f freq) sharp() freq {
	return f * root12
}

// flat returns a frequency one semitone lower in pitch
func (f freq) flat() freq {
	return f / root12
}

// semi is pitch, expressed in semitones above/below the tonic
type semi int

// sharp returns a semitone one semitone higher in pitch
func (s semi) sharp() semi {
	return s + 1
}

// flat returns a semitone one semitone lower in pitch
func (s semi) flat() semi {
	return s - 1
}

// toSemi converts a freq to a semi, given a particular tonic freq
func toSemi(f freq, tonic freq) semi {
	return semi(math.Log(float64(f-tonic)) / math.Log(root12))
}

// toFreq converts a semi to a freq, given a particular tonic freq
func toFreq(s semi, tonic freq) freq {
	return tonic * freq(math.Pow(root12, float64(s)))
}

// duration is the integer length of a musical syllable
type duration uint

// note wraps semi and length
type note struct {
	semi
	duration
}

// scale is a grouping of semitones (canonically, first is 0 == the tonic )
type scale []semi

// key takes a scale and roots it at a given tonic frequnecy
type key struct {
	tonic freq
	scale
}
