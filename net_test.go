package ghw

import (
	"testing"
)

func TestNet(t *testing.T) {
	info, err := Network()

	if err != nil {
		t.Fatalf("Expected nil err, but got %v", err)
	}
	if info == nil {
		t.Fatalf("Expected non-nil NetworkInfo, but got nil")
	}

	if len(info.NICs) > 0 {
		for _, n := range info.NICs {
			if n.Name == "" {
				t.Fatalf("Expected a NIC name but got \"\".")
			}
		}
	}
}
