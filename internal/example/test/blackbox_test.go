package blackbox

import (
	"errors"
	"fmt"
	"math/rand"
	"testing"
	"time"
	"unsafe"

	"github.com/jonny-burkholder/go-test-tutorial/internal/example"
)

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// in a standard, non table-driven test, each test case must
// be written out prodecurally
func TestExportedFunction(t *testing.T) {
	var tr, fa bool = true, false
	if _, err := example.ExportedFunction(nil); !errors.Is(err, example.ErrUserError) {
		t.Errorf("expected %v, got %v", example.ErrUserError, err)
	}
	if _, err := example.ExportedFunction(&fa); !errors.Is(err, example.ErrThisFailed) {
		t.Errorf("expected %v, got %v", example.ErrThisFailed, err)
	}
	if _, err := example.ExportedFunction(&tr); err != nil {
		t.Errorf("expected %v, got %v", example.Message, err)
	}
}

// in a table-driven test, testcases are structs that provide
// input and expected output, greatly reducing the amount
// of code that needs to be writte for each test case
func TestTableTest(t *testing.T) {

	// define the test case struct
	type testCase struct {
		input    *bool
		expected error
	}

	var tr, fa bool = true, false

	// create test cases
	testCases := []testCase{
		{&tr, nil},
		{&fa, example.ErrThisFailed},
		{nil, example.ErrUserError},
	}

	// iterate through each test case and run its test
	for testNumber, testCase := range testCases {
		name := fmt.Sprintf("test %d", testNumber)
		// t.Run allows us to output a name for each test case,
		// as well as define the test function that will run for
		// each case. This is similar to if we wrote a discrete
		// Test function for each test case
		t.Run(name, func(t *testing.T) {
			if _, err := example.ExportedFunction(testCase.input); err != testCase.expected {
				t.Errorf("expected %v, got %v", testCase.expected, err)
			}
		})
	}

}

// let's test our IsFoo() function. Since we don't want to
// manually generate each test case, we can instead use fuzz
// testing to auto-generate the input
func FuzzIsFoo(f *testing.F) {
	// first we seed our fuzzer. This tells the fuzzer the
	// type of input that it is expected to generate
	f.Add("foo")
	f.Fuzz(func(t *testing.T, word string) {
		// if the word is not foo, but IsFoo() says it is, we have a problem
		if word != example.Foo && example.IsFoo(word) {
			t.Errorf("this function thinks %s is equal to %s", word, example.Foo)
		}
	})
}

func BenchmarkIsFoo(b *testing.B) {
	data := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = randomString(rand.Intn(100) + 1)
	}
	b.Run("IsFoo", func(b *testing.B) {
		for _, value := range data {
			example.IsFoo(value)
		}
	})
	b.Run("IsFoo2", func(b *testing.B) {
		for _, value := range data {
			example.IsFoo2(value)
		}
	})
	b.Run("IsFoo3", func(b *testing.B) {
		for _, value := range data {
			example.IsFoo3(value)
		}
	})
}

func ExampleIsFoo() {
	valid := example.IsFoo("foo")
	invalid := example.IsFoo("bar")
	fmt.Println(valid, invalid)
	// Output: true false
}

// https://stackoverflow.com/a/31832326
func randomString(n int) string {
	src := rand.NewSource(time.Now().UnixNano())

	result := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			result[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&result))
}
