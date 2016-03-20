package neuralnetwork

import (
	"errors"
	"reflect"
	"sync"
)

func (net *NeuralNetworkType) Learn(Input []float64, y []float64) (bool, error) {
	Output := net.Cogitate(Input)
	if (len(Output) != len(y)) {
		return false, errors.New("Output len does not match")
	}

	if (reflect.DeepEqual(y, Output)) != true {
		// Retro propagate
		//fmt.Println("#### RETROPROPAGATE")
		for j := len(net.Layers) - 1; j >= 0; j-- {
			if j == len(net.Layers) - 1 {
				//Output layer
				for k, _ := range net.Layers[j].Neurals {
					net.Layers[j].Neurals[k].Delta = y[k] - Output[k]
					for l, _ := range net.Layers[j].Neurals[k].Weight {
						//fmt.Println("w", l, " av modif : ", net.Layers[j].Neurals[k].Weight[l])
						net.Layers[j].Neurals[k].Weight[l] = net.Layers[j].Neurals[k].Weight[l] + Ɛ * net.Layers[j].Neurals[k].Input[l] * net.Layers[j].Neurals[k].Delta
						//fmt.Println("w", l, " ap modif : ", net.Layers[j].Neurals[k].Weight[l])
					}

				}
			} else {
				for k, _ := range net.Layers[j].Neurals {
					temp := 0.0
					for l, _ := range net.Layers[j + 1].Neurals {
						temp += net.Layers[j + 1].Neurals[l].Weight[k] * net.Layers[j + 1].Neurals[l].Delta
					}
					net.Layers[j].Neurals[k].Delta = net.Layers[j].Neurals[k].Output * (1 - net.Layers[j].Neurals[k].Output) * temp
					for l, _ := range net.Layers[j].Neurals[k].Weight {
						//fmt.Println("w", l, " av modif : ", net.Layers[j].Neurals[k].Weight[l])
						net.Layers[j].Neurals[k].Weight[l] = net.Layers[j].Neurals[k].Weight[l] + Ɛ * net.Layers[j].Neurals[k].Input[l] * net.Layers[j].Neurals[k].Delta
						//fmt.Println("w", l, " ap modif : ", net.Layers[j].Neurals[k].Weight[l])
					}

				}
			}
		}
		return false, nil
	} else {
		//fmt.Println("#### GOOD")
		return true, nil
	}

	return false, nil
}


func (net *NeuralNetworkType) LearnNew(Input []float64, y []float64) (bool, error) {
	Output := net.Cogitate(Input)
	if (len(Output) != len(y)) {
		return false, errors.New("Output len does not match")
	}

	if (reflect.DeepEqual(y, Output)) != true {
		// Retro propagate
		//fmt.Println("#### RETROPROPAGATE")


		//calcul delta
		j := len(net.Layers) - 1
		var wFistCouchDelta sync.WaitGroup
		wFistCouchDelta.Add(len(net.Layers[j].Neurals))
		for k := range net.Layers[j].Neurals {
			go func(j int, k int) {
				defer wFistCouchDelta.Done()
				net.Layers[j].Neurals[k].Delta = y[k] - Output[k]
			}(j, k)
		}
		wFistCouchDelta.Wait()

		for j := len(net.Layers) - 2; j >= 0; j-- {
			for k, _ := range net.Layers[j].Neurals {
				func(j int, k int) {

				}(j, k)
				temp := 0.0
				var wDelta sync.WaitGroup
				wDelta.Add(len(net.Layers[j + 1].Neurals))
				for l, _ := range net.Layers[j + 1].Neurals {
					go func(j int, k int, l int) {
						defer wDelta.Done()
						temp += net.Layers[j + 1].Neurals[l].Weight[k] * net.Layers[j + 1].Neurals[l].Delta
					}(j, k, l)
				}
				wDelta.Wait()
				net.Layers[j].Neurals[k].Delta = net.Layers[j].Neurals[k].Output * (1 - net.Layers[j].Neurals[k].Output) * temp

			}

		}
		//calcul delta

		var wDelta sync.WaitGroup
		for j := len(net.Layers) - 1; j >= 0; j-- {
			if j == len(net.Layers) - 1 {
				//Output layer
				for k, _ := range net.Layers[j].Neurals {
					for l, _ := range net.Layers[j].Neurals[k].Weight {
						wDelta.Add(1)
						go func(j int, k int, l int) {
							defer wDelta.Done()
							net.Layers[j].Neurals[k].Weight[l] = net.Layers[j].Neurals[k].Weight[l] + Ɛ * net.Layers[j].Neurals[k].Input[l] * net.Layers[j].Neurals[k].Delta
						}(j, k, l)
					}
				}

			} else {

				for k, _ := range net.Layers[j].Neurals {
					for l, _ := range net.Layers[j].Neurals[k].Weight {
						wDelta.Add(1)
						go func(j int, k int, l int) {
							defer wDelta.Done()
							net.Layers[j].Neurals[k].Weight[l] = net.Layers[j].Neurals[k].Weight[l] + Ɛ * net.Layers[j].Neurals[k].Input[l] * net.Layers[j].Neurals[k].Delta
						}(j, k, l)
					}
				}
			}
		}
		wDelta.Wait()

		return false, nil
	} else {
		//xfmt.Println("#### GOOD")
		return true, nil
	}
	return false, nil
}