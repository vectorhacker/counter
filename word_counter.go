package counter

import (
	"fmt"
	"strings"
)

type Words struct {
	Count int
}

func (w Words) String() string {
	return fmt.Sprintf("Words: %d", w.Count)
}

func WordCounter(sentences <-chan string) <-chan fmt.Stringer {
	count := make(chan fmt.Stringer)

	go func() {
		for s := range sentences {
			words := strings.Split(s, " ")

			count <- Words{Count: len(words)}
		}
		close(count)
	}()

	return count
}
