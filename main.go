package main

type example struct {
	Input  []float64
	Output []float64
	Good   bool
}
type examples struct {
	Ex []example
}

func main() {

	Alphabet()
	//exampleAlphaNum()


	//Fish()


	//Xor := neuralnetwork.New(2, 3)
	//
	//Xor.Init()
	//
	//type example struct {
	//	Input  []float64
	//	Output []float64
	//	Good   bool
	//}
	//type examples struct {
	//	Ex []example
	//}
	//
	////xorX
	//
	//EX := new(examples)
	//ex := example{[]float64{0, 1}, []float64{0, 1, 0}, false} //0
	//EX.Ex = append(EX.Ex, ex)
	//ex = example{[]float64{1, 0}, []float64{1, 0, 0}, false} //1
	//EX.Ex = append(EX.Ex, ex)
	//ex = example{[]float64{0, 0}, []float64{1, 0, 0}, false} //1
	//EX.Ex = append(EX.Ex, ex)
	//ex = example{[]float64{0, 0}, []float64{1, 0, 1}, false} //1
	//EX.Ex = append(EX.Ex, ex)
	//rand.Seed(time.Now().Unix())
	////fmt.Println(EXalpha)
	//
	//k := 0
	//fmt.Println(time.Now())
	//for i := 0; i < 100000000000; i++ {
	//	j := rand.Intn(len(EX.Ex))
	//	b, err := Xor.Learn(EX.Ex[j].Input, EX.Ex[j].Output)
	//	EX.Ex[j].Good = b
	//
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	B := true
	//	for k := range EX.Ex {
	//		B = (B && EX.Ex[k].Good)
	//	}
	//	if B {
	//		k += 1
	//	}else {
	//		k = 0
	//	}
	//	if k > 1 {
	//		fmt.Println("END:", i)
	//		fmt.Println(Xor)
	//		fmt.Println(time.Now())
	//		//Xor.Sav("NUM" + strconv.Itoa(i))
	//		break
	//	}
	//	//neuralnetwork.Ɛ -= 0.0000001
	//}
	//
	//for i := range EX.Ex {
	//	if reflect.DeepEqual(Xor.Cogitate(EX.Ex[i].Input), EX.Ex[i].Output) {
	//		fmt.Println(i, "=", i)
	//	}
	//}


	//Xor := neuralnetwork.New(2, 1)
	//
	//Xor.Init()
	//
	//type example struct {
	//	Input  []float64
	//	Output []float64
	//	Good   bool
	//}
	//type examples struct {
	//	Ex []example
	//}
	//
	////xorX
	//
	//EX := new(examples)
	//ex := example{[]float64{0., 0.}, []float64{0.}, false}
	//EX.Ex = append(EX.Ex, ex)
	//ex = example{[]float64{1., 0.}, []float64{1.}, false}
	//EX.Ex = append(EX.Ex, ex)
	//ex = example{[]float64{0., 1.}, []float64{1.}, false}
	//EX.Ex = append(EX.Ex, ex)
	//ex = example{[]float64{1., 1.}, []float64{0.}, false}
	//EX.Ex = append(EX.Ex, ex)
	//
	//rand.Seed(time.Now().Unix())
	//fmt.Println(EX)
	//
	//for i := 0; i < 10000000000000000; i++ {
	//	fmt.Println("Ɛ:", neuralnetwork.Ɛ, " i:", i * 4, Xor)
	//
	//	j := rand.Intn(len(EX.Ex))
	//	b, err := Xor.Learn(EX.Ex[j].Input, EX.Ex[j].Output)
	//	EX.Ex[j].Good = b
	//
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	B := true
	//	for k := range EX.Ex {
	//		B = (B && EX.Ex[k].Good)
	//	}
	//	if B {
	//		fmt.Println("END:",i)
	//		Xor.Sav("XORFUN" + strconv.Itoa(i))
	//		break
	//	}
	//	//neuralnetwork.Ɛ -= 0.0000001
	//
	//}
	//
	//fmt.Println(time.Now())


}
