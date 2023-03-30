package mock

import "github.com/stretchr/testify/mock"

type Generator struct {
	mock.Mock
}

func (g *Generator) UniqueID() string {
	args := g.Called()
	return args.String(0)
}

func NewGenerator() *Generator {
	return new(Generator)
}
