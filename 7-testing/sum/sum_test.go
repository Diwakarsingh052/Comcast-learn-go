package sum

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// go test ./... // run all the tests for the project

// function names must start with word Test to signal it is a test
// file name must have _test.go to register a file as a test
func TestSumInt(t *testing.T) {
	// three thing we need to identify while testing
	// how to call the function
	// parameters
	// return values
	input := []int{1, 2, 3, 4, 5}
	want := 15
	got := SumInt(input)
	if got != want {
		// test would continue on if test case fail
		t.Errorf("sum of 1 to 5 should be %v; got %v", want, got)
		//Fatalf stop the test if it fails at this point.
		//t.Fatalf()
	}

	want = 0
	got = SumInt(nil)
	if got != want {
		// test would continue on if test case fail
		t.Errorf("sum of nil should be %v; got %v", want, got)
		//Fatalf stop the test if it fails at this point.
		//t.Fatalf()
	}

}

func TestSumIntV2(t *testing.T) {
	// table driven testing
	//columns = [name, input, want]
	// 			[	rows = values		]

	type args struct {
		input  []int
		input2 int
		input3 int
	}
	tt := []struct {
		// columns
		name  string
		input []int
		//args // embed args struct if you have multiple args to be passed to a function
		want int
	}{
		// test cases
		// using subtest we get isolated environment for the test cases
		{
			name:  "one to five numbers",
			input: []int{1, 2, 3, 4, 5},
			want:  15,
		},
		{
			name:  "nil slice",
			input: nil,
			want:  0,
		},
	}

	// ranging over individual test cases
	for _, tc := range tt {
		//t.Run starts a sub test
		t.Run(tc.name, func(t *testing.T) {
			got := SumInt(tc.input)
			if got != tc.want {
				// fatalf would not affect other subtests from running
				t.Fatalf("sum of %v  should be %v; got %v", tc.input, tc.want, got)
			}
		})
	}

}

func TestSumIntV3(t *testing.T) {

	//type args struct {
	//	vs []int
	//}
	// each index in slice would represent one test case
	tt := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "one to five numbers",
			input: []int{1, 2, 3, 4, 5},
			want:  15,
		},
		{
			name:  "nil slice",
			input: nil,
			want:  0,
		},
	}

	for _, tc := range tt {

		// t.Run creates a subtest, you can run each index of slice as a subtest
		t.Run(tc.name, func(t *testing.T) {
			got := SumInt(tc.input)
			require.Equal(t, tc.want, got)
		})

	}

}
