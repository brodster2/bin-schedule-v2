package helpers

import (
	"reflect"
	"testing"
)

func TestLoadFromCSV(t *testing.T) {
	csv := `Date, Bins
17-11-21,Orange
24-11-21,Grey:Blue
01-12-21,Orange
08-12-21, Green`

	got, err := LoadFromCSV(csv)
	if err != nil {
		t.Fatalf("Failure loading CSV: %v", err)
	}
	want := map[string][]string{
		"17-11-21": {"Orange"},
		"24-11-21": {"Grey", "Blue"},
		"01-12-21": {"Orange"},
		"08-12-21": {"Green"},
	}
	if equal := reflect.DeepEqual(want, got); !equal {
		t.Fatalf("Want: %v, Got: %v", want, got)
	}
}
