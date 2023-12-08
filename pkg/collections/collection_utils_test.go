package collections

import "testing"

func TestCopyMap(t *testing.T) {
	src := map[int]string{1: "value 1", 2: "value 2", 3: "value 3"}
	cpy := CopyMap(src)
	if len(src) == 0 {
		t.Fatal("invalid map size")
	}
	if len(cpy) != len(src) {
		t.Fatalf("expected: %v, actual: %v", len(src), len(cpy))
	}
	for k, v := range src {
		if cpy[k] != v {
			t.Fatalf("expected: %v, actual: %v", v, cpy[k])
		}
	}
}
