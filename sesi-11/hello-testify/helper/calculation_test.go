package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSum(t *testing.T) {
	result := Sum(10, 10)

	assert.Equal(t, 20, result, "Result has to be 20")

	fmt.Println("TestSum Eksekusi Terhenti")
}

func TestFailSum(t *testing.T) {
	result := Sum(10, 10)

	assert.Equal(t, 20, result, "Result has to be 20")

	fmt.Println("TestFailSum Eksekusi Terhenti")
}

func TestTableSum(t *testing.T) {
	testCases := []struct {
		request  int
		expected int
		errMsg   string
	}{
		{
			request:  Sum(10, 10),
			expected: 20,
			errMsg:   "Result has to be 20",
		},
		{
			request:  Sum(20, 20),
			expected: 41,
			errMsg:   "Result has to be 40",
		},
		{
			request:  Sum(25, 25),
			expected: 50,
			errMsg:   "Result has to be 50",
		},
	}
	for i, tC := range testCases {
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			require.Equal(t, tC.expected, tC.request, tC.errMsg)
		})
	}
}
