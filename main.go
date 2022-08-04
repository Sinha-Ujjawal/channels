package main

import (
	"channels/channelutils"
	"channels/pair"
	"fmt"
)

func fibs() <-chan uint64 {
	fibStep := func(f0 uint64, f1 uint64) (uint64, uint64) {
		return f1, f0 + f1
	}
	iterator := func(p pair.Pair[uint64, uint64]) pair.Pair[uint64, uint64] {
		return pair.Map(fibStep, p)
	}
	return channelutils.MapOn(
		pair.First[uint64, uint64],
		channelutils.Iterate(
			pair.MkPair(uint64(0), uint64(1)),
			iterator,
		),
	)
}

func collatz(n uint64) <-chan uint64 {
	collatzStep := func(m uint64) uint64 {
		if m%2 == 0 {
			return m / 2
		} else {
			return m + m + m + 1
		}
	}
	return channelutils.Iterate(n, collatzStep)
}

func main() {
	fs := channelutils.MapOn(func(x uint64) string { return fmt.Sprintf("Fibonacci: %d", x) }, fibs())
	cs := channelutils.MapOn(func(x uint64) string { return fmt.Sprintf("Collatz: %d", x) }, collatz(456))
	os := channelutils.Take(30, channelutils.FanIn(fs, cs))
	printToConsole := func(s string) {
		println(s)
	}
	for {
		x, err := (<-os).Unwrap()
		if err != nil {
			break
		}
		x.Do(
			printToConsole,
			printToConsole,
		)
	}
}
