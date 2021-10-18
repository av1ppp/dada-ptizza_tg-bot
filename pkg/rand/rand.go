package rand

import "math/rand"

// IntMinMax - генерация случайного числа от min до max
func IntMinMax(min, max int) int {
	return rand.Intn(max-min) + min
}
