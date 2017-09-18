package main

import (
	"github.com/vectorhacker/counter"
)

func run(in <-chan string) <-chan struct{} {
	out1, out2 := counter.Spliter(in)

	letters := counter.LetterCounter(out1)
	words := counter.WordCounter(out2)

	join := counter.Joiner(letters, words)
	return counter.Printer(join)
}

func main() {
	in := make(chan string)
	done := run(in)

	in <- "I never put off till tomorrow what I can do the day after."
	in <- "Fashion is a form of ugliness so intolerable that we have to alter it every six months."
	in <- "Life is too important to be taken seriously."

	close(in)

	<-done
}
