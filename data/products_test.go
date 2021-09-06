package data 

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name : "aditya",
		Price: 20,
		SKU:"abs-abc-acb",
	}

	err := p.Validate()

	if err != nil{
		t.Fatal(err)
	}

}