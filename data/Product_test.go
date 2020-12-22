package data

import "testing"

func TestCheckProductValidation(t *testing.T) {
	p := &Product{ Name:"sdf", Price: 2, SKU: "abc-ddbc-acb"}
	got := p.Validate()

	if(got != nil){
		t.Fatal(got)
	}
}