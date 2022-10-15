package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// method 1
// func TestSum(t *testing.T) {
// 	result := Sum(10, 10)

// 	if result != 20 {
// 		panic("Result should be 20")
// 	}
// }

// method 2
// func TestSum(t *testing.T) {
// 	result := Sum(10, 10)

// 	if result != 20 {
// 		t.FailNow()
// 	}
// 	fmt.Println("Result should be 20")
// }

// method fail 1
// func TestFailSum(t *testing.T) {
// 	result := Sum(10, 10)

// 	if result != 40 {
// 		t.Fail()
// 	}

// 	fmt.Println("TestFailSum Eksekusi Terhenti")
// }

// method 3
func TestSum(t *testing.T) {
	result := Sum(10, 10)

	assert.Equal(t, 20, result, "Result has to be 20")

	fmt.Println("TestSum Eksekusi Terhenti")
}

// method fail 2
func TestFailSum(t *testing.T) {
	result := Sum(10, 10)

	assert.Equal(t, 20, result, "Result has to be 20")

	fmt.Println("TestFailSum Eksekusi Terhenti")
}
