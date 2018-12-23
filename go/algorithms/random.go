package algorithms

import (
	"math/rand"
	"time"
)

func RandInt(min,max int) int {
	rand.Seed(time.Now().Unix())
	if max <= min || max == 0{
		return max
	}
	return rand.Intn(max-min) + min
}
