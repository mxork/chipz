package chip

import (
	"math"
	"math/rand"
)

type tone []float64
func (t tone) play(samples []float64) {
	for i := range samples {
		samples[i] += t[i % len(t)]
	}
}

func rotate(x, y float64, m [2][2]float64) (float64, float64) {
	return (x*m[0][0] + y*m[0][1]), (x*m[1][0] + y*m[1][1])
}

func rotateMatrix(th float64) [2][2]float64 {
	c, s := math.Cos(th), math.Sin(th)
	return [2][2]float64{{c, -s}, {s, c}}
}

func sinWaveR(freq, sampleRate float64) tone {
	var sx, sy, x, y float64
	sx, sy = float64(math.MaxInt16), 0.0

	step := 2 * math.Pi * freq / sampleRate
	rm := rotateMatrix(step)

	// generate full waveform, initialize
	samples := make([]float64, 0, int(sampleRate/freq))
	samples = append(samples, float64(sy))

	// output enough samples that the wave naturally returns to the start
	// value, within an error
	e := 100.0
	for x, y = rotate(sx, sy, rm); !(x < sx+e && x > sx-e); x, y = rotate(x, y, rm) {
		samples = append(samples, float64(y))
	}

	return tone(samples)
}

func sinWave(freq, sampleRate float64) tone {
	// generate full waveform, initialize
	samples := make([]float64, int(sampleRate/freq))

	// output enough samples that the wave naturally returns to the start
	// value, within an error
	a := float64(math.MaxInt16)
	step := 2 * math.Pi * freq / sampleRate
	x := 0.0

	for i := range samples {
		samples[i] = float64(a * math.Sin(x))
		x += step
	}

	return tone(samples)
}

// TODO cleanup
func triWave(freq, sampleRate float64) tone {
	x, y := 0.75, 0.0
	step := freq / sampleRate

	// generate full waveform, initialize
	samples := make([]float64, int(sampleRate/freq))
	period := func(a float64) float64 { return math.Abs(4.0*math.Mod(a, 1.0)-2.0) - 1.0 }

	for i := range samples {
		y = period(x)
		samples[i] = y * float64(math.MaxInt16)
		x += step
	}

	return tone(samples)
}

func sawWave(freq, sampleRate float64) tone {
	x, y := 0.0, 0.0
	step := freq / sampleRate

	// generate full waveform, initialize
	samples := make([]float64, int(sampleRate/freq))
	period := func(a float64) float64 { return 2.0*math.Mod(a, 1.0) - 1.0 }

	for i := range samples {
		y = period(x)
		samples[i] = y * float64(math.MaxInt16)
		x += step
	}

	return tone(samples)
}

type noise struct {
	volume float64
}
func (n noise) play(samps []float64) {
	for i := range samps {
		samps[i] += (2.0*rand.Float64()-1.0)*n.volume
	}
}
