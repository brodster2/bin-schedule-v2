package main_test

import (
	bs "bin_schedule_v2"
	"testing"
)

func TestCSVDataSource_LoadData(t *testing.T) {
	ds := new(bs.CSVDataSource)
	got, err := ds.Load("test_schedule.csv")
	if err != nil {
		t.Fatalf("Failure loading CSV: %v", err)
	}
	want := []bs.Collection{
		{Date: "17/11/21", Bins: []string{"Orange"}},
		{Date: "24/11/21", Bins: []string{"Grey", "Blue"}},
		{Date: "01/12/21", Bins: []string{"Orange"}},
		{Date: "08/12/21", Bins: []string{"Green"}},
	}
	for i := 0; i < len(want); i++ {
		if want[i].Date != got[i].Date {
			t.Errorf("Wanted: %s, Got: %s", want[i].Date, got[i].Date)
		}
		if len(want[i].Bins) != len(got[i].Bins) {
			t.Errorf("Wanted: %v, Got: %v", want[i].Bins, got[i].Bins)
		}
		for g := 0; g < len(want[i].Bins); g++ {
			if want[i].Bins[g] != got[i].Bins[g] {
				t.Fatalf("Wanted: %s, Got: %s", want[i].Bins[g], got[i].Bins[g])
			}
		}
	}
}
