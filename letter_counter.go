package counter

import (
	"fmt"
	"regexp"
)

// Letter represents a letter count
type Letters struct {
	Count int
}

func (l Letters) String() string {
	return fmt.Sprintf("Letters: %d", l.Count)
}

// LetterCounter is a process that counts letters in a string
func LetterCounter(sentences <-chan string) <-chan fmt.Stringer {
	count := make(chan fmt.Stringer)

	re := regexp.MustCompile("[a-zA-Z]")

	go func() {
		for s := range sentences {
			letters := re.FindAllString(s, -1)

			count <- Letters{Count: len(letters)}
		}
		close(count)
	}()

	return count
}
