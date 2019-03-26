package prosopagnosia

import (
	"testing"
)

func TestUpload(t *testing.T) {
	err := uploadPhoto("faces.hilden", "c:\\temp\\lucas.png")
	if err != nil {
		t.Error("Testcase failed")
	}
}
