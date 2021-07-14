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
	ass.NotEmpty(patDoc.Title)
	ass.NotEmpty(patDoc.Claims)
	ass.NotEmpty(patDoc.Citations)
}

func TestProcessXMLSimple2(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("../../dtds/1-5-1.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.NotEmpty(patDoc.Title)
	ass.Equal("de", patDoc.Title[0].Language)
	ass.Equal("KOMMUNIKATIONSVORRICHTUNG UND -VERFAHREN", patDoc.Title[0].Text)
	ass.Equal("en", patDoc.Title[1].Language)
	ass.Equal("COMMUNICATIONS DEVICE AND METHOD", patDoc.Title[1].Text)
	ass.Equal("fr", patDoc.Title[2].Language)
	ass.Equal("DISPOSITIF ET PROCÉDÉ DE COMMUNICATIONS", patDoc.Title[2].Text)
	// claims
	ass.NotEmpty(patDoc.Claims)
	ass.NotEmpty(patDoc.Citations)
	// citations
	ass.Equal("US", patDoc.Citations[0].Country)
	ass.Equal("20120281566", patDoc.Citations[0].DocNumber)
	ass.Equal("A", patDoc.Citations[0].Kind)
	ass.Equal("GB", patDoc.Citations[1].Country)
	ass.Equal("1307186", patDoc.Citations[1].DocNumber)
	ass.Equal("A", patDoc.Citations[1].Kind)
	// inventors
	ass.NotEmpty(patDoc.Inventors)
	ass.Equal("MARTIN, Brian Alexander", patDoc.Inventors[0].Name)
	ass.Equal("c/o Sony Europe IP Europe Jays Close, Viables", patDoc.Inventors[0].Street)
	ass.Equal("Basingstoke, Hampshire RG22 4SB", patDoc.Inventors[0].City)
	ass.Equal("GB", patDoc.Inventors[0].Country)
	// owners
	ass.NotEmpty(patDoc.Owners)
	ass.Equal("SONY Corporation", patDoc.Owners[0].Name)
	ass.Equal("101825199", patDoc.Owners[0].IID)
	ass.Equal("P103307EP1AA", patDoc.Owners[0].IRF)
	ass.Equal("1-7-1 Konan, Minato-ku", patDoc.Owners[0].Street)
	ass.Equal("Tokyo 108-0075", patDoc.Owners[0].City)
	ass.Equal("JP", patDoc.Owners[0].Country)
}
