package abbrev

import (
	"testing"
)

func TestMatchString(t *testing.T) {
	var a *Abbrev

	t.Logf("checking t/est")
	a = MustCompile("t/est")
	if !a.MatchString("test") {
		t.Errorf("error: 't/est' should match 'test'")
	}
	if a.MatchString("tst") {
		t.Errorf("error: 't/est' shouldn't match 'tst'")
	}

	t.Logf("checking s/aveas/ar/clm")
	a = MustCompile("s/aveas/ar/clm")
	if !a.MatchString("saveasarclm") {
		t.Errorf("error: 's/aveas/ar/clm' should match 'saveasarclm'")
	}
	if !a.MatchString("sar") {
		t.Errorf("error: 's/aveas/ar/clm' should match 'sar'")
	}

	t.Logf("checking s/ave/ave/nue")
	a = MustCompile("s/ave/ave/nue")
	if !a.MatchString("save") {
		t.Errorf("error: 's/ave/ave/nue' should match 'save'")
	}
}
