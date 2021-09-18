package fiboService

import (
	"math/big"
	"testing"
)

func TestFiboNumber(t *testing.T) {
	var n uint = 10
	var out big.Int = *big.NewInt(55)

	result := FibonacciNumber(n).Int64()

	if result != out.Int64() {
		t.Error("Incorrect result", result, "not equal", out)
	}

}

func TestFiboSlice(t *testing.T) {
	var x, y uint
	x, y = 0, 3
	var arr = FiboSlice(x, y)
	var arrTest = []big.Int{*big.NewInt(0), *big.NewInt(1), *big.NewInt(1), *big.NewInt(2)}

	for i := 0; i < len(arr); i++ {
		if arr[i].Int64() != arrTest[i].Int64() {
			t.Error("Incorrect result", arr[i], "not equal", arrTest[i])
		}
	}
}

// In progress...
// func TestGetDiscription(t *testing.T) {

// 	var want []byte
// 	str := "str"

// 	want = []byte(str)

// 	handler := http.HandlerFunc(GetDiscription)

// 	rec := httptest.NewRecorder()

// 	req := httptest.NewRequest("GET", "/", nil)

// 	handler.ServeHTTP(rec, req)
// 	re := rec.Body.Bytes()

// 	for i := 0; i < len(want); i++ {
// 		if want[i] != re[i] {
// 			fmt.Print(want)
// 			t.Error("Incorrect result", want[i], "not equal", re[i])

// 		}
// 	}
// }
