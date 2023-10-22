package idgen

import "math/rand"

// In order to generate determined sequence of "random" numbers, we'll
// set a fixed seed
const seed = 1

var abc = []byte("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM")

type Generator struct {
	idLen int
	src   *rand.Rand
}

func NewGenerator(idLen int) Generator {
	return Generator{
		idLen: idLen,
		src:   rand.New(rand.NewSource(seed)),
	}
}

func (g Generator) ID() string {
	buff := make([]byte, g.idLen)

	for i := 0; i < g.idLen; i++ {
		buff[i] = abc[g.src.Intn(len(abc))]
	}

	return string(buff)
}
