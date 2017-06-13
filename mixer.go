package main

// mixers are non-leafs in the play tree

// addMixer is a simple mixer that mixes by simply adding the signals together
type addMixer struct {
	players []player // children
	buf     []tf     // buffer for intermediate operations
}

// implements player
func (m addMixer) play(out []tf) {
	for j := 0; j < len(out); j += len(m.buf) {
		n := max(len(out[j:]), len(m.buf))

		subbuf := m.buf[:n]
		subout := out[j : j+n]

		for _, p := range m.players {
			p.play(subbuf)

			for i := range subbuf {
				subout[i] += subbuf[i]
			}
		}
	}
}

// volMixer takes a linear combination of the input players
type volMixer struct {
	players []player // children
	volumes []tf     // volume coefficients for i'th player
	buf     []tf     // buffer for intermediate operations
}

// implements player
func (m volMixer) play(out []tf) {
	assert(len(m.players) == len(m.volumes))

	for j := 0; j < len(out); j += len(m.buf) {
		n := max(len(out[j:]), len(m.buf))

		subbuf := m.buf[:n]
		subout := out[j : j+n]

		for pi, p := range m.players {
			vol := m.volumes[pi]
			p.play(subbuf)

			for i := range subbuf {
				subout[i] += vol * subbuf[i]
			}
		}
	}
}
