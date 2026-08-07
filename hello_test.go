package hello_test

import (
	"testing"

	"github.com/asd794/hello"
)

func TestSayHello(t *testing.T) {
	// tests := []struct {
	// 	name string // description of this test case
	// 	// Named input parameters for target function.
	// 	name string
	// 	want string
	// }{
	// 	// TODO: Add test cases.
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		got := hello.SayHello(tt.name)
	// 		// TODO: update the condition below to compare got with tt.want.
	// 		if true {
	// 			t.Errorf("SayHello() = %v, want %v", got, tt.want)
	// 		}
	// 	})
	// }
	got := hello.SayHello("Go")
	want := "Hello, Go!"
	t.Log(got)
	if got != want {
		t.Errorf("SayHello() = %v, want %v", got, want)
	}
}
