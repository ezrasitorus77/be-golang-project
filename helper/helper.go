package helper

import (
	"be-golang-project/consts"
	"math"
	"math/rand"
	"time"
)

func GenerateRandomKey(size float64) string {
	var key string

	rand.Seed(time.Now().UnixNano())
	num := []rune("0123456789")

	for i := 0.; i < (math.Ceil(consts.MinSecretKeySize) / 3); i++ {
		key += string([]rune{'A' + rune(rand.Intn(26)), 'b' + rune(rand.Intn(26)), num[rand.Intn(len(num))]})
	}

	return key
}
