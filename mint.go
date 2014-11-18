package main

// modulo-int ⇒ mint. Get it?
type mint struct {
	int
	N int
}

func (m *mint) inc() {
	m.int++
	if m.int == m.N {
		m.int = 0
	}
}

func (m *mint) val() {
	return m.int
}

func (m *mint) adv() int {
	ret := m.int
	m.inc()
	return ret
}
