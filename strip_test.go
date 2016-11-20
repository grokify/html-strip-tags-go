package strip

import (
	"testing"
)

var stripTests = []struct {
	v    string
	want string
}{
	{"<h1>Hello World</h1>", "Hello World"}
}

func TestStripTags(t *testing.T) {
	for _, tt := range stripTests {
		got := StripTags(tt.v)
		if got != tt.want {
			t.Errorf("StripTags(%v): want %v, got %v", tt.v, tt.want, got)
		}
	}
}
