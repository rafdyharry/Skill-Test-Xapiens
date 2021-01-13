// package main
package main

// import fmt
import "fmt"

// Main function
func main() {

	a, b, c, d, e, f, g := 1, 2, 3, 4, 5, 6, 7 //input angka yang di berikan
	a *= 1000000                               // Kali anga dengan yang di minta input
	b *= 100000
	c *= 10000
	d *= 1000
	e *= 100
	f *= 10
	g *= 1
	// using channel

	chnl := make(chan int) //bikin pengulangan loop
	go func() {
		chnl <- a
		chnl <- b
		chnl <- c
		chnl <- d
		chnl <- e
		chnl <- f
		chnl <- g
		close(chnl)
	}()
	for i := range chnl { // lakukan looping
		fmt.Println(i)
	}

}
