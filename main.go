package main

import (
	"./neuralnetwork"
	"fmt"
	"time"
	"math/rand"
	"strconv"
)

func main() {
	Xor := neuralnetwork.New(2, 1)

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
	ex := example{[]float64{0., 0.}, []float64{0.}, false}
	EX.Ex = append(EX.Ex, ex)
	ex = example{[]float64{1., 0.}, []float64{1.}, false}
	EX.Ex = append(EX.Ex, ex)
	ex = example{[]float64{0., 1.}, []float64{1.}, false}
	EX.Ex = append(EX.Ex, ex)
	ex = example{[]float64{1., 1.}, []float64{0.}, false}
	EX.Ex = append(EX.Ex, ex)

	rand.Seed(time.Now().Unix())
	fmt.Println(EX)

	for i := 0; i < 10000000000000000; i++ {
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
			fmt.Println("END:",i)
			Xor.Sav("XORFUN" + strconv.Itoa(i))
			break
		}
		//neuralnetwork.Ɛ -= 0.0000001

	}

	fmt.Println(time.Now())


	////OU
	//xorX1 := make([]float64, 2)
	//xorX1[0] = 0.0
	//xorX1[1] = 0.0
	//ouC1 := 0.0
	//
	//xorX2 := make([]float64, 2)
	//xorX2[0] = 0.0
	//xorX2[1] = 1.0
	//ouC2 := 1.0
	//
	//xorX3 := make([]float64, 2)
	//xorX3[0] = 1.0
	//xorX3[1] = 0.0
	//ouC3 := 1.0
	//
	//xorX4 := make([]float64, 2)
	//xorX4[0] = 1.0
	//xorX4[1] = 1.0
	//ouC4 := 1.0
	//
	//
	//ou := neural.New(2)
	//ou.Init()
	//log.Println("########-OU-########")
	//for i := 0; i < 2000; i++ {
	//	ou.LearnIt(xorX1, ouC1)
	//	ou.LearnIt(xorX2, ouC2)
	//	ou.LearnIt(xorX3, ouC3)
	//	ou.LearnIt(xorX4, ouC4)
	//}
	//fmt.Println()
	//
	//etX1 := make([]float64, 2)
	//etX1[0] = 2.0
	//etX1[1] = 0.0
	//etC1 := 1.0
	//
	//etX2 := make([]float64, 2)
	//etX2[0] = 0.0
	//etX2[1] = 3.0
	//etC2 := 0.0
	//
	//etX3 := make([]float64, 2)
	//etX3[0] = 3.0
	//etX3[1] = 0.0
	//etC3 := 0.0
	//
	//etX4 := make([]float64, 2)
	//etX4[0] = 1.0
	//etX4[1] = 1.0
	//etC4 := 1.0
	//
	//et := neural.New(2)
	//et.Init()
	//log.Println("########-ET-########")
	//for i := 0; i < 0; i++ {
	//	et.LearnIt(etX1, etC1)
	//	et.LearnIt(etX2, etC2)
	//	et.LearnIt(etX3, etC3)
	//	et.LearnIt(etX4, etC4)
	//}
	//fmt.Println(time.Now())
	//ET
	//etX1 := make([]float64, 2)
	//etX1[0] = 0.0
	//etX1[1] = 0.0
	//etC1 := 0.0
	//
	//etX2 := make([]float64, 2)
	//etX2[0] = 1.0
	//etX2[1] = 0.0
	//etC2 := 0.0
	//
	//etX3 := make([]float64, 2)
	//etX3[0] = 1.0
	//etX3[1] = 1.0
	//etC3 := 1.0
	//
	//etX4 := make([]float64, 2)
	//etX4[0] = 0.0
	//etX4[1] = 1.0
	//etC4 := 0.0
	//
	//
	//et := neural.New(2)
	//et.Init()
	//log.Println("########-ET-########")
	//for i := 0; i < 10; i++ {
	//	et.LearnIt(etX1, etC1)
	//	et.LearnIt(etX2, etC2)
	//	et.LearnIt(etX3, etC3)
	//	et.LearnIt(etX4, etC4)
	//}
	//et.Input[1] = 1.
	//et.Input[2] = 1.
	//fmt.Println(et.Input[1],".",et.Input[2],"=",et.Cogitate())
	//
	//et.Input[1] = 0.
	//et.Input[2] = 1.
	//fmt.Println(et.Input[1],".",et.Input[2],"=",et.Cogitate())

}
