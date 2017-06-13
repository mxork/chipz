package chip

import (
	"math"
)

const semi float64 = 1.059463

type player interface {
	play([]float64)
}

type scale []note
type key []note
type chord []note

type note struct {
	pitch int // abs. pos
	id    int // 'A', 'B'
	key
	player
}

/*
func makeKey(freq float64, mode []int) key {
	o := chromatic(freq)
	m := make(key, 4*7+1) //TODO
	for i := range m {
		m[i] = note{
			i, i % ds, m,
			triWave(freq*math.Pow(semi, 1.0+float64(semis)), 44100),
		}
		semis += ds[i%ds]
	}
	return m
}
*/

func (n note) majorChord() chord {
	return chord{n.key[pitch],
		n.key[n.pitch+2], n.key[n.pitch+4],
		n.key[n.pitch+6],
	}
}

func chromatic(freq float64) key {
	s := make(key, 13)
	for i := range s {
		s[i] = note{i, i % 12, s,
			&triWave(freq*math.Pow(semi, 1.0+float64(i)), 44100),
		}
	}
	return s
}
