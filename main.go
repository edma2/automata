package main

import (
	"fmt"

	"github.com/edma2/coursera/automata/dfa"
)

func main() {
	// DFA with no consecutive 1's
	d := dfa.NewDFA("A", []string{"A", "B"})
	d.NewTransition("A", '0', "A")
	d.NewTransition("A", '1', "B")
	d.NewTransition("B", '0', "A")
	d.NewTransition("B", '1', "C")
	d.NewTransition("C", '0', "C")
	d.NewTransition("C", '1', "C")

	fmt.Println(d.Execute("011"))
	fmt.Println(d.Execute("0101"))
	fmt.Println(d.Execute(""))
	fmt.Println(d.Execute("01010100001010111"))
	fmt.Println(d.Execute("000000000000"))
}
