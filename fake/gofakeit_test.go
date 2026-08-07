package fake

import "testing"

func TestName(t *testing.T) {
	for i := 0; i < 10; i++ {
		got := Name()
		t.Logf("Test iteration %d, %s", i+1, got)
		if got == "" {
			t.Fatalf("Name() returned an empty string")
		}
	}
}
