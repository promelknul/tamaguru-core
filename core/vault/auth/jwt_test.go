package auth

import "testing"

func TestGenerate(t *testing.T) {
	if _, err := Generate("dev"); err != nil {
		t.Fatalf("cannot sign token: %v", err)
	}
}
