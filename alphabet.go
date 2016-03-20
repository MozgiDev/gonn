package main

import (
	"github.com/nfnt/resize"
	"image/jpeg"
	"log"
	"os"
	"fmt"
	"./neuralnetwork"
	"time"
	"math/rand"
	"reflect"
)

func Alphabet() {

	for char := 'a'; char < 'a' + 26; char++ {
		test := GetCharByte(string(char))
		for i := range test {
			fmt.Print(test[i])
			if (i + 1) % 16 == 0 {
				fmt.Print("\n")
			}
		}
	}

	EXalpha := new(examples)
	ex := example{make([]float64, 256), []float64{0, 0, 0, 0, 0, 1}, false} //a
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{GetChar("a"), []float64{0, 0, 0, 0, 0, 1}, false} //a
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{GetChar("AA"), []float64{0, 0, 0, 0, 0, 1}, false} //A
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{GetChar("AAA"), []float64{0, 0, 0, 0, 0, 1}, false} //A
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{GetChar("AAAA"), []float64{0, 0, 0, 0, 0, 1}, false} //A
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{GetChar("AAAA"), []float64{0, 0, 0, 0, 0, 1}, false} //A
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{GetChar("AAAAA"), []float64{0, 0, 0, 0, 0, 1}, false} //A
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{GetChar("AAAAAA"), []float64{0, 0, 0, 0, 0, 1}, false} //A
	EXalpha.Ex = append(EXalpha.Ex, ex)

	ex = example{GetChar("A_M"), []float64{0, 0, 0, 0, 0, 1}, false} //A
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{GetChar("b"), []float64{0, 0, 0, 0, 1, 0}, false} //b
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{GetChar("B_M"), []float64{0, 0, 0, 0, 1, 0}, false} //B
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{GetChar("c"), []float64{0, 0, 0, 0, 1, 1}, false} //c
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{GetChar("C_M"), []float64{0, 0, 0, 0, 1, 1}, false} //C
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{GetChar("d"), []float64{0, 0, 0, 1, 0, 0}, false} //d
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{GetChar("D_M"), []float64{0, 0, 0, 1, 0, 0}, false} //D
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{GetChar("e"), []float64{0, 0, 0, 1, 0, 1}, false} //e
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{GetChar("E_M"), []float64{0, 0, 0, 1, 0, 1}, false} //E
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{GetChar("f"), []float64{0, 0, 0, 1, 1, 0}, false} //f
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{GetChar("F_M"), []float64{0, 0, 0, 1, 1, 0}, false} //F
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{GetChar("g"), []float64{0, 0, 0, 1, 1, 1}, false} //g 
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{GetChar("G_M"), []float64{0, 0, 0, 1, 1, 1}, false} //G
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{GetChar("h"), []float64{0, 0, 1, 0, 0, 0}, false} //h
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{GetChar("H_M"), []float64{0, 0, 1, 0, 0, 0}, false} //H
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{GetChar("i"), []float64{0, 0, 1, 0, 0, 1}, false} //i
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{GetChar("I_M"), []float64{0, 0, 1, 0, 0, 1}, false} //I
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{GetChar("j"), []float64{0, 0, 1, 0, 1, 0}, false} //j
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{GetChar("J_M"), []float64{0, 0, 1, 0, 1, 0}, false} //J
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{GetChar("k"), []float64{0, 0, 1, 0, 1, 1}, false} //k
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{GetChar("K_M"), []float64{0, 0, 1, 0, 1, 1}, false} //K
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{GetChar("l"), []float64{0, 0, 1, 1, 0, 0}, false} //l
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{GetChar("L_M"), []float64{0, 0, 1, 1, 0, 0}, false} //L
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{GetChar("m"), []float64{0, 0, 1, 1, 0, 1}, false} //m
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{GetChar("M_M"), []float64{0, 0, 1, 1, 0, 1}, false} //M
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{GetChar("n"), []float64{0, 0, 1, 1, 1, 0}, false} //n
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{GetChar("N_M"), []float64{0, 0, 1, 1, 1, 0}, false} //N
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{GetChar("o"), []float64{0, 0, 1, 1, 1, 1}, false} //o
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{GetChar("O_M"), []float64{0, 0, 1, 1, 1, 1}, false} //O
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{GetChar("p"), []float64{0, 1, 0, 0, 0, 0}, false} //p
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{GetChar("P_M"), []float64{0, 1, 0, 0, 0, 0}, false} //P
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{GetChar("q"), []float64{0, 1, 0, 0, 0, 1}, false} //q
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{GetChar("Q_M"), []float64{0, 1, 0, 0, 0, 1}, false} //Q
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{GetChar("r"), []float64{0, 1, 0, 0, 1, 0}, false} //r
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{GetChar("R_M"), []float64{0, 1, 0, 0, 1, 0}, false} //R
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{GetChar("s"), []float64{0, 1, 0, 0, 0, 1}, false} //s
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{GetChar("S_M"), []float64{0, 1, 0, 0, 0, 1}, false} //S
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{GetChar("t"), []float64{0, 1, 0, 0, 1, 0}, false} //t
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{GetChar("T_M"), []float64{0, 1, 0, 0, 1, 0}, false} //T
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{GetChar("u"), []float64{0, 1, 0, 0, 1, 1}, false} //u
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{GetChar("UU"), []float64{0, 1, 0, 0, 1, 1}, false} //u
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{GetChar("UUU"), []float64{0, 1, 0, 0, 1, 1}, false} //u
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{GetChar("UUUU"), []float64{0, 1, 0, 0, 1, 1}, false} //u
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{GetChar("UUUUU"), []float64{0, 1, 0, 0, 1, 1}, false} //u
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{GetChar("UUUUUU"), []float64{0, 1, 0, 0, 1, 1}, false} //u
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{GetChar("U_M"), []float64{0, 1, 0, 0, 1, 1}, false} //U
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{GetChar("v"), []float64{0, 1, 0, 1, 0, 0}, false} //v
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{GetChar("v_M"), []float64{0, 1, 0, 1, 0, 0}, false} //V
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{GetChar("w"), []float64{0, 1, 0, 1, 0, 1}, false} //w
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{GetChar("W_M"), []float64{0, 1, 0, 1, 0, 1}, false} //W
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{GetChar("x"), []float64{0, 1, 0, 1, 1, 0}, false} //x
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{GetChar("X_M"), []float64{0, 1, 0, 1, 1, 0}, false} //X
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{GetChar("y"), []float64{0, 1, 0, 1, 1, 1}, false} //y
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{GetChar("Y_M"), []float64{0, 1, 0, 1, 1, 1}, false} //Y
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{GetChar("z"), []float64{0, 1, 1, 0, 0, 0}, false} //z
	EXalpha.Ex = append(EXalpha.Ex, ex)
	ex = example{GetChar("Z_M"), []float64{0, 1, 1, 0, 0, 0}, false} //Z
	EXalpha.Ex = append(EXalpha.Ex, ex)

	Alphabet := neuralnetwork.New(256, 6)

	Alphabet.Init()
	rand.Seed(time.Now().Unix())

	k := 0
	fmt.Println(time.Now())
	neuralnetwork.Ɛ = 0.01
	for i := 0; i < 1000000000000000000; i++ {
		j := rand.Intn(len(EXalpha.Ex))

		b, err := Alphabet.Learn(EXalpha.Ex[j].Input, EXalpha.Ex[j].Output)
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
		//fmt.Println(neuralnetwork.Ɛ)

	}

	var ok int
	for i := range EXalpha.Ex {
		if reflect.DeepEqual(Alphabet.Cogitate(EXalpha.Ex[i].Input), EXalpha.Ex[i].Output) {
			ok++
			fmt.Println(i, "=", i, "ok")
		}else {
			fmt.Println(i, "≠", i, "err")
		}
	}
	fmt.Println((ok / len(EXalpha.Ex)) * 100)
	fmt.Println(Alphabet.Cogitate(GetChar("testA")))
	fmt.Println(Alphabet.Cogitate(GetChar("testU")))
}

