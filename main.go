package main

import (
	"os"
)

// plays two sine waves in 8-bit RAW encoding to stdout
func main() {
	a := playableSample{&mint{}, sineSample(200, 44100)}
	b := playableSample{&mint{}, sineSample(400, 44100)}

	// will disappear behind a constructor, eventually
	a.mint.N = len(a.sample)
	b.mint.N = len(b.sample)

	ab := volMixer{
		[]player{a, b},
		[]tf{0.7, 0.3}, // volumes
		make([]tf, 1024),
	}

	bp := bytePlayer{ab, make([]tf, 1024), make([]byte, 1024)}
	for {
		bp.WriteTo(os.Stdout)
	}
}
