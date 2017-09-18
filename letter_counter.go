package counter

import (
	"fmt"
	"regexp"
)

type Letters struct {
	Count int
}

func (l Letters) String() string {
	return fmt.Sprintf("Letters: %d", l.Count)
}

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
