package prosopagnosia

import (
	"testing"
)

func TestUpload(t *testing.T) {
	err := uploadPhoto("faces.hilden", "c:\\temp\\addison.png", "addison.png")
	if err != nil {
		t.Error("Testcase failed")
	}
}

func TestUploadCollectionFacesTammy(t *testing.T) {
	err := uploadPhoto("faces.collection", ".\\test\\collection\\tammy.jpg", "tammy.jpg")
	if err != nil {
		t.Error("Testcase failed")
	}
}

func TestUploadCollectionFacesAddison(t *testing.T) {
	err := uploadPhoto("faces.collection", ".\\test\\collection\\addison.jpg", "addison.jpg")
	if err != nil {
		t.Error("Testcase failed")
	}
}
