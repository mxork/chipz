package chip

// returns the index of the sample beginning at time t
func (w wav) timeIndex(t float64) int {
	return int(t * float64(w.rate))
}

// returns slice of samples beginning at t, and ending at u
func (w wav) timeSlice(t, u float64) []float64 {
	return w.samples[w.timeIndex(t) : w.timeIndex(u)]
}

type accents []float64
type metronome struct {
	time, bpm, pause float64
	count int
	accents
}

func (m *metronome) beat(count int) (float64, float64, float64, float64) {
	dur := float64(count)*60.0/m.bpm
	start, end, next := m.time, m.time + dur - m.pause, m.time + dur
	accent := m.accents[ m.count ]

	m.time = next
	m.count = ( m.count+count) % len(m.accents)
	return start, end, next, accent
}
