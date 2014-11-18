package main

// mixer takes several players and mixes them into a single output
type mixer interface {
	register(player)
	player
}

// addmixer is a simple mixer that mixes by simply adding the signals together
type addmixer struct {
	players []player
	scratch []tf
}

func (m *addmixer) register(p player) {
	m.players = append(m.players, p)
}

func (m *addmixer) play(out []tf) {
	// admittedly, some slice magic going on here.
	// summary: slicing should just ensure that
	// the scratch space and the chunk of the output
	// being worked on are aligned, and of the same size.

	for j := 0; j < len(out); j += n {
		n := max(len(out[j:]), len(m.scratch))

		schunk := m.scratch[:n]
		ochunk := out[j : j+n]

		for _, p := range players {
			p.play(schunk)

			for i := range schunk {
				ochunk[i] += schunk[i]
			}
		}
	}
}

// locmixer is a little more sophisticated: on registration,
// it assigns the player a polar coordinate in the plane
// centered on the listeners head.

// speed of sound-ish
const SoS = 340

// head width in centimeters
const headw = 0.20
