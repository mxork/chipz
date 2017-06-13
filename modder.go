package chip

import (
	"math"
	"math/rand"
)

type modder func([]float64)

func maxVol(samps []float64) float64 {
	var maxvol float64
	for i := range samps {
		if samps[i] > maxvol {
			maxvol = samps[i]
		}
	}
	return maxvol
}

func setvol(targvol int16, samps []float64) {
	maxvol := maxVol(samps)
	ratio := float64(targvol) / maxvol

	for i := range samps {
		samps[i] = ratio * samps[i]
	}
}

func setRelativeVol(ratio float64, samps []float64) {
	for i := range samps {
		samps[i] *= ratio
	}
}

func fadeIn(samps []float64) {
	r := math.Pow(0.01, 1.0/float64(len(samps)))
	rr := r
	for i := range samps {
		samps[i] *= 1.0 - rr
		rr *= r
	}
}

func fadeOut(samps []float64) {
	fadeTo(0.01, samps)
}

func fadeTo(ratio float64, samps []float64) {
	r := math.Pow(ratio, 1.0/float64(len(samps)))
	rr := r
	for i := range samps {
		samps[i] *= rr
		rr *= r
	}
}

func throw(samps []float64) {
	maxvol := maxVol(samps)
	for i := range samps {
		samps[i] *= rand.Float64()*maxvol*0.1
	}
}
