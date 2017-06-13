package chip

import ()

// errything mono right now
type wav struct {
	samples []float64
	rate    uint16
}

func (w wav) dur() float64 {
	return float64(len(w.samples)) / float64(w.rate)
}

func newSilence(dur float64, rate uint16) *wav {
	samps := make([]float64, int(dur)*int(rate)+1)
	w := wav{samps, rate}
	return &w
}

func (w wav) playBeat(m *metronome, count int, n ...player) (playtime []float64) {
	s, e, _, a := m.beat(count)
	playtime = w.timeSlice(s, e)
	multiplay(playtime, n...)
	setRelativeVol(a, playtime)
	fadeTo(0.2, playtime)
	return
}

// HACK
func (w wav) playDrum(m *metronome, n ...player) (playtime []float64) {
	s, e, _, a := m.beat(1)
	m.time = s
	playtime = w.timeSlice(s, e)
	multiplay(playtime, n...)
	setRelativeVol(a+(a-1.0)*2, playtime)
	fadeTo(0.2, playtime)
	throw(playtime)
	return
	
}

func multiplay(samps []float64, ps ...player) {
	for _, n := range ps {
		n.play(samps)
	}
}
