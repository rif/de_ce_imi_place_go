package main

import (
	"github.com/rs/zerolog/log"
)

func main() {

	log.Debug().
		Str("Scale", "833 cents").
		Float64("Interval", 833.09).
		Msg("Fibonacci is everywhere")
}
