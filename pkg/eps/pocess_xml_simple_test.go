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
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20210630", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("B1", patDoc.Kind)
	ass.Equal("n", patDoc.Status)
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
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20210707", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("B1", patDoc.Kind)
	ass.Equal("n", patDoc.Status)
	ass.Equal(Country("EP"), patDoc.Country)
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
	ass.Equal(Country("US"), patDoc.Citations[0].Country)
	ass.Equal("20120281566", patDoc.Citations[0].DocNumber)
	ass.Equal("A", patDoc.Citations[0].Kind)
	ass.Equal(Country("GB"), patDoc.Citations[1].Country)
	ass.Equal("1307186", patDoc.Citations[1].DocNumber)
	ass.Equal("A", patDoc.Citations[1].Kind)
	// inventors
	ass.NotEmpty(patDoc.Inventors)
	ass.Equal("MARTIN, Brian Alexander", patDoc.Inventors[0].Name)
	ass.Equal("c/o Sony Europe IP Europe Jays Close, Viables", patDoc.Inventors[0].Street)
	ass.Equal("Basingstoke, Hampshire RG22 4SB", patDoc.Inventors[0].City)
	ass.Equal(Country("GB"), patDoc.Inventors[0].Country)
	// owners
	ass.NotEmpty(patDoc.Owners)
	ass.Equal("SONY Corporation", patDoc.Owners[0].Name)
	ass.Equal("101825199", patDoc.Owners[0].IID)
	ass.Equal("P103307EP1AA", patDoc.Owners[0].IRF)
	ass.Equal("1-7-1 Konan, Minato-ku", patDoc.Owners[0].Street)
	ass.Equal("Tokyo 108-0075", patDoc.Owners[0].City)
	ass.Equal(Country("JP"), patDoc.Owners[0].Country)
	// representative
	ass.NotEmpty(patDoc.Representatives)
	ass.Equal("D Young & Co LLP", patDoc.Representatives[0].Name)
	ass.Equal("101533551", patDoc.Representatives[0].IID)
	ass.Equal("120 Holborn", patDoc.Representatives[0].Street)
	ass.Equal("London EC1N 2DY", patDoc.Representatives[0].City)
	ass.Equal(Country("GB"), patDoc.Representatives[0].Country)
	// ContractingStates
	ass.NotEmpty(patDoc.ContractingStates)
	ass.Equal(Country("AL"), patDoc.ContractingStates[0])
	ass.Equal(Country("AT"), patDoc.ContractingStates[1])
	// Classifications
	ass.NotEmpty(patDoc.Classifications)
	ass.Equal("H04W 76/28 20180101AFI20201221BHEP ", patDoc.Classifications[0].Text)
	ass.Equal("H", patDoc.Classifications[0].Section)
	ass.Equal("04", patDoc.Classifications[0].Class)
	ass.Equal("W", patDoc.Classifications[0].SubClass)
	ass.Equal("76", patDoc.Classifications[0].MainGroup)
	ass.Equal("28", patDoc.Classifications[0].SubGroup)
	ass.Equal("H04W 52/02 20090101ALI20201221BHEP ", patDoc.Classifications[1].Text)
	ass.Equal("H", patDoc.Classifications[1].Section)
	ass.Equal("04", patDoc.Classifications[1].Class)
	ass.Equal("W", patDoc.Classifications[1].SubClass)
	ass.Equal("52", patDoc.Classifications[1].MainGroup)
	ass.Equal("02", patDoc.Classifications[1].SubGroup)
}
