package main

import (
	"github.com/googollee/go-socket.io"
	"log"
	"net/http"
	"./neuralnetwork"
	"strings"
	"strconv"
	"time"
	"fmt"
	"math/rand"
	"sort"
)

const (
	options string = "OPTIONS"
	allow_origin string = "Access-Control-Allow-Origin"
	allow_methods string = "Access-Control-Allow-Methods"
	allow_headers string = "Access-Control-Allow-Headers"
	allow_credentials string = "Access-Control-Allow-Credentials"
	credentials string = "true"
	origin string = "Origin"
	methods string = "POST, GET, OPTIONS, PUT, DELETE, HEAD, PATCH"
	headers string = "Accept, Accept-Encoding, Authorization, Content-Length, Content-Type, X-CSRF-Token"
)

type corsHandler struct {
	h http.Handler
}

func CORS() func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return &corsHandler{h}
	}
}

func (c *corsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if o := r.Header.Get(origin); o != "" {
		w.Header().Set(allow_origin, o)
	} else {
		w.Header().Set(allow_origin, "*")
	}

	w.Header().Set(allow_headers, headers)
	w.Header().Set(allow_methods, methods)
	w.Header().Set(allow_credentials, credentials)

	if r.Method == options {
		w.WriteHeader(http.StatusOK)
		w.Write(nil)
		return
	}

	c.h.ServeHTTP(w, r)
}

var socket socketio.Socket

func NewFish(fish neuralnetwork.NeuralNetworkType) *neuralnetwork.NeuralNetworkType {
	rand.Seed(time.Now().Unix())
	newFish := neuralnetwork.New(1, 3)

	for i, _ := range fish.Layers {
		for j, _ := range fish.Layers[i].Neurals {
			for k, _ := range fish.Layers[i].Neurals[j].Weight {
				newFish.Layers[i].Neurals[j].Weight[k] = fish.Layers[i].Neurals[j].Weight[k]
			}
		}
	}
	i := rand.Intn(len(fish.Layers))
	j := rand.Intn(len(fish.Layers[i].Neurals))
	k := rand.Intn(len(fish.Layers[i].Neurals[j].Weight))
	E := rand.Intn(2)
	if E > 2 {
		newFish.Layers[i].Neurals[j].Weight[k] = (rand.Float64() * 1.)
	}else {
		newFish.Layers[i].Neurals[j].Weight[k] = (rand.Float64() * -1.)
	}

	return newFish
}
func ReproductionFish(fish1 neuralnetwork.NeuralNetworkType, fish2 neuralnetwork.NeuralNetworkType) *neuralnetwork.NeuralNetworkType {
	rand.Seed(time.Now().Unix())
	newFish := neuralnetwork.New(1, 3)

	for i, _ := range fish1.Layers {
		for j, _ := range fish1.Layers[i].Neurals {
			for k, _ := range fish1.Layers[i].Neurals[j].Weight {
				E := rand.Intn(2)
				if E > 2 {
					newFish.Layers[i].Neurals[j].Weight[k] = fish1.Layers[i].Neurals[j].Weight[k]
				}else {
					newFish.Layers[i].Neurals[j].Weight[k] = fish2.Layers[i].Neurals[j].Weight[k]
				}
			}
		}
	}

	return newFish
}

type FishType struct {
	Score int64
	Brain neuralnetwork.NeuralNetworkType
}

type ByScore []FishType

func (b ByScore)Len() int {
	return len(b)
}
func (a ByScore) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a ByScore) Less(i, j int) bool {
	return a[i].Score > a[j].Score
}

func Fish() {
	var score int64
	curFish := neuralnetwork.New(1, 3)
	curFish.Init()

	var allFishes []FishType

	go func() {
		gen := 0
		c := time.Tick(5 * time.Second)
		for range c {
			gen += 1

			if gen % 5 == 0 {
				fmt.Println("\n\n###### new Generation ######")
				allFishes = append(allFishes, FishType{score, *curFish})
				sort.Sort(ByScore(allFishes))

				fmt.Println("Max Score: ", allFishes[0].Score)

				score = 0

				curFish = ReproductionFish(allFishes[0].Brain,allFishes[1].Brain)

				for i, _ := range curFish.Layers {
					for j, _ := range curFish.Layers[i].Neurals {
						fmt.Println("\nNeurals:", j)
						for k, _ := range curFish.Layers[i].Neurals[j].Weight {
							fmt.Print("Weight:", i, curFish.Layers[i].Neurals[j].Weight[k], ";")
						}
					}
				}

			}else {
				fmt.Println("\n\n###### new Mutation ######")
				allFishes = append(allFishes, FishType{score, *curFish})
				sort.Sort(ByScore(allFishes))

				fmt.Println("Max Score: ", allFishes[0].Score)

				score = 0

				curFish = NewFish(allFishes[rand.Intn(len(allFishes))].Brain)

				for i, _ := range curFish.Layers {
					for j, _ := range curFish.Layers[i].Neurals {
						fmt.Println("\nNeurals:", j)
						for k, _ := range curFish.Layers[i].Neurals[j].Weight {
							fmt.Print("Weight:", i, curFish.Layers[i].Neurals[j].Weight[k], ";")
						}
					}
				}
			}
		}
	}()

	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	server.On("connection", func(so socketio.Socket) {
		socket = so
		log.Println("on connection")
		so.Join("canal")
		so.On("canal score", func(msg string) {
			i, _ := strconv.ParseInt(msg, 10, 64)
			score += i
			fmt.Println("Miam")
		})
		so.On("canal a", func(msg string) {
			param := strings.Split(msg, ";")
			input := make([]float64, 1)
			for i := range param {
				input[i], err = strconv.ParseFloat(param[i], 64)
				if err != nil {
					fmt.Println(" err input", err)
				}

			}
			//fmt.Println("input", input)
			out := curFish.Cogitate(input)

			so.Emit("canal a", strconv.FormatFloat(out[0], 'f', 6, 64) + ";" + strconv.FormatFloat(out[1], 'f', 6, 64) + ";" + strconv.FormatFloat(out[2], 'f', 6, 64))

			//log.Println("emit:", msg)
		})
		so.On("disconnection", func() {
			log.Println("on disconnect")
		})
	})
	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})

	mux := http.NewServeMux()
	mux.Handle("/socket.io/", server)
	mux.Handle("/", http.FileServer(http.Dir("./asset")))


	//http.Handle("/socket.io/", server)
	//http.Handle("/", http.FileServer(http.Dir("./asset")))
	http.Handle("/", CORS()(mux))
	log.Println("Serving at localhost:5000...")
	log.Fatal(http.ListenAndServe(":5000", nil))
}