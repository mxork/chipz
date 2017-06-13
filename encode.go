package chip

//TODO write own wav encoder.

import (
	fl "github.com/happyalu/goflite"
	"io"
)

func (w *wav) encodeRiff(o io.Writer) {
	gof := fl.Wave{}
	gof.SampleRate = w.rate
	gof.NumSamples = uint32(len(w.samples))

	gof.NumChannels = 1

	gof.Samples = make([]uint16, len(w.samples) )
	for i, v := range w.samples {
		gof.Samples[i] = uint16(v)
	}

	gof.EncodeRIFF(o)
}
