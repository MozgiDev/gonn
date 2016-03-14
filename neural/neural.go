package neural

import (
	"math/rand"
	"time"
	"fmt"
)

type NeuralType struct {
	Input  []float64
	weight []float64
}

var e = 1.

func New(nbInput int) (*NeuralType) {
	neu := new(NeuralType)
	neu.Input = make([]float64, nbInput + 1)
	neu.Input[0] = 0 //biais
	neu.weight = make([]float64, nbInput + 1)
	return neu
}

func (neu *NeuralType) Init() {
	rand.Seed(time.Now().Unix())
	//for i := 0; i < len(neu.weight); i++ {
	//	neu.weight[i] = rand.Float64() * 2 - 1
	//}
	for i := 0; i < len(neu.weight); i++ {
		neu.weight[i] = float64(rand.Intn(2)-1)
	}
}

func (neu *NeuralType) LearnIt(x []float64, c float64) {
	for i := 0; i < len(x); i++ {
		neu.Input[i + 1] = x[i]
	}
	neu.Input[0] = 1//biais

	somme := 0.0
	for i := 0; i < len(neu.Input); i++ {
		somme += neu.Input[i] * neu.weight[i]
	}
	output := Sigmoid(somme)

	if (output != c) {
		for j := 0; j < len(neu.weight); j++ {
			fmt.Println("w", j, "=", neu.weight[j], "+", e, "*(", c, "-", output, ")*", neu.Input[j])
			neu.weight[j] = neu.weight[j] + e * (c - output) * neu.Input[j]

		}
	}else {
		for j := 0; j < len(neu.weight); j++ {
			fmt.Println("w", j, "=", neu.weight[j])
		}
	}

	fmt.Print("|")
	for k := 0; k < len(neu.weight); k++ {
		fmt.Print(int(neu.weight[k]), "|")
	}
	for l := 0; l < len(neu.Input); l++ {
		fmt.Print(neu.Input[l])
	}
	fmt.Print("|", output, "|", c)
	fmt.Println()
}

func (neu *NeuralType) Cogitate ()(output float64) {
	somme := 0.0
	for i := 0; i < len(neu.Input); i++ {
		somme += neu.Input[i] * neu.weight[i]
	}
	return Sigmoid(somme)
}


func Sigmoid(x float64) (float64) {
	//y := 5.0
	//return 1 / (1 + math.Pow(math.E, -(y) * x))
	//y :=  2 / (1 + math.Pow(math.E,-x) - 1)

	if x > 0. {
		return 1.0
	}
	return 0.0
}
