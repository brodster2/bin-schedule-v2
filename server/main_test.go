package main_test

import (
	bs "bin_schedule_v2"
	"reflect"
	"testing"
)

func TestLoadFromCSV(t *testing.T) {
	got, err := bs.LoadFromCSV("test_schedule.csv")
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
