package counter

import (
	"fmt"
	"strings"
)

// Words represents a word count
type Words struct {
	Count int
}

func (w Words) String() string {
	return fmt.Sprintf("Words: %d", w.Count)
}

// WordCounter is a step that counts the words in a string
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
