package validate

import "testing"

func TestValidateNameEmptyError(t *testing.T) {
	err := ValidateName("")
	if err == nil {
		t.Fatal("there must be a error here!")
	}
	want := "empty name"
	if err.Error() != want {
		t.Errorf("we got: %s, but we wanted: %s", err, want)
	}
}

func TestValidateNameError(t *testing.T) {
	err := ValidateName("Nina")
	if err != nil {
		t.Errorf("error: %s. But error must be nil!", err)
	}
}
