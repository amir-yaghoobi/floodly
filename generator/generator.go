package generator

import (
	"sync"

	"github.com/amir-yaghoobi/floodly/reporter"
)

type Generator interface {
	Generate(total int, results chan<- reporter.Result, done *sync.WaitGroup)
}
