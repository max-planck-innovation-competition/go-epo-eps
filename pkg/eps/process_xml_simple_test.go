package eps

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestProcessXMLSimple15A1(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("test-data/application/v1-5-A1.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("EP18877305A1", patDoc.ID)
	ass.Equal("EP18877305NWA1.xml", patDoc.File)
	ass.Equal("en", patDoc.Lang)
	ass.Equal(Country("EP"), patDoc.Country)
	ass.Equal("3782854", patDoc.DocNumber)
	ass.Equal("A1", patDoc.Kind)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20210224", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("n", patDoc.Status)
	ass.Equal("ep-patent-document-v1-5", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("FAHRZEUGMONTIERTE ANZEIGEANORDNUNG UND FAHRZEUG", patDoc.Title[0].Text)
	ass.Equal("de", patDoc.Title[0].Language)
	ass.Equal("VEHICLE-MOUNTED DISPLAY ASSEMBLY AND VEHICLE", patDoc.Title[1].Text)
	ass.Equal("en", patDoc.Title[1].Language)
	ass.Equal("ENSEMBLE D'AFFICHAGE MONTÉ SUR VÉHICULE ET VÉHICULE", patDoc.Title[2].Text)
	ass.Equal("fr", patDoc.Title[2].Language)

	// abstract
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(447, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// claims
	ass.NotEmpty(patDoc.Claims)
	ass.Equal(1, len(patDoc.Claims))
	ass.Equal(2562, len(patDoc.Claims[0].Text))
	ass.Equal("en", patDoc.Claims[0].Language)
	ass.Equal("claims01", patDoc.Claims[0].Id)

	// description
	ass.NotEmpty(patDoc.Description)
	ass.Equal(21342, len(patDoc.Description[0].Text))
	ass.Equal("en", patDoc.Description[0].Language)

	// citations
	ass.NotEmpty(patDoc.Citations)
	ass.Equal(Country("CN"), patDoc.Citations[0].Country)
	ass.Equal("201810358967", patDoc.Citations[0].DocNumber)
	ass.Equal("", patDoc.Citations[0].Kind)

	// Inventors
	ass.NotEmpty(patDoc.Inventors)
	ass.Equal(Country("CN"), patDoc.Inventors[0].Country)
	ass.Equal("Beijing 100176", patDoc.Inventors[0].City)
	ass.Equal("No. 9 Dize Rd.\nBDA", patDoc.Inventors[0].Street)
	ass.Equal("FENG, Yuhsiung", patDoc.Inventors[0].Name)

	ass.Equal(Country("CN"), patDoc.Inventors[1].Country)
	ass.Equal("Beijing 100176", patDoc.Inventors[1].City)
	ass.Equal("No. 9 Dize Rd.\nBDA", patDoc.Inventors[1].Street)
	ass.Equal("GAO, Wenbao", patDoc.Inventors[1].Name)

	ass.Equal(Country("CN"), patDoc.Inventors[2].Country)
	ass.Equal("Beijing 100176", patDoc.Inventors[2].City)
	ass.Equal("No. 9 Dize Rd.\nBDA", patDoc.Inventors[2].Street)
	ass.Equal("ZENG, Jia", patDoc.Inventors[2].Name)

	// representatives
	ass.NotEmpty(patDoc.Representatives)
	ass.Equal(Country("DE"), patDoc.Representatives[0].Country)
	ass.Equal("101201236", patDoc.Representatives[0].IID)
	ass.Equal("40009 Düsseldorf", patDoc.Representatives[0].City)
	ass.Equal("Patent- & Rechtsanwälte \nPartnerschaftsgesellschaft mbB \nPostfach 10 18 30", patDoc.Representatives[0].Street)
	ass.Equal("Cohausz & Florack", patDoc.Representatives[0].Name)

	// contracting states
	ass.NotEmpty(patDoc.ContractingStates)
	ass.Equal(38, len(patDoc.ContractingStates))
	for i := 0; i <= 37; i++ {
		ass.Equal(2, len(patDoc.ContractingStates[i]))
	}

	// classifications
	ass.Equal("B60R   7/06        20060101AFI20191025BHEP        ", patDoc.Classifications[0].Text)
	ass.Equal(IPC, patDoc.Classifications[0].System)
	ass.Equal(1, patDoc.Classifications[0].Sequence)
	ass.Equal("B", patDoc.Classifications[0].Section)
	ass.Equal("60", patDoc.Classifications[0].Class)
	ass.Equal("R", patDoc.Classifications[0].SubClass)
	ass.Equal("7", patDoc.Classifications[0].MainGroup)
	ass.Equal("06", patDoc.Classifications[0].SubGroup)
	ass.Equal("20060101", patDoc.Classifications[0].Version)
	ass.Equal("A", patDoc.Classifications[0].ClassificationLevel)
	ass.Equal("F", patDoc.Classifications[0].FirstLater)
	ass.Equal("I", patDoc.Classifications[0].ClassificationValue)
	ass.Equal("20191025", patDoc.Classifications[0].ActionDate)
	ass.Equal("B", patDoc.Classifications[0].OriginalOrReclassified)
	ass.Equal("H", patDoc.Classifications[0].Source)
	ass.Equal("EP", patDoc.Classifications[0].GeneratingOffice)

	ass.Equal("B60N   3/12        20060101ALI20191025BHEP        ", patDoc.Classifications[1].Text)
	ass.Equal(IPC, patDoc.Classifications[1].System)
	ass.Equal(2, patDoc.Classifications[1].Sequence)
	ass.Equal("B", patDoc.Classifications[1].Section)
	ass.Equal("60", patDoc.Classifications[1].Class)
	ass.Equal("N", patDoc.Classifications[1].SubClass)
	ass.Equal("3", patDoc.Classifications[1].MainGroup)
	ass.Equal("12", patDoc.Classifications[1].SubGroup)
	ass.Equal("20060101", patDoc.Classifications[1].Version)
	ass.Equal("A", patDoc.Classifications[1].ClassificationLevel)
	ass.Equal("L", patDoc.Classifications[1].FirstLater)
	ass.Equal("I", patDoc.Classifications[1].ClassificationValue)
	ass.Equal("20191025", patDoc.Classifications[1].ActionDate)
	ass.Equal("B", patDoc.Classifications[1].OriginalOrReclassified)
	ass.Equal("H", patDoc.Classifications[1].Source)
	ass.Equal("EP", patDoc.Classifications[1].GeneratingOffice)

	fmt.Println(patDoc.Title[0])
	fmt.Println(patDoc.Claims[0])
	fmt.Println(patDoc.Description[0])
	fmt.Println(patDoc.Citations[0])
}

func TestProcessXMLSimple15B1(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("test-data/grant/v1-5-B1.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("EP17171508B1", patDoc.ID)
	ass.Equal("EP17171508NWB1.xml", patDoc.File)
	ass.Equal("en", patDoc.Lang)
	ass.Equal(Country("EP"), patDoc.Country)
	ass.Equal("3404678", patDoc.DocNumber)
	ass.Equal("B1", patDoc.Kind)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20210630", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("n", patDoc.Status)
	ass.Equal("ep-patent-document-v1-5", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("HOCHSPANNUNGSANORDNUNG UND VERFAHREN ZUM BETREIBEN DER HOCHSPANNUNGSANORDNUNG", patDoc.Title[0].Text)
	ass.Equal("de", patDoc.Title[0].Language)
	ass.Equal("HIGH VOLTAGE ASSEMBLY AND METHOD TO OPERATE THE HIGH VOLTAGE ASSEMBLY", patDoc.Title[1].Text)
	ass.Equal("en", patDoc.Title[1].Language)
	ass.Equal("ENSEMBLE À HAUTE TENSION ET PROCÉDÉ DE FONCTIONNEMENT DE L'ENSEMBLE À HAUTE TENSION", patDoc.Title[2].Text)
	ass.Equal("fr", patDoc.Title[2].Language)

	// abstract !! Empty
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(0, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// claims
	ass.NotEmpty(patDoc.Claims)
	ass.Equal(3, len(patDoc.Claims))
	ass.Equal(5061, len(patDoc.Claims[0].Text))
	ass.Equal("en", patDoc.Claims[0].Language)
	ass.Equal("claims01", patDoc.Claims[0].Id)
	ass.Equal(5301, len(patDoc.Claims[1].Text))
	ass.Equal("de", patDoc.Claims[1].Language)
	ass.Equal("claims02", patDoc.Claims[1].Id)
	ass.Equal(5736, len(patDoc.Claims[2].Text))
	ass.Equal("fr", patDoc.Claims[2].Language)
	ass.Equal("claims03", patDoc.Claims[2].Id)

	// description
	ass.NotEmpty(patDoc.Description)
	ass.Equal(13662, len(patDoc.Description[0].Text))
	ass.Equal("en", patDoc.Description[0].Language)

	// citations
	ass.NotEmpty(patDoc.Citations)
	ass.Equal(Country("EP"), patDoc.Citations[0].Country)
	ass.Equal("3109871", patDoc.Citations[0].DocNumber)
	ass.Equal("A1", patDoc.Citations[0].Kind)
	ass.Equal(Country("US"), patDoc.Citations[1].Country)
	ass.Equal("2008197955", patDoc.Citations[1].DocNumber)
	ass.Equal("A1", patDoc.Citations[1].Kind)
	ass.Equal(Country("GB"), patDoc.Citations[2].Country)
	ass.Equal("693448", patDoc.Citations[2].DocNumber)
	ass.Equal("A", patDoc.Citations[2].Kind)

	// Inventors
	ass.NotEmpty(patDoc.Inventors)
	ass.Equal(Country("DE"), patDoc.Inventors[0].Country)
	ass.Equal("50374 Erftstadt", patDoc.Inventors[0].City)
	ass.Equal("Elly-Heuss-Knapp-Weg 20", patDoc.Inventors[0].Street)
	ass.Equal("HARTHUM, Jörg", patDoc.Inventors[0].Name)

	ass.Equal(Country("DE"), patDoc.Inventors[1].Country)
	ass.Equal("47918 Tönisvorst", patDoc.Inventors[1].City)
	ass.Equal("Pappelallee 52", patDoc.Inventors[1].Street)
	ass.Equal("STIRL, Tobias", patDoc.Inventors[1].Name)

	// representatives
	ass.NotEmpty(patDoc.Representatives)
	ass.Equal(Country("CH"), patDoc.Representatives[0].Country)
	ass.Equal("101820002", patDoc.Representatives[0].IID)
	ass.Equal("5400 Baden", patDoc.Representatives[0].City)
	ass.Equal("Hahnrainweg 4", patDoc.Representatives[0].Street)
	ass.Equal("Freigutpartners IP Law Firm", patDoc.Representatives[0].Name)

	// contracting states
	ass.NotEmpty(patDoc.ContractingStates)
	ass.Equal(38, len(patDoc.ContractingStates))
	for i := 0; i <= 37; i++ {
		ass.Equal(2, len(patDoc.ContractingStates[i]))
	}

	// classifications
	ass.Equal("H01F  27/14        20060101AFI20171122BHEP        ", patDoc.Classifications[0].Text)
	ass.Equal(IPC, patDoc.Classifications[0].System)
	ass.Equal(1, patDoc.Classifications[0].Sequence)
	ass.Equal("H", patDoc.Classifications[0].Section)
	ass.Equal("01", patDoc.Classifications[0].Class)
	ass.Equal("F", patDoc.Classifications[0].SubClass)
	ass.Equal("27", patDoc.Classifications[0].MainGroup)
	ass.Equal("14", patDoc.Classifications[0].SubGroup)
	ass.Equal("20060101", patDoc.Classifications[0].Version)
	ass.Equal("A", patDoc.Classifications[0].ClassificationLevel)
	ass.Equal("F", patDoc.Classifications[0].FirstLater)
	ass.Equal("I", patDoc.Classifications[0].ClassificationValue)
	ass.Equal("20171122", patDoc.Classifications[0].ActionDate)
	ass.Equal("B", patDoc.Classifications[0].OriginalOrReclassified)
	ass.Equal("H", patDoc.Classifications[0].Source)
	ass.Equal("EP", patDoc.Classifications[0].GeneratingOffice)

	ass.Equal("H01F  27/40        20060101ALI20171122BHEP        ", patDoc.Classifications[1].Text)
	ass.Equal(IPC, patDoc.Classifications[1].System)
	ass.Equal(2, patDoc.Classifications[1].Sequence)
	ass.Equal("H", patDoc.Classifications[1].Section)
	ass.Equal("01", patDoc.Classifications[1].Class)
	ass.Equal("F", patDoc.Classifications[1].SubClass)
	ass.Equal("27", patDoc.Classifications[1].MainGroup)
	ass.Equal("40", patDoc.Classifications[1].SubGroup)
	ass.Equal("20060101", patDoc.Classifications[1].Version)
	ass.Equal("A", patDoc.Classifications[1].ClassificationLevel)
	ass.Equal("L", patDoc.Classifications[1].FirstLater)
	ass.Equal("I", patDoc.Classifications[1].ClassificationValue)
	ass.Equal("20171122", patDoc.Classifications[1].ActionDate)
	ass.Equal("B", patDoc.Classifications[1].OriginalOrReclassified)
	ass.Equal("H", patDoc.Classifications[1].Source)
	ass.Equal("EP", patDoc.Classifications[1].GeneratingOffice)

	fmt.Println(patDoc.Title[0])
	fmt.Println(patDoc.Claims[0])
	fmt.Println(patDoc.Description[0])
	fmt.Println(patDoc.Citations[0])
}

func TestProcessXMLSimple151B1(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("test-data/grant/v1-5-1-B1.xml")
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
