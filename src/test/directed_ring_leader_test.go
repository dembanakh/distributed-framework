package test

import (
	"io/ioutil"
	"leader/directed_ring"
	"log"
	"testing"
)

func TestDirectedRingChangRoberts(t *testing.T) {
	checkLogOutput()
	directed_ring.RunChangRoberts(1000)
}

func TestDirectedRingItaiRodeh(t *testing.T) {
	checkLogOutput()
	directed_ring.RunItaiRodeh(1000)
}

func BenchmarkDirectedRingChangRoberts(b *testing.B) {
	log.SetOutput(ioutil.Discard)
	for iteration := 0; iteration < b.N; iteration++ {
		directed_ring.RunChangRoberts(1000)
	}
}

func BenchmarkDirectedRingItaiRodeh(b *testing.B) {
	log.SetOutput(ioutil.Discard)
	for iteration := 0; iteration < b.N; iteration++ {
		directed_ring.RunItaiRodeh(1000)
	}
}
