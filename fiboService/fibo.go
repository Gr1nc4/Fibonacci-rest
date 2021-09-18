package fiboService

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strconv"

	"github.com/Gr1nc4/Fibonacci-rest.git/models"
)

var (
	cache map[uint]*big.Int
)

func init() {
	cache = make(map[uint]*big.Int)
}

//Calculate fibonacci number
func FibonacciNumber(n uint) *big.Int {
	if n <= 1 {
		return big.NewInt(int64(n))
	}
	if r, ok := cache[n]; ok {
		return r
	}
	var n2, n1 = big.NewInt(0), big.NewInt(1)

	for i := uint(1); i < n; i++ {
		n2.Add(n2, n1)
		n1, n2 = n2, n1
	}
	cache[n] = n1
	return n1
}

//Сreate a slice of fibonacci numbers
func FiboSlice(x, y uint) []big.Int {
	var arr []big.Int
	for i := x; i <= y; i++ {
		arr = append(arr, *FibonacciNumber(uint(i)))
	}
	return arr
}

//Sending a range and receiving fibonacci sequence on rang from X to Y
func GetFiboSlice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	var fr models.FibRange

	err := json.NewDecoder(r.Body).Decode(&fr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fbX, err := strconv.Atoi(fr.X)
	if err != nil {
		log.Fatal(err)
	}
	fbY, err := strconv.Atoi(fr.Y)
	if err != nil {
		log.Fatal(err)
	}
	var sl = FiboSlice(uint(fbX), uint(fbY))
	json.NewEncoder(w).Encode(sl)

}

//Get appllication description
func GetDiscription(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `This application of fibonacci sequence. For get the sequence go to "/fibo" and set the range for the sequence in the format:
	{
		"x":"10"
		"y":"20"
	}
	
	Where "x" is the beginning of the range, and "y" is end of range.
	Then send a POST request to the above endpoint`)
}
