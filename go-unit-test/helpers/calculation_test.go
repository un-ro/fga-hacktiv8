package helpers

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFailSum(t *testing.T) {
	result := Sum(10, 10)

	require.Equal(t, 40, result, "Sum should be 40")

	fmt.Println("TestFailSum done")
}

func TestSum(t *testing.T) {
	result := Sum(10, 10)

	assert.Equal(t, 20, result, "Sum should be 20")

	fmt.Println("TestSum done")
}

func TestTableSum(t *testing.T) {
	tests := []struct {
		request  int
		expected int
		errMsg   string
	}{
		{Sum(10, 10), 20, "Sum should be 20"},
		{Sum(20, 20), 40, "Sum should be 40"},
		{Sum(25, 25), 50, "Sum should be 50"},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			require.Equal(t, test.expected, test.request, test.errMsg)
		})
	}
}
