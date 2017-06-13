package chip

type jnote struct {
	Semi, Beat, Dur int
}

func playJtune(w *wav, s scale, jtune []jnote) {
	seg := 44100
	for _, jn := range jtune {
		s[jn.Semi].play( w.samples[jn.Beat*seg:seg*(jn.Beat+jn.Dur)] )
	}
}
