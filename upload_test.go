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

func TestUploadCollectionFacesAddison1(t *testing.T) {
	err := uploadPhoto("faces.hilden", ".\\test\\facestest\\addison1.jpg", "addison1.jpg")
	if err != nil {
		t.Error("Testcase failed")
	}
}

func TestUploadCollectionFacesAddison2(t *testing.T) {
	err := uploadPhoto("faces.hilden", ".\\test\\facestest\\addison2.jpg", "addison2.jpg")
	if err != nil {
		t.Error("Testcase failed")
	}
}

func TestUploadCollectionFacesLucas(t *testing.T) {
	err := uploadPhoto("faces.hilden", ".\\test\\facestest\\lucas.jpg", "lucas.jpg")
	if err != nil {
		t.Error("Testcase failed")
	}
}

func TestUploadCollectionFacesTammy1(t *testing.T) {
	err := uploadPhoto("faces.hilden", ".\\test\\facestest\\tammy1.jpg", "tammy1.jpg")
	if err != nil {
		t.Error("Testcase failed")
	}
}