func GetChar(char string) []float64 {
	// open "test.jpg"
	file, err := os.Open("alphabet" + string(os.PathSeparator) + char + ".jpg")
	if err != nil {
		log.Fatal(err)
	}

	// decode jpeg into image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	new_image := resize.Resize(16, 16, img, resize.Lanczos3)

	bimg := new_image.Bounds()

	var Input []float64
	for y := bimg.Min.Y; y < bimg.Max.Y; y++ {
		for x := bimg.Min.X; x < bimg.Max.X; x++ {
			r, g, b, _ := new_image.At(x, y).RGBA()
			Input = append(Input, float64((r + g + b) / 3))
		}
	}
	return Input
}

func GetCharByte(char string) []float64 {
	// open "test.jpg"
	file, err := os.Open("alphabet" + string(os.PathSeparator) + char + ".jpg")
	if err != nil {
		log.Fatal(err)
	}

	// decode jpeg into image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	// resize to width 1000 using Lanczos resampling
	// 0 preserve aspect ratio
	new_image := resize.Resize(16, 16, img, resize.Lanczos3)

	bimg := new_image.Bounds()

	var Input []float64
	for y := bimg.Min.Y; y < bimg.Max.Y; y++ {
		for x := bimg.Min.X; x < bimg.Max.X; x++ {
			r, g, b, _ := new_image.At(x, y).RGBA()
			if ((r + g + b) / 3) > 32767 {
				Input = append(Input, float64(0))
			}else {
				Input = append(Input, float64(1))
			}
		}
	}
	return Input
}