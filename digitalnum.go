package main

import (
	"./neuralnetwork"
	"fmt"
	"time"
	"math/rand"
	"reflect"
)

func exampleDigitalNum(){


	Xor := neuralnetwork.New(7, 10)

	Xor.Init()

	type example struct {
		Input  []float64
		Output []float64
		Good   bool
	}
	type examples struct {
		Ex []example
	}

	//xorX

	EX := new(examples)
	ex := example{[]float64{0., 0., 0., 0., 0., 0., 0.}, []float64{0., 0., 0., 0., 0., 0., 0., 0., 0., 1.}, false} //0
	EX.Ex = append(EX.Ex, ex)
	ex = example{[]float64{0., 1., 0., 1., 0., 0., 0.}, []float64{1., 0., 0., 0., 0., 0., 0., 0., 0., 0.}, false} //1
	EX.Ex = append(EX.Ex, ex)
	ex = example{[]float64{0., 1., 1., 0., 1., 1., 1.}, []float64{0., 1., 0., 0., 0., 0., 0., 0., 0., 0.}, false} //2
	EX.Ex = append(EX.Ex, ex)
	ex = example{[]float64{0., 1., 0., 1., 1., 1., 1.}, []float64{0., 0., 1., 0., 0., 0., 0., 0., 0., 0.}, false} //3
	EX.Ex = append(EX.Ex, ex)
	ex = example{[]float64{1., 1., 0., 1., 0., 1., 0.}, []float64{0., 0., 0., 1., 0., 0., 0., 0., 0., 0.}, false}//4
	EX.Ex = append(EX.Ex, ex)
	ex = example{[]float64{1., 0., 0., 1., 1., 1., 1.}, []float64{0., 0., 0., 0., 1., 0., 0., 0., 0., 0.}, false} //5
	EX.Ex = append(EX.Ex, ex)
	ex = example{[]float64{1., 0., 1., 1., 1., 1., 1.}, []float64{0., 0., 0., 0., 0., 1., 0., 0., 0., 0.}, false} //6
	EX.Ex = append(EX.Ex, ex)
	ex = example{[]float64{0., 1., 0., 1., 1., 0., 0.}, []float64{0., 0., 0., 0., 0., 0., 1., 0., 0., 0.}, false} //7
	EX.Ex = append(EX.Ex, ex)
	ex = example{[]float64{1., 1., 1., 1., 1., 0., 1.}, []float64{0., 0., 0., 0., 0., 0., 0., 1., 0., 0.}, false} //8
	EX.Ex = append(EX.Ex, ex)
	ex = example{[]float64{1., 1., 0., 1., 1., 1., 1.}, []float64{0., 0., 0., 0., 0., 0., 0., 0., 1., 0.}, false} //9
	EX.Ex = append(EX.Ex, ex)

	rand.Seed(time.Now().Unix())
	fmt.Println(EX)

	k := 0
	for i := 0; i < 10000; i++ {

		fmt.Println("Ɛ:", neuralnetwork.Ɛ, " i:", i * 4, Xor)

		j := rand.Intn(len(EX.Ex))
		b, err := Xor.Learn(EX.Ex[j].Input, EX.Ex[j].Output)
		EX.Ex[j].Good = b

		if err != nil {
			fmt.Println(err)
		}
		B := true
		for k := range EX.Ex {
			B = (B && EX.Ex[k].Good)
		}
		if B {
			k += 1
		}
		if k > 10 {
			fmt.Println("END:", i)
			fmt.Println(time.Now())
			//Xor.Sav("NUM" + strconv.Itoa(i))
			break
		}
		//neuralnetwork.Ɛ -= 0.0000001

	}

	for i := range EX.Ex {
		if reflect.DeepEqual(Xor.Cogitate(EX.Ex[i].Input), EX.Ex[i].Output) {
			fmt.Println(i, "=", i)
		}
	}

}
