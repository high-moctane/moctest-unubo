package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strings"
)

const jackpotProbability = 0.3

var fireProbability = calcFireProbability()

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, firePenguin())
}

func calcFireProbability() float64 {
	return math.Pow(jackpotProbability, 1.0/8.0)
}

func genEmoji() string {
	if rand.Float64() < fireProbability {
		return "🔥"
	}
	return "❄️"
}

func firePenguin() string {
	fireSnow := make([]string, 8)
	for i := 0; i < 8; i++ {
		fireSnow[i] = genEmoji()
	}

	return strings.Join(fireSnow[:3], "") + "\n" +
		fireSnow[3] + "🐧" + fireSnow[4] + "\n" +
		strings.Join(fireSnow[5:], "")
}
