package fakr_test

import (
	"github.com/dtomasi/fakr"
	"testing"
)

func TestNew(t *testing.T) {
	fk := fakr.New()

	if s := fk.GetSink(); s == nil {
		t.Errorf("sink is nil")
	}
}

// nolint:godox,nolintlint
// TODO add more tests
