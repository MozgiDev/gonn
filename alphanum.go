package main

import (
	"fmt"
	"time"
	"math/rand"
	"./neuralnetwork"
	"reflect"
)

func exampleAlphaNum() {

	EXalpha := new(examples)
	ex := example{[]float64{
		0, 0, 0, 0, 0,
		0, 0, 0, 0, 0,
		0, 0, 0, 0, 0,
		0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, }, []float64{0, 0, 0, 0, 0, 0}, false} //nil
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{[]float64{
		0, 0, 1, 0, 0,
		0, 1, 0, 1, 0,
		1, 0, 0, 0, 1,
		1, 1, 1, 1, 1,
		1, 0, 0, 0, 1, }, []float64{0, 0, 0, 0, 0, 1}, false} //A
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{[]float64{
		1, 1, 1, 1, 0,
		1, 0, 0, 0, 1,
		1, 1, 1, 1, 0,
		1, 0, 0, 0, 1,
		1, 1, 1, 1, 0, }, []float64{0, 0, 0, 0, 1, 0}, false} //B
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{[]float64{
		0, 1, 1, 1, 1,
		1, 0, 0, 0, 0,
		1, 0, 0, 0, 0,
		1, 0, 0, 0, 0,
		0, 1, 1, 1, 1, }, []float64{0, 0, 0, 0, 1, 1}, false} //C
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{[]float64{
		1, 1, 1, 1, 0,
		1, 0, 0, 0, 1,
		1, 0, 0, 0, 1,
		1, 0, 0, 0, 1,
		1, 1, 1, 1, 0, }, []float64{0, 0, 0, 1, 0, 0}, false} //D
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{[]float64{
		1, 1, 1, 1, 1,
		1, 0, 0, 0, 0,
		1, 1, 1, 1, 0,
		1, 0, 0, 0, 0,
		1, 1, 1, 1, 1, }, []float64{0, 0, 0, 1, 0, 1}, false} //E
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{[]float64{
		1, 1, 1, 1, 1,
		1, 0, 0, 0, 0,
		1, 1, 1, 1, 0,
		1, 0, 0, 0, 0,
		1, 0, 0, 0, 0, }, []float64{0, 0, 0, 1, 1, 0}, false} //F
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{[]float64{
		0, 1, 1, 1, 1,
		1, 0, 0, 0, 0,
		1, 0, 0, 1, 1,
		1, 0, 0, 0, 1,
		0, 1, 1, 1, 1, }, []float64{0, 0, 0, 1, 1, 1}, false} //G
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{[]float64{
		1, 0, 0, 0, 1,
		1, 0, 0, 0, 1,
		1, 1, 1, 1, 1,
		1, 0, 0, 0, 1,
		1, 0, 0, 0, 1, }, []float64{0, 0, 1, 0, 0, 0}, false} //H
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{[]float64{
		1, 1, 1, 1, 1,
		0, 0, 1, 0, 0,
		0, 0, 1, 0, 0,
		0, 0, 1, 0, 0,
		1, 1, 1, 1, 1, }, []float64{0, 0, 1, 0, 0, 1}, false} //I
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{[]float64{
		1, 1, 1, 1, 1,
		0, 0, 0, 1, 0,
		0, 0, 0, 1, 0,
		1, 0, 0, 1, 0,
		0, 1, 1, 0, 0, }, []float64{0, 0, 1, 0, 1, 0}, false} //J
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{[]float64{
		1, 0, 0, 0, 1,
		1, 0, 0, 1, 0,
		1, 1, 1, 0, 0,
		1, 0, 0, 1, 0,
		1, 0, 0, 0, 1, }, []float64{0, 0, 1, 0, 1, 1}, false} //K
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{[]float64{
		1, 0, 0, 0, 0,
		1, 0, 0, 0, 0,
		1, 0, 0, 0, 0,
		1, 0, 0, 0, 0,
		1, 1, 1, 1, 1, }, []float64{0, 0, 1, 1, 0, 0}, false} //L
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{[]float64{
		1, 0, 0, 0, 1,
		1, 1, 0, 1, 1,
		1, 0, 1, 0, 1,
		1, 0, 0, 0, 1,
		1, 0, 0, 0, 1, }, []float64{0, 0, 1, 1, 0, 1}, false} //M
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{[]float64{
		1, 0, 0, 0, 1,
		1, 1, 0, 0, 1,
		1, 0, 1, 0, 1,
		1, 0, 0, 1, 1,
		1, 0, 0, 0, 1, }, []float64{0, 0, 1, 1, 1, 0}, false}//N
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{[]float64{
		0, 1, 1, 1, 0,
		1, 0, 0, 0, 1,
		1, 0, 0, 0, 1,
		1, 0, 0, 0, 1,
		0, 1, 1, 1, 0, }, []float64{0, 0, 1, 1, 1, 1}, false} //O
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{[]float64{
		1, 1, 1, 1, 0,
		1, 0, 0, 0, 1,
		1, 1, 1, 1, 0,
		1, 0, 0, 0, 0,
		1, 0, 0, 0, 0, }, []float64{0, 1, 0, 0, 0, 0}, false} //P
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{[]float64{
		0, 1, 1, 1, 0,
		1, 0, 0, 0, 1,
		1, 0, 1, 0, 1,
		1, 0, 0, 1, 0,
		0, 1, 1, 0, 1, }, []float64{0, 1, 0, 0, 0, 1}, false} //Q
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{[]float64{
		1, 1, 1, 1, 0,
		1, 0, 0, 0, 1,
		1, 1, 1, 1, 0,
		1, 0, 0, 0, 1,
		1, 0, 0, 0, 1, }, []float64{0, 1, 0, 0, 1, 0}, false} //R
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{[]float64{
		0, 1, 1, 1, 1,
		1, 0, 0, 0, 0,
		0, 1, 1, 1, 0,
		0, 0, 0, 0, 1,
		1, 1, 1, 1, 0, }, []float64{0, 1, 0, 0, 1, 1}, false} //S
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{[]float64{
		1, 1, 1, 1, 1,
		0, 0, 1, 0, 0,
		0, 0, 1, 0, 0,
		0, 0, 1, 0, 0,
		0, 0, 1, 0, 0, }, []float64{0, 1, 0, 1, 0, 0}, false} //T
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{[]float64{
		1, 0, 0, 0, 1,
		1, 0, 0, 0, 1,
		1, 0, 0, 0, 1,
		1, 0, 0, 0, 1,
		0, 1, 1, 1, 0, }, []float64{0, 1, 0, 1, 0, 1}, false} //U
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{[]float64{
		1, 0, 0, 0, 1,
		1, 0, 0, 0, 1,
		1, 0, 0, 0, 1,
		0, 1, 0, 1, 0,
		0, 0, 1, 0, 0, }, []float64{0, 1, 0, 1, 1, 0}, false} //V
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{[]float64{
		1, 0, 0, 0, 1,
		1, 0, 0, 0, 1,
		1, 0, 1, 0, 1,
		1, 1, 0, 1, 1,
		1, 0, 0, 0, 1, }, []float64{0, 1, 0, 1, 1, 1}, false} //W
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{[]float64{
		1, 0, 0, 0, 1,
		0, 1, 0, 1, 0,
		0, 0, 1, 0, 0,
		0, 1, 0, 1, 0,
		1, 0, 0, 0, 1, }, []float64{0, 1, 1, 0, 0, 0}, false}//X
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{[]float64{
		1, 0, 0, 0, 1,
		0, 1, 0, 1, 0,
		0, 0, 1, 0, 0,
		0, 0, 1, 0, 0,
		0, 0, 1, 0, 0, }, []float64{0, 1, 1, 0, 0, 1}, false} //Y
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{[]float64{
		1, 1, 1, 1, 1,
		0, 0, 0, 1, 0,
		0, 0, 1, 0, 0,
		0, 1, 0, 0, 0,
		1, 1, 1, 1, 1, }, []float64{0, 1, 1, 0, 1, 0}, false} //Z
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{[]float64{
		0, 1, 1, 1, 0,
		1, 0, 0, 1, 1,
		1, 0, 1, 0, 1,
		1, 1, 0, 0, 1,
		0, 1, 1, 1, 0, }, []float64{0, 1, 1, 0, 1, 1}, false} //0
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{[]float64{
		0, 1, 1, 0, 0,
		0, 0, 1, 0, 0,
		0, 0, 1, 0, 0,
		0, 0, 1, 0, 0,
		0, 1, 1, 1, 0, }, []float64{0, 1, 1, 1, 0, 0}, false} //1
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{[]float64{
		1, 1, 1, 1, 0,
		0, 0, 0, 0, 1,
		0, 1, 1, 1, 0,
		1, 0, 0, 0, 0,
		1, 1, 1, 1, 1, }, []float64{0, 1, 1, 1, 0, 1}, false} //2
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{[]float64{
		1, 1, 1, 1, 0,
		0, 0, 0, 0, 1,
		0, 0, 1, 1, 0,
		0, 0, 0, 0, 1,
		1, 1, 1, 1, 1, }, []float64{0, 1, 1, 1, 1, 1}, false} //3
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{[]float64{
		0, 0, 1, 1, 0,
		0, 1, 0, 1, 0,
		1, 0, 0, 1, 0,
		1, 1, 1, 1, 1,
		0, 0, 0, 1, 0, }, []float64{1, 0, 0, 0, 0, 0}, false} //4
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{[]float64{
		1, 1, 1, 1, 1,
		1, 0, 0, 0, 0,
		1, 1, 1, 1, 0,
		0, 0, 0, 0, 1,
		1, 1, 1, 1, 0, }, []float64{1, 0, 0, 0, 0, 1}, false} //5
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{[]float64{
		0, 1, 1, 1, 1,
		1, 0, 0, 0, 0,
		1, 1, 1, 1, 0,
		1, 0, 0, 0, 1,
		0, 1, 1, 1, 0, }, []float64{1, 0, 0, 0, 1, 0}, false} //6
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{[]float64{
		1, 1, 1, 1, 1,
		0, 0, 0, 0, 1,
		0, 0, 0, 1, 0,
		0, 0, 1, 0, 0,
		0, 1, 0, 0, 0, }, []float64{1, 0, 0, 0, 1, 1}, false} //7
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{[]float64{
		0, 1, 1, 1, 0,
		1, 0, 0, 0, 1,
		0, 1, 1, 1, 0,
		1, 0, 0, 0, 1,
		0, 1, 1, 1, 0, }, []float64{1, 0, 0, 1, 0, 0}, false} //8
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{[]float64{
		0, 1, 1, 1, 0,
		1, 0, 0, 0, 1,
		0, 1, 1, 1, 1,
		0, 0, 0, 0, 1,
		1, 1, 1, 1, 0, }, []float64{1, 0, 0, 1, 0, 1}, false} //9
	EXalpha.Ex = append(EXalpha.Ex, ex)

	Xor := neuralnetwork.New(25, 25, 6)

	Xor.Init()
	rand.Seed(time.Now().Unix())
	//fmt.Println(EXalpha)

	k := 0
	fmt.Println(time.Now())
	neuralnetwork.Ɛ = 0.1
	for i := 0; i < 100000000000; i++ {

		//fmt.Println("Ɛ:", neuralnetwork.Ɛ, " i:", i * 4, Xor)

		j := rand.Intn(len(EXalpha.Ex))
		b, err := Xor.Learn(EXalpha.Ex[j].Input, EXalpha.Ex[j].Output)
		EXalpha.Ex[j].Good = b

		if err != nil {
			fmt.Println(err)
		}

		B := true
		for k := range EXalpha.Ex {
			B = (B && EXalpha.Ex[k].Good)
		}
		if B {
			k += 1
		}else {
			k = 0
		}
		if k > 100 {
			fmt.Println("END:", i)
			fmt.Println(time.Now())
			//Xor.Sav("NUM" + strconv.Itoa(i))
			break
		}
		//neuralnetwork.Ɛ -= 0.0000001

	}

	for i := range EXalpha.Ex {
		if reflect.DeepEqual(Xor.Cogitate(EXalpha.Ex[i].Input), EXalpha.Ex[i].Output) {
			fmt.Println(i, "=", i)
		}
	}
}
