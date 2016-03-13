package neuralnetwork

import (
	"math/rand"
	"time"
	"fmt"
	"errors"
)

type NeuralType  struct {
	input  []float64
	weight []float64
	delta  float64
	outup  float64
}

type LayerType struct {
	neurals []NeuralType
}

type NeuralNetworkType struct {
	layers []LayerType
}

var ε = 1.0

func New(param ...int) *NeuralNetworkType {
	net := new(NeuralNetworkType)
	net.layers = make([]LayerType, len(param))
	if len(param) > 1 {
		for i, p := range param {
			net.layers[i].neurals = make([]NeuralType, p)
			for j, _ := range net.layers[i].neurals {
				if i == 0 {
					net.layers[i].neurals[j].input = make([]float64, p + 1)
					net.layers[i].neurals[j].weight = make([]float64, p + 1)
					net.layers[i].neurals[j].input[0] = 1. //biais
					fmt.Println(net.layers[i].neurals[j])
					continue
				}
				net.layers[i].neurals[j].input = make([]float64, len(net.layers[i - 1].neurals) + 1)
				net.layers[i].neurals[j].weight = make([]float64, len(net.layers[i - 1].neurals) + 1)
				net.layers[i].neurals[j].input[0] = 1 //biais
			}
		}
	}
	return net
}

func (net *NeuralNetworkType) Init() {
	rand.Seed(time.Now().Unix())

	for i, _ := range net.layers {
		for j, _ := range net.layers[i].neurals {
			for k, _ := range net.layers[i].neurals[j].weight {
				net.layers[i].neurals[j].weight[k] = float64(rand.Intn(2) - 1)
			}
		}
	}
}

func (net *NeuralNetworkType) Learn(input []float64, y []float64) error {
	output := net.Cogitate(input)
	if (len(output) != len(y)) {
		return errors.New("output len does not match")
	}

	for i, _ := range output {
		if (output[i] != y[i]) {
			// Retro propagate
			fmt.Println("#### RETROPROPAGATE")
			for j := len(net.layers)-1; j >= 0; j-- {
				if j == len(net.layers) - 1 {
					//output layer
					for k, _ := range net.layers[j].neurals {
						net.layers[j].neurals[k].delta = y[k] - output[k]
						for l, _ := range net.layers[j].neurals[k].weight {
							fmt.Println("w", l, " av modif : ", net.layers[j].neurals[k].weight[l])
							net.layers[j].neurals[k].weight[l] = net.layers[j].neurals[k].weight[l] + ε * net.layers[j].neurals[k].input[l] * net.layers[j].neurals[k].delta
							fmt.Println("w", l, " ap modif : ", net.layers[j].neurals[k].weight[l])
						}

					}
				}
			}
		}
	}
	return nil
}

func (net *NeuralNetworkType) Cogitate(input []float64) (output []float64) {
	for i := range net.layers {
		for j := range net.layers[i].neurals {
			if i == 0 {
				for k := range input {
					net.layers[i].neurals[j].input[k + 1] = input[k]
				}
				net.layers[i].neurals[j].outup = net.layers[i].neurals[j].sum()
				continue
			}else {
				for k := range net.layers[i - 1].neurals {
					net.layers[i].neurals[j].input[k + 1] = net.layers[i - 1].neurals[k].outup
				}
				net.layers[i].neurals[j].outup = net.layers[i].neurals[j].sum()
			}
		}
	}

	output = make([]float64, len(net.layers[len(net.layers) - 1].neurals))

	for i := range net.layers[len(net.layers) - 1].neurals {
		output[i] = net.layers[len(net.layers) - 1].neurals[i].outup
	}

	return output

}

func (neu *NeuralType)sum() float64 {
	sum := 0.
	for i := range neu.input {
		sum += neu.input[i] * neu.weight[i]
	}
	return activited(sum)
}

func (net *NeuralNetworkType)View() {
	fmt.Println("Layers nb: ", len(net.layers))
	for i, lay := range net.layers {
		fmt.Println("Layer", i + 1, ":")
		fmt.Println("Neurals nb: ", len(lay.neurals))
		for j, neu := range lay.neurals {
			fmt.Println("Neural", j + 1, ":")
			for k, x := range neu.input {

				fmt.Print("x", k, "=", x, ";")
			}
			fmt.Print("\n")
			for k, w := range neu.weight {
				fmt.Print("w", k, "=", w, ";")
			}
			fmt.Print("\n")
		}

	}
}

func activited(x float64) (float64) {
	//y := 5.0
	//return 1 / (1 + math.Pow(math.E, -(y) * x))
	//y :=  2 / (1 + math.Pow(math.E,-x) - 1)

	if x > 0. {
		return 1.0
	}
	return 0.0
}