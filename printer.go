package counter

import "fmt"

// Printer prints the values sent to it
func Printer(values <-chan fmt.Stringer) <-chan struct{} {
	done := make(chan struct{})

	go func() {
		for v := range values {
			fmt.Println(v)
		}
		done <- struct{}{}
	}()

	return done
}
