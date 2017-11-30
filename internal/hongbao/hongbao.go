package hongbao

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
)

var generate bool = false

//https://github.com/haraldkoch/golang-play/blob/master/simulation/boxmuller.go
func generateGaussianNoise(mu, sigma float32) float32 {
	epsilon := float32(math.SmallestNonzeroFloat32)
	two_pi := 2.0 * math.Pi

	var z0, z1 float32

	generate = !generate

	if !generate {
		return z1*sigma + mu
	}
	var u1, u2 float32
	for ok := true; ok; ok = (u1 <= epsilon) {
		u1 = rand.Float32()
		u2 = rand.Float32()
	}

	z0 = float32(math.Sqrt(-2.0*math.Log(float64(u1))) * math.Cos(two_pi*float64(u2)))
	z1 = float32(math.Sqrt(-2.0*math.Log(float64(u1))) * math.Sin(two_pi*float64(u2)))
	return z0*sigma + mu
}

func GenerateMoneyVector(money float32, num int) (result []float32) {
	moneyLeft := money - float32(num)*0.01
	var mu float32 = 0
	var sigma float32 = 0
	for i := 0; i < num; i = i + 1 {
		mu = moneyLeft / (float32(num) - 1)
		sigma = mu / 2.0
		noiseValue := generateGaussianNoise(mu, sigma)
		if noiseValue < 0 {
			noiseValue = 0
		}
		if noiseValue > moneyLeft {
			noiseValue = moneyLeft
		}
		if i == num-1 {
			var sum float32
			for _, v := range result {
				sum += v
			}
			result = append(result, money-sum)
			return result
		}
		appendmoney, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", noiseValue+0.01), 32)
		result = append(result, float32(appendmoney))
		moneyLeft = moneyLeft - noiseValue + 0.01
	}
	return
}
