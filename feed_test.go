package googlenews

import (
	"os"
	"testing"
)

func TestFeed(t *testing.T) {
	b, err := os.ReadFile("testdata/worldnews.xml")
	if err != nil {
		t.Fail()
	}
	f, err := createFeed(string(b))
	if err != nil {
		t.Fail()
	}

	if len(f.Items) == 0 {
		t.Fail()
	}
}
