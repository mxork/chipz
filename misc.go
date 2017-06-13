package main

import (
	"fmt"
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func ck(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func must(e error) {
	if e != nil {
		panic(e)
	}
}

func assert(cond bool) {
	if !cond {
		panic("assert fail")
	}
}

// modulo-int ⇒ mint. We use a lot
// of ring buffers, so this is a conenience
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

func (m *mint) val() int {
	return m.int
}

func (m *mint) adv() int {
	ret := m.int
	m.inc()
	return ret
}
