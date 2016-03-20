package neuralnetwork

import (
	"math/rand"
	"time"
	"fmt"
	"math"
	"encoding/json"
	"os"
)

type NeuralType  struct {
	Input  []float64
	Weight []float64
	Delta  float64
	Output float64
}

type LayerType struct {
	Neurals []NeuralType
}

type NeuralNetworkType struct {
	Layers []LayerType
}

var Ɛ = 1.

func New(param ...int) *NeuralNetworkType {
	net := new(NeuralNetworkType)
	net.Layers = make([]LayerType, len(param))
	if len(param) > 1 {
		for i, p := range param {
			net.Layers[i].Neurals = make([]NeuralType, p)
			for j, _ := range net.Layers[i].Neurals {
				if i == 0 {
					net.Layers[i].Neurals[j].Input = make([]float64, p + 1)
					net.Layers[i].Neurals[j].Weight = make([]float64, p + 1)
					net.Layers[i].Neurals[j].Input[0] = 1. //biais
					continue
				}
				net.Layers[i].Neurals[j].Input = make([]float64, len(net.Layers[i - 1].Neurals) + 1)
				net.Layers[i].Neurals[j].Weight = make([]float64, len(net.Layers[i - 1].Neurals) + 1)
				net.Layers[i].Neurals[j].Input[0] = 1 //biais
			}
		}
	}
	return net
}

func (net *NeuralNetworkType) Init() {
	rand.Seed(time.Now().Unix())

	for i, _ := range net.Layers {
		for j, _ := range net.Layers[i].Neurals {
			for k, _ := range net.Layers[i].Neurals[j].Weight {
				net.Layers[i].Neurals[j].Weight[k] = rand.Float64() * 6 - 3
			}
		}
	}
}


func (net *NeuralNetworkType) Cogitate(Input []float64) (Output []float64) {
	for i := range net.Layers {
		for j := range net.Layers[i].Neurals {
			if i == 0 {
				for k := range Input {
					net.Layers[i].Neurals[j].Input[k + 1] = Input[k]
				}
				net.Layers[i].Neurals[j].Output = net.Layers[i].Neurals[j].sum()
				continue
			}else {
				for k := range net.Layers[i - 1].Neurals {
					net.Layers[i].Neurals[j].Input[k + 1] = net.Layers[i - 1].Neurals[k].Output
				}
				net.Layers[i].Neurals[j].Output = net.Layers[i].Neurals[j].sum()
			}
		}
	}

	Output = make([]float64, len(net.Layers[len(net.Layers) - 1].Neurals))

	for i := range net.Layers[len(net.Layers) - 1].Neurals {
		//Output[i] = net.Layers[len(net.Layers) - 1].Neurals[i].Output
		//test
		if (net.Layers[len(net.Layers) - 1].Neurals[i].Output > 0.5) {
			Output[i] = 1;
		} else {
			Output[i] = 0;
		}
		//fin test
	}

	return Output

}

func (neu *NeuralType)sum() float64 {
	sum := 0.
	for i := range neu.Input {
		sum += neu.Input[i] * neu.Weight[i]
	}
	return activited(sum)
}

func (net *NeuralNetworkType)View() {
	fmt.Println("Layers nb: ", len(net.Layers))
	for i, lay := range net.Layers {
		fmt.Println("Layer", i + 1, ":")
		fmt.Println("Neurals nb: ", len(lay.Neurals))
		for j, neu := range lay.Neurals {
			fmt.Println("Neural", j + 1, ":")
			for k, x := range neu.Input {

				fmt.Print("x", k, "=", x, ";")
			}
			fmt.Print("\n")
			for k, w := range neu.Weight {
				fmt.Print("w", k, "=", w, ";")
			}
			fmt.Print("\n")
		}

	}
}

func (neu *NeuralNetworkType) Sav(name string) error {
	jsonneu, err := json.MarshalIndent(neu, "", "\t")
	if err != nil {
		return err
	}

	f, err := os.Create("E:\\neural_netgo\\" + name)
	defer f.Close()
	if err != nil {
		return err
	}
	f.Write(jsonneu)
	f.Close()
	return nil
}

func activited(x float64) (float64) {
	//y := 5.0
	//return 1 / (1 + math.Pow(math.E, -(y) * x))
	//y :=  2 / (1 + math.Pow(math.E,-x) - 1)

	//if x > 0. {
	//	return 1.0
	//}
	//return 0.0


	// La logistic
	y := 1 / (1 + math.Pow(math.E, -x));
	return y

}