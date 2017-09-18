package counter

import (
	"fmt"
	"sync"
)

// Joiner joins two or more Stringer channels into a single Stringer channel
func Joiner(channels ...<-chan fmt.Stringer) <-chan fmt.Stringer {
	join := make(chan fmt.Stringer)

	wg := sync.WaitGroup{}

	for _, c := range channels {
		wg.Add(1)
		go func(c <-chan fmt.Stringer) {
			for v := range c {
				join <- v
			}
			wg.Done()
		}(c)
	}

	go func() {
		wg.Wait()
		close(join)
	}()

	return join
}
