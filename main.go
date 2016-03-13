package main

import (
	"./neuralnetwork"
	"fmt"
	"time"
)

func main() {
	Xor := neuralnetwork.New(2, 1)

	Xor.Init()
	fmt.Println(Xor)
	//xorX
	xorX1 := make([]float64, 2)
	xorX1[0] = 0.0
	xorX1[1] = 0.0
	xorY1 := make([]float64, 1)
	xorY1[0] = 0.

	xorX2 := make([]float64, 2)
	xorX2[0] = 0.0
	xorX2[1] = 1.0
	xorY2 := make([]float64, 1)
	xorY2[0] = 1.

	xorX3 := make([]float64, 2)
	xorX3[0] = 1.0
	xorX3[1] = 0.0
	xorY3 := make([]float64, 1)
	xorY3[0] = 1.

	xorX4 := make([]float64, 2)
	xorX4[0] = 1.0
	xorX4[1] = 1.0
	xorY4 := make([]float64, 1)
	xorY4[0] = 0.

	for i := 0; i < 1; i++ {
		Xor.Learn(xorX1,xorY1)
		Xor.Learn(xorX2,xorY2)
		Xor.Learn(xorX3,xorY3)
		Xor.Learn(xorX4,xorY4)
	}

	fmt.Println()
	fmt.Println()
	fmt.Println()
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
