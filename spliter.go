package counter

func Spliter(sentences <-chan string) (<-chan string, <-chan string) {
	out1, out2 := make(chan string), make(chan string)

	go func() {
		for s := range sentences {
			out1 <- s
			out2 <- s
		}
		close(out1)
		close(out2)
	}()

	return out1, out2
}
