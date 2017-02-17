package lib

import (
	"log"
	"time"
)

func Benchmark(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s\n", name, elapsed)
}
