package hello

import (
	"rsc.io/quote"
	quoteV3 "rsc.io/quote/v3"
)

// Hello returns Hello World
func Hello() string { return quote.Hello() }

// Greet return a greeting
func Greet() string { return "Howzit?" }

// Proverb a Go concurency proverb
func Proverb() string {
	return quoteV3.Concurrency()
}
