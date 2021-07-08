package eps

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestProcessXMLSimple(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("../../dtds/1-5-example.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)
	ass.Equal("de", patDoc.Title[0].Language)
	ass.Equal("BREMSÜBERWACHUNG", patDoc.Title[0].Text)
	ass.Equal("en", patDoc.Title[1].Language)
	ass.Equal("BRAKE MONITORING", patDoc.Title[1].Text)
	ass.Equal("fr", patDoc.Title[2].Language)
	ass.Equal("SURVEILLANCE DE FREIN", patDoc.Title[2].Text)
}

func TestProcessXMLSimple2(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("../../dtds/1-5-1.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)
	ass.Equal("de", patDoc.Title[0].Language)
	ass.Equal("KOMMUNIKATIONSVORRICHTUNG UND -VERFAHREN", patDoc.Title[0].Text)
	ass.Equal("en", patDoc.Title[1].Language)
	ass.Equal("COMMUNICATIONS DEVICE AND METHOD", patDoc.Title[1].Text)
	ass.Equal("fr", patDoc.Title[2].Language)
	ass.Equal("DISPOSITIF ET PROCÉDÉ DE COMMUNICATIONS", patDoc.Title[2].Text)
}
