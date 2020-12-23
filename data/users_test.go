package data

import "testing"

func TestChecksValidation(t *testing.T) {
	u := &User{Name: "Esto"}
	err := u.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
