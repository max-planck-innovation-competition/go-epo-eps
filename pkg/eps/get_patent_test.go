package eps

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const testID = "EP2921808NWB1"

func TestGetPatentHTML(t *testing.T) {
	ass := assert.New(t)
	res, err := GetPatentHTML(testID)
	ass.NoError(err)
	ass.NotNil(res)
	ass.NotEmpty(res)
	err = SaveFile(res, "./test-data/", string(testID)+".html")
	ass.NoError(err)
}

func TestGetPatentZIP(t *testing.T) {
	ass := assert.New(t)
	res, err := GetPatentZIP(testID)
	ass.NoError(err)
	ass.NotNil(res)
	ass.NotEmpty(res)
	err = SaveFile(res, "./test-data/", string(testID)+".zip")
	ass.NoError(err)
}

func TestGetPatentXML(t *testing.T) {
	ass := assert.New(t)
	res, err := GetPatentXML(testID)
	ass.NoError(err)
	ass.NotNil(res)
	ass.NotEmpty(res)
	err = SaveFile(res, "./test-data/", string(testID)+".xml")
	ass.NoError(err)
}

func TestGetPatentPDF(t *testing.T) {
	ass := assert.New(t)
	res, err := GetPatentPDF(testID)
	ass.NoError(err)
	ass.NotNil(res)
	ass.NotEmpty(res)
	err = SaveFile(res, "./test-data/", string(testID)+".pdf")
	ass.NoError(err)
}
