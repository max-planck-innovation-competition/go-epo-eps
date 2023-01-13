package eps

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

/*
tests:   A1 A2 B1 B2
v 1.0    2  2  2  2
v 1.01   1  1  X  X
v 1.1    1  2  2  2
v 1.2    2  2  2  2
v 1.3	 2  2  3  2
v 1.4    2  1  2  2
v 1.5    2  2  2  2
v 1.5.1  4  3  2  2
*/

//v 1.0
func TestProcessXMLSimple10A1(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("test-data/application/v1-0-A1.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("EP01963450A1", patDoc.ID)
	ass.Equal("01963450.xml", patDoc.File)
	ass.Equal("en", patDoc.Lang)
	ass.Equal(Country("EP"), patDoc.Country)
	ass.Equal("1325900", patDoc.DocNumber)
	ass.Equal("A1", patDoc.Kind)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20030709", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("n", patDoc.Status)
	ass.Equal("ep-patent-document-v1-0", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("VERFAHREN ZUR HERSTELLUNG VON FLUORALKANOL", patDoc.Title[0].Text)
	ass.Equal("de", patDoc.Title[0].Language)
	ass.Equal("PROCESS FOR PRODUCING FLUOROALKANOL", patDoc.Title[1].Text)
	ass.Equal("en", patDoc.Title[1].Language)
	ass.Equal("PROCEDE DE PRODUCTION DE FLUOROALCANOL", patDoc.Title[2].Text)
	ass.Equal("fr", patDoc.Title[2].Language)

	// abstract
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(441, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// claims
	ass.NotEmpty(patDoc.Claims)
	ass.Equal(1, len(patDoc.Claims))
	ass.Equal(1807, len(patDoc.Claims[0].Text))
	ass.Equal("en", patDoc.Claims[0].Language)
	ass.Equal("claims01", patDoc.Claims[0].Id)

	// description
	ass.NotEmpty(patDoc.Description)
	ass.Equal(17802, len(patDoc.Description[0].Text))
	ass.Equal("en", patDoc.Description[0].Language)

	// citations
	ass.Empty(patDoc.Citations)

	// Inventors
	ass.NotEmpty(patDoc.Inventors)
	ass.Equal(Country("JP"), patDoc.Inventors[0].Country)
	ass.Equal("Ichihara-shi,\nChiba 290-8566", patDoc.Inventors[0].City)
	ass.Equal("10, Goikaigan", patDoc.Inventors[0].Street)
	ass.Equal("TOHMA, Toshihiko", patDoc.Inventors[0].Name)

	ass.Equal(Country("JP"), patDoc.Inventors[1].Country)
	ass.Equal("Ichihara-shi,\nChiba 290-8566", patDoc.Inventors[1].City)
	ass.Equal("10, Goikaigan", patDoc.Inventors[1].Street)
	ass.Equal("WADA, Akihiro", patDoc.Inventors[1].Name)

	// representatives
	ass.NotEmpty(patDoc.Representatives)
	ass.Equal(Country("DE"), patDoc.Representatives[0].Country)
	ass.Equal("00100651", patDoc.Representatives[0].IID)
	ass.Equal("81671 München", patDoc.Representatives[0].City)
	ass.Equal("Grafinger Strasse 2", patDoc.Representatives[0].Street)
	ass.Equal("Müller-Boré & Partner\nPatentanwälte", patDoc.Representatives[0].Name)

	// contracting states
	ass.NotEmpty(patDoc.ContractingStates)
	ass.Equal(20, len(patDoc.ContractingStates))
	for i := 0; i <= 19; i++ {
		ass.Equal(2, len(patDoc.ContractingStates[i]))
	}

	// classifications
	ass.Empty(patDoc.Classifications)
}

// dtd version = ep bulletin v1-0?
func TestProcessXMLSimple10A11(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("test-data/application/v1-0-A1-1.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("EP04715346A1", patDoc.ID)
	ass.Equal("04715346.5", patDoc.File)
	ass.Equal("fr", patDoc.Lang)
	ass.Equal(Country("EP"), patDoc.Country)
	ass.Equal("1605745", patDoc.DocNumber)
	ass.Equal("A1", patDoc.Kind)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20051221", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("n", patDoc.Status)
	ass.Equal("ep-bulletin-v1-0", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("VORRICHTUNG ZUM FÜLLEN VON KERNBOHRUNGSLÖCHERN IN EINEM RASEN", patDoc.Title[0].Text)
	ass.Equal("de", patDoc.Title[0].Language)
	ass.Equal("DEVICE FOR FILLING CARROTING HOLES IN A LAWN", patDoc.Title[1].Text)
	ass.Equal("en", patDoc.Title[1].Language)
	ass.Equal("DISPOSITIF POUR LE REMPLISSAGE DES TROUS DE CAROTTAGE D'UNE PELOUSE", patDoc.Title[2].Text)
	ass.Equal("fr", patDoc.Title[2].Language)

	// abstract zero
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(0, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// claims empty
	ass.Empty(patDoc.Claims)

	// description zero
	ass.NotEmpty(patDoc.Description)
	ass.Equal(0, len(patDoc.Description[0].Text))
	ass.Equal("en", patDoc.Description[0].Language)

	// citations empty
	ass.Empty(patDoc.Citations)

	// Inventors
	ass.NotEmpty(patDoc.Inventors)
	ass.Equal(Country("FR"), patDoc.Inventors[0].Country)
	ass.Equal("66180 Villeneuve de la Raho", patDoc.Inventors[0].City)
	ass.Equal("2, allée des Romarins", patDoc.Inventors[0].Street)
	ass.Equal("Cuadrado, Jean-Claude", patDoc.Inventors[0].Name)

	// representatives
	ass.NotEmpty(patDoc.Representatives)
	ass.Equal(Country("FR"), patDoc.Representatives[0].Country)
	ass.Equal("00015907", patDoc.Representatives[0].IID)
	ass.Equal("34000 Montpellier", patDoc.Representatives[0].City)
	ass.Equal("Cabinet Claude BES 2 Bis Rue de Verdun", patDoc.Representatives[0].Street)
	ass.Equal("Gosse, Michel", patDoc.Representatives[0].Name)

	// contracting states
	ass.NotEmpty(patDoc.ContractingStates)
	ass.Equal(27, len(patDoc.ContractingStates))
	for i := 0; i <= 26; i++ {
		ass.Equal(2, len(patDoc.ContractingStates[i]))
	}

	// classifications empty
	ass.Empty(patDoc.Classifications)
}

func TestProcessXMLSimple10A2(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("test-data/application/v1-0-A2.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("EP02019404A2", patDoc.ID)
	ass.Equal("02019404.xml", patDoc.File)
	ass.Equal("de", patDoc.Lang)
	ass.Equal(Country("EP"), patDoc.Country)
	ass.Equal("1326188", patDoc.DocNumber)
	ass.Equal("A2", patDoc.Kind)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20030709", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("n", patDoc.Status)
	ass.Equal("ep-patent-document-v1-0", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("Verfahren und System zur Ermittlung von infolge der Nutzung einer Anlage anfallenden Nutzungsgebühren", patDoc.Title[0].Text)
	ass.Equal("de", patDoc.Title[0].Language)
	ass.Equal("Method and system for establishing the usage costs for the use of an apparatus", patDoc.Title[1].Text)
	ass.Equal("en", patDoc.Title[1].Language)
	ass.Equal("Procédé et appareil pour la génération des coûts d' utilisation pour l' utilisation d' un appareil", patDoc.Title[2].Text)
	ass.Equal("fr", patDoc.Title[2].Language)

	// abstract
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(1085, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// claims
	ass.NotEmpty(patDoc.Claims)
	ass.Equal(1, len(patDoc.Claims))
	ass.Equal(6334, len(patDoc.Claims[0].Text))
	ass.Equal("de", patDoc.Claims[0].Language)
	ass.Equal("claims01", patDoc.Claims[0].Id)

	// description
	ass.NotEmpty(patDoc.Description)
	ass.Equal(22464, len(patDoc.Description[0].Text))
	ass.Equal("en", patDoc.Description[0].Language)

	// citations
	ass.Empty(patDoc.Citations)

	// Inventors
	ass.NotEmpty(patDoc.Inventors)
	ass.Equal(Country("DE"), patDoc.Inventors[0].Country)
	ass.Equal("32760 Detmold", patDoc.Inventors[0].City)
	ass.Equal("Auf dem Brinke 18", patDoc.Inventors[0].Street)
	ass.Equal("Bibelhausen, Volker, Dipl.-Ing.", patDoc.Inventors[0].Name)

	// representatives
	ass.Empty(patDoc.Representatives)

	// contracting states
	ass.NotEmpty(patDoc.ContractingStates)
	ass.Equal(24, len(patDoc.ContractingStates))
	for i := 0; i <= 23; i++ {
		ass.Equal(2, len(patDoc.ContractingStates[i]))
	}

	// classifications
	ass.Empty(patDoc.Classifications)
}

// again dtd version ep bulletin v1-0
func TestProcessXMLSimple10A21(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("test-data/application/v1-0-A2-1.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("EP04718673A2", patDoc.ID)
	ass.Equal("04718673.9", patDoc.File)
	ass.Equal("fr", patDoc.Lang)
	ass.Equal(Country("EP"), patDoc.Country)
	ass.Equal("1605748", patDoc.DocNumber)
	ass.Equal("A2", patDoc.Kind)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20051221", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("n", patDoc.Status)
	ass.Equal("ep-bulletin-v1-0", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("VORRICHTUNG UM PRODUKTE WIE GRAS ZU GRUPPIEREN", patDoc.Title[0].Text)
	ass.Equal("de", patDoc.Title[0].Language)
	ass.Equal("MACHINE FOR GATHERING PRODUCTS SUCH AS GRASS", patDoc.Title[1].Text)
	ass.Equal("en", patDoc.Title[1].Language)
	ass.Equal("MACHINE POUR GROUPER DES PRODUITS TELS QUE DE L'HERBE", patDoc.Title[2].Text)
	ass.Equal("fr", patDoc.Title[2].Language)

	// abstract zero
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(0, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// claims empty
	ass.Empty(patDoc.Claims)

	// description zero
	ass.NotEmpty(patDoc.Description)
	ass.Equal(0, len(patDoc.Description[0].Text))
	ass.Equal("en", patDoc.Description[0].Language)

	// citations
	ass.Empty(patDoc.Citations)

	// Inventors
	ass.NotEmpty(patDoc.Inventors)
	ass.Equal(Country("DE"), patDoc.Inventors[0].Country)
	ass.Equal("38173 Dettum", patDoc.Inventors[0].City)
	ass.Equal("Elmstrasse No 9", patDoc.Inventors[0].Street)
	ass.Equal("GEISER, Jens", patDoc.Inventors[0].Name)

	// representatives
	ass.NotEmpty(patDoc.Representatives)
	ass.Equal(Country("FR"), patDoc.Representatives[0].Country)
	ass.Equal("67700 Saverne", patDoc.Representatives[0].City)
	ass.Equal("00044261", patDoc.Representatives[0].IID)
	ass.Equal("KUHN S.A. 4, Impasse des Fabriques", patDoc.Representatives[0].Street)
	ass.Equal("Andres, Jean-Claude", patDoc.Representatives[0].Name)

	// contracting states
	ass.NotEmpty(patDoc.ContractingStates)
	ass.Equal(28, len(patDoc.ContractingStates))
	for i := 0; i <= 27; i++ {
		ass.Equal(2, len(patDoc.ContractingStates[i]))
	}

	// classifications empty
	ass.Empty(patDoc.Classifications)
}

func TestProcessXMLSimple10B1(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("test-data/grant/v1-0-B1.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("EP02779063B1", patDoc.ID)
	ass.Equal("02779063.xml", patDoc.File)
	ass.Equal("en", patDoc.Lang)
	ass.Equal(Country("EP"), patDoc.Country)
	ass.Equal("1442058", patDoc.DocNumber)
	ass.Equal("B1", patDoc.Kind)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20060719", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("n", patDoc.Status)
	ass.Equal("ep-patent-document-v1-0", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("PEPTIDE ZUM EINSATZ IN DER BEHANDLUNG VON TUMOREN UND ANDEREN ZUSTÄNDEN, DIE DAS ENTFERNEN ODER ZERSTÖREN VON ZELLEN ERFORDERN", patDoc.Title[0].Text)
	ass.Equal("de", patDoc.Title[0].Language)
	ass.Equal("PEPTIDES EFFECTIVE IN THE TREATMENT OF TUMORS AND OTHER CONDITIONS REQUIRING THE REMOVAL OR DESTRUCTION OF CELLS", patDoc.Title[1].Text)
	ass.Equal("en", patDoc.Title[1].Language)
	ass.Equal("PEPTIDES SERVANT A TRAITER DES TUMEURS OU D'AUTRES ETATS NECESSITANT LA SUPPRESSION OU LA DESTRUCTION DE CELLULES", patDoc.Title[2].Text)
	ass.Equal("fr", patDoc.Title[2].Language)

	// abstract
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(0, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// claims
	ass.NotEmpty(patDoc.Claims)
	ass.Equal(3, len(patDoc.Claims))
	ass.Equal(775, len(patDoc.Claims[0].Text))
	ass.Equal("en", patDoc.Claims[0].Language)
	ass.Equal("claims01", patDoc.Claims[0].Id)
	ass.Equal(790, len(patDoc.Claims[1].Text))
	ass.Equal("de", patDoc.Claims[1].Language)
	ass.Equal("claims02", patDoc.Claims[1].Id)
	ass.Equal(798, len(patDoc.Claims[2].Text))
	ass.Equal("fr", patDoc.Claims[2].Language)
	ass.Equal("claims03", patDoc.Claims[2].Id)

	// description
	ass.NotEmpty(patDoc.Description)
	ass.Equal(79384, len(patDoc.Description[0].Text))
	ass.Equal("en", patDoc.Description[0].Language)

	// citations
	ass.Empty(patDoc.Citations)

	// Inventors
	ass.NotEmpty(patDoc.Inventors)
	ass.Equal(Country("CA"), patDoc.Inventors[0].Country)
	ass.Equal("Beaconsfield, Québec H9W 4J1", patDoc.Inventors[0].City)
	ass.Equal("383 Lakeshore Road", patDoc.Inventors[0].Street)
	ass.Equal("AVERBACK, Paul", patDoc.Inventors[0].Name)

	// representatives
	ass.NotEmpty(patDoc.Representatives)
	ass.Equal(1, len(patDoc.Representatives))
	ass.Equal(Country("DK"), patDoc.Representatives[0].Country)
	ass.Equal("Plougmann & Vingtoft A/S", patDoc.Representatives[0].Name)
	ass.Equal("00101177", patDoc.Representatives[0].IID)
	ass.Equal("2100 Copenhagen", patDoc.Representatives[0].City)
	ass.Equal("Sundkrogsgade 9, \nP.O. Box 831", patDoc.Representatives[0].Street)

	// contracting states
	ass.NotEmpty(patDoc.ContractingStates)
	ass.Equal(24, len(patDoc.ContractingStates))
	for i := 0; i <= 23; i++ {
		ass.Equal(2, len(patDoc.ContractingStates[i]))
	}

	// classifications
	ass.NotEmpty(patDoc.Classifications)
	ass.Equal(2, len(patDoc.Classifications))

	ass.Equal("C07K  14/47        20060101AFI20050713BHEP        ", patDoc.Classifications[0].Text)
	ass.Equal(IPC, patDoc.Classifications[0].System)
	ass.Equal(1, patDoc.Classifications[0].Sequence)
	ass.Equal("C", patDoc.Classifications[0].Section)
	ass.Equal("07", patDoc.Classifications[0].Class)
	ass.Equal("K", patDoc.Classifications[0].SubClass)
	ass.Equal("14", patDoc.Classifications[0].MainGroup)
	ass.Equal("47", patDoc.Classifications[0].SubGroup)
	ass.Equal("20060101", patDoc.Classifications[0].Version)
	ass.Equal("A", patDoc.Classifications[0].ClassificationLevel)
	ass.Equal("F", patDoc.Classifications[0].FirstLater)
	ass.Equal("I", patDoc.Classifications[0].ClassificationValue)
	ass.Equal("20050713", patDoc.Classifications[0].ActionDate)
	ass.Equal("B", patDoc.Classifications[0].OriginalOrReclassified)
	ass.Equal("H", patDoc.Classifications[0].Source)
	ass.Equal("EP", patDoc.Classifications[0].GeneratingOffice)

	ass.Equal("A61K  38/17        20060101ALI20050713BHEP        ", patDoc.Classifications[1].Text)
	ass.Equal(IPC, patDoc.Classifications[1].System)
	ass.Equal(2, patDoc.Classifications[1].Sequence)
	ass.Equal("A", patDoc.Classifications[1].Section)
	ass.Equal("61", patDoc.Classifications[1].Class)
	ass.Equal("K", patDoc.Classifications[1].SubClass)
	ass.Equal("38", patDoc.Classifications[1].MainGroup)
	ass.Equal("17", patDoc.Classifications[1].SubGroup)
	ass.Equal("20060101", patDoc.Classifications[1].Version)
	ass.Equal("A", patDoc.Classifications[1].ClassificationLevel)
	ass.Equal("L", patDoc.Classifications[1].FirstLater)
	ass.Equal("I", patDoc.Classifications[1].ClassificationValue)
	ass.Equal("20050713", patDoc.Classifications[1].ActionDate)
	ass.Equal("B", patDoc.Classifications[1].OriginalOrReclassified)
	ass.Equal("H", patDoc.Classifications[1].Source)
	ass.Equal("EP", patDoc.Classifications[1].GeneratingOffice)
}

func TestProcessXMLSimple10B11(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("test-data/grant/v1-0-B1-1.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("EP94113456B1", patDoc.ID)
	ass.Equal("94113456.xml", patDoc.File)
	ass.Equal("en", patDoc.Lang)
	ass.Equal(Country("EP"), patDoc.Country)
	ass.Equal("0640886", patDoc.DocNumber)
	ass.Equal("B1", patDoc.Kind)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20060426", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("n", patDoc.Status)
	ass.Equal("ep-patent-document-v1-0", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("Elektrophotografisches Gerät und Prozesseinheit", patDoc.Title[0].Text)
	ass.Equal("de", patDoc.Title[0].Language)
	ass.Equal("Electrophotographic apparatus and process cartridge", patDoc.Title[1].Text)
	ass.Equal("en", patDoc.Title[1].Language)
	ass.Equal("Appareil électrophotographique et unité de traitement", patDoc.Title[2].Text)
	ass.Equal("fr", patDoc.Title[2].Language)

	// abstract
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(0, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// claims
	ass.NotEmpty(patDoc.Claims)
	ass.Equal(3, len(patDoc.Claims))
	ass.Equal(3492, len(patDoc.Claims[0].Text))
	ass.Equal("en", patDoc.Claims[0].Language)
	ass.Equal("claims01", patDoc.Claims[0].Id)
	ass.Equal(4049, len(patDoc.Claims[1].Text))
	ass.Equal("fr", patDoc.Claims[1].Language)
	ass.Equal("claims02", patDoc.Claims[1].Id)
	ass.Equal(3952, len(patDoc.Claims[2].Text))
	ass.Equal("de", patDoc.Claims[2].Language)
	ass.Equal("claims03", patDoc.Claims[2].Id)

	// description
	ass.NotEmpty(patDoc.Description)
	ass.Equal(22371, len(patDoc.Description[0].Text))
	ass.Equal("en", patDoc.Description[0].Language)

	// citations
	ass.Empty(patDoc.Citations)

	// Inventors
	ass.NotEmpty(patDoc.Inventors)
	ass.Equal(6, len(patDoc.Inventors))

	for i := 0; i <= 5; i++ {
		ass.Equal(Country("JP"), patDoc.Inventors[i].Country)
		ass.Equal("Tokyo 146", patDoc.Inventors[i].City)
		ass.Equal("30-2, 3-chome, Shimomaruko,\nOhta-ku", patDoc.Inventors[i].Street)
	}

	ass.Equal("Yoshihara, Toshiyuki,\nc/o Canon Kabushiki Kaisha", patDoc.Inventors[0].Name)
	ass.Equal("Anayama, Hideki,\nc/o Canon Kabushiki Kaisha", patDoc.Inventors[1].Name)
	ass.Equal("Yamazaki, Itaru,\nc/o Canon Kabushiki Kaisha", patDoc.Inventors[2].Name)
	ass.Equal("Ainoya, Hideyuki,\nc/o Canon Kabushiki Kaisha", patDoc.Inventors[3].Name)
	ass.Equal("Hirano, Hidetoshi,\nc/o Canon Kabushiki Kaisha", patDoc.Inventors[4].Name)
	ass.Equal("Kimura, Mayumi,\nc/o Canon Kabushiki Kaisha", patDoc.Inventors[5].Name)

	// representatives
	ass.NotEmpty(patDoc.Representatives)
	ass.Equal(1, len(patDoc.Representatives))
	ass.Equal(Country("DE"), patDoc.Representatives[0].Country)
	ass.Equal("TBK-Patent", patDoc.Representatives[0].Name)
	ass.Equal("00102381", patDoc.Representatives[0].IID)
	ass.Equal("80336 München", patDoc.Representatives[0].City)
	ass.Equal("Bavariaring 4-6", patDoc.Representatives[0].Street)

	// contracting states
	ass.NotEmpty(patDoc.ContractingStates)
	ass.Equal(4, len(patDoc.ContractingStates))
	for i := 0; i <= 3; i++ {
		ass.Equal(2, len(patDoc.ContractingStates[i]))
	}

	// classifications
	ass.NotEmpty(patDoc.Classifications)
	ass.Equal(2, len(patDoc.Classifications))

	ass.Equal("G03G  15/08        20060101AFI19941209BHEP        ", patDoc.Classifications[0].Text)
	ass.Equal(IPC, patDoc.Classifications[0].System)
	ass.Equal(1, patDoc.Classifications[0].Sequence)
	ass.Equal("G", patDoc.Classifications[0].Section)
	ass.Equal("03", patDoc.Classifications[0].Class)
	ass.Equal("G", patDoc.Classifications[0].SubClass)
	ass.Equal("15", patDoc.Classifications[0].MainGroup)
	ass.Equal("08", patDoc.Classifications[0].SubGroup)
	ass.Equal("20060101", patDoc.Classifications[0].Version)
	ass.Equal("A", patDoc.Classifications[0].ClassificationLevel)
	ass.Equal("F", patDoc.Classifications[0].FirstLater)
	ass.Equal("I", patDoc.Classifications[0].ClassificationValue)
	ass.Equal("19941209", patDoc.Classifications[0].ActionDate)
	ass.Equal("B", patDoc.Classifications[0].OriginalOrReclassified)
	ass.Equal("H", patDoc.Classifications[0].Source)
	ass.Equal("EP", patDoc.Classifications[0].GeneratingOffice)

	ass.Equal("G03G   5/147       20060101ALI19941209BHEP        ", patDoc.Classifications[1].Text)
	ass.Equal(IPC, patDoc.Classifications[1].System)
	ass.Equal(2, patDoc.Classifications[1].Sequence)
	ass.Equal("G", patDoc.Classifications[1].Section)
	ass.Equal("03", patDoc.Classifications[1].Class)
	ass.Equal("G", patDoc.Classifications[1].SubClass)
	ass.Equal("5", patDoc.Classifications[1].MainGroup)
	ass.Equal("147", patDoc.Classifications[1].SubGroup)
	ass.Equal("20060101", patDoc.Classifications[1].Version)
	ass.Equal("A", patDoc.Classifications[1].ClassificationLevel)
	ass.Equal("L", patDoc.Classifications[1].FirstLater)
	ass.Equal("I", patDoc.Classifications[1].ClassificationValue)
	ass.Equal("19941209", patDoc.Classifications[1].ActionDate)
	ass.Equal("B", patDoc.Classifications[1].OriginalOrReclassified)
	ass.Equal("H", patDoc.Classifications[1].Source)
	ass.Equal("EP", patDoc.Classifications[1].GeneratingOffice)
}

func TestProcessXMLSimple10B2(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("test-data/grant/v1-0-B2.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("EP96939832B2", patDoc.ID)
	ass.Equal("96939832.xml", patDoc.File)
	ass.Equal("en", patDoc.Lang)
	ass.Equal(Country("EP"), patDoc.Country)
	ass.Equal("0874807", patDoc.DocNumber)
	ass.Equal("B2", patDoc.Kind)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20060802", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("n", patDoc.Status)
	ass.Equal("ep-patent-document-v1-0", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("VERFAHREN ZUR HERSTELLUNG VON METHOXIMINOPHENYLGLYOXYLSÄURE-DERIVATEN", patDoc.Title[0].Text)
	ass.Equal("de", patDoc.Title[0].Language)
	ass.Equal("PROCESS FOR THE PREPARATION OF METHOXYMINOPHENYLGLYOXYLIC ACID DERIVATIVES", patDoc.Title[1].Text)
	ass.Equal("en", patDoc.Title[1].Language)
	ass.Equal("PROCEDE DE PREPARATION DE DERIVES DE L'ACIDE METHOXIMINOPHENYLGLYOXYLIQUE", patDoc.Title[2].Text)
	ass.Equal("fr", patDoc.Title[2].Language)

	// abstract
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(0, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// claims
	ass.NotEmpty(patDoc.Claims)
	ass.Equal(3, len(patDoc.Claims))
	ass.Equal(858, len(patDoc.Claims[0].Text))
	ass.Equal("en", patDoc.Claims[0].Language)
	ass.Equal("claims01", patDoc.Claims[0].Id)
	ass.Equal(893, len(patDoc.Claims[1].Text))
	ass.Equal("de", patDoc.Claims[1].Language)
	ass.Equal("claims02", patDoc.Claims[1].Id)
	ass.Equal(1008, len(patDoc.Claims[2].Text))
	ass.Equal("fr", patDoc.Claims[2].Language)
	ass.Equal("claims03", patDoc.Claims[2].Id)

	// description
	ass.NotEmpty(patDoc.Description)
	ass.Equal(10357, len(patDoc.Description[0].Text))
	ass.Equal("en", patDoc.Description[0].Language)

	// citations
	ass.Empty(patDoc.Citations)

	// Inventors
	ass.NotEmpty(patDoc.Inventors)
	ass.Equal(Country("CH"), patDoc.Inventors[0].Country)
	ass.Equal("CH-1870 Monthey", patDoc.Inventors[0].City)
	ass.Equal("Avenue de l'Europe 38A", patDoc.Inventors[0].Street)
	ass.Equal("ASSERCQ, Jean-Marie", patDoc.Inventors[0].Name)

	// representatives
	ass.Empty(patDoc.Representatives)

	// contracting states
	ass.NotEmpty(patDoc.ContractingStates)
	ass.Equal(17, len(patDoc.ContractingStates))
	for i := 0; i <= 16; i++ {
		ass.Equal(2, len(patDoc.ContractingStates[i]))
	}

	// classifications
	ass.NotEmpty(patDoc.Classifications)
	ass.Equal(5, len(patDoc.Classifications))

	for i := 0; i <= 2; i++ {
		ass.Equal(IPC, patDoc.Classifications[i].System)
		ass.Equal(i+1, patDoc.Classifications[i].Sequence)
		ass.Equal("C", patDoc.Classifications[i].Section)
		ass.Equal("07", patDoc.Classifications[i].Class)
		ass.Equal("C", patDoc.Classifications[i].SubClass)

		ass.Equal("20060101", patDoc.Classifications[i].Version)
		ass.Equal("A", patDoc.Classifications[i].ClassificationLevel)
		ass.Equal("19970721", patDoc.Classifications[i].ActionDate)
		ass.Equal("B", patDoc.Classifications[i].OriginalOrReclassified)
		ass.Equal("H", patDoc.Classifications[i].Source)
		ass.Equal("EP", patDoc.Classifications[i].GeneratingOffice)
	}

	ass.Equal("C07C 249/08        20060101AFI19970721BHEP        ", patDoc.Classifications[0].Text)
	ass.Equal("249", patDoc.Classifications[0].MainGroup)
	ass.Equal("08", patDoc.Classifications[0].SubGroup)
	ass.Equal("F", patDoc.Classifications[0].FirstLater)
	ass.Equal("I", patDoc.Classifications[0].ClassificationValue)

	ass.Equal("C07C 231/06        20060101ALI19970721BHEP        ", patDoc.Classifications[1].Text)
	ass.Equal("231", patDoc.Classifications[1].MainGroup)
	ass.Equal("06", patDoc.Classifications[1].SubGroup)
	ass.Equal("L", patDoc.Classifications[1].FirstLater)
	ass.Equal("I", patDoc.Classifications[1].ClassificationValue)

	ass.Equal("C07C 235/78        20060101ALI19970721BHEP        ", patDoc.Classifications[2].Text)
	ass.Equal("235", patDoc.Classifications[2].MainGroup)
	ass.Equal("78", patDoc.Classifications[2].SubGroup)
	ass.Equal("I", patDoc.Classifications[2].ClassificationValue)
	ass.Equal("L", patDoc.Classifications[2].FirstLater)

	ass.Equal("C07C 251/48        20060101ALI19970721BHEP        ", patDoc.Classifications[3].Text)
	ass.Equal("251", patDoc.Classifications[3].MainGroup)
	ass.Equal("48", patDoc.Classifications[3].SubGroup)
	ass.Equal("I", patDoc.Classifications[3].ClassificationValue)
	ass.Equal("L", patDoc.Classifications[3].FirstLater)

	ass.Equal("C07C 251/60        20060101ALI19970721BHEP        ", patDoc.Classifications[4].Text)
	ass.Equal("251", patDoc.Classifications[4].MainGroup)
	ass.Equal("60", patDoc.Classifications[4].SubGroup)
	ass.Equal("I", patDoc.Classifications[4].ClassificationValue)
	ass.Equal("L", patDoc.Classifications[4].FirstLater)

}

func TestProcessXMLSimple10B21(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("test-data/grant/v1-0-B2-1.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("EP95912446B2", patDoc.ID)
	ass.Equal("95912446.xml", patDoc.File)
	ass.Equal("en", patDoc.Lang)
	ass.Equal(Country("EP"), patDoc.Country)
	ass.Equal("0716240", patDoc.DocNumber)
	ass.Equal("B2", patDoc.Kind)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20060426", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("n", patDoc.Status)
	ass.Equal("ep-patent-document-v1-0", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("GLEITLAGER", patDoc.Title[0].Text)
	ass.Equal("de", patDoc.Title[0].Language)
	ass.Equal("SLIDING BEARING", patDoc.Title[1].Text)
	ass.Equal("en", patDoc.Title[1].Language)
	ass.Equal("PALIER A GLISSEMENT", patDoc.Title[2].Text)
	ass.Equal("fr", patDoc.Title[2].Language)

	// abstract
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(0, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// claims
	ass.NotEmpty(patDoc.Claims)
	ass.Equal(3, len(patDoc.Claims))
	ass.Equal(2356, len(patDoc.Claims[0].Text))
	ass.Equal("en", patDoc.Claims[0].Language)
	ass.Equal("claims01", patDoc.Claims[0].Id)
	ass.Equal(2517, len(patDoc.Claims[1].Text))
	ass.Equal("de", patDoc.Claims[1].Language)
	ass.Equal("claims02", patDoc.Claims[1].Id)
	ass.Equal(2665, len(patDoc.Claims[2].Text))
	ass.Equal("fr", patDoc.Claims[2].Language)
	ass.Equal("claims03", patDoc.Claims[2].Id)

	// description
	ass.NotEmpty(patDoc.Description)
	ass.Equal(18525, len(patDoc.Description[0].Text))
	ass.Equal("en", patDoc.Description[0].Language)

	// citations
	ass.Empty(patDoc.Citations)

	// Inventors
	ass.NotEmpty(patDoc.Inventors)
	ass.Equal(Country("JP"), patDoc.Inventors[0].Country)
	ass.Equal("Aichi 471", patDoc.Inventors[0].City)
	ass.Equal("65, Midorigaoka 3-chome\nToyota-shi", patDoc.Inventors[0].Street)
	ass.Equal("KUMADA, Yoshio\nTaiho Kogyo Co., Ltd.", patDoc.Inventors[0].Name)

	// representatives
	ass.NotEmpty(patDoc.Representatives)
	ass.Equal(Country("GB"), patDoc.Representatives[0].Country)
	ass.Equal("00034621", patDoc.Representatives[0].IID)
	ass.Equal("London WC2B 6HP", patDoc.Representatives[0].City)
	ass.Equal("Paget, Hugh Charles Edward", patDoc.Representatives[0].Name)
	ass.Equal("Mewburn Ellis LLP \nYork House \n23 Kingsway", patDoc.Representatives[0].Street)

	// contracting states
	ass.NotEmpty(patDoc.ContractingStates)
	ass.Equal(2, len(patDoc.ContractingStates))
	for i := 0; i <= 1; i++ {
		ass.Equal(2, len(patDoc.ContractingStates[i]))
	}

	// classifications
	ass.NotEmpty(patDoc.Classifications)
	ass.Equal(2, len(patDoc.Classifications))

	ass.Equal("F16C  33/12        20060101AFI19951016BHEP        ", patDoc.Classifications[0].Text)
	ass.Equal(IPC, patDoc.Classifications[0].System)
	ass.Equal(1, patDoc.Classifications[0].Sequence)
	ass.Equal("F", patDoc.Classifications[0].Section)
	ass.Equal("16", patDoc.Classifications[0].Class)
	ass.Equal("C", patDoc.Classifications[0].SubClass)
	ass.Equal("33", patDoc.Classifications[0].MainGroup)
	ass.Equal("12", patDoc.Classifications[0].SubGroup)
	ass.Equal("20060101", patDoc.Classifications[0].Version)
	ass.Equal("A", patDoc.Classifications[0].ClassificationLevel)
	ass.Equal("F", patDoc.Classifications[0].FirstLater)
	ass.Equal("I", patDoc.Classifications[0].ClassificationValue)
	ass.Equal("19951016", patDoc.Classifications[0].ActionDate)
	ass.Equal("B", patDoc.Classifications[0].OriginalOrReclassified)
	ass.Equal("H", patDoc.Classifications[0].Source)
	ass.Equal("EP", patDoc.Classifications[0].GeneratingOffice)

	ass.Equal("F16C  33/24        20060101ALI19960611BHEP        ", patDoc.Classifications[1].Text)
	ass.Equal(IPC, patDoc.Classifications[1].System)
	ass.Equal(2, patDoc.Classifications[1].Sequence)
	ass.Equal("F", patDoc.Classifications[1].Section)
	ass.Equal("16", patDoc.Classifications[1].Class)
	ass.Equal("C", patDoc.Classifications[1].SubClass)
	ass.Equal("33", patDoc.Classifications[1].MainGroup)
	ass.Equal("24", patDoc.Classifications[1].SubGroup)
	ass.Equal("20060101", patDoc.Classifications[1].Version)
	ass.Equal("A", patDoc.Classifications[1].ClassificationLevel)
	ass.Equal("L", patDoc.Classifications[1].FirstLater)
	ass.Equal("I", patDoc.Classifications[1].ClassificationValue)
	ass.Equal("19960611", patDoc.Classifications[1].ActionDate)
	ass.Equal("B", patDoc.Classifications[1].OriginalOrReclassified)
	ass.Equal("H", patDoc.Classifications[1].Source)
	ass.Equal("EP", patDoc.Classifications[1].GeneratingOffice)

}

//v 1.01
func TestProcessXMLSimple101A1(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("test-data/application/v1-01-A1.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("EP04768329A1", patDoc.ID)
	ass.Equal("04768329.7", patDoc.File)
	ass.Equal("en", patDoc.Lang)
	ass.Equal(Country("EP"), patDoc.Country)
	ass.Equal("1679948", patDoc.DocNumber)
	ass.Equal("A1", patDoc.Kind)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20060719", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("n", patDoc.Status)
	ass.Equal("ep-patent-document-v1-01", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("BAHNABDECKUNGEN FÜR FÖRDERSIEBE", patDoc.Title[0].Text)
	ass.Equal("de", patDoc.Title[0].Language)
	ass.Equal("WEB COVERS FOR CONVEYOR SCREENS", patDoc.Title[1].Text)
	ass.Equal("en", patDoc.Title[1].Language)
	ass.Equal("GARNITURES POUR TAPIS TRANSPORTEURS", patDoc.Title[2].Text)
	ass.Equal("fr", patDoc.Title[2].Language)

	// abstract
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(0, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// claims
	ass.Empty(patDoc.Claims)

	// description
	ass.NotEmpty(patDoc.Description)
	ass.Equal(0, len(patDoc.Description[0].Text))
	ass.Equal("en", patDoc.Description[0].Language)

	// citations
	ass.Empty(patDoc.Citations)

	// Inventors
	ass.NotEmpty(patDoc.Inventors)
	ass.Equal(Country("GB"), patDoc.Inventors[0].Country)
	ass.Equal("Bangor BT19 7QT,\nCounty Down", patDoc.Inventors[0].City)
	ass.Equal("Unit 35, \n2-4 Ballo Avenue", patDoc.Inventors[0].Street)
	ass.Equal("Rankin, David", patDoc.Inventors[0].Name)

	ass.Equal(Country("GB"), patDoc.Inventors[1].Country)
	ass.Equal("Bangor BT19 7QT,\nCounty Down", patDoc.Inventors[1].City)
	ass.Equal("Unit 35, \nUnit 2-4 Balloo Avenue", patDoc.Inventors[1].Street)
	ass.Equal("Shanks, Anne", patDoc.Inventors[1].Name)

	// representatives
	ass.NotEmpty(patDoc.Representatives)
	ass.Equal(Country("IE"), patDoc.Representatives[0].Country)
	ass.Equal("00143541", patDoc.Representatives[0].IID)
	ass.Equal("Dublin 2", patDoc.Representatives[0].City)
	ass.Equal("MacLachlan & Donaldson \n47 Merrion Square", patDoc.Representatives[0].Street)
	ass.Equal("Hanna, John Philip", patDoc.Representatives[0].Name)

	// contracting states
	ass.NotEmpty(patDoc.ContractingStates)
	ass.Equal(28, len(patDoc.ContractingStates))
	for i := 0; i <= 27; i++ {
		ass.Equal(2, len(patDoc.ContractingStates[i]))
	}

	// classifications
	ass.NotEmpty(patDoc.Classifications)
	ass.Equal(1, len(patDoc.Classifications))
	ass.Equal("A01D  17/10        19680901AFI20050322BHEP        ", patDoc.Classifications[0].Text)
	ass.Equal(IPC, patDoc.Classifications[0].System)
	ass.Equal(1, patDoc.Classifications[0].Sequence)
	ass.Equal("A", patDoc.Classifications[0].Section)
	ass.Equal("01", patDoc.Classifications[0].Class)
	ass.Equal("D", patDoc.Classifications[0].SubClass)
	ass.Equal("17", patDoc.Classifications[0].MainGroup)
	ass.Equal("10", patDoc.Classifications[0].SubGroup)
	ass.Equal("19680901", patDoc.Classifications[0].Version)
	ass.Equal("A", patDoc.Classifications[0].ClassificationLevel)
	ass.Equal("F", patDoc.Classifications[0].FirstLater)
	ass.Equal("I", patDoc.Classifications[0].ClassificationValue)
	ass.Equal("20050322", patDoc.Classifications[0].ActionDate)
	ass.Equal("B", patDoc.Classifications[0].OriginalOrReclassified)
	ass.Equal("H", patDoc.Classifications[0].Source)
	ass.Equal("EP", patDoc.Classifications[0].GeneratingOffice)
}

func TestProcessXMLSimple101A2(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("test-data/application/v1-01-A2.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("EP04794205A2", patDoc.ID)
	ass.Equal("04794205.7", patDoc.File)
	ass.Equal("en", patDoc.Lang)
	ass.Equal(Country("EP"), patDoc.Country)
	ass.Equal("1680538", patDoc.DocNumber)
	ass.Equal("A2", patDoc.Kind)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20060719", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("n", patDoc.Status)
	ass.Equal("ep-patent-document-v1-01", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("VERBUNDGARN UND ERZEUGNISSE DARAUS", patDoc.Title[0].Text)
	ass.Equal("de", patDoc.Title[0].Language)
	ass.Equal("COMPOSITE YARN AND PRODUCTS MADE THEREFROM", patDoc.Title[1].Text)
	ass.Equal("en", patDoc.Title[1].Language)
	ass.Equal("FIL COMPOSITE ET PRODUITS FABRIQUES A PARTIR D'UN TEL FIL", patDoc.Title[2].Text)
	ass.Equal("fr", patDoc.Title[2].Language)

	// abstract
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(0, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// claims
	ass.Empty(patDoc.Claims)

	// description
	ass.NotEmpty(patDoc.Description)
	ass.Equal(0, len(patDoc.Description[0].Text))
	ass.Equal("en", patDoc.Description[0].Language)

	// citations
	ass.Empty(patDoc.Citations)

	// Inventors
	ass.NotEmpty(patDoc.Inventors)
	ass.Equal(1, len(patDoc.Inventors))
	ass.Equal(Country("US"), patDoc.Inventors[0].Country)
	ass.Equal("Hickory, NC 28601", patDoc.Inventors[0].City)
	ass.Equal("1740 5th Street Drive N.W.", patDoc.Inventors[0].Street)
	ass.Equal("KOLMES, Nathaniel", patDoc.Inventors[0].Name)

	// representatives
	ass.NotEmpty(patDoc.Representatives)
	ass.Equal(1, len(patDoc.Representatives))
	ass.Equal(Country("FR"), patDoc.Representatives[0].Country)
	ass.Equal("75008 Paris", patDoc.Representatives[0].City)
	ass.Equal("Cabinet ORES, \n36, rue de Saint Pétersbourg", patDoc.Representatives[0].Street)
	ass.Equal("Corizzi, Valérie", patDoc.Representatives[0].Name)
	ass.Equal("00086752", patDoc.Representatives[0].IID)

	// contracting states
	ass.NotEmpty(patDoc.ContractingStates)
	ass.Equal(28, len(patDoc.ContractingStates))
	for i := 0; i <= 27; i++ {
		ass.Equal(2, len(patDoc.ContractingStates[i]))
	}

	// classifications
	ass.NotEmpty(patDoc.Classifications)
	ass.Equal(1, len(patDoc.Classifications))
	ass.Equal("D02G   3/02        19680901AFI20051108BHEP        ", patDoc.Classifications[0].Text)
	ass.Equal(IPC, patDoc.Classifications[0].System)
	ass.Equal(1, patDoc.Classifications[0].Sequence)
	ass.Equal("D", patDoc.Classifications[0].Section)
	ass.Equal("02", patDoc.Classifications[0].Class)
	ass.Equal("G", patDoc.Classifications[0].SubClass)
	ass.Equal("3", patDoc.Classifications[0].MainGroup)
	ass.Equal("02", patDoc.Classifications[0].SubGroup)
	ass.Equal("19680901", patDoc.Classifications[0].Version)
	ass.Equal("A", patDoc.Classifications[0].ClassificationLevel)
	ass.Equal("F", patDoc.Classifications[0].FirstLater)
	ass.Equal("I", patDoc.Classifications[0].ClassificationValue)
	ass.Equal("20051108", patDoc.Classifications[0].ActionDate)
	ass.Equal("B", patDoc.Classifications[0].OriginalOrReclassified)
	ass.Equal("H", patDoc.Classifications[0].Source)
	ass.Equal("EP", patDoc.Classifications[0].GeneratingOffice)

}

//v 1.1
func TestProcessXMLSimple11A1(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("test-data/application/v1-1-A1.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("EP78100007A1", patDoc.ID)
	ass.Equal("EP78100007NWA1.xml", patDoc.File)
	ass.Equal("de", patDoc.Lang)
	ass.Equal(Country("EP"), patDoc.Country)
	ass.Equal("0000002", patDoc.DocNumber)
	ass.Equal("A1", patDoc.Kind)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("19781220", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("n", patDoc.Status)
	ass.Equal("ep-patent-document-v1-1", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("Tetrahydrofuran-Derivate, Verfahren zu ihrer Herstellung sowie ihre Verwendung als Herbizide.", patDoc.Title[0].Text)
	ass.Equal("de", patDoc.Title[0].Language)
	ass.Equal("Tetrahydrofurane derivatives, processes for their preparation and their use as herbicides", patDoc.Title[1].Text)
	ass.Equal("en", patDoc.Title[1].Language)
	ass.Equal("Dérivés du tétrahydrofuranne, leurs procédés de préparation et leur utilisation comme herbicides", patDoc.Title[2].Text)
	ass.Equal("fr", patDoc.Title[2].Language)

	// abstract
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(526, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// claims
	ass.NotEmpty(patDoc.Claims)
	ass.Equal(1, len(patDoc.Claims))
	ass.Equal(6893, len(patDoc.Claims[0].Text))
	ass.Equal("de", patDoc.Claims[0].Language)
	ass.Equal("claims01", patDoc.Claims[0].Id)

	// description
	ass.NotEmpty(patDoc.Description)
	ass.Equal(52416, len(patDoc.Description[0].Text))
	ass.Equal("en", patDoc.Description[0].Language)

	// citations
	ass.Empty(patDoc.Citations)

	// Inventors
	ass.NotEmpty(patDoc.Inventors)
	ass.Equal(4, len(patDoc.Inventors))
	ass.Equal(Country("DE"), patDoc.Inventors[0].Country)
	ass.Equal("D-5600 Wuppertal 1", patDoc.Inventors[0].City)
	ass.Equal("Wormser strasse 23", patDoc.Inventors[0].Street)
	ass.Equal("Schmidt, Thomas, Dr.", patDoc.Inventors[0].Name)

	ass.Equal(Country("DE"), patDoc.Inventors[1].Country)
	ass.Equal("D-5600 Wuppertal 1", patDoc.Inventors[1].City)
	ass.Equal("In den Birken 81", patDoc.Inventors[1].Street)
	ass.Equal("Draber, Wilfried, Dr.", patDoc.Inventors[1].Name)

	ass.Equal(Country("DE"), patDoc.Inventors[2].Country)
	ass.Equal("D- 090 Leverkusen 1", patDoc.Inventors[2].City)
	ass.Equal("Paul-Klee-Strasse 36", patDoc.Inventors[2].Street)
	ass.Equal("Eue, Ludwig, Dr.", patDoc.Inventors[2].Name)

	ass.Equal(Country("DE"), patDoc.Inventors[3].Country)
	ass.Equal("D-5000 Köln 80", patDoc.Inventors[3].City)
	ass.Equal("Hahnenweg 5", patDoc.Inventors[3].Street)
	ass.Equal("Schmidt, Robert Rudolf, Dr.", patDoc.Inventors[3].Name)

	// representatives
	ass.Empty(patDoc.Representatives)

	// contracting states
	ass.NotEmpty(patDoc.ContractingStates)
	ass.Equal(7, len(patDoc.ContractingStates))
	for i := 0; i <= 6; i++ {
		ass.Equal(2, len(patDoc.ContractingStates[i]))
	}

	// classifications
	ass.Empty(patDoc.Classifications)
}

func TestProcessXMLSimple11A2(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("test-data/application/v1-1-A2.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("EP04795906A2", patDoc.ID)
	ass.Equal("04795906.9", patDoc.File)
	ass.Equal("en", patDoc.Lang)
	ass.Equal(Country("EP"), patDoc.Country)
	ass.Equal("1792486", patDoc.DocNumber)
	ass.Equal("A2", patDoc.Kind)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20070606", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("n", patDoc.Status)
	ass.Equal("ep-patent-document-v1-1", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("VERFAHREN  ZUR VEREINHEITLICHUNG VON HETEROGENEN MULTIMEDA TUNER", patDoc.Title[0].Text)
	ass.Equal("de", patDoc.Title[0].Language)
	ass.Equal("SYSTEMS FOR UNIFYING HETEROGENEOUS MULTIMEDIA TUNERS", patDoc.Title[1].Text)
	ass.Equal("en", patDoc.Title[1].Language)
	ass.Equal("SYSTEMES D'UNIFICATION DE TUNERS MULTIMEDIA HETEROGENES", patDoc.Title[2].Text)
	ass.Equal("fr", patDoc.Title[2].Language)

	// abstract
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(0, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// claims Empty
	ass.Empty(patDoc.Claims)

	// description
	ass.NotEmpty(patDoc.Description)
	ass.Equal(0, len(patDoc.Description[0].Text))
	ass.Equal("en", patDoc.Description[0].Language)

	// citations
	ass.Empty(patDoc.Citations)

	// Inventors
	ass.NotEmpty(patDoc.Inventors)
	ass.Equal(1, len(patDoc.Inventors))
	ass.Equal(Country("US"), patDoc.Inventors[0].Country)
	ass.Equal("Redmond, WA 98052", patDoc.Inventors[0].City)
	ass.Equal("One Microsoft Way", patDoc.Inventors[0].Street)
	ass.Equal("POTREBIC, Peter, J.", patDoc.Inventors[0].Name)

	// representatives
	ass.NotEmpty(patDoc.Representatives)
	ass.Equal(Country("DE"), patDoc.Representatives[0].Country)
	ass.Equal("00100721", patDoc.Representatives[0].IID)
	ass.Equal("80538 München", patDoc.Representatives[0].City)
	ass.Equal("Maximilianstrasse 58", patDoc.Representatives[0].Street)
	ass.Equal("Grünecker, Kinkeldey, \nStockmair & Schwanhäusser \nAnwaltssozietät", patDoc.Representatives[0].Name)

	// contracting states
	ass.NotEmpty(patDoc.ContractingStates)
	ass.Equal(28, len(patDoc.ContractingStates))
	for i := 0; i <= 27; i++ {
		ass.Equal(2, len(patDoc.ContractingStates[i]))
	}

	// classifications
	ass.NotEmpty(patDoc.Classifications)
	ass.Equal(1, len(patDoc.Classifications))

	ass.Equal("H04N   7/173       20060101AFI20070423BHEP        ", patDoc.Classifications[0].Text)
	ass.Equal(IPC, patDoc.Classifications[0].System)
	ass.Equal(1, patDoc.Classifications[0].Sequence)
	ass.Equal("H", patDoc.Classifications[0].Section)
	ass.Equal("04", patDoc.Classifications[0].Class)
	ass.Equal("N", patDoc.Classifications[0].SubClass)
	ass.Equal("7", patDoc.Classifications[0].MainGroup)
	ass.Equal("173", patDoc.Classifications[0].SubGroup)
	ass.Equal("20060101", patDoc.Classifications[0].Version)
	ass.Equal("A", patDoc.Classifications[0].ClassificationLevel)
	ass.Equal("F", patDoc.Classifications[0].FirstLater)
	ass.Equal("I", patDoc.Classifications[0].ClassificationValue)
	ass.Equal("20070423", patDoc.Classifications[0].ActionDate)
	ass.Equal("B", patDoc.Classifications[0].OriginalOrReclassified)
	ass.Equal("H", patDoc.Classifications[0].Source)
	ass.Equal("EP", patDoc.Classifications[0].GeneratingOffice)

}

func TestProcessXMLSimple11A21(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("test-data/application/v1-1-A2-1.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("EP94109198A2", patDoc.ID)
	ass.Equal("EP94109198NWA2.xml", patDoc.File)
	ass.Equal("en", patDoc.Lang)
	ass.Equal(Country("EP"), patDoc.Country)
	ass.Equal("0630155", patDoc.DocNumber)
	ass.Equal("A2", patDoc.Kind)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("19941221", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("n", patDoc.Status)
	ass.Equal("ep-patent-document-v1-1", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("Gerät zur Aufzeichnung von digitalen Signalen", patDoc.Title[0].Text)
	ass.Equal("de", patDoc.Title[0].Language)
	ass.Equal("Digital signal recording apparatus", patDoc.Title[1].Text)
	ass.Equal("en", patDoc.Title[1].Language)
	ass.Equal("Appareil d'enregistrement de signaux numériques", patDoc.Title[2].Text)
	ass.Equal("fr", patDoc.Title[2].Language)

	// abstract
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(875, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// claims Empty
	ass.NotEmpty(patDoc.Claims)
	ass.Equal(1, len(patDoc.Claims))
	ass.Equal(4188, len(patDoc.Claims[0].Text))
	ass.Equal("en", patDoc.Claims[0].Language)
	ass.Equal("claims01", patDoc.Claims[0].Id)

	// description
	ass.NotEmpty(patDoc.Description)
	ass.Equal(20979, len(patDoc.Description[0].Text))
	ass.Equal("en", patDoc.Description[0].Language)

	// citations
	ass.Empty(patDoc.Citations)

	// Inventors
	ass.NotEmpty(patDoc.Inventors)
	ass.Equal(3, len(patDoc.Inventors))
	ass.Equal(Country("JP"), patDoc.Inventors[0].Country)
	ass.Equal("Osaka-shi,\nOsaka-fu 534", patDoc.Inventors[0].City)
	ass.Equal("1-5-8-2804, Tomobuchi-cho,\nMiykojima-ku", patDoc.Inventors[0].Street)
	ass.Equal("Juri, Tatsuro", patDoc.Inventors[0].Name)

	ass.Equal(Country("JP"), patDoc.Inventors[1].Country)
	ass.Equal("Higashi-osaka-shi,\nOsaka-fu 577", patDoc.Inventors[1].City)
	ass.Equal("2-4-28, Nishitsutsumi hondori higashi", patDoc.Inventors[1].Street)
	ass.Equal("Iketani, Akira", patDoc.Inventors[1].Name)

	ass.Equal(Country("JP"), patDoc.Inventors[2].Country)
	ass.Equal("Suita-shi,\nOsaka-fu 565", patDoc.Inventors[2].City)
	ass.Equal("3-1- A3-202, Tsugumodai", patDoc.Inventors[2].Street)
	ass.Equal("Matsumi, Chiyoko", patDoc.Inventors[2].Name)

	// representatives
	ass.NotEmpty(patDoc.Representatives)
	ass.Equal(Country("DE"), patDoc.Representatives[0].Country)
	ass.Equal("00100151", patDoc.Representatives[0].IID)
	ass.Equal("28195 Bremen", patDoc.Representatives[0].City)
	ass.Equal("Martinistrasse 24", patDoc.Representatives[0].Street)
	ass.Equal("Eisenführ, Speiser & Partner", patDoc.Representatives[0].Name)

	// contracting states
	ass.NotEmpty(patDoc.ContractingStates)
	ass.Equal(4, len(patDoc.ContractingStates))
	for i := 0; i <= 3; i++ {
		ass.Equal(2, len(patDoc.ContractingStates[i]))
	}

	// classifications
	ass.Empty(patDoc.Classifications)
}

func TestProcessXMLSimple11B1(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("test-data/grant/v1-1-B1.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("EP98948476B1", patDoc.ID)
	ass.Equal("EP98948476NWB1.xml", patDoc.File)
	ass.Equal("en", patDoc.Lang)
	ass.Equal(Country("EP"), patDoc.Country)
	ass.Equal("1019261", patDoc.DocNumber)
	ass.Equal("B1", patDoc.Kind)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20030709", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("n", patDoc.Status)
	ass.Equal("ep-patent-document-v1-1", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("INTELLIGENTES BREMSSYSTEM FÜR MATERIALHANDHABUNGSFAHRZEUGE", patDoc.Title[0].Text)
	ass.Equal("de", patDoc.Title[0].Language)
	ass.Equal("INTELLIGENT BRAKING SYSTEM FOR MATERIALS HANDLING VEHICLES", patDoc.Title[1].Text)
	ass.Equal("en", patDoc.Title[1].Language)
	ass.Equal("SYSTEME DE FREINAGE INTELLIGENT POUR VEHICULES DE MANUTENTION DE MATIERES", patDoc.Title[2].Text)
	ass.Equal("fr", patDoc.Title[2].Language)

	// abstract
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(0, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// claims
	ass.NotEmpty(patDoc.Claims)
	ass.Equal(3, len(patDoc.Claims))
	ass.Equal(10506, len(patDoc.Claims[0].Text))
	ass.Equal("en", patDoc.Claims[0].Language)
	ass.Equal("claims01", patDoc.Claims[0].Id)
	ass.Equal(11253, len(patDoc.Claims[1].Text))
	ass.Equal("de", patDoc.Claims[1].Language)
	ass.Equal("claims02", patDoc.Claims[1].Id)
	ass.Equal(13369, len(patDoc.Claims[2].Text))
	ass.Equal("fr", patDoc.Claims[2].Language)
	ass.Equal("claims03", patDoc.Claims[2].Id)

	// description
	ass.NotEmpty(patDoc.Description)
	ass.Equal(35365, len(patDoc.Description[0].Text))
	ass.Equal("en", patDoc.Description[0].Language)

	// citations
	ass.Empty(patDoc.Citations)

	// Inventors
	ass.NotEmpty(patDoc.Inventors)
	ass.Equal(Country("US"), patDoc.Inventors[0].Country)
	ass.Equal("Minster, OH 45865", patDoc.Inventors[0].City)
	ass.Equal("276 South Frankfort Street", patDoc.Inventors[0].Street)
	ass.Equal("SHERMAN, Nicholas, J.", patDoc.Inventors[0].Name)

	// owners
	ass.NotEmpty(patDoc.Owners)
	ass.Equal(Country("US"), patDoc.Owners[0].Country)
	ass.Equal("01416510", patDoc.Owners[0].IID)
	ass.Equal("SPC/P6132EP", patDoc.Owners[0].IRF)
	ass.Equal("New Bremen,\nOhio 45869", patDoc.Owners[0].City)
	ass.Equal("40-44 South Washington Street", patDoc.Owners[0].Street)
	ass.Equal("CROWN EQUIPMENT CORPORATION", patDoc.Owners[0].Name)

	// representatives
	ass.NotEmpty(patDoc.Representatives)
	ass.Equal(Country("GB"), patDoc.Representatives[0].Country)
	ass.Equal("00072881", patDoc.Representatives[0].IID)
	ass.Equal("London EC4M 7ET", patDoc.Representatives[0].City)
	ass.Equal("David Keltie Associates\nFleet Place House\n2 Fleet Place", patDoc.Representatives[0].Street)
	ass.Equal("Cummings, Sean Patrick", patDoc.Representatives[0].Name)

	// contracting states
	ass.NotEmpty(patDoc.ContractingStates)
	ass.Equal(19, len(patDoc.ContractingStates))
	for i := 0; i <= 18; i++ {
		ass.Equal(2, len(patDoc.ContractingStates[i]))
	}

	// classifications
	ass.Empty(patDoc.Classifications)
}

func TestProcessXMLSimple11B11(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("test-data/grant/v1-1-B1-1.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("EP84305989B1", patDoc.ID)
	ass.Equal("EP84305989NWB1.xml", patDoc.File)
	ass.Equal("en", patDoc.Lang)
	ass.Equal(Country("EP"), patDoc.Country)
	ass.Equal("0138357", patDoc.DocNumber)
	ass.Equal("B1", patDoc.Kind)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("19941221", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("n", patDoc.Status)
	ass.Equal("ep-patent-document-v1-1", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("Markiertes DNA", patDoc.Title[0].Text)
	ass.Equal("de", patDoc.Title[0].Language)
	ass.Equal("Labeled DNA", patDoc.Title[1].Text)
	ass.Equal("en", patDoc.Title[1].Language)
	ass.Equal("DNA marqué", patDoc.Title[2].Text)
	ass.Equal("fr", patDoc.Title[2].Language)

	// abstract
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(0, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// claims
	ass.NotEmpty(patDoc.Claims)
	ass.Equal(3, len(patDoc.Claims))
	ass.Equal(8953, len(patDoc.Claims[0].Text))
	ass.Equal("en", patDoc.Claims[0].Language)
	ass.Equal("claims01", patDoc.Claims[0].Id)
	ass.Equal(10100, len(patDoc.Claims[1].Text))
	ass.Equal("de", patDoc.Claims[1].Language)
	ass.Equal("claims02", patDoc.Claims[1].Id)
	ass.Equal(10631, len(patDoc.Claims[2].Text))
	ass.Equal("fr", patDoc.Claims[2].Language)
	ass.Equal("claims03", patDoc.Claims[2].Id)

	// description
	ass.NotEmpty(patDoc.Description)
	ass.Equal(24201, len(patDoc.Description[0].Text))
	ass.Equal("en", patDoc.Description[0].Language)

	// citations
	ass.Empty(patDoc.Citations)

	// Inventors
	ass.NotEmpty(patDoc.Inventors)
	ass.Equal(Country("US"), patDoc.Inventors[0].Country)
	ass.Equal("Natick\nMassachusetts", patDoc.Inventors[0].City)
	ass.Equal("45 Village Brook Lane No 7", patDoc.Inventors[0].Street)
	ass.Equal("Landes, Greg M.", patDoc.Inventors[0].Name)

	// owners
	ass.NotEmpty(patDoc.Owners)
	ass.Equal(Country("US"), patDoc.Owners[0].Country)
	ass.Equal("00697103", patDoc.Owners[0].IID)
	ass.Equal("", patDoc.Owners[0].IRF)
	ass.Equal("Cambridge,\nMassachusetts 02139-1562", patDoc.Owners[0].City)
	ass.Equal("One Kendall Square", patDoc.Owners[0].Street)
	ass.Equal("GENZYME CORPORATION", patDoc.Owners[0].Name)

	// representatives
	ass.NotEmpty(patDoc.Representatives)
	ass.Equal(Country("GB"), patDoc.Representatives[0].Country)
	ass.Equal("00030021", patDoc.Representatives[0].IID)
	ass.Equal("London WC1A 1LW", patDoc.Representatives[0].City)
	ass.Equal("Lloyd Wise, Tregear & Co.,\nCommonwealth House,\n1-19 New Oxford Street", patDoc.Representatives[0].Street)
	ass.Equal("Deans, Michael John Percy", patDoc.Representatives[0].Name)

	// contracting states
	ass.NotEmpty(patDoc.ContractingStates)
	ass.Equal(11, len(patDoc.ContractingStates))
	for i := 0; i <= 10; i++ {
		ass.Equal(2, len(patDoc.ContractingStates[i]))
	}

	// classifications
	ass.Empty(patDoc.Classifications)
}

func TestProcessXMLSimple11B2(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("test-data/grant/v1-1-B2.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("EP91121229B2", patDoc.ID)
	ass.Equal("EP91121229NWB2.xml", patDoc.File)
	ass.Equal("en", patDoc.Lang)
	ass.Equal(Country("EP"), patDoc.Country)
	ass.Equal("0546210", patDoc.DocNumber)
	ass.Equal("B2", patDoc.Kind)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20030709", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("n", patDoc.Status)
	ass.Equal("ep-patent-document-v1-1", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("Verfahren und Vorrichtung zur Wärmebehandlung von Gussstücken", patDoc.Title[0].Text)
	ass.Equal("de", patDoc.Title[0].Language)
	ass.Equal("Method and apparatus for heat treating metal castings", patDoc.Title[1].Text)
	ass.Equal("en", patDoc.Title[1].Language)
	ass.Equal("Procédé et dispositif pour le traitement thermique de pièces coulées", patDoc.Title[2].Text)
	ass.Equal("fr", patDoc.Title[2].Language)

	// abstract
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(0, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// claims
	ass.NotEmpty(patDoc.Claims)
	ass.Equal(3, len(patDoc.Claims))
	ass.Equal(9576, len(patDoc.Claims[0].Text))
	ass.Equal("en", patDoc.Claims[0].Language)
	ass.Equal("claims01", patDoc.Claims[0].Id)
	ass.Equal(10999, len(patDoc.Claims[1].Text))
	ass.Equal("de", patDoc.Claims[1].Language)
	ass.Equal("claims02", patDoc.Claims[1].Id)
	ass.Equal(11191, len(patDoc.Claims[2].Text))
	ass.Equal("fr", patDoc.Claims[2].Language)
	ass.Equal("claims03", patDoc.Claims[2].Id)

	// description
	ass.NotEmpty(patDoc.Description)
	ass.Equal(51465, len(patDoc.Description[0].Text))
	ass.Equal("en", patDoc.Description[0].Language)

	// citations
	ass.Empty(patDoc.Citations)

	// Inventors
	ass.NotEmpty(patDoc.Inventors)
	ass.Equal(Country("US"), patDoc.Inventors[0].Country)
	ass.Equal("Kennesaw,\nGeorgia 30144", patDoc.Inventors[0].City)
	ass.Equal("743 Edgewater Lane", patDoc.Inventors[0].Street)
	ass.Equal("Crafton, Scott P.", patDoc.Inventors[0].Name)

	ass.Equal(Country("US"), patDoc.Inventors[1].Country)
	ass.Equal("Marietta,\nGeorgia 30062", patDoc.Inventors[1].City)
	ass.Equal("1976 Kramer Way", patDoc.Inventors[1].Street)
	ass.Equal("Crafton, Paul M.", patDoc.Inventors[1].Name)

	// owners
	ass.NotEmpty(patDoc.Owners)
	ass.Equal(Country("US"), patDoc.Owners[0].Country)
	ass.Equal("01442570", patDoc.Owners[0].IID)
	ass.Equal("23111/jst:si", patDoc.Owners[0].IRF)
	ass.Equal("Kennesaw,\nGeorgia 30144", patDoc.Owners[0].City)
	ass.Equal("2871 McCollum Parkway, N.W.", patDoc.Owners[0].Street)
	ass.Equal("CONSOLIDATED ENGINEERING COMPANY, INC.", patDoc.Owners[0].Name)

	// representatives
	ass.NotEmpty(patDoc.Representatives)
	ass.Equal(Country("DE"), patDoc.Representatives[0].Country)
	ass.Equal("00088023", patDoc.Representatives[0].IID)
	ass.Equal("81541 München", patDoc.Representatives[0].City)
	ass.Equal("Hofstetter, Schurack & Skora\nPatentanwälte\nBalanstrasse 57", patDoc.Representatives[0].Street)
	ass.Equal("Schurack, Eduard F.", patDoc.Representatives[0].Name)

	// contracting states
	ass.NotEmpty(patDoc.ContractingStates)
	ass.Equal(8, len(patDoc.ContractingStates))
	for i := 0; i <= 7; i++ {
		ass.Equal(2, len(patDoc.ContractingStates[i]))
	}

	// classifications
	ass.Empty(patDoc.Classifications)
}

func TestProcessXMLSimple11B21(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("test-data/grant/v1-1-B2-1.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("EP85106915B2", patDoc.ID)
	ass.Equal("EP85106915NWB2.xml", patDoc.File)
	ass.Equal("en", patDoc.Lang)
	ass.Equal(Country("EP"), patDoc.Country)
	ass.Equal("0166996", patDoc.DocNumber)
	ass.Equal("B2", patDoc.Kind)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("19941221", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("n", patDoc.Status)
	ass.Equal("ep-patent-document-v1-1", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("Verfahren zur Stabilisierung einer physiologisch aktiven Substanz, die durch rekombinante DNA-Technik hergestellt wurde und eine stabile wässrige Lösung oder ein Pulver, die diese Substanz enthält", patDoc.Title[0].Text)
	ass.Equal("de", patDoc.Title[0].Language)
	ass.Equal("A method for stabilizing a physiologically active substance produced by recombinant DNA technique and a stable aqueous solution or powder containing the same", patDoc.Title[1].Text)
	ass.Equal("en", patDoc.Title[1].Language)
	ass.Equal("Méthode pour la stabilisation d'une substance physiologiquement active, préparée par technique d'ADN recombinant et une solution aqueuse stable ou poudre contenant la dite substance", patDoc.Title[2].Text)
	ass.Equal("fr", patDoc.Title[2].Language)

	// abstract
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(0, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// claims
	ass.NotEmpty(patDoc.Claims)
	ass.Equal(3, len(patDoc.Claims))
	ass.Equal(5340, len(patDoc.Claims[0].Text))
	ass.Equal("en", patDoc.Claims[0].Language)
	ass.Equal("claims01", patDoc.Claims[0].Id)
	ass.Equal(5469, len(patDoc.Claims[1].Text))
	ass.Equal("de", patDoc.Claims[1].Language)
	ass.Equal("claims02", patDoc.Claims[1].Id)
	ass.Equal(5887, len(patDoc.Claims[2].Text))
	ass.Equal("fr", patDoc.Claims[2].Language)
	ass.Equal("claims03", patDoc.Claims[2].Id)

	// description
	ass.NotEmpty(patDoc.Description)
	ass.Equal(96152, len(patDoc.Description[0].Text))
	ass.Equal("en", patDoc.Description[0].Language)

	// citations
	ass.Empty(patDoc.Citations)

	// Inventors
	ass.NotEmpty(patDoc.Inventors)
	ass.Equal(Country("JP"), patDoc.Inventors[0].Country)
	ass.Equal("Fuji-shi\nShizuoka-ken", patDoc.Inventors[0].City)
	ass.Equal("3984-19 Imaizumi", patDoc.Inventors[0].Street)
	ass.Equal("Sakamoto, Hajimu", patDoc.Inventors[0].Name)

	ass.Equal(Country("JP"), patDoc.Inventors[1].Country)
	ass.Equal("Fuji-shi\nShizuoka-ken", patDoc.Inventors[1].City)
	ass.Equal("125-8 Mitojima", patDoc.Inventors[1].Street)
	ass.Equal("Kiyota, Takao", patDoc.Inventors[1].Name)

	// owners
	ass.NotEmpty(patDoc.Owners)
	ass.Equal(Country("JP"), patDoc.Owners[0].Country)
	ass.Equal("00219570", patDoc.Owners[0].IID)
	ass.Equal("", patDoc.Owners[0].IRF)
	ass.Equal("Osaka-shi\nOsaka 530", patDoc.Owners[0].City)
	ass.Equal("2-6, Dojimahama 1-chome\nKita-ku", patDoc.Owners[0].Street)
	ass.Equal("Asahi Kasei Kogyo Kabushiki Kaisha", patDoc.Owners[0].Name)

	// representatives
	ass.NotEmpty(patDoc.Representatives)
	ass.Equal(Country("DE"), patDoc.Representatives[0].Country)
	ass.Equal("00002193", patDoc.Representatives[0].IID)
	ass.Equal("81541 München", patDoc.Representatives[0].City)
	ass.Equal("Patentanwälte Boeters & Bauer,\nBereiteranger 15", patDoc.Representatives[0].Street)
	ass.Equal("Boeters, Hans Dietrich, Dr.", patDoc.Representatives[0].Name)

	// contracting states
	ass.NotEmpty(patDoc.ContractingStates)
	ass.Equal(9, len(patDoc.ContractingStates))
	for i := 0; i <= 8; i++ {
		ass.Equal(2, len(patDoc.ContractingStates[i]))
	}

	// classifications
	ass.Empty(patDoc.Classifications)
}

//v 1.2
func TestProcessXMLSimple12A1(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("test-data/application/v1-2-A1.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("EP06381048A1", patDoc.ID)
	ass.Equal("EP06381048NWA1.xml", patDoc.File)
	ass.Equal("en", patDoc.Lang)
	ass.Equal(Country("EP"), patDoc.Country)
	ass.Equal("1921219", patDoc.DocNumber)
	ass.Equal("A1", patDoc.Kind)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20080514", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("n", patDoc.Status)
	ass.Equal("ep-patent-document-v1-2", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("Schalldämpfungs- und Brandschutzisolierplatte und Verfahren zu ihrer Herstellung", patDoc.Title[0].Text)
	ass.Equal("de", patDoc.Title[0].Language)
	ass.Equal("Sound damping and fire restistant insulating panel and its manufacturing procedure", patDoc.Title[1].Text)
	ass.Equal("en", patDoc.Title[1].Language)
	ass.Equal("Panneau isolant insonorisant et résistant au feu et son procédé de fabrication", patDoc.Title[2].Text)
	ass.Equal("fr", patDoc.Title[2].Language)

	// abstract
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(644, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// claims
	ass.NotEmpty(patDoc.Claims)
	ass.Equal(1, len(patDoc.Claims))
	ass.Equal(4340, len(patDoc.Claims[0].Text))
	ass.Equal("en", patDoc.Claims[0].Language)
	ass.Equal("claims01", patDoc.Claims[0].Id)

	// description
	ass.NotEmpty(patDoc.Description)
	ass.Equal(11388, len(patDoc.Description[0].Text))
	ass.Equal("en", patDoc.Description[0].Language)

	// citations
	ass.NotEmpty(patDoc.Citations)
	ass.Equal(Country("ES"), patDoc.Citations[0].Country)
	ass.Equal("2115900", patDoc.Citations[0].DocNumber)
	ass.Equal("", patDoc.Citations[0].Kind)

	// Inventors
	ass.NotEmpty(patDoc.Inventors)
	ass.Equal(1, len(patDoc.Inventors))
	ass.Equal(Country("ES"), patDoc.Inventors[0].Country)
	ass.Equal("Barcelona", patDoc.Inventors[0].City)
	ass.Equal("C/Industria nau 4 Pol. Ind. Pla de la Costa 08182 Sant Feliu de Codines", patDoc.Inventors[0].Street)
	ass.Equal("Lloveras Calvo, Juan", patDoc.Inventors[0].Name)

	// representatives
	ass.NotEmpty(patDoc.Representatives)
	ass.Equal(1, len(patDoc.Representatives))
	ass.Equal("28040 Madrid", patDoc.Representatives[0].City)
	ass.Equal("Esteban Perez-Serrano, Maria Isabel", patDoc.Representatives[0].Name)
	ass.Equal("00158882", patDoc.Representatives[0].IID)
	ass.Equal(Country("ES"), patDoc.Representatives[0].Country)
	ass.Equal("UDAPI & ASOCIADOS Explanada, 8", patDoc.Representatives[0].Street)

	// contracting states
	ass.NotEmpty(patDoc.ContractingStates)
	ass.Equal(31, len(patDoc.ContractingStates))
	for i := 0; i <= 30; i++ {
		ass.Equal(2, len(patDoc.ContractingStates[i]))
	}

	// classifications
	ass.NotEmpty(patDoc.Classifications)
	ass.Equal(3, len(patDoc.Classifications))

	for i := 0; i <= 2; i++ {
		ass.Equal(IPC, patDoc.Classifications[i].System)
		ass.Equal(i+1, patDoc.Classifications[i].Sequence)
		ass.Equal("E", patDoc.Classifications[i].Section)
		ass.Equal("04", patDoc.Classifications[i].Class)
		ass.Equal("B", patDoc.Classifications[i].SubClass)
		ass.Equal("1", patDoc.Classifications[i].MainGroup)
		ass.Equal("20060101", patDoc.Classifications[i].Version)
		ass.Equal("A", patDoc.Classifications[i].ClassificationLevel)
		ass.Equal("20070316", patDoc.Classifications[i].ActionDate)
		ass.Equal("B", patDoc.Classifications[i].OriginalOrReclassified)
		ass.Equal("H", patDoc.Classifications[i].Source)
		ass.Equal("EP", patDoc.Classifications[i].GeneratingOffice)
	}

	ass.Equal("E04B 1/88 20060101AFI20070316BHEP ", patDoc.Classifications[0].Text)
	ass.Equal("88", patDoc.Classifications[0].SubGroup)
	ass.Equal("F", patDoc.Classifications[0].FirstLater)
	ass.Equal("I", patDoc.Classifications[0].ClassificationValue)

	ass.Equal("E04B 1/90 20060101ALI20070316BHEP ", patDoc.Classifications[1].Text)
	ass.Equal("90", patDoc.Classifications[1].SubGroup)
	ass.Equal("L", patDoc.Classifications[1].FirstLater)
	ass.Equal("I", patDoc.Classifications[1].ClassificationValue)

	ass.Equal("E04B 1/84 20060101ALN20070316BHEP ", patDoc.Classifications[2].Text)
	ass.Equal("84", patDoc.Classifications[2].SubGroup)
	ass.Equal("N", patDoc.Classifications[2].ClassificationValue)
	ass.Equal("L", patDoc.Classifications[2].FirstLater)
}

func TestProcessXMLSimple12A11(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("test-data/application/v1-2-A1-1.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("EP07012401A1", patDoc.ID)
	ass.Equal("EP07012401NWA1.xml", patDoc.File)
	ass.Equal("en", patDoc.Lang)
	ass.Equal(Country("EP"), patDoc.Country)
	ass.Equal("1874102", patDoc.DocNumber)
	ass.Equal("A1", patDoc.Kind)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20080102", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("n", patDoc.Status)
	ass.Equal("ep-patent-document-v1-2", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("Leiterplatte mit eingebautem Kondensator", patDoc.Title[0].Text)
	ass.Equal("de", patDoc.Title[0].Language)
	ass.Equal("Wiring board with built-in capacitor", patDoc.Title[1].Text)
	ass.Equal("en", patDoc.Title[1].Language)
	ass.Equal("Carte de circuit avec condensateur intégré", patDoc.Title[2].Text)
	ass.Equal("fr", patDoc.Title[2].Language)

	// abstract
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(955, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// claims
	ass.NotEmpty(patDoc.Claims)
	ass.Equal(1, len(patDoc.Claims))
	ass.Equal(2728, len(patDoc.Claims[0].Text))
	ass.Equal("en", patDoc.Claims[0].Language)
	ass.Equal("claims01", patDoc.Claims[0].Id)

	// description
	ass.NotEmpty(patDoc.Description)
	ass.Equal(42158, len(patDoc.Description[0].Text))
	ass.Equal("en", patDoc.Description[0].Language)

	// citations
	ass.NotEmpty(patDoc.Citations)
	ass.Equal(1, len(patDoc.Citations))
	ass.Equal(Country("JP"), patDoc.Citations[0].Country)
	ass.Equal("2005191266", patDoc.Citations[0].DocNumber)
	ass.Equal("A", patDoc.Citations[0].Kind)

	// Inventors
	ass.NotEmpty(patDoc.Inventors)
	ass.Equal(1, len(patDoc.Inventors))
	ass.Equal(Country("JP"), patDoc.Inventors[0].Country)
	ass.Equal("Ibi-gun\nGifu 501-0695", patDoc.Inventors[0].City)
	ass.Equal("Ibiden Co., Ltd.\nOgki Kita Plant\n1-1, Kitagata\nIbigawa-cho", patDoc.Inventors[0].Street)
	ass.Equal("Tanaka, Hironori", patDoc.Inventors[0].Name)

	// representatives
	ass.NotEmpty(patDoc.Representatives)
	ass.Equal(1, len(patDoc.Representatives))
	ass.Equal("80538 München", patDoc.Representatives[0].City)
	ass.Equal("Grünecker, Kinkeldey, \nStockmair & Schwanhäusser \nAnwaltssozietät", patDoc.Representatives[0].Name)
	ass.Equal("00100721", patDoc.Representatives[0].IID)
	ass.Equal(Country("DE"), patDoc.Representatives[0].Country)
	ass.Equal("Maximilianstrasse 58", patDoc.Representatives[0].Street)

	// contracting states
	ass.NotEmpty(patDoc.ContractingStates)
	ass.Equal(32, len(patDoc.ContractingStates))
	for i := 0; i <= 31; i++ {
		ass.Equal(2, len(patDoc.ContractingStates[i]))
	}

	// classifications
	ass.NotEmpty(patDoc.Classifications)
	ass.Equal(2, len(patDoc.Classifications))

	ass.Equal("H05K   1/16        20060101AFI20070924BHEP        ", patDoc.Classifications[0].Text)
	ass.Equal(IPC, patDoc.Classifications[0].System)
	ass.Equal(1, patDoc.Classifications[0].Sequence)
	ass.Equal("H", patDoc.Classifications[0].Section)
	ass.Equal("05", patDoc.Classifications[0].Class)
	ass.Equal("K", patDoc.Classifications[0].SubClass)
	ass.Equal("1", patDoc.Classifications[0].MainGroup)
	ass.Equal("16", patDoc.Classifications[0].SubGroup)
	ass.Equal("20060101", patDoc.Classifications[0].Version)
	ass.Equal("A", patDoc.Classifications[0].ClassificationLevel)
	ass.Equal("F", patDoc.Classifications[0].FirstLater)
	ass.Equal("I", patDoc.Classifications[0].ClassificationValue)
	ass.Equal("20070924", patDoc.Classifications[0].ActionDate)
	ass.Equal("B", patDoc.Classifications[0].OriginalOrReclassified)
	ass.Equal("H", patDoc.Classifications[0].Source)
	ass.Equal("EP", patDoc.Classifications[0].GeneratingOffice)

	ass.Equal("H01L  23/50        20060101ALI20070924BHEP        ", patDoc.Classifications[1].Text)
	ass.Equal(IPC, patDoc.Classifications[1].System)
	ass.Equal(2, patDoc.Classifications[1].Sequence)
	ass.Equal("H", patDoc.Classifications[1].Section)
	ass.Equal("01", patDoc.Classifications[1].Class)
	ass.Equal("L", patDoc.Classifications[1].SubClass)
	ass.Equal("23", patDoc.Classifications[1].MainGroup)
	ass.Equal("50", patDoc.Classifications[1].SubGroup)
	ass.Equal("20060101", patDoc.Classifications[1].Version)
	ass.Equal("A", patDoc.Classifications[1].ClassificationLevel)
	ass.Equal("L", patDoc.Classifications[1].FirstLater)
	ass.Equal("I", patDoc.Classifications[1].ClassificationValue)
	ass.Equal("20070924", patDoc.Classifications[1].ActionDate)
	ass.Equal("B", patDoc.Classifications[1].OriginalOrReclassified)
	ass.Equal("H", patDoc.Classifications[1].Source)
	ass.Equal("EP", patDoc.Classifications[1].GeneratingOffice)

}

func TestProcessXMLSimple12A2(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("test-data/application/v1-2-A2.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("EP07012363A2", patDoc.ID)
	ass.Equal("EP07012363NWA2.xml", patDoc.File)
	ass.Equal("de", patDoc.Lang)
	ass.Equal(Country("EP"), patDoc.Country)
	ass.Equal("1873405", patDoc.DocNumber)
	ass.Equal("A2", patDoc.Kind)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20080102", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("n", patDoc.Status)
	ass.Equal("ep-patent-document-v1-2", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("Spanplattenschraube", patDoc.Title[0].Text)
	ass.Equal("de", patDoc.Title[0].Language)
	ass.Equal("Chipboard screw", patDoc.Title[1].Text)
	ass.Equal("en", patDoc.Title[1].Language)
	ass.Equal("Vis pour panneau aggloméré", patDoc.Title[2].Text)
	ass.Equal("fr", patDoc.Title[2].Language)

	// abstract
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(848, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// claims
	ass.NotEmpty(patDoc.Claims)
	ass.Equal(3018, len(patDoc.Claims[0].Text))
	ass.Equal("de", patDoc.Claims[0].Language)
	// description
	ass.NotEmpty(patDoc.Description)
	ass.Equal(15085, len(patDoc.Description[0].Text))
	ass.Equal("en", patDoc.Description[0].Language)

	// citations
	ass.NotEmpty(patDoc.Citations)
	ass.Equal(Country("DE"), patDoc.Citations[0].Country)
	ass.Equal("9216791", patDoc.Citations[0].DocNumber)
	ass.Equal("", patDoc.Citations[0].Kind)

	// Inventors
	ass.NotEmpty(patDoc.Inventors)
	ass.Equal(1, len(patDoc.Inventors))
	ass.Equal(Country("DE"), patDoc.Inventors[0].Country)
	ass.Equal("74632 Neuenstein", patDoc.Inventors[0].City)
	ass.Equal("Haller Strasse 19", patDoc.Inventors[0].Street)
	ass.Equal("Heim, Markus", patDoc.Inventors[0].Name)

	// representatives
	ass.NotEmpty(patDoc.Representatives)
	ass.Equal(Country("DE"), patDoc.Representatives[0].Country)
	ass.Equal("00100169", patDoc.Representatives[0].IID)
	ass.Equal("70174 Stuttgart", patDoc.Representatives[0].City)
	ass.Equal("Kronenstrasse 30", patDoc.Representatives[0].Street)
	ass.Equal("Patentanwälte \nRuff, Wilhelm, Beier, Dauster & Partner", patDoc.Representatives[0].Name)

	// contracting states
	ass.NotEmpty(patDoc.ContractingStates)
	ass.Equal(32, len(patDoc.ContractingStates))
	for i := 0; i <= 31; i++ {
		ass.Equal(2, len(patDoc.ContractingStates[i]))
	}

	// classifications
	ass.NotEmpty(patDoc.Classifications)
	ass.Equal(2, len(patDoc.Classifications))

	ass.Equal("F16B  25/10        20060101AFI20071009BHEP        ", patDoc.Classifications[0].Text)
	ass.Equal(IPC, patDoc.Classifications[0].System)
	ass.Equal(1, patDoc.Classifications[0].Sequence)
	ass.Equal("F", patDoc.Classifications[0].Section)
	ass.Equal("16", patDoc.Classifications[0].Class)
	ass.Equal("B", patDoc.Classifications[0].SubClass)
	ass.Equal("25", patDoc.Classifications[0].MainGroup)
	ass.Equal("10", patDoc.Classifications[0].SubGroup)
	ass.Equal("20060101", patDoc.Classifications[0].Version)
	ass.Equal("A", patDoc.Classifications[0].ClassificationLevel)
	ass.Equal("F", patDoc.Classifications[0].FirstLater)
	ass.Equal("I", patDoc.Classifications[0].ClassificationValue)
	ass.Equal("20071009", patDoc.Classifications[0].ActionDate)
	ass.Equal("B", patDoc.Classifications[0].OriginalOrReclassified)
	ass.Equal("H", patDoc.Classifications[0].Source)
	ass.Equal("EP", patDoc.Classifications[0].GeneratingOffice)

	ass.Equal("F16B  35/06        20060101ALN20071009BHEP        ", patDoc.Classifications[1].Text)
	ass.Equal(IPC, patDoc.Classifications[1].System)
	ass.Equal(2, patDoc.Classifications[1].Sequence)
	ass.Equal("F", patDoc.Classifications[1].Section)
	ass.Equal("16", patDoc.Classifications[1].Class)
	ass.Equal("B", patDoc.Classifications[1].SubClass)
	ass.Equal("35", patDoc.Classifications[1].MainGroup)
	ass.Equal("06", patDoc.Classifications[1].SubGroup)
	ass.Equal("20060101", patDoc.Classifications[1].Version)
	ass.Equal("A", patDoc.Classifications[1].ClassificationLevel)
	ass.Equal("L", patDoc.Classifications[1].FirstLater)
	ass.Equal("N", patDoc.Classifications[1].ClassificationValue)
	ass.Equal("20071009", patDoc.Classifications[1].ActionDate)
	ass.Equal("B", patDoc.Classifications[1].OriginalOrReclassified)
	ass.Equal("H", patDoc.Classifications[1].Source)
	ass.Equal("EP", patDoc.Classifications[1].GeneratingOffice)
}

func TestProcessXMLSimple12A21(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("test-data/application/v1-2-A2-1.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("EP07012827A2", patDoc.ID)
	ass.Equal("EP07012827NWA2.xml", patDoc.File)
	ass.Equal("en", patDoc.Lang)
	ass.Equal(Country("EP"), patDoc.Country)
	ass.Equal("1874103", patDoc.DocNumber)
	ass.Equal("A2", patDoc.Kind)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20080102", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("n", patDoc.Status)
	ass.Equal("ep-patent-document-v1-2", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("Motorsteuerung", patDoc.Title[0].Text)
	ass.Equal("de", patDoc.Title[0].Language)
	ass.Equal("Engine controller", patDoc.Title[1].Text)
	ass.Equal("en", patDoc.Title[1].Language)
	ass.Equal("Contrôleur de moteur", patDoc.Title[2].Text)
	ass.Equal("fr", patDoc.Title[2].Language)

	// abstract
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(989, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// claims
	ass.NotEmpty(patDoc.Claims)
	ass.Equal(3779, len(patDoc.Claims[0].Text))
	ass.Equal("en", patDoc.Claims[0].Language)
	// description
	ass.NotEmpty(patDoc.Description)
	ass.Equal(22521, len(patDoc.Description[0].Text))
	ass.Equal("en", patDoc.Description[0].Language)

	// citations
	ass.NotEmpty(patDoc.Citations)
	ass.Equal(Country("JP"), patDoc.Citations[0].Country)
	ass.Equal("2001085858", patDoc.Citations[0].DocNumber)
	ass.Equal("A", patDoc.Citations[0].Kind)

	// Inventors
	ass.NotEmpty(patDoc.Inventors)
	ass.Equal(4, len(patDoc.Inventors))
	ass.Equal(Country("JP"), patDoc.Inventors[0].Country)
	ass.Equal("Hitachinaka-shi,\nIbaraki 312-0062", patDoc.Inventors[0].City)
	ass.Equal("c/o Hitachi Car Engineering Co., Ltd.,\n2477, Takaba", patDoc.Inventors[0].Street)
	ass.Equal("Kamoshida, Masaru", patDoc.Inventors[0].Name)
	ass.Equal(Country("JP"), patDoc.Inventors[1].Country)
	ass.Equal("Hitachinaka-shi,\nIbaraki 312-0062", patDoc.Inventors[1].City)
	ass.Equal("c/o Hitachi Car Engineering Co., Ltd.,\n2477, Takaba", patDoc.Inventors[1].Street)
	ass.Equal("Sasaki, Masahiro", patDoc.Inventors[1].Name)
	ass.Equal(Country("JP"), patDoc.Inventors[2].Country)
	ass.Equal("Chiyoda-ku,\nTokyo 100-8220", patDoc.Inventors[2].City)
	ass.Equal("c/o Hitachi, Ltd.,\n12th Floor, Marunouchi Center Bldg.\n6-1, Marunouchi 1-chome", patDoc.Inventors[2].Street)
	ass.Equal("Igarashi, Takeshi", patDoc.Inventors[2].Name)
	ass.Equal(Country("JP"), patDoc.Inventors[3].Country)
	ass.Equal("Hitachinaka-shi,\nIbaraki 312-0062", patDoc.Inventors[3].City)
	ass.Equal("c/o Hitachi Car Engineering Co., Ltd.,\n2477, Takaba", patDoc.Inventors[3].Street)
	ass.Equal("Asano, Masahiko", patDoc.Inventors[3].Name)

	// representatives
	ass.NotEmpty(patDoc.Representatives)
	ass.Equal(Country("DE"), patDoc.Representatives[0].Country)
	ass.Equal("00103261", patDoc.Representatives[0].IID)
	ass.Equal("80538 München", patDoc.Representatives[0].City)
	ass.Equal("Patentanwälte, \nSteinsdorfstrasse 10", patDoc.Representatives[0].Street)
	ass.Equal("Beetz & Partner", patDoc.Representatives[0].Name)

	// contracting states
	ass.NotEmpty(patDoc.ContractingStates)
	ass.Equal(32, len(patDoc.ContractingStates))
	for i := 0; i <= 31; i++ {
		ass.Equal(2, len(patDoc.ContractingStates[i]))
	}

	// classifications
	ass.NotEmpty(patDoc.Classifications)
	ass.Equal(2, len(patDoc.Classifications))

	ass.Equal("H05K   5/00        20060101AFI20071018BHEP        ", patDoc.Classifications[0].Text)
	ass.Equal(IPC, patDoc.Classifications[0].System)
	ass.Equal(1, patDoc.Classifications[0].Sequence)
	ass.Equal("H", patDoc.Classifications[0].Section)
	ass.Equal("05", patDoc.Classifications[0].Class)
	ass.Equal("K", patDoc.Classifications[0].SubClass)
	ass.Equal("5", patDoc.Classifications[0].MainGroup)
	ass.Equal("00", patDoc.Classifications[0].SubGroup)
	ass.Equal("20060101", patDoc.Classifications[0].Version)
	ass.Equal("A", patDoc.Classifications[0].ClassificationLevel)
	ass.Equal("F", patDoc.Classifications[0].FirstLater)
	ass.Equal("I", patDoc.Classifications[0].ClassificationValue)
	ass.Equal("20071018", patDoc.Classifications[0].ActionDate)
	ass.Equal("B", patDoc.Classifications[0].OriginalOrReclassified)
	ass.Equal("H", patDoc.Classifications[0].Source)
	ass.Equal("EP", patDoc.Classifications[0].GeneratingOffice)

	ass.Equal("H05K   5/06        20060101ALI20071018BHEP        ", patDoc.Classifications[1].Text)
	ass.Equal(IPC, patDoc.Classifications[1].System)
	ass.Equal(2, patDoc.Classifications[1].Sequence)
	ass.Equal("H", patDoc.Classifications[1].Section)
	ass.Equal("05", patDoc.Classifications[1].Class)
	ass.Equal("K", patDoc.Classifications[1].SubClass)
	ass.Equal("5", patDoc.Classifications[1].MainGroup)
	ass.Equal("06", patDoc.Classifications[1].SubGroup)
	ass.Equal("20060101", patDoc.Classifications[1].Version)
	ass.Equal("A", patDoc.Classifications[1].ClassificationLevel)
	ass.Equal("L", patDoc.Classifications[1].FirstLater)
	ass.Equal("I", patDoc.Classifications[1].ClassificationValue)
	ass.Equal("20071018", patDoc.Classifications[1].ActionDate)
	ass.Equal("B", patDoc.Classifications[1].OriginalOrReclassified)
	ass.Equal("H", patDoc.Classifications[1].Source)
	ass.Equal("EP", patDoc.Classifications[1].GeneratingOffice)
}

func TestProcessXMLSimple12B1(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("test-data/grant/v1-2-B1.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("EP92922494B1", patDoc.ID)
	ass.Equal("EP92922494NWB1.xml", patDoc.File)
	ass.Equal("en", patDoc.Lang)
	ass.Equal(Country("EP"), patDoc.Country)
	ass.Equal("0610335", patDoc.DocNumber)
	ass.Equal("B1", patDoc.Kind)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20080102", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("n", patDoc.Status)
	ass.Equal("ep-patent-document-v1-2", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("Verfahren zum Entwurf von Rekombitopen Peptiden", patDoc.Title[0].Text)
	ass.Equal("de", patDoc.Title[0].Language)
	ass.Equal("Methods for designing recombitope peptides", patDoc.Title[1].Text)
	ass.Equal("en", patDoc.Title[1].Language)
	ass.Equal("Méthodes pour le design de peptides d'epitope recombiné", patDoc.Title[2].Text)
	ass.Equal("fr", patDoc.Title[2].Language)

	// abstract
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(0, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// claims
	ass.NotEmpty(patDoc.Claims)
	ass.Equal(3, len(patDoc.Claims))
	ass.Equal(4319, len(patDoc.Claims[0].Text))
	ass.Equal("en", patDoc.Claims[0].Language)
	ass.Equal("claims01", patDoc.Claims[0].Id)
	ass.Equal(4688, len(patDoc.Claims[1].Text))
	ass.Equal("de", patDoc.Claims[1].Language)
	ass.Equal("claims02", patDoc.Claims[1].Id)
	ass.Equal(5191, len(patDoc.Claims[2].Text))
	ass.Equal("fr", patDoc.Claims[2].Language)
	ass.Equal("claims03", patDoc.Claims[2].Id)

	// description
	ass.NotEmpty(patDoc.Description)
	ass.Equal(99270, len(patDoc.Description[0].Text))
	ass.Equal("en", patDoc.Description[0].Language)

	// citations
	ass.NotEmpty(patDoc.Citations)
	ass.Equal(6, len(patDoc.Citations))
	ass.Equal(Country("WO"), patDoc.Citations[0].Country)
	ass.Equal("9106571", patDoc.Citations[0].DocNumber)
	ass.Equal("A", patDoc.Citations[0].Kind)
	ass.Equal(Country("US"), patDoc.Citations[1].Country)
	ass.Equal("4939239", patDoc.Citations[1].DocNumber)
	ass.Equal("A", patDoc.Citations[1].Kind)
	ass.Equal(Country("US"), patDoc.Citations[2].Country)
	ass.Equal("662276", patDoc.Citations[2].DocNumber)
	ass.Equal("A", patDoc.Citations[2].Kind)
	ass.Equal(Country("US"), patDoc.Citations[3].Country)
	ass.Equal("9208694", patDoc.Citations[3].DocNumber)
	ass.Equal("B", patDoc.Citations[3].Kind)
	ass.Equal(Country("US"), patDoc.Citations[4].Country)
	ass.Equal("777859", patDoc.Citations[4].DocNumber)
	ass.Equal("B", patDoc.Citations[4].Kind)
	ass.Equal(Country("US"), patDoc.Citations[5].Country)
	ass.Equal("807529", patDoc.Citations[5].DocNumber)
	ass.Equal("B", patDoc.Citations[5].Kind)

	// Inventors
	ass.NotEmpty(patDoc.Inventors)
	ass.Equal(Country("US"), patDoc.Inventors[0].Country)
	ass.Equal("Belmont, MA 02178", patDoc.Inventors[0].City)
	ass.Equal("54 Richardson Road", patDoc.Inventors[0].Street)
	ass.Equal("ROGERS, Bruce, L.", patDoc.Inventors[0].Name)

	// owners
	ass.NotEmpty(patDoc.Owners)
	ass.Equal(Country("DE"), patDoc.Owners[0].Country)
	ass.Equal("07257080", patDoc.Owners[0].IID)
	ass.Equal("VAP/JID/P56870P", patDoc.Owners[0].IRF)
	ass.Equal("64293 Darmstadt", patDoc.Owners[0].City)
	ass.Equal("Frankfurter Strasse 250", patDoc.Owners[0].Street)
	ass.Equal("Merck Patent GmbH", patDoc.Owners[0].Name)

	// representatives
	ass.Empty(patDoc.Representatives)

	// contracting states
	ass.NotEmpty(patDoc.ContractingStates)
	ass.Equal(16, len(patDoc.ContractingStates))
	for i := 0; i <= 15; i++ {
		ass.Equal(2, len(patDoc.ContractingStates[i]))
	}

	// classifications
	ass.NotEmpty(patDoc.Classifications)
	ass.Equal(4, len(patDoc.Classifications))
	for i := 0; i <= 3; i++ {
		ass.Equal(IPC, patDoc.Classifications[i].System)
		ass.Equal(i+1, patDoc.Classifications[i].Sequence)
		ass.Equal("20060101", patDoc.Classifications[i].Version)
		ass.Equal("20070823", patDoc.Classifications[i].ActionDate)
		ass.Equal("A", patDoc.Classifications[i].ClassificationLevel)
		ass.Equal("I", patDoc.Classifications[i].ClassificationValue)
		ass.Equal("B", patDoc.Classifications[i].OriginalOrReclassified)
		ass.Equal("H", patDoc.Classifications[i].Source)
		ass.Equal("EP", patDoc.Classifications[i].GeneratingOffice)
	}

	ass.Equal("C12N  15/62        20060101AFI20070823BHEP        ", patDoc.Classifications[0].Text)
	ass.Equal("C", patDoc.Classifications[0].Section)
	ass.Equal("12", patDoc.Classifications[0].Class)
	ass.Equal("N", patDoc.Classifications[0].SubClass)
	ass.Equal("15", patDoc.Classifications[0].MainGroup)
	ass.Equal("62", patDoc.Classifications[0].SubGroup)
	ass.Equal("F", patDoc.Classifications[0].FirstLater)

	ass.Equal("C12N  15/62        20060101AFI20070823BHEP        ", patDoc.Classifications[0].Text)
	ass.Equal("A", patDoc.Classifications[1].Section)
	ass.Equal("61", patDoc.Classifications[1].Class)
	ass.Equal("K", patDoc.Classifications[1].SubClass)
	ass.Equal("39", patDoc.Classifications[1].MainGroup)
	ass.Equal("35", patDoc.Classifications[1].SubGroup)
	ass.Equal("L", patDoc.Classifications[1].FirstLater)

	ass.Equal("G01N  33/53        20060101ALI20070823BHEP        ", patDoc.Classifications[2].Text)
	ass.Equal("G", patDoc.Classifications[2].Section)
	ass.Equal("01", patDoc.Classifications[2].Class)
	ass.Equal("N", patDoc.Classifications[2].SubClass)
	ass.Equal("33", patDoc.Classifications[2].MainGroup)
	ass.Equal("53", patDoc.Classifications[2].SubGroup)
	ass.Equal("L", patDoc.Classifications[2].FirstLater)

	ass.Equal("A61K  39/36        20060101ALI20070823BHEP        ", patDoc.Classifications[3].Text)
	ass.Equal("A", patDoc.Classifications[3].Section)
	ass.Equal("61", patDoc.Classifications[3].Class)
	ass.Equal("K", patDoc.Classifications[3].SubClass)
	ass.Equal("39", patDoc.Classifications[3].MainGroup)
	ass.Equal("36", patDoc.Classifications[3].SubGroup)
	ass.Equal("L", patDoc.Classifications[3].FirstLater)

}

func TestProcessXMLSimple12B11(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("test-data/grant/v1-2-B1-1.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("EP97201413B1", patDoc.ID)
	ass.Equal("EP97201413NWB1.xml", patDoc.File)
	ass.Equal("en", patDoc.Lang)
	ass.Equal(Country("EP"), patDoc.Country)
	ass.Equal("0809299", patDoc.DocNumber)
	ass.Equal("B1", patDoc.Kind)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20080423", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("n", patDoc.Status)
	ass.Equal("ep-patent-document-v1-2", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("Aktiver Pixelsensor mit Durchbruch-Rücksetzstruktur und Unterdrückung des Übersprechsignales", patDoc.Title[0].Text)
	ass.Equal("de", patDoc.Title[0].Language)
	ass.Equal("Active pixel sensor with punch-through reset and cross-talk suppression", patDoc.Title[1].Text)
	ass.Equal("en", patDoc.Title[1].Language)
	ass.Equal("Capteur à pixel actif comprenant un système de remise à niveau à effet de perçage ainsi qu'une suppression de diaphonie", patDoc.Title[2].Text)
	ass.Equal("fr", patDoc.Title[2].Language)

	// abstract
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(0, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// claims
	ass.NotEmpty(patDoc.Claims)
	ass.Equal(3, len(patDoc.Claims))
	ass.Equal(1272, len(patDoc.Claims[0].Text))
	ass.Equal("en", patDoc.Claims[0].Language)
	ass.Equal("claims01", patDoc.Claims[0].Id)
	ass.Equal(1368, len(patDoc.Claims[1].Text))
	ass.Equal("de", patDoc.Claims[1].Language)
	ass.Equal("claims02", patDoc.Claims[1].Id)
	ass.Equal(1542, len(patDoc.Claims[2].Text))
	ass.Equal("fr", patDoc.Claims[2].Language)
	ass.Equal("claims03", patDoc.Claims[2].Id)

	// description
	ass.NotEmpty(patDoc.Description)
	ass.Equal(15824, len(patDoc.Description[0].Text))
	ass.Equal("en", patDoc.Description[0].Language)

	// citations
	ass.NotEmpty(patDoc.Citations)
	ass.Equal(1, len(patDoc.Citations))
	ass.Equal(Country("US"), patDoc.Citations[0].Country)
	ass.Equal("5298778", patDoc.Citations[0].DocNumber)
	ass.Equal("A", patDoc.Citations[0].Kind)

	// Inventors
	ass.NotEmpty(patDoc.Inventors)
	ass.Equal(Country("US"), patDoc.Inventors[0].Country)
	ass.Equal("Rochester,\nNew York 14650-2201", patDoc.Inventors[0].City)
	ass.Equal("Eastman Kodak Company,\n343 State Street", patDoc.Inventors[0].Street)
	ass.Equal("Guidash, Robert Michael", patDoc.Inventors[0].Name)

	// owners
	ass.NotEmpty(patDoc.Owners)
	ass.Equal(Country("US"), patDoc.Owners[0].Country)
	ass.Equal("00201214", patDoc.Owners[0].IID)
	ass.Equal("73880", patDoc.Owners[0].IRF)
	ass.Equal("Rochester, New York 14650-2201", patDoc.Owners[0].City)
	ass.Equal("343 State Street", patDoc.Owners[0].Street)
	ass.Equal("Eastman Kodak Company", patDoc.Owners[0].Name)

	// representatives
	ass.NotEmpty(patDoc.Representatives)
	ass.Equal(Country("FR"), patDoc.Representatives[0].Country)
	ass.Equal("00091684", patDoc.Representatives[0].IID)
	ass.Equal("71102 Chalon sur Saône Cedex", patDoc.Representatives[0].City)
	ass.Equal("Kodak Industrie \nDépartement Brevets - CRT \nZone Industrielle", patDoc.Representatives[0].Street)
	ass.Equal("Weber, Etienne Nicolas", patDoc.Representatives[0].Name)

	// contracting states
	ass.NotEmpty(patDoc.ContractingStates)
	ass.Equal(4, len(patDoc.ContractingStates))
	for i := 0; i <= 3; i++ {
		ass.Equal(2, len(patDoc.ContractingStates[i]))
	}

	// classifications
	ass.NotEmpty(patDoc.Classifications)
	ass.Equal(1, len(patDoc.Classifications))
	ass.Equal("H01L  27/146       20060101AFI19970908BHEP        ", patDoc.Classifications[0].Text)
	ass.Equal(IPC, patDoc.Classifications[0].System)
	ass.Equal(1, patDoc.Classifications[0].Sequence)
	ass.Equal("H", patDoc.Classifications[0].Section)
	ass.Equal("01", patDoc.Classifications[0].Class)
	ass.Equal("L", patDoc.Classifications[0].SubClass)
	ass.Equal("27", patDoc.Classifications[0].MainGroup)
	ass.Equal("146", patDoc.Classifications[0].SubGroup)
	ass.Equal("20060101", patDoc.Classifications[0].Version)
	ass.Equal("A", patDoc.Classifications[0].ClassificationLevel)
	ass.Equal("F", patDoc.Classifications[0].FirstLater)
	ass.Equal("I", patDoc.Classifications[0].ClassificationValue)
	ass.Equal("19970908", patDoc.Classifications[0].ActionDate)
	ass.Equal("B", patDoc.Classifications[0].OriginalOrReclassified)
	ass.Equal("H", patDoc.Classifications[0].Source)
	ass.Equal("EP", patDoc.Classifications[0].GeneratingOffice)

}

func TestProcessXMLSimple12B2(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("test-data/grant/v1-2-B2.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("EP90308718B2", patDoc.ID)
	ass.Equal("EP90308718NWB2.xml", patDoc.File)
	ass.Equal("en", patDoc.Lang)
	ass.Equal(Country("EP"), patDoc.Country)
	ass.Equal("0430402", patDoc.DocNumber)
	ass.Equal("B2", patDoc.Kind)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20080305", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("n", patDoc.Status)
	ass.Equal("ep-patent-document-v1-2", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("Verfahren und Zusammensetzungen für chromosomenspezifische Färbung", patDoc.Title[0].Text)
	ass.Equal("de", patDoc.Title[0].Language)
	ass.Equal("Methods and compositions for chromosome-specific staining", patDoc.Title[1].Text)
	ass.Equal("en", patDoc.Title[1].Language)
	ass.Equal("Méthodes et compositions pour la coloration de chromosomes particuliers", patDoc.Title[2].Text)
	ass.Equal("fr", patDoc.Title[2].Language)

	// abstract
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(0, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// claims
	ass.NotEmpty(patDoc.Claims)
	ass.Equal(3, len(patDoc.Claims))
	ass.Equal(1286, len(patDoc.Claims[0].Text))
	ass.Equal("en", patDoc.Claims[0].Language)
	ass.Equal("claims01", patDoc.Claims[0].Id)
	ass.Equal(1550, len(patDoc.Claims[1].Text))
	ass.Equal("de", patDoc.Claims[1].Language)
	ass.Equal("claims02", patDoc.Claims[1].Id)
	ass.Equal(1617, len(patDoc.Claims[2].Text))
	ass.Equal("fr", patDoc.Claims[2].Language)
	ass.Equal("claims03", patDoc.Claims[2].Id)

	// description
	ass.NotEmpty(patDoc.Description)
	ass.Equal(205082, len(patDoc.Description[0].Text))
	ass.Equal("en", patDoc.Description[0].Language)

	// citations
	ass.NotEmpty(patDoc.Citations)
	ass.Equal(4, len(patDoc.Citations))
	ass.Equal(Country("US"), patDoc.Citations[0].Country)
	ass.Equal("4683195", patDoc.Citations[0].DocNumber)
	ass.Equal("A", patDoc.Citations[0].Kind)
	ass.Equal(Country("US"), patDoc.Citations[1].Country)
	ass.Equal("4683202", patDoc.Citations[1].DocNumber)
	ass.Equal("A", patDoc.Citations[1].Kind)
	ass.Equal(Country("US"), patDoc.Citations[2].Country)
	ass.Equal("4800159", patDoc.Citations[2].DocNumber)
	ass.Equal("A", patDoc.Citations[2].Kind)
	ass.Equal(Country("US"), patDoc.Citations[3].Country)
	ass.Equal("389224", patDoc.Citations[3].DocNumber)
	ass.Equal("A", patDoc.Citations[3].Kind)

	// Inventors
	ass.NotEmpty(patDoc.Inventors)
	ass.Equal(Country("US"), patDoc.Inventors[0].Country)
	ass.Equal("Livermore, CA 94550", patDoc.Inventors[0].City)
	ass.Equal("1760 Lomitas", patDoc.Inventors[0].Street)
	ass.Equal("Gray, Joe W.", patDoc.Inventors[0].Name)

	ass.Equal(Country("US"), patDoc.Inventors[1].Country)
	ass.Equal("Livermore, CA 94550", patDoc.Inventors[1].City)
	ass.Equal("540 Shelley Street", patDoc.Inventors[1].Street)
	ass.Equal("Tkachuk, Douglas", patDoc.Inventors[1].Name)

	// owners
	ass.NotEmpty(patDoc.Owners)
	ass.Equal(Country("US"), patDoc.Owners[0].Country)
	ass.Equal("00221072", patDoc.Owners[0].IID)
	ass.Equal("", patDoc.Owners[0].IRF)
	ass.Equal("Oakland,\nCalifornia 94612-3550", patDoc.Owners[0].City)
	ass.Equal("300 Lakeside Drive, \n22nd Floor", patDoc.Owners[0].Street)
	ass.Equal("The Regents of the University \nof California", patDoc.Owners[0].Name)

	// representatives
	ass.NotEmpty(patDoc.Representatives)
	ass.Equal(Country("GB"), patDoc.Representatives[0].Country)
	ass.Equal("00052611", patDoc.Representatives[0].IID)
	ass.Equal("London WC1A 2RA", patDoc.Representatives[0].City)
	ass.Equal("CARPMAELS & RANSFORD \n43-45 Bloomsbury Square", patDoc.Representatives[0].Street)
	ass.Equal("Fisher, Adrian John", patDoc.Representatives[0].Name)

	// contracting states
	ass.NotEmpty(patDoc.ContractingStates)
	ass.Equal(14, len(patDoc.ContractingStates))
	for i := 0; i <= 13; i++ {
		ass.Equal(2, len(patDoc.ContractingStates[i]))
	}

	// classifications
	ass.NotEmpty(patDoc.Classifications)
	ass.Equal(1, len(patDoc.Classifications))
	ass.Equal("C12Q   1/68        20060101AFI19910325BHEP        ", patDoc.Classifications[0].Text)
	ass.Equal(IPC, patDoc.Classifications[0].System)
	ass.Equal(1, patDoc.Classifications[0].Sequence)
	ass.Equal("C", patDoc.Classifications[0].Section)
	ass.Equal("12", patDoc.Classifications[0].Class)
	ass.Equal("Q", patDoc.Classifications[0].SubClass)
	ass.Equal("1", patDoc.Classifications[0].MainGroup)
	ass.Equal("68", patDoc.Classifications[0].SubGroup)
	ass.Equal("20060101", patDoc.Classifications[0].Version)
	ass.Equal("A", patDoc.Classifications[0].ClassificationLevel)
	ass.Equal("F", patDoc.Classifications[0].FirstLater)
	ass.Equal("I", patDoc.Classifications[0].ClassificationValue)
	ass.Equal("19910325", patDoc.Classifications[0].ActionDate)
	ass.Equal("B", patDoc.Classifications[0].OriginalOrReclassified)
	ass.Equal("H", patDoc.Classifications[0].Source)
	ass.Equal("EP", patDoc.Classifications[0].GeneratingOffice)

}

func TestProcessXMLSimple12B21(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("test-data/grant/v1-2-B2-1.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("EP96308101B2", patDoc.ID)
	ass.Equal("EP96308101NWB2.xml", patDoc.File)
	ass.Equal("en", patDoc.Lang)
	ass.Equal(Country("EP"), patDoc.Country)
	ass.Equal("0773297", patDoc.DocNumber)
	ass.Equal("B2", patDoc.Kind)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20080423", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("n", patDoc.Status)
	ass.Equal("ep-patent-document-v1-2", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("Verfahren zur Herstellung einer alpha-Hydroxy-säure oder eines alpha-Hydroxy-amids mittels Mikroorganismen", patDoc.Title[0].Text)
	ass.Equal("de", patDoc.Title[0].Language)
	ass.Equal("Process for producing alpha-hydroxy acid or alpha-hydroxyamide by microorganism", patDoc.Title[1].Text)
	ass.Equal("en", patDoc.Title[1].Language)
	ass.Equal("Procédé pour la préparation d'un acide alpha-hydroxy ou d'un amide alpha-hydroxy en utilisant des microorganismes", patDoc.Title[2].Text)
	ass.Equal("fr", patDoc.Title[2].Language)

	// abstract
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(0, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// claims
	ass.NotEmpty(patDoc.Claims)
	ass.Equal(3, len(patDoc.Claims))
	ass.Equal(2305, len(patDoc.Claims[0].Text))
	ass.Equal("en", patDoc.Claims[0].Language)
	ass.Equal("claims01", patDoc.Claims[0].Id)
	ass.Equal(2575, len(patDoc.Claims[1].Text))
	ass.Equal("de", patDoc.Claims[1].Language)
	ass.Equal("claims02", patDoc.Claims[1].Id)
	ass.Equal(2469, len(patDoc.Claims[2].Text))
	ass.Equal("fr", patDoc.Claims[2].Language)
	ass.Equal("claims03", patDoc.Claims[2].Id)

	// description
	ass.NotEmpty(patDoc.Description)

	ass.Equal(24300, len(patDoc.Description[0].Text))
	ass.Equal("en", patDoc.Description[0].Language)

	// citations
	ass.NotEmpty(patDoc.Citations)
	ass.Equal(21, len(patDoc.Citations))
	for i := 0; i <= 20; i++ {
		ass.Equal("A", patDoc.Citations[i].Kind)
	}
	ass.Equal(Country("JP"), patDoc.Citations[0].Country)
	ass.Equal("2084198", patDoc.Citations[0].DocNumber)
	ass.Equal(Country("JP"), patDoc.Citations[1].Country)
	ass.Equal("3224496", patDoc.Citations[1].DocNumber)
	ass.Equal(Country("JP"), patDoc.Citations[2].Country)
	ass.Equal("3277292", patDoc.Citations[2].DocNumber)
	ass.Equal(Country("JP"), patDoc.Citations[3].Country)
	ass.Equal("4099495", patDoc.Citations[3].DocNumber)
	ass.Equal(Country("JP"), patDoc.Citations[4].Country)
	ass.Equal("4099496", patDoc.Citations[4].DocNumber)
	ass.Equal(Country("JP"), patDoc.Citations[5].Country)
	ass.Equal("4218385", patDoc.Citations[5].DocNumber)
	ass.Equal(Country("US"), patDoc.Citations[6].Country)
	ass.Equal("5223416", patDoc.Citations[6].DocNumber)
	ass.Equal(Country("JP"), patDoc.Citations[7].Country)
	ass.Equal("4099497", patDoc.Citations[7].DocNumber)
	ass.Equal(Country("US"), patDoc.Citations[8].Country)
	ass.Equal("5234826", patDoc.Citations[8].DocNumber)
	ass.Equal(Country("JP"), patDoc.Citations[9].Country)
	ass.Equal("5095795", patDoc.Citations[9].DocNumber)
	ass.Equal(Country("US"), patDoc.Citations[10].Country)
	ass.Equal("5296373", patDoc.Citations[10].DocNumber)
	ass.Equal(Country("US"), patDoc.Citations[10].Country)
	ass.Equal("5296373", patDoc.Citations[10].DocNumber)
	ass.Equal(Country("JP"), patDoc.Citations[11].Country)
	ass.Equal("5021987", patDoc.Citations[11].DocNumber)
	ass.Equal(Country("JP"), patDoc.Citations[12].Country)
	ass.Equal("5192189", patDoc.Citations[12].DocNumber)
	ass.Equal(Country("US"), patDoc.Citations[13].Country)
	ass.Equal("5326702", patDoc.Citations[13].DocNumber)
	ass.Equal(Country("JP"), patDoc.Citations[14].Country)
	ass.Equal("6237789", patDoc.Citations[14].DocNumber)
	ass.Equal(Country("EP"), patDoc.Citations[15].Country)
	ass.Equal("0610048", patDoc.Citations[15].DocNumber)
	ass.Equal(Country("JP"), patDoc.Citations[16].Country)
	ass.Equal("6284899", patDoc.Citations[16].DocNumber)
	ass.Equal(Country("EP"), patDoc.Citations[17].Country)
	ass.Equal("0610049", patDoc.Citations[17].DocNumber)
	ass.Equal(Country("JP"), patDoc.Citations[18].Country)
	ass.Equal("7213296", patDoc.Citations[18].DocNumber)
	ass.Equal(Country("US"), patDoc.Citations[19].Country)
	ass.Equal("5508181", patDoc.Citations[19].DocNumber)
	ass.Equal(Country("JP"), patDoc.Citations[20].Country)
	ass.Equal("4222591", patDoc.Citations[20].DocNumber)

	// Inventors
	ass.NotEmpty(patDoc.Inventors)
	ass.Equal(Country("JP"), patDoc.Inventors[0].Country)
	ass.Equal("Yokohama-shi, Kanagawa", patDoc.Inventors[0].City)
	ass.Equal("10-1 Daikoku-cho,\nTsurumi-ku", patDoc.Inventors[0].Street)
	ass.Equal("Yamaguchi, Yasumasa, \nNitto Chem. Ind. Co., Ltd.", patDoc.Inventors[0].Name)

	ass.Equal(Country("JP"), patDoc.Inventors[1].Country)
	ass.Equal("Yokohama-shi,\nKanagawa", patDoc.Inventors[1].City)
	ass.Equal("10-1 Daikoku-cho,\nTsurumi-ku", patDoc.Inventors[1].Street)
	ass.Equal("Ushigome, Masahiro ,\nNitto Chem. Ind. Co., Ltd.", patDoc.Inventors[1].Name)

	// owners
	ass.NotEmpty(patDoc.Owners)
	ass.Equal(Country("JP"), patDoc.Owners[0].Country)
	ass.Equal("00223389", patDoc.Owners[0].IID)
	ass.Equal("Q036087PEP", patDoc.Owners[0].IRF)
	ass.Equal("Minato-ku,\nTokyo 108-8506", patDoc.Owners[0].City)
	ass.Equal("6-41, Konan 1-chome", patDoc.Owners[0].Street)
	ass.Equal("MITSUBISHI RAYON CO., LTD.", patDoc.Owners[0].Name)

	// representatives
	ass.NotEmpty(patDoc.Representatives)
	ass.Equal(Country("GB"), patDoc.Representatives[0].Country)
	ass.Equal("00034741", patDoc.Representatives[0].IID)
	ass.Equal("Birmingham B1 1TT", patDoc.Representatives[0].City)
	ass.Equal("MARKS & CLERK, \nAlpha Tower, \nSuffolk Street Queensway", patDoc.Representatives[0].Street)
	ass.Equal("Pearce, Anthony Richmond", patDoc.Representatives[0].Name)

	// contracting states
	ass.NotEmpty(patDoc.ContractingStates)
	ass.Equal(3, len(patDoc.ContractingStates))
	for i := 0; i <= 2; i++ {
		ass.Equal(2, len(patDoc.ContractingStates[i]))
	}

	// classifications
	ass.NotEmpty(patDoc.Classifications)
	ass.Equal(2, len(patDoc.Classifications))
	ass.Equal("C12P   7/42        20060101AFI19970220BHEP        ", patDoc.Classifications[0].Text)
	ass.Equal(IPC, patDoc.Classifications[0].System)
	ass.Equal(1, patDoc.Classifications[0].Sequence)
	ass.Equal("C", patDoc.Classifications[0].Section)
	ass.Equal("12", patDoc.Classifications[0].Class)
	ass.Equal("P", patDoc.Classifications[0].SubClass)
	ass.Equal("7", patDoc.Classifications[0].MainGroup)
	ass.Equal("42", patDoc.Classifications[0].SubGroup)
	ass.Equal("20060101", patDoc.Classifications[0].Version)
	ass.Equal("A", patDoc.Classifications[0].ClassificationLevel)
	ass.Equal("F", patDoc.Classifications[0].FirstLater)
	ass.Equal("I", patDoc.Classifications[0].ClassificationValue)
	ass.Equal("19970220", patDoc.Classifications[0].ActionDate)
	ass.Equal("B", patDoc.Classifications[0].OriginalOrReclassified)
	ass.Equal("H", patDoc.Classifications[0].Source)
	ass.Equal("EP", patDoc.Classifications[0].GeneratingOffice)

	ass.Equal("C12P  13/02        20060101ALI19970220BHEP        ", patDoc.Classifications[1].Text)
	ass.Equal(IPC, patDoc.Classifications[1].System)
	ass.Equal(2, patDoc.Classifications[1].Sequence)
	ass.Equal("C", patDoc.Classifications[1].Section)
	ass.Equal("12", patDoc.Classifications[1].Class)
	ass.Equal("P", patDoc.Classifications[1].SubClass)
	ass.Equal("13", patDoc.Classifications[1].MainGroup)
	ass.Equal("02", patDoc.Classifications[1].SubGroup)
	ass.Equal("20060101", patDoc.Classifications[1].Version)
	ass.Equal("A", patDoc.Classifications[1].ClassificationLevel)
	ass.Equal("L", patDoc.Classifications[1].FirstLater)
	ass.Equal("I", patDoc.Classifications[1].ClassificationValue)
	ass.Equal("19970220", patDoc.Classifications[1].ActionDate)
	ass.Equal("B", patDoc.Classifications[1].OriginalOrReclassified)
	ass.Equal("H", patDoc.Classifications[1].Source)
	ass.Equal("EP", patDoc.Classifications[1].GeneratingOffice)

}

//v 1.3
func TestProcessXMLSimple13A1(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("test-data/application/v1-3-A1.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("EP07704233A1", patDoc.ID)
	ass.Equal("07704233.1", patDoc.File)
	ass.Equal("en", patDoc.Lang)
	ass.Equal(Country("EP"), patDoc.Country)
	ass.Equal("1981358", patDoc.DocNumber)
	ass.Equal("A1", patDoc.Kind)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20081022", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("n", patDoc.Status)
	ass.Equal("ep-patent-document-v1-3", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("NÄHRSTOFFZUSAMMENSETZUNG FÜR SÄUGLINGE MIT NIEDRIGEM GEBURTSGEWICHT", patDoc.Title[0].Text)
	ass.Equal("de", patDoc.Title[0].Language)
	ass.Equal("NUTRITIONAL COMPOSITION FOR LOW BIRTH WEIGHT INFANTS", patDoc.Title[1].Text)
	ass.Equal("en", patDoc.Title[1].Language)
	ass.Equal("PRÉPARATION NUTRITIONNELLE POUR NOURRISSONS PRÉSENTANT UN POIDS FAIBLE À LA NAISSANCE", patDoc.Title[2].Text)
	ass.Equal("fr", patDoc.Title[2].Language)

	// abstract Empty
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(0, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// claims
	ass.Empty(patDoc.Claims)

	// description
	ass.NotEmpty(patDoc.Description)
	ass.Equal(0, len(patDoc.Description[0].Text))
	ass.Equal("en", patDoc.Description[0].Language)

	// citations Empty
	ass.Empty(patDoc.Citations)

	// Inventors
	ass.NotEmpty(patDoc.Inventors)
	ass.Equal(Country("CH"), patDoc.Inventors[0].Country)
	ass.Equal("CH-1814 La Tour De Peilz", patDoc.Inventors[0].City)
	ass.Equal("240 Route St. Maurice", patDoc.Inventors[0].Street)
	ass.Equal("HASCHKE, Ferdinand", patDoc.Inventors[0].Name)

	// representatives
	ass.NotEmpty(patDoc.Representatives)
	ass.Equal(Country("CH"), patDoc.Representatives[0].Country)
	ass.Equal("09331551", patDoc.Representatives[0].IID)
	ass.Equal("CH-1800 Vevey", patDoc.Representatives[0].City)
	ass.Equal("Nestec S.A. \nAvenue Nestlé 55", patDoc.Representatives[0].Street)
	ass.Equal("Corticchiato, Olivier", patDoc.Representatives[0].Name)

	// contracting states
	ass.NotEmpty(patDoc.ContractingStates)
	ass.Equal(31, len(patDoc.ContractingStates))
	for i := 0; i <= 30; i++ {
		ass.Equal(2, len(patDoc.ContractingStates[i]))
	}

	// classifications
	ass.NotEmpty(patDoc.Classifications)
	ass.Equal(1, len(patDoc.Classifications))
	ass.Equal("A23L   1/29        20060101AFI20070911BHEP        ", patDoc.Classifications[0].Text)
	ass.Equal(IPC, patDoc.Classifications[0].System)
	ass.Equal(1, patDoc.Classifications[0].Sequence)
	ass.Equal("A", patDoc.Classifications[0].Section)
	ass.Equal("23", patDoc.Classifications[0].Class)
	ass.Equal("L", patDoc.Classifications[0].SubClass)
	ass.Equal("1", patDoc.Classifications[0].MainGroup)
	ass.Equal("29", patDoc.Classifications[0].SubGroup)
	ass.Equal("20060101", patDoc.Classifications[0].Version)
	ass.Equal("A", patDoc.Classifications[0].ClassificationLevel)
	ass.Equal("F", patDoc.Classifications[0].FirstLater)
	ass.Equal("I", patDoc.Classifications[0].ClassificationValue)
	ass.Equal("20070911", patDoc.Classifications[0].ActionDate)
	ass.Equal("B", patDoc.Classifications[0].OriginalOrReclassified)
	ass.Equal("H", patDoc.Classifications[0].Source)
	ass.Equal("EP", patDoc.Classifications[0].GeneratingOffice)

}

func TestProcessXMLSimple13A11(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("test-data/application/v1-3-A1-1.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("EP06798269A1", patDoc.ID)
	ass.Equal("EP06798269NWA1.xml", patDoc.File)
	ass.Equal("en", patDoc.Lang)
	ass.Equal(Country("EP"), patDoc.Country)
	ass.Equal("2009972", patDoc.DocNumber)
	ass.Equal("A1", patDoc.Kind)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20081231", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("n", patDoc.Status)
	ass.Equal("ep-patent-document-v1-3", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("MOBILES ENDGERÄT", patDoc.Title[0].Text)
	ass.Equal("de", patDoc.Title[0].Language)
	ass.Equal("MOBILE TERMINAL", patDoc.Title[1].Text)
	ass.Equal("en", patDoc.Title[1].Language)
	ass.Equal("Terminal mobile", patDoc.Title[2].Text)
	ass.Equal("fr", patDoc.Title[2].Language)

	// abstract
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(796, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// claims
	ass.NotEmpty(patDoc.Claims)
	ass.Equal(4291, len(patDoc.Claims[0].Text))
	ass.Equal("en", patDoc.Claims[0].Language)
	ass.Equal("claims01", patDoc.Claims[0].Id)

	// description
	ass.NotEmpty(patDoc.Description)
	ass.Equal(48259, len(patDoc.Description[0].Text))
	ass.Equal("en", patDoc.Description[0].Language)

	// citations
	ass.NotEmpty(patDoc.Citations)
	ass.Equal("3241669", patDoc.Citations[0].DocNumber)
	ass.Equal(Country("JP"), patDoc.Citations[0].Country)
	ass.Equal("B", patDoc.Citations[0].Kind)

	// Inventors
	ass.NotEmpty(patDoc.Inventors)
	ass.Equal(Country("JP"), patDoc.Inventors[0].Country)
	ass.Equal("1-3-7 Shiromi, Chuo-ku, Osaka 540-6319", patDoc.Inventors[0].City)
	ass.Equal("I.P.R.,\nIP Development Center, 19F Matsushita IMP Bldg.,", patDoc.Inventors[0].Street)
	ass.Equal("KOUNO, Kazunori,\nc/o Matsushita Elec. Industrial Co., Ltd.", patDoc.Inventors[0].Name)

	// representatives
	ass.NotEmpty(patDoc.Representatives)
	ass.Equal(Country("DE"), patDoc.Representatives[0].Country)
	ass.Equal("00100721", patDoc.Representatives[0].IID)
	ass.Equal("80802 München", patDoc.Representatives[0].City)
	ass.Equal("Leopoldstrasse 4", patDoc.Representatives[0].Street)
	ass.Equal("Grünecker, Kinkeldey, \nStockmair & Schwanhäusser \nAnwaltssozietät", patDoc.Representatives[0].Name)

	// contracting states
	ass.NotEmpty(patDoc.ContractingStates)
	ass.Equal(31, len(patDoc.ContractingStates))
	for i := 0; i <= 30; i++ {
		ass.Equal(2, len(patDoc.ContractingStates[i]))
	}

	// classifications
	ass.NotEmpty(patDoc.Classifications)
	ass.Equal(2, len(patDoc.Classifications))

	ass.Equal("H05K   5/02        20060101AFI20080415BHEP        ", patDoc.Classifications[0].Text)
	ass.Equal(IPC, patDoc.Classifications[0].System)
	ass.Equal(1, patDoc.Classifications[0].Sequence)
	ass.Equal("H", patDoc.Classifications[0].Section)
	ass.Equal("05", patDoc.Classifications[0].Class)
	ass.Equal("K", patDoc.Classifications[0].SubClass)
	ass.Equal("5", patDoc.Classifications[0].MainGroup)
	ass.Equal("02", patDoc.Classifications[0].SubGroup)
	ass.Equal("20060101", patDoc.Classifications[0].Version)
	ass.Equal("A", patDoc.Classifications[0].ClassificationLevel)
	ass.Equal("F", patDoc.Classifications[0].FirstLater)
	ass.Equal("I", patDoc.Classifications[0].ClassificationValue)
	ass.Equal("20080415", patDoc.Classifications[0].ActionDate)
	ass.Equal("B", patDoc.Classifications[0].OriginalOrReclassified)
	ass.Equal("H", patDoc.Classifications[0].Source)
	ass.Equal("EP", patDoc.Classifications[0].GeneratingOffice)

	ass.Equal("H04M   1/02        20060101ALI20080415BHEP        ", patDoc.Classifications[1].Text)
	ass.Equal(IPC, patDoc.Classifications[1].System)
	ass.Equal(2, patDoc.Classifications[1].Sequence)
	ass.Equal("H", patDoc.Classifications[1].Section)
	ass.Equal("04", patDoc.Classifications[1].Class)
	ass.Equal("M", patDoc.Classifications[1].SubClass)
	ass.Equal("1", patDoc.Classifications[1].MainGroup)
	ass.Equal("02", patDoc.Classifications[1].SubGroup)
	ass.Equal("20060101", patDoc.Classifications[1].Version)
	ass.Equal("A", patDoc.Classifications[1].ClassificationLevel)
	ass.Equal("L", patDoc.Classifications[1].FirstLater)
	ass.Equal("I", patDoc.Classifications[1].ClassificationValue)
	ass.Equal("20080415", patDoc.Classifications[1].ActionDate)
	ass.Equal("B", patDoc.Classifications[1].OriginalOrReclassified)
	ass.Equal("H", patDoc.Classifications[1].Source)
	ass.Equal("EP", patDoc.Classifications[1].GeneratingOffice)

}

func TestProcessXMLSimple13A2(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("test-data/application/v1-3-A2.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("EP07738344A2", patDoc.ID)
	ass.Equal("EP07738344NWA2.xml", patDoc.File)
	ass.Equal("en", patDoc.Lang)
	ass.Equal(Country("EP"), patDoc.Country)
	ass.Equal("2007181", patDoc.DocNumber)
	ass.Equal("A2", patDoc.Kind)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20081224", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("n", patDoc.Status)
	ass.Equal("ep-patent-document-v1-3", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("STEUERUNG FÜR EINEN INDUSTRIEROBOTER", patDoc.Title[0].Text)
	ass.Equal("de", patDoc.Title[0].Language)
	ass.Equal("CONTROLLER FOR INDUSTRIAL ROBOT", patDoc.Title[1].Text)
	ass.Equal("en", patDoc.Title[1].Language)
	ass.Equal("CONTROLEUR POUR ROBOT INDUSTRIEL", patDoc.Title[2].Text)
	ass.Equal("fr", patDoc.Title[2].Language)

	// abstract Empty
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(723, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// claims
	ass.NotEmpty(patDoc.Claims)
	ass.Equal(1, len(patDoc.Claims))
	ass.Equal(1055, len(patDoc.Claims[0].Text))
	ass.Equal("en", patDoc.Claims[0].Language)
	ass.Equal("claims01", patDoc.Claims[0].Id)

	// description
	ass.NotEmpty(patDoc.Description)
	ass.Equal(18042, len(patDoc.Description[0].Text))
	ass.Equal("en", patDoc.Description[0].Language)

	// citations
	ass.NotEmpty(patDoc.Citations)
	ass.Equal(Country("JP"), patDoc.Citations[0].Country)
	ass.Equal("58181591", patDoc.Citations[0].DocNumber)
	ass.Equal("A", patDoc.Citations[0].Kind)
	ass.Equal(Country("JP"), patDoc.Citations[1].Country)
	ass.Equal("4365581", patDoc.Citations[1].DocNumber)
	ass.Equal("A", patDoc.Citations[1].Kind)

	// Inventors
	ass.NotEmpty(patDoc.Inventors)
	ass.Equal(Country("JP"), patDoc.Inventors[0].Country)
	ass.Equal("Kitakyushu-shi, Fukuoka 806-0004", patDoc.Inventors[0].City)
	ass.Equal("2-1, Kurosaki-Shiroishi, Yahatanishi-ku,", patDoc.Inventors[0].Street)
	ass.Equal("OHFUCHI, Yoshitaka\nc/o KABUSHIKI KAISHA YASKAWA DENKI", patDoc.Inventors[0].Name)

	// representatives
	ass.NotEmpty(patDoc.Representatives)
	ass.Equal(Country("DE"), patDoc.Representatives[0].Country)
	ass.Equal("00100721", patDoc.Representatives[0].IID)
	ass.Equal("80802 München", patDoc.Representatives[0].City)
	ass.Equal("Leopoldstrasse 4", patDoc.Representatives[0].Street)
	ass.Equal("Grünecker, Kinkeldey, \nStockmair & Schwanhäusser \nAnwaltssozietät", patDoc.Representatives[0].Name)

	// contracting states
	ass.NotEmpty(patDoc.ContractingStates)
	ass.Equal(2, len(patDoc.ContractingStates))
	for i := 0; i <= 1; i++ {
		ass.Equal(2, len(patDoc.ContractingStates[i]))
	}

	// classifications
	ass.NotEmpty(patDoc.Classifications)
	ass.Equal(2, len(patDoc.Classifications))

	ass.Equal("H05K   5/02        20060101AFI20081020BHEP        ", patDoc.Classifications[0].Text)
	ass.Equal(IPC, patDoc.Classifications[0].System)
	ass.Equal(1, patDoc.Classifications[0].Sequence)
	ass.Equal("H", patDoc.Classifications[0].Section)
	ass.Equal("05", patDoc.Classifications[0].Class)
	ass.Equal("K", patDoc.Classifications[0].SubClass)
	ass.Equal("5", patDoc.Classifications[0].MainGroup)
	ass.Equal("02", patDoc.Classifications[0].SubGroup)
	ass.Equal("20060101", patDoc.Classifications[0].Version)
	ass.Equal("A", patDoc.Classifications[0].ClassificationLevel)
	ass.Equal("F", patDoc.Classifications[0].FirstLater)
	ass.Equal("I", patDoc.Classifications[0].ClassificationValue)
	ass.Equal("20081020", patDoc.Classifications[0].ActionDate)
	ass.Equal("B", patDoc.Classifications[0].OriginalOrReclassified)
	ass.Equal("H", patDoc.Classifications[0].Source)
	ass.Equal("EP", patDoc.Classifications[0].GeneratingOffice)

	ass.Equal("B25J  13/06        20060101ALI20081020BHEP        ", patDoc.Classifications[1].Text)
	ass.Equal(IPC, patDoc.Classifications[1].System)
	ass.Equal(2, patDoc.Classifications[1].Sequence)
	ass.Equal("B", patDoc.Classifications[1].Section)
	ass.Equal("25", patDoc.Classifications[1].Class)
	ass.Equal("J", patDoc.Classifications[1].SubClass)
	ass.Equal("13", patDoc.Classifications[1].MainGroup)
	ass.Equal("06", patDoc.Classifications[1].SubGroup)
	ass.Equal("20060101", patDoc.Classifications[1].Version)
	ass.Equal("A", patDoc.Classifications[1].ClassificationLevel)
	ass.Equal("L", patDoc.Classifications[1].FirstLater)
	ass.Equal("I", patDoc.Classifications[1].ClassificationValue)
	ass.Equal("20081020", patDoc.Classifications[1].ActionDate)
	ass.Equal("B", patDoc.Classifications[1].OriginalOrReclassified)
	ass.Equal("H", patDoc.Classifications[1].Source)
	ass.Equal("EP", patDoc.Classifications[1].GeneratingOffice)
}

func TestProcessXMLSimple13A21(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("test-data/application/v1-3-A2-1.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("EP08011147A2", patDoc.ID)
	ass.Equal("EP08011147NWA2.xml", patDoc.File)
	ass.Equal("de", patDoc.Lang)
	ass.Equal(Country("EP"), patDoc.Country)
	ass.Equal("2009973", patDoc.DocNumber)
	ass.Equal("A2", patDoc.Kind)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20081231", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("n", patDoc.Status)
	ass.Equal("ep-patent-document-v1-3", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("Einrichtung für die Kontrolle eines sanften Anlaufs oder Auslaufs von Drehstrommotoren, sogenannte Soft Starter", patDoc.Title[0].Text)
	ass.Equal("de", patDoc.Title[0].Language)
	ass.Equal("Device for controlling a soft start or stop for AC motors, so-called soft starters", patDoc.Title[1].Text)
	ass.Equal("en", patDoc.Title[1].Language)
	ass.Equal("Dispositif pour le contrôle d'un démarrage ou d'un arrêt doux de moteurs triphasés, le démarrage progressif", patDoc.Title[2].Text)
	ass.Equal("fr", patDoc.Title[2].Language)

	// abstract Empty
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(942, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// claims
	ass.NotEmpty(patDoc.Claims)
	ass.Equal(1, len(patDoc.Claims))
	ass.Equal(4807, len(patDoc.Claims[0].Text))
	ass.Equal("de", patDoc.Claims[0].Language)
	ass.Equal("claims01", patDoc.Claims[0].Id)

	// description
	ass.NotEmpty(patDoc.Description)
	ass.Equal(21544, len(patDoc.Description[0].Text))
	ass.Equal("en", patDoc.Description[0].Language)

	// citations
	ass.NotEmpty(patDoc.Citations)
	ass.Equal(Country("WO"), patDoc.Citations[0].Country)
	ass.Equal("2006034977", patDoc.Citations[0].DocNumber)
	ass.Equal("A", patDoc.Citations[0].Kind)

	// Inventors
	ass.NotEmpty(patDoc.Inventors)
	ass.Equal(Country("DE"), patDoc.Inventors[0].Country)
	ass.Equal("47661 Issum", patDoc.Inventors[0].City)
	ass.Equal("Girlitzweg 20", patDoc.Inventors[0].Street)
	ass.Equal("Siemes, Andreas", patDoc.Inventors[0].Name)

	// representatives
	ass.NotEmpty(patDoc.Representatives)
	ass.Equal(Country("DE"), patDoc.Representatives[0].Country)
	ass.Equal("09227581", patDoc.Representatives[0].IID)
	ass.Equal("D-47229 Duisburg", patDoc.Representatives[0].City)
	ass.Equal("Am Kruppsee 16", patDoc.Representatives[0].Street)
	ass.Equal("Flaig, Siegfried", patDoc.Representatives[0].Name)

	// contracting states
	ass.NotEmpty(patDoc.ContractingStates)
	ass.Equal(34, len(patDoc.ContractingStates))
	for i := 0; i <= 33; i++ {
		ass.Equal(2, len(patDoc.ContractingStates[i]))
	}

	// classifications
	ass.NotEmpty(patDoc.Classifications)
	ass.Equal(1, len(patDoc.Classifications))
	ass.Equal("H05K   7/14        20060101AFI20081015BHEP        ", patDoc.Classifications[0].Text)
	ass.Equal(IPC, patDoc.Classifications[0].System)
	ass.Equal(1, patDoc.Classifications[0].Sequence)
	ass.Equal("H", patDoc.Classifications[0].Section)
	ass.Equal("05", patDoc.Classifications[0].Class)
	ass.Equal("K", patDoc.Classifications[0].SubClass)
	ass.Equal("7", patDoc.Classifications[0].MainGroup)
	ass.Equal("14", patDoc.Classifications[0].SubGroup)
	ass.Equal("20060101", patDoc.Classifications[0].Version)
	ass.Equal("A", patDoc.Classifications[0].ClassificationLevel)
	ass.Equal("F", patDoc.Classifications[0].FirstLater)
	ass.Equal("I", patDoc.Classifications[0].ClassificationValue)
	ass.Equal("20081015", patDoc.Classifications[0].ActionDate)
	ass.Equal("B", patDoc.Classifications[0].OriginalOrReclassified)
	ass.Equal("H", patDoc.Classifications[0].Source)
	ass.Equal("EP", patDoc.Classifications[0].GeneratingOffice)
}

func TestProcessXMLSimple13B1(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("test-data/grant/v1-3-B1.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("EP04777385B1", patDoc.ID)
	ass.Equal("EP04777385NWB1.xml", patDoc.File)
	ass.Equal("en", patDoc.Lang)
	ass.Equal(Country("EP"), patDoc.Country)
	ass.Equal("1654642", patDoc.DocNumber)
	ass.Equal("B1", patDoc.Kind)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20081224", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("n", patDoc.Status)
	ass.Equal("ep-patent-document-v1-3", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("VERFAHREN UND VORRICHTUNGEN ZUM VERIFIZIEREN VON KONTEXTTEILNEHMERN IN EINEM\n                    KONTEXTVERWALTUNGSSYSTEM IN EINER VERNETZTEN UMGEBUNG", patDoc.Title[0].Text)
	ass.Equal("de", patDoc.Title[0].Language)
	ass.Equal("METHODS AND APPARATUS FOR VERIFYING CONTEXT PARTICIPANTS IN A CONTEXT MANAGEMENT SYSTEM IN A\n                    NETWORKED ENVIRONMENT", patDoc.Title[1].Text)
	ass.Equal("en", patDoc.Title[1].Language)
	ass.Equal("PROCEDES ET DISPOSITIFS POUR VERIFIER DES PARTICIPANTS DE CONTEXTE DANS UN SYSTEME DE GESTION DE\n                    CONTEXTE DANS UN ENVIRONNEMENT RESEAU", patDoc.Title[2].Text)
	ass.Equal("fr", patDoc.Title[2].Language)

	// abstract Empty
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(0, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// claims
	ass.NotEmpty(patDoc.Claims)
	ass.Equal(3, len(patDoc.Claims))
	ass.Equal(23391, len(patDoc.Claims[0].Text))
	ass.Equal("en", patDoc.Claims[0].Language)
	ass.Equal("claims01", patDoc.Claims[0].Id)
	ass.Equal(24747, len(patDoc.Claims[1].Text))
	ass.Equal("de", patDoc.Claims[1].Language)
	ass.Equal("claims02", patDoc.Claims[1].Id)
	ass.Equal(27090, len(patDoc.Claims[2].Text))
	ass.Equal("fr", patDoc.Claims[2].Language)
	ass.Equal("claims03", patDoc.Claims[2].Id)

	// description
	ass.NotEmpty(patDoc.Description)
	ass.Equal(84229, len(patDoc.Description[0].Text))
	ass.Equal("en", patDoc.Description[0].Language)

	// citations
	ass.NotEmpty(patDoc.Citations)
	ass.Equal(Country("US"), patDoc.Citations[0].Country)
	ass.Equal("545396", patDoc.Citations[0].DocNumber)
	ass.Equal("A", patDoc.Citations[0].Kind)
	ass.Equal(Country("WO"), patDoc.Citations[1].Country)
	ass.Equal("0244858", patDoc.Citations[1].DocNumber)
	ass.Equal("A2", patDoc.Citations[1].Kind)
	ass.Equal(Country("US"), patDoc.Citations[2].Country)
	ass.Equal("583301", patDoc.Citations[2].DocNumber)
	ass.Equal("A", patDoc.Citations[2].Kind)

	// Inventors
	ass.NotEmpty(patDoc.Inventors)
	ass.Equal(Country("US"), patDoc.Inventors[0].Country)
	ass.Equal("Groton, MA 01450", patDoc.Inventors[0].City)
	ass.Equal("331 Riverbend Drive", patDoc.Inventors[0].Street)
	ass.Equal("FUSARI, David", patDoc.Inventors[0].Name)

	// representatives
	ass.NotEmpty(patDoc.Representatives)
	ass.Equal(Country("DE"), patDoc.Representatives[0].Country)
	ass.Equal("00101511", patDoc.Representatives[0].IID)
	ass.Equal("81925 München", patDoc.Representatives[0].City)
	ass.Equal("Patent- und Rechtsanwälte Arabellastrasse 4", patDoc.Representatives[0].Street)
	ass.Equal("HOFFMANN EITLE", patDoc.Representatives[0].Name)

	// contracting states
	ass.NotEmpty(patDoc.ContractingStates)
	ass.Equal(28, len(patDoc.ContractingStates))
	for i := 0; i <= 27; i++ {
		ass.Equal(2, len(patDoc.ContractingStates[i]))
	}

	// classifications
	ass.Equal("G06F 9/455 20060101AFI20051108BHEP", patDoc.Classifications[0].Text)
	ass.Equal(IPC, patDoc.Classifications[0].System)
	ass.Equal(1, patDoc.Classifications[0].Sequence)
	ass.Equal("G", patDoc.Classifications[0].Section)
	ass.Equal("06", patDoc.Classifications[0].Class)
	ass.Equal("F", patDoc.Classifications[0].SubClass)
	ass.Equal("9", patDoc.Classifications[0].MainGroup)
	ass.Equal("455", patDoc.Classifications[0].SubGroup)
	ass.Equal("20060101", patDoc.Classifications[0].Version)
	ass.Equal("A", patDoc.Classifications[0].ClassificationLevel)
	ass.Equal("F", patDoc.Classifications[0].FirstLater)
	ass.Equal("I", patDoc.Classifications[0].ClassificationValue)
	ass.Equal("20051108", patDoc.Classifications[0].ActionDate)
	ass.Equal("B", patDoc.Classifications[0].OriginalOrReclassified)
	ass.Equal("H", patDoc.Classifications[0].Source)
	ass.Equal("EP", patDoc.Classifications[0].GeneratingOffice)

}

func TestProcessXMLSimple13B11(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("test-data/grant/v1-3-B1-1.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("EP95937855B1", patDoc.ID)
	ass.Equal("EP95937855NWB1.xml", patDoc.File)
	ass.Equal("en", patDoc.Lang)
	ass.Equal(Country("EP"), patDoc.Country)
	ass.Equal("0791067", patDoc.DocNumber)
	ass.Equal("B1", patDoc.Kind)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20081217", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("n", patDoc.Status)
	ass.Equal("ep-patent-document-v1-3", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("PROTEINKINASE NPK-110", patDoc.Title[0].Text)
	ass.Equal("de", patDoc.Title[0].Language)
	ass.Equal("PROTEIN KINASE NPK-110", patDoc.Title[1].Text)
	ass.Equal("en", patDoc.Title[1].Language)
	ass.Equal("PROTEINE KINASE NPK-110", patDoc.Title[2].Text)
	ass.Equal("fr", patDoc.Title[2].Language)

	// abstract Empty
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(0, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// claims
	ass.NotEmpty(patDoc.Claims)
	ass.Equal(3, len(patDoc.Claims))
	ass.Equal(4113, len(patDoc.Claims[0].Text))
	ass.Equal("en", patDoc.Claims[0].Language)
	ass.Equal("claims01", patDoc.Claims[0].Id)
	ass.Equal(4407, len(patDoc.Claims[1].Text))
	ass.Equal("de", patDoc.Claims[1].Language)
	ass.Equal("claims02", patDoc.Claims[1].Id)
	ass.Equal(4599, len(patDoc.Claims[2].Text))
	ass.Equal("fr", patDoc.Claims[2].Language)
	ass.Equal("claims03", patDoc.Claims[2].Id)

	// description
	ass.NotEmpty(patDoc.Description)
	ass.Equal(84350, len(patDoc.Description[0].Text))
	ass.Equal("en", patDoc.Description[0].Language)

	// citations
	ass.NotEmpty(patDoc.Citations)
	ass.Equal(Country("EP"), patDoc.Citations[0].Country)
	ass.Equal("9202829", patDoc.Citations[0].DocNumber)
	ass.Equal("W", patDoc.Citations[0].Kind)

	// Inventors
	ass.NotEmpty(patDoc.Inventors)
	ass.Equal(Country("DE"), patDoc.Inventors[0].Country)
	ass.Equal("D-22607 Hamburg", patDoc.Inventors[0].City)
	ass.Equal("Baron-Voght-Strasse 212a", patDoc.Inventors[0].Street)
	ass.Equal("MANDELKOW, Eckhard", patDoc.Inventors[0].Name)

	// representatives
	ass.NotEmpty(patDoc.Representatives)
	ass.Equal(Country("DE"), patDoc.Representatives[0].Country)
	ass.Equal("00100311", patDoc.Representatives[0].IID)
	ass.Equal("81634 München", patDoc.Representatives[0].City)
	ass.Equal("P.O. Box 86 07 67", patDoc.Representatives[0].Street)
	ass.Equal("Vossius & Partner", patDoc.Representatives[0].Name)

	// contracting states
	ass.NotEmpty(patDoc.ContractingStates)
	ass.Equal(17, len(patDoc.ContractingStates))
	for i := 0; i <= 16; i++ {
		ass.Equal(2, len(patDoc.ContractingStates[i]))
	}

	// classifications
	ass.Equal(10, len(patDoc.Classifications))
	ass.NotEmpty(patDoc.Classifications)

	for i := 0; i <= 9; i++ {
		ass.Equal(IPC, patDoc.Classifications[i].System)
		ass.Equal(i+1, patDoc.Classifications[i].Sequence)
		ass.Equal("20060101", patDoc.Classifications[i].Version)
		ass.Equal("A", patDoc.Classifications[i].ClassificationLevel)
		ass.Equal("I", patDoc.Classifications[i].ClassificationValue)
		ass.Equal("19960619", patDoc.Classifications[i].ActionDate)
		ass.Equal("B", patDoc.Classifications[i].OriginalOrReclassified)
		ass.Equal("H", patDoc.Classifications[i].Source)
		ass.Equal("EP", patDoc.Classifications[i].GeneratingOffice)
	}
	ass.Equal("C12N  15/54        20060101AFI19960619BHEP        ", patDoc.Classifications[0].Text)
	ass.Equal("C", patDoc.Classifications[0].Section)
	ass.Equal("12", patDoc.Classifications[0].Class)
	ass.Equal("N", patDoc.Classifications[0].SubClass)
	ass.Equal("15", patDoc.Classifications[0].MainGroup)
	ass.Equal("54", patDoc.Classifications[0].SubGroup)

	ass.Equal("A61K  39/395       20060101ALI19960619BHEP        ", patDoc.Classifications[1].Text)
	ass.Equal("A", patDoc.Classifications[1].Section)
	ass.Equal("61", patDoc.Classifications[1].Class)
	ass.Equal("K", patDoc.Classifications[1].SubClass)
	ass.Equal("39", patDoc.Classifications[1].MainGroup)
	ass.Equal("395", patDoc.Classifications[1].SubGroup)
	ass.Equal("L", patDoc.Classifications[1].FirstLater)

	ass.Equal("C12N   9/12        20060101ALI19960619BHEP        ", patDoc.Classifications[2].Text)
	ass.Equal("C", patDoc.Classifications[2].Section)
	ass.Equal("12", patDoc.Classifications[2].Class)
	ass.Equal("N", patDoc.Classifications[2].SubClass)
	ass.Equal("9", patDoc.Classifications[2].MainGroup)
	ass.Equal("12", patDoc.Classifications[2].SubGroup)
	ass.Equal("L", patDoc.Classifications[2].FirstLater)

	ass.Equal("A61K  38/46        20060101ALI19960619BHEP        ", patDoc.Classifications[3].Text)
	ass.Equal("A", patDoc.Classifications[3].Section)
	ass.Equal("61", patDoc.Classifications[3].Class)
	ass.Equal("K", patDoc.Classifications[3].SubClass)
	ass.Equal("38", patDoc.Classifications[3].MainGroup)
	ass.Equal("46", patDoc.Classifications[3].SubGroup)
	ass.Equal("L", patDoc.Classifications[3].FirstLater)

	ass.Equal("C07C 229/08        20060101ALI19960619BHEP        ", patDoc.Classifications[4].Text)
	ass.Equal("C", patDoc.Classifications[4].Section)
	ass.Equal("07", patDoc.Classifications[4].Class)
	ass.Equal("C", patDoc.Classifications[4].SubClass)
	ass.Equal("229", patDoc.Classifications[4].MainGroup)
	ass.Equal("08", patDoc.Classifications[4].SubGroup)
	ass.Equal("L", patDoc.Classifications[4].FirstLater)

	ass.Equal("A61K  38/16        20060101ALI19960619BHEP        ", patDoc.Classifications[5].Text)
	ass.Equal("A", patDoc.Classifications[5].Section)
	ass.Equal("61", patDoc.Classifications[5].Class)
	ass.Equal("K", patDoc.Classifications[5].SubClass)
	ass.Equal("38", patDoc.Classifications[5].MainGroup)
	ass.Equal("16", patDoc.Classifications[5].SubGroup)
	ass.Equal("L", patDoc.Classifications[5].FirstLater)

	ass.Equal("C07K   4/12        20060101ALI19960619BHEP        ", patDoc.Classifications[6].Text)
	ass.Equal("C", patDoc.Classifications[6].Section)
	ass.Equal("07", patDoc.Classifications[6].Class)
	ass.Equal("K", patDoc.Classifications[6].SubClass)
	ass.Equal("4", patDoc.Classifications[6].MainGroup)
	ass.Equal("12", patDoc.Classifications[6].SubGroup)
	ass.Equal("L", patDoc.Classifications[6].FirstLater)

	ass.Equal("C12Q   1/42        20060101ALI19960619BHEP        ", patDoc.Classifications[7].Text)
	ass.Equal("C", patDoc.Classifications[7].Section)
	ass.Equal("12", patDoc.Classifications[7].Class)
	ass.Equal("Q", patDoc.Classifications[7].SubClass)
	ass.Equal("1", patDoc.Classifications[7].MainGroup)
	ass.Equal("42", patDoc.Classifications[7].SubGroup)
	ass.Equal("L", patDoc.Classifications[7].FirstLater)

	ass.Equal("C07K  16/40        20060101ALI19960619BHEP        ", patDoc.Classifications[8].Text)
	ass.Equal("C", patDoc.Classifications[8].Section)
	ass.Equal("07", patDoc.Classifications[8].Class)
	ass.Equal("K", patDoc.Classifications[8].SubClass)
	ass.Equal("16", patDoc.Classifications[8].MainGroup)
	ass.Equal("40", patDoc.Classifications[8].SubGroup)
	ass.Equal("L", patDoc.Classifications[8].FirstLater)

	ass.Equal("C12Q   1/48        20060101ALI19960619BHEP        ", patDoc.Classifications[9].Text)
	ass.Equal("C", patDoc.Classifications[9].Section)
	ass.Equal("12", patDoc.Classifications[9].Class)
	ass.Equal("Q", patDoc.Classifications[9].SubClass)
	ass.Equal("1", patDoc.Classifications[9].MainGroup)
	ass.Equal("48", patDoc.Classifications[9].SubGroup)
	ass.Equal("L", patDoc.Classifications[9].FirstLater)

}

func TestProcessXMLSimple13B12(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("test-data/grant/v1-3-B1-2.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("EP98400796B1", patDoc.ID)
	ass.Equal("EP98400796NWB1.xml", patDoc.File)
	ass.Equal("en", patDoc.Lang)
	ass.Equal(Country("EP"), patDoc.Country)
	ass.Equal("0874348", patDoc.DocNumber)
	ass.Equal("B1", patDoc.Kind)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20081231", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("n", patDoc.Status)
	ass.Equal("ep-patent-document-v1-3", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("Verfahren und Schaltungsanordnung zur Adressierung einer Plasmaanzeigetafel mit verschiedenen Koden", patDoc.Title[0].Text)
	ass.Equal("de", patDoc.Title[0].Language)
	ass.Equal("Process and device for addressing of a plasma display with diverse codes", patDoc.Title[1].Text)
	ass.Equal("en", patDoc.Title[1].Language)
	ass.Equal("Procédé et dispositif d'addressage à code changeant d'un écran à plasma", patDoc.Title[2].Text)
	ass.Equal("fr", patDoc.Title[2].Language)

	// abstract Empty
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(0, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// claims
	ass.NotEmpty(patDoc.Claims)
	ass.Equal(3, len(patDoc.Claims))
	ass.Equal(2250, len(patDoc.Claims[0].Text))
	ass.Equal("en", patDoc.Claims[0].Language)
	ass.Equal("claims01", patDoc.Claims[0].Id)
	ass.Equal(2637, len(patDoc.Claims[1].Text))
	ass.Equal("de", patDoc.Claims[1].Language)
	ass.Equal("claims02", patDoc.Claims[1].Id)
	ass.Equal(2498, len(patDoc.Claims[2].Text))
	ass.Equal("fr", patDoc.Claims[2].Language)
	ass.Equal("claims03", patDoc.Claims[2].Id)

	// description
	ass.NotEmpty(patDoc.Description)
	ass.Equal(20643, len(patDoc.Description[0].Text))
	ass.Equal("en", patDoc.Description[0].Language)

	// citations
	ass.NotEmpty(patDoc.Citations)
	ass.Equal(Country("EP"), patDoc.Citations[0].Country)
	ass.Equal("0837441", patDoc.Citations[0].DocNumber)
	ass.Equal("A", patDoc.Citations[0].Kind)
	ass.Equal(Country("US"), patDoc.Citations[1].Country)
	ass.Equal("5187578", patDoc.Citations[1].DocNumber)
	ass.Equal("A", patDoc.Citations[1].Kind)
	ass.Equal(Country("JP"), patDoc.Citations[2].Country)
	ass.Equal("7271325", patDoc.Citations[2].DocNumber)
	ass.Equal("A", patDoc.Citations[2].Kind)

	// Inventors
	ass.NotEmpty(patDoc.Inventors)
	ass.Equal(Country("FR"), patDoc.Inventors[0].Country)
	ass.Equal("92100 Boulogne Billancourt", patDoc.Inventors[0].City)
	ass.Equal("THOMSON Multimedia,\n46 quai Alphonse Le Gallo", patDoc.Inventors[0].Street)
	ass.Equal("Doyen, Didier", patDoc.Inventors[0].Name)

	// representatives
	ass.NotEmpty(patDoc.Representatives)
	ass.Equal(Country("FR"), patDoc.Representatives[0].Country)
	ass.Equal("00047343", patDoc.Representatives[0].IID)
	ass.Equal("92648 Boulogne Cedex", patDoc.Representatives[0].City)
	ass.Equal("Thomson multimedia \nPatent Department \n46 Quai A. Le Gallo", patDoc.Representatives[0].Street)
	ass.Equal("Ruellan-Lemonnier, Brigitte", patDoc.Representatives[0].Name)

	// contracting states
	ass.NotEmpty(patDoc.ContractingStates)
	ass.Equal(4, len(patDoc.ContractingStates))
	for i := 0; i <= 3; i++ {
		ass.Equal(2, len(patDoc.ContractingStates[i]))
	}

	// classifications
	ass.Equal(1, len(patDoc.Classifications))
	ass.Equal("G09G   3/28        20060101AFI19980714BHEP        ", patDoc.Classifications[0].Text)
	ass.Equal(IPC, patDoc.Classifications[0].System)
	ass.Equal(1, patDoc.Classifications[0].Sequence)
	ass.Equal("G", patDoc.Classifications[0].Section)
	ass.Equal("09", patDoc.Classifications[0].Class)
	ass.Equal("G", patDoc.Classifications[0].SubClass)
	ass.Equal("3", patDoc.Classifications[0].MainGroup)
	ass.Equal("28", patDoc.Classifications[0].SubGroup)
	ass.Equal("20060101", patDoc.Classifications[0].Version)
	ass.Equal("A", patDoc.Classifications[0].ClassificationLevel)
	ass.Equal("F", patDoc.Classifications[0].FirstLater)
	ass.Equal("I", patDoc.Classifications[0].ClassificationValue)
	ass.Equal("19980714", patDoc.Classifications[0].ActionDate)
	ass.Equal("B", patDoc.Classifications[0].OriginalOrReclassified)
	ass.Equal("H", patDoc.Classifications[0].Source)
	ass.Equal("EP", patDoc.Classifications[0].GeneratingOffice)

}

func TestProcessXMLSimple13B2(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("test-data/grant/v1-3-B2.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("EP03721225B2", patDoc.ID)
	ass.Equal("EP03721225NWB2.xml", patDoc.File)
	ass.Equal("en", patDoc.Lang)
	ass.Equal(Country("EP"), patDoc.Country)
	ass.Equal("1497510", patDoc.DocNumber)
	ass.Equal("B2", patDoc.Kind)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20081224", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("n", patDoc.Status)
	ass.Equal("ep-patent-document-v1-3", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("FUSSBODEN", patDoc.Title[0].Text)
	ass.Equal("de", patDoc.Title[0].Language)
	ass.Equal("FLOORING", patDoc.Title[1].Text)
	ass.Equal("en", patDoc.Title[1].Language)
	ass.Equal("PLANCHER", patDoc.Title[2].Text)
	ass.Equal("fr", patDoc.Title[2].Language)

	// abstract Empty
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(0, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// claims
	ass.NotEmpty(patDoc.Claims)
	ass.Equal(3, len(patDoc.Claims))
	ass.Equal(1583, len(patDoc.Claims[0].Text))
	ass.Equal("en", patDoc.Claims[0].Language)
	ass.Equal("claims01", patDoc.Claims[0].Id)
	ass.Equal(1886, len(patDoc.Claims[1].Text))
	ass.Equal("de", patDoc.Claims[1].Language)
	ass.Equal("claims02", patDoc.Claims[1].Id)
	ass.Equal(2053, len(patDoc.Claims[2].Text))
	ass.Equal("fr", patDoc.Claims[2].Language)
	ass.Equal("claims03", patDoc.Claims[2].Id)

	// description
	ass.NotEmpty(patDoc.Description)
	ass.Equal(53875, len(patDoc.Description[0].Text))
	ass.Equal("en", patDoc.Description[0].Language)

	// citations
	ass.NotEmpty(patDoc.Citations)
	ass.Equal(Country("WO"), patDoc.Citations[0].Country)
	ass.Equal("030253087", patDoc.Citations[0].DocNumber)
	ass.Equal("A", patDoc.Citations[0].Kind)
	ass.Equal(Country("WO"), patDoc.Citations[1].Country)
	ass.Equal("0166877", patDoc.Citations[1].DocNumber)
	ass.Equal("A", patDoc.Citations[1].Kind)
	ass.Equal(Country("WO"), patDoc.Citations[2].Country)
	ass.Equal("0020705", patDoc.Citations[2].DocNumber)
	ass.Equal("A", patDoc.Citations[2].Kind)

	// Inventors
	ass.NotEmpty(patDoc.Inventors)
	ass.Equal(Country("SE"), patDoc.Inventors[0].Country)
	ass.Equal("S-260 40 Viken", patDoc.Inventors[0].City)
	ass.Equal("Kyrkogränden 1", patDoc.Inventors[0].Street)
	ass.Equal("PERVAN, Darko", patDoc.Inventors[0].Name)

	// representatives
	ass.NotEmpty(patDoc.Representatives)
	ass.Equal(Country("SE"), patDoc.Representatives[0].Country)
	ass.Equal("09247041", patDoc.Representatives[0].IID)
	ass.Equal("251 10 Helsingborg", patDoc.Representatives[0].City)
	ass.Equal("Awapatent AB \nBox 1066", patDoc.Representatives[0].Street)
	ass.Equal("Åkesson, Sten Jan-Åke", patDoc.Representatives[0].Name)

	// contracting states
	ass.NotEmpty(patDoc.ContractingStates)
	ass.Equal(27, len(patDoc.ContractingStates))
	for i := 0; i <= 26; i++ {
		ass.Equal(2, len(patDoc.ContractingStates[i]))
	}

	// classifications
	ass.NotEmpty(patDoc.Classifications)
	ass.Equal(1, len(patDoc.Classifications))
	ass.Equal("E04F  15/04        20060101AFI20031106BHEP        ", patDoc.Classifications[0].Text)
	ass.Equal(IPC, patDoc.Classifications[0].System)
	ass.Equal(1, patDoc.Classifications[0].Sequence)
	ass.Equal("E", patDoc.Classifications[0].Section)
	ass.Equal("04", patDoc.Classifications[0].Class)
	ass.Equal("F", patDoc.Classifications[0].SubClass)
	ass.Equal("15", patDoc.Classifications[0].MainGroup)
	ass.Equal("04", patDoc.Classifications[0].SubGroup)
	ass.Equal("20060101", patDoc.Classifications[0].Version)
	ass.Equal("A", patDoc.Classifications[0].ClassificationLevel)
	ass.Equal("F", patDoc.Classifications[0].FirstLater)
	ass.Equal("I", patDoc.Classifications[0].ClassificationValue)
	ass.Equal("20031106", patDoc.Classifications[0].ActionDate)
	ass.Equal("B", patDoc.Classifications[0].OriginalOrReclassified)
	ass.Equal("H", patDoc.Classifications[0].Source)
	ass.Equal("EP", patDoc.Classifications[0].GeneratingOffice)

}

func TestProcessXMLSimple13B21(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("test-data/grant/v1-3-B2-1.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("EP93110982B2", patDoc.ID)
	ass.Equal("EP93110982NWB2.xml", patDoc.File)
	ass.Equal("de", patDoc.Lang)
	ass.Equal(Country("EP"), patDoc.Country)
	ass.Equal("0588006", patDoc.DocNumber)
	ass.Equal("B2", patDoc.Kind)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20081217", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("n", patDoc.Status)
	ass.Equal("ep-patent-document-v1-3", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("Rundsteuerverfahren und Rundsteuerempfänger zur Ausführung des Rundsteuerverfahrens", patDoc.Title[0].Text)
	ass.Equal("de", patDoc.Title[0].Language)
	ass.Equal("Remote-control method and associated receiver", patDoc.Title[1].Text)
	ass.Equal("en", patDoc.Title[1].Language)
	ass.Equal("Procédé de télécommande et récepteur pour sa mise en place", patDoc.Title[2].Text)
	ass.Equal("fr", patDoc.Title[2].Language)

	// abstract Empty
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(0, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// claims
	ass.NotEmpty(patDoc.Claims)
	ass.Equal(3, len(patDoc.Claims))
	ass.Equal(11419, len(patDoc.Claims[0].Text))
	ass.Equal("de", patDoc.Claims[0].Language)
	ass.Equal("claims01", patDoc.Claims[0].Id)
	ass.Equal(11884, len(patDoc.Claims[1].Text))
	ass.Equal("en", patDoc.Claims[1].Language)
	ass.Equal("claims02", patDoc.Claims[1].Id)
	ass.Equal(14516, len(patDoc.Claims[2].Text))
	ass.Equal("fr", patDoc.Claims[2].Language)
	ass.Equal("claims03", patDoc.Claims[2].Id)

	// description
	ass.NotEmpty(patDoc.Description)
	ass.Equal(58030, len(patDoc.Description[0].Text))
	ass.Equal("en", patDoc.Description[0].Language)

	// citations
	ass.NotEmpty(patDoc.Citations)
	ass.Equal(Country("WO"), patDoc.Citations[0].Country)
	ass.Equal("8802583", patDoc.Citations[0].DocNumber)
	ass.Equal("A", patDoc.Citations[0].Kind)
	ass.Equal(Country("EP"), patDoc.Citations[1].Country)
	ass.Equal("0153856", patDoc.Citations[1].DocNumber)
	ass.Equal("A", patDoc.Citations[1].Kind)

	// Inventors
	ass.NotEmpty(patDoc.Inventors)
	ass.Equal(Country("CH"), patDoc.Inventors[0].Country)
	ass.Equal("CH-6312 Steinhausen", patDoc.Inventors[0].City)
	ass.Equal("Albisstr. 58", patDoc.Inventors[0].Street)
	ass.Equal("Lorenzi, Marcel", patDoc.Inventors[0].Name)

	// representatives
	ass.NotEmpty(patDoc.Representatives)
	ass.Equal(Country("CH"), patDoc.Representatives[0].Country)
	ass.Equal("00106161", patDoc.Representatives[0].IID)
	ass.Equal("8022 Zürich", patDoc.Representatives[0].City)
	ass.Equal("Fraumünsterstrasse 9 \nPostfach 2441", patDoc.Representatives[0].Street)
	ass.Equal("Rentsch & Partner", patDoc.Representatives[0].Name)

	// contracting states
	ass.NotEmpty(patDoc.ContractingStates)
	ass.Equal(8, len(patDoc.ContractingStates))
	for i := 0; i <= 7; i++ {
		ass.Equal(2, len(patDoc.ContractingStates[i]))
	}

	// classifications
	ass.NotEmpty(patDoc.Classifications)
	ass.Equal(2, len(patDoc.Classifications))
	ass.Equal("H02J  13/00        20060101AFI19940107BHEP        ", patDoc.Classifications[0].Text)
	ass.Equal(IPC, patDoc.Classifications[0].System)
	ass.Equal(1, patDoc.Classifications[0].Sequence)
	ass.Equal("H", patDoc.Classifications[0].Section)
	ass.Equal("02", patDoc.Classifications[0].Class)
	ass.Equal("J", patDoc.Classifications[0].SubClass)
	ass.Equal("13", patDoc.Classifications[0].MainGroup)
	ass.Equal("00", patDoc.Classifications[0].SubGroup)
	ass.Equal("20060101", patDoc.Classifications[0].Version)
	ass.Equal("A", patDoc.Classifications[0].ClassificationLevel)
	ass.Equal("F", patDoc.Classifications[0].FirstLater)
	ass.Equal("I", patDoc.Classifications[0].ClassificationValue)
	ass.Equal("19940107", patDoc.Classifications[0].ActionDate)
	ass.Equal("B", patDoc.Classifications[0].OriginalOrReclassified)
	ass.Equal("H", patDoc.Classifications[0].Source)
	ass.Equal("EP", patDoc.Classifications[0].GeneratingOffice)

	ass.Equal("H04Q   9/14        20060101ALI19940107BHEP        ", patDoc.Classifications[1].Text)
	ass.Equal(IPC, patDoc.Classifications[1].System)
	ass.Equal(2, patDoc.Classifications[1].Sequence)
	ass.Equal("H", patDoc.Classifications[1].Section)
	ass.Equal("04", patDoc.Classifications[1].Class)
	ass.Equal("Q", patDoc.Classifications[1].SubClass)
	ass.Equal("9", patDoc.Classifications[1].MainGroup)
	ass.Equal("14", patDoc.Classifications[1].SubGroup)
	ass.Equal("20060101", patDoc.Classifications[1].Version)
	ass.Equal("A", patDoc.Classifications[1].ClassificationLevel)
	ass.Equal("L", patDoc.Classifications[1].FirstLater)
	ass.Equal("I", patDoc.Classifications[1].ClassificationValue)
	ass.Equal("19940107", patDoc.Classifications[0].ActionDate)
	ass.Equal("B", patDoc.Classifications[1].OriginalOrReclassified)
	ass.Equal("H", patDoc.Classifications[1].Source)
	ass.Equal("EP", patDoc.Classifications[1].GeneratingOffice)
}

//v 1.4
func TestProcessXMLSimple14A1(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("test-data/application/v1-4-A1.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("EP08018756A1", patDoc.ID)
	ass.Equal("EP08018756NWA1.xml", patDoc.File)
	ass.Equal("en", patDoc.Lang)
	ass.Equal(Country("EP"), patDoc.Country)
	ass.Equal("2055205", patDoc.DocNumber)
	ass.Equal("A1", patDoc.Kind)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20090506", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("n", patDoc.Status)
	ass.Equal("ep-patent-document-v1-4", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("Haarspange", patDoc.Title[0].Text)
	ass.Equal("de", patDoc.Title[0].Language)
	ass.Equal("Bobby pin", patDoc.Title[1].Text)
	ass.Equal("en", patDoc.Title[1].Language)
	ass.Equal("Pince à cheveux", patDoc.Title[2].Text)
	ass.Equal("fr", patDoc.Title[2].Language)

	// abstract
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(793, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// claims
	ass.NotEmpty(patDoc.Claims)
	ass.Equal(1, len(patDoc.Claims))
	ass.Equal(3542, len(patDoc.Claims[0].Text))
	ass.Equal("en", patDoc.Claims[0].Language)
	ass.Equal("claims01", patDoc.Claims[0].Id)

	// description
	ass.NotEmpty(patDoc.Description)
	ass.Equal(12087, len(patDoc.Description[0].Text))
	ass.Equal("en", patDoc.Description[0].Language)

	// citations
	ass.Empty(patDoc.Citations)

	// Inventors
	ass.NotEmpty(patDoc.Inventors)
	ass.Equal(1, len(patDoc.Inventors))
	ass.Equal(Country("TW"), patDoc.Inventors[0].Country)
	ass.Equal("Tainan City", patDoc.Inventors[0].City)
	ass.Equal("No. 28, Chun-Te 23th St.", patDoc.Inventors[0].Street)
	ass.Equal("Tu, Jui-Ming", patDoc.Inventors[0].Name)

	// representatives
	ass.NotEmpty(patDoc.Representatives)
	ass.Equal(1, len(patDoc.Representatives))
	ass.Equal("80469 München", patDoc.Representatives[0].City)
	ass.Equal("Kador & Partner", patDoc.Representatives[0].Name)
	ass.Equal("00100211", patDoc.Representatives[0].IID)
	ass.Equal(Country("DE"), patDoc.Representatives[0].Country)
	ass.Equal("Corneliusstrasse 15", patDoc.Representatives[0].Street)

	// contracting states
	ass.NotEmpty(patDoc.ContractingStates)
	ass.Equal(34, len(patDoc.ContractingStates))
	for i := 0; i <= 6; i++ {
		ass.Equal(2, len(patDoc.ContractingStates[i]))
	}

	// classifications
	ass.NotEmpty(patDoc.Classifications)
	ass.Equal(2, len(patDoc.Classifications))

	for i := 0; i <= 1; i++ {
		ass.Equal(IPC, patDoc.Classifications[i].System)
		ass.Equal(i+1, patDoc.Classifications[i].Sequence)
		ass.Equal("A", patDoc.Classifications[i].Section)
		ass.Equal("45", patDoc.Classifications[i].Class)
		ass.Equal("D", patDoc.Classifications[i].SubClass)
		ass.Equal("8", patDoc.Classifications[i].MainGroup)
		ass.Equal("20060101", patDoc.Classifications[i].Version)
		ass.Equal("A", patDoc.Classifications[i].ClassificationLevel)
		ass.Equal("I", patDoc.Classifications[i].ClassificationValue)
		ass.Equal("20090304", patDoc.Classifications[i].ActionDate)
		ass.Equal("B", patDoc.Classifications[i].OriginalOrReclassified)
		ass.Equal("H", patDoc.Classifications[i].Source)
		ass.Equal("EP", patDoc.Classifications[i].GeneratingOffice)
	}

	ass.Equal("A45D 8/20 20060101AFI20090304BHEP", patDoc.Classifications[0].Text)
	ass.Equal("20", patDoc.Classifications[0].SubGroup)
	ass.Equal("F", patDoc.Classifications[0].FirstLater)

	ass.Equal("A45D 8/34 20060101ALI20090304BHEP", patDoc.Classifications[1].Text)
	ass.Equal("34", patDoc.Classifications[1].SubGroup)
	ass.Equal("L", patDoc.Classifications[1].FirstLater)

}

func TestProcessXMLSimple14A2(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("test-data/application/v1-4-A2.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("EP09007746A2", patDoc.ID)
	ass.Equal("EP09007746NWA2.xml", patDoc.File)
	ass.Equal("de", patDoc.Lang)
	ass.Equal(Country("EP"), patDoc.Country)
	ass.Equal("2136607", patDoc.DocNumber)
	ass.Equal("A2", patDoc.Kind)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20091223", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("n", patDoc.Status)
	ass.Equal("ep-patent-document-v1-4", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("Verfahren zum Betreiben einer Akku-Handlampe", patDoc.Title[0].Text)
	ass.Equal("de", patDoc.Title[0].Language)
	ass.Equal("Method for operating a battery powered torch", patDoc.Title[1].Text)
	ass.Equal("en", patDoc.Title[1].Language)
	ass.Equal("Procédé de fonctionnement d'une lampe manuelle à batterie", patDoc.Title[2].Text)
	ass.Equal("fr", patDoc.Title[2].Language)

	// abstract Empty
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(775, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// claims
	ass.NotEmpty(patDoc.Claims)
	ass.Equal(1, len(patDoc.Claims))
	ass.Equal(5712, len(patDoc.Claims[0].Text))
	ass.Equal("de", patDoc.Claims[0].Language)
	ass.Equal("claims01", patDoc.Claims[0].Id)

	// description
	ass.NotEmpty(patDoc.Description)
	ass.Equal(12570, len(patDoc.Description[0].Text))
	ass.Equal("en", patDoc.Description[0].Language)

	// citations
	ass.NotEmpty(patDoc.Citations)
	ass.Equal(Country("DE"), patDoc.Citations[0].Country)
	ass.Equal("9412976", patDoc.Citations[0].DocNumber)
	ass.Equal("U1", patDoc.Citations[0].Kind)
	ass.Equal(Country("DE"), patDoc.Citations[1].Country)
	ass.Equal("2446958", patDoc.Citations[1].DocNumber)
	ass.Equal("B2", patDoc.Citations[1].Kind)

	// Inventors
	ass.NotEmpty(patDoc.Inventors)
	ass.Equal(Country("DE"), patDoc.Inventors[0].Country)
	ass.Equal("73266 Bissingen", patDoc.Inventors[0].City)
	ass.Equal("Gartenstraße 3", patDoc.Inventors[0].Street)
	ass.Equal("Gräber, Jochen", patDoc.Inventors[0].Name)

	// representatives
	ass.NotEmpty(patDoc.Representatives)
	ass.Equal(Country("DE"), patDoc.Representatives[0].Country)
	ass.Equal("00100861", patDoc.Representatives[0].IID)
	ass.Equal("70032 Stuttgart", patDoc.Representatives[0].City)
	ass.Equal("Patentanwälte \nPostfach 10 37 62", patDoc.Representatives[0].Street)
	ass.Equal("Dreiss", patDoc.Representatives[0].Name)

	// contracting states
	ass.NotEmpty(patDoc.ContractingStates)
	ass.Equal(35, len(patDoc.ContractingStates))
	for i := 0; i <= 34; i++ {
		ass.Equal(2, len(patDoc.ContractingStates[i]))
	}

	// classifications
	ass.NotEmpty(patDoc.Classifications)
	ass.Equal(1, len(patDoc.Classifications))

	ass.Equal("H05B  37/02        20060101AFI20091022BHEP        ", patDoc.Classifications[0].Text)
	ass.Equal(IPC, patDoc.Classifications[0].System)
	ass.Equal(1, patDoc.Classifications[0].Sequence)
	ass.Equal("H", patDoc.Classifications[0].Section)
	ass.Equal("05", patDoc.Classifications[0].Class)
	ass.Equal("B", patDoc.Classifications[0].SubClass)
	ass.Equal("37", patDoc.Classifications[0].MainGroup)
	ass.Equal("02", patDoc.Classifications[0].SubGroup)
	ass.Equal("20060101", patDoc.Classifications[0].Version)
	ass.Equal("A", patDoc.Classifications[0].ClassificationLevel)
	ass.Equal("F", patDoc.Classifications[0].FirstLater)
	ass.Equal("I", patDoc.Classifications[0].ClassificationValue)
	ass.Equal("20091022", patDoc.Classifications[0].ActionDate)
	ass.Equal("B", patDoc.Classifications[0].OriginalOrReclassified)
	ass.Equal("H", patDoc.Classifications[0].Source)
	ass.Equal("EP", patDoc.Classifications[0].GeneratingOffice)
}

func TestProcessXMLSimple14B1(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("test-data/grant/v1-4-B1.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("EP91302636B1", patDoc.ID)
	ass.Equal("EP91302636NWB1.xml", patDoc.File)
	ass.Equal("en", patDoc.Lang)
	ass.Equal(Country("EP"), patDoc.Country)
	ass.Equal("0449582", patDoc.DocNumber)
	ass.Equal("B1", patDoc.Kind)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20091021", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("n", patDoc.Status)
	ass.Equal("ep-patent-document-v1-4", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("Messverfahren und -vorrichtung", patDoc.Title[0].Text)
	ass.Equal("de", patDoc.Title[0].Language)
	ass.Equal("Measuring method and apparatus", patDoc.Title[1].Text)
	ass.Equal("en", patDoc.Title[1].Language)
	ass.Equal("Méthode et appareil de mesure", patDoc.Title[2].Text)
	ass.Equal("fr", patDoc.Title[2].Language)

	// abstract Empty
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(0, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// claims
	ass.NotEmpty(patDoc.Claims)
	ass.Equal(3, len(patDoc.Claims))
	ass.Equal(9380, len(patDoc.Claims[0].Text))
	ass.Equal("en", patDoc.Claims[0].Language)
	ass.Equal("claims01", patDoc.Claims[0].Id)
	ass.Equal(10389, len(patDoc.Claims[1].Text))
	ass.Equal("de", patDoc.Claims[1].Language)
	ass.Equal("claims02", patDoc.Claims[1].Id)
	ass.Equal(11395, len(patDoc.Claims[2].Text))
	ass.Equal("fr", patDoc.Claims[2].Language)
	ass.Equal("claims03", patDoc.Claims[2].Id)

	// description
	ass.NotEmpty(patDoc.Description)
	ass.Equal(59625, len(patDoc.Description[0].Text))
	ass.Equal("en", patDoc.Description[0].Language)

	// citations
	ass.NotEmpty(patDoc.Citations)
	ass.Equal(Country("US"), patDoc.Citations[0].Country)
	ass.Equal("4710026", patDoc.Citations[0].DocNumber)
	ass.Equal("A", patDoc.Citations[0].Kind)

	// Inventors
	ass.NotEmpty(patDoc.Inventors)
	ass.Equal(Country("JP"), patDoc.Inventors[0].Country)
	ass.Equal("Ohta-ku,\nTokyo", patDoc.Inventors[0].City)
	ass.Equal("c/o Canon Kabushiki Kaisha,\n3-30-2 Shimomaruko", patDoc.Inventors[0].Street)
	ass.Equal("Nose, Noriyuki", patDoc.Inventors[0].Name)

	// representatives
	ass.NotEmpty(patDoc.Representatives)
	ass.Equal(Country("DE"), patDoc.Representatives[0].Country)
	ass.Equal("00102381", patDoc.Representatives[0].IID)
	ass.Equal("80336 München", patDoc.Representatives[0].City)
	ass.Equal("Bavariaring 4-6", patDoc.Representatives[0].Street)
	ass.Equal("TBK-Patent", patDoc.Representatives[0].Name)

	// contracting states
	ass.NotEmpty(patDoc.ContractingStates)
	ass.Equal(6, len(patDoc.ContractingStates))
	for i := 0; i <= 5; i++ {
		ass.Equal(2, len(patDoc.ContractingStates[i]))
	}

	// classifications
	ass.NotEmpty(patDoc.Classifications)
	ass.Equal(2, len(patDoc.Classifications))
	ass.Equal("G03F   9/00        20060101AFI19920327BHEP        ", patDoc.Classifications[0].Text)
	ass.Equal(IPC, patDoc.Classifications[0].System)
	ass.Equal(1, patDoc.Classifications[0].Sequence)
	ass.Equal("G", patDoc.Classifications[0].Section)
	ass.Equal("03", patDoc.Classifications[0].Class)
	ass.Equal("F", patDoc.Classifications[0].SubClass)
	ass.Equal("9", patDoc.Classifications[0].MainGroup)
	ass.Equal("00", patDoc.Classifications[0].SubGroup)
	ass.Equal("20060101", patDoc.Classifications[0].Version)
	ass.Equal("A", patDoc.Classifications[0].ClassificationLevel)
	ass.Equal("F", patDoc.Classifications[0].FirstLater)
	ass.Equal("I", patDoc.Classifications[0].ClassificationValue)
	ass.Equal("19920327", patDoc.Classifications[0].ActionDate)
	ass.Equal("B", patDoc.Classifications[0].OriginalOrReclassified)
	ass.Equal("H", patDoc.Classifications[0].Source)
	ass.Equal("EP", patDoc.Classifications[0].GeneratingOffice)

	ass.Equal("G03F   7/20        20060101ALI19920327BHEP        ", patDoc.Classifications[1].Text)
	ass.Equal(IPC, patDoc.Classifications[1].System)
	ass.Equal(2, patDoc.Classifications[1].Sequence)
	ass.Equal("G", patDoc.Classifications[1].Section)
	ass.Equal("03", patDoc.Classifications[1].Class)
	ass.Equal("F", patDoc.Classifications[1].SubClass)
	ass.Equal("7", patDoc.Classifications[1].MainGroup)
	ass.Equal("20", patDoc.Classifications[1].SubGroup)
	ass.Equal("20060101", patDoc.Classifications[1].Version)
	ass.Equal("A", patDoc.Classifications[1].ClassificationLevel)
	ass.Equal("L", patDoc.Classifications[1].FirstLater)
	ass.Equal("I", patDoc.Classifications[1].ClassificationValue)
	ass.Equal("19920327", patDoc.Classifications[1].ActionDate)
	ass.Equal("B", patDoc.Classifications[1].OriginalOrReclassified)
	ass.Equal("H", patDoc.Classifications[1].Source)
	ass.Equal("EP", patDoc.Classifications[1].GeneratingOffice)
}

func TestProcessXMLSimple14B11(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("test-data/grant/v1-4-B1-1.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("EP95305416B1", patDoc.ID)
	ass.Equal("EP95305416NWB1.xml", patDoc.File)
	ass.Equal("en", patDoc.Lang)
	ass.Equal(Country("EP"), patDoc.Country)
	ass.Equal("0697382", patDoc.DocNumber)
	ass.Equal("B1", patDoc.Kind)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20091216", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("n", patDoc.Status)
	ass.Equal("ep-patent-document-v1-4", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("Eine Gips/Zellulosefaser enthaltende Zusammensetzung für Schallziegel", patDoc.Title[0].Text)
	ass.Equal("de", patDoc.Title[0].Language)
	ass.Equal("A gypsum/cellulosic fiber acoustical tile composition", patDoc.Title[1].Text)
	ass.Equal("en", patDoc.Title[1].Language)
	ass.Equal("Une composition pour dalle acoustique en plâtre et fibre cellulosique", patDoc.Title[2].Text)
	ass.Equal("fr", patDoc.Title[2].Language)

	// abstract Empty
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(0, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// claims
	ass.NotEmpty(patDoc.Claims)
	ass.Equal(3, len(patDoc.Claims))
	ass.Equal(7741, len(patDoc.Claims[0].Text))
	ass.Equal("en", patDoc.Claims[0].Language)
	ass.Equal("claims01", patDoc.Claims[0].Id)
	ass.Equal(7787, len(patDoc.Claims[1].Text))
	ass.Equal("de", patDoc.Claims[1].Language)
	ass.Equal("claims02", patDoc.Claims[1].Id)
	ass.Equal(9462, len(patDoc.Claims[2].Text))
	ass.Equal("fr", patDoc.Claims[2].Language)
	ass.Equal("claims03", patDoc.Claims[2].Id)

	// description
	ass.NotEmpty(patDoc.Description)
	ass.Equal(54145, len(patDoc.Description[0].Text))
	ass.Equal("en", patDoc.Description[0].Language)

	// citations
	ass.NotEmpty(patDoc.Citations)
	ass.Equal(Country("US"), patDoc.Citations[0].Country)
	ass.Equal("1769519", patDoc.Citations[0].DocNumber)
	ass.Equal("A", patDoc.Citations[0].Kind)

	// Inventors
	ass.NotEmpty(patDoc.Inventors)
	ass.Equal(Country("US"), patDoc.Inventors[0].Country)
	ass.Equal("Des Plaines,\nIllinois 60016", patDoc.Inventors[0].City)
	ass.Equal("9415 Meadow Lane", patDoc.Inventors[0].Street)
	ass.Equal("Baig, Mirza A.", patDoc.Inventors[0].Name)

	// representatives
	ass.NotEmpty(patDoc.Representatives)
	ass.Equal(Country("GB"), patDoc.Representatives[0].Country)
	ass.Equal("00042961", patDoc.Representatives[0].IID)
	ass.Equal("London\nWC2E 9RA", patDoc.Representatives[0].City)
	ass.Equal("Marks & Clerk LLP \n90 Long Acre", patDoc.Representatives[0].Street)
	ass.Equal("Johnson, Terence Leslie", patDoc.Representatives[0].Name)

	// contracting states
	ass.NotEmpty(patDoc.ContractingStates)
	ass.Equal(17, len(patDoc.ContractingStates))
	for i := 0; i <= 16; i++ {
		ass.Equal(2, len(patDoc.ContractingStates[i]))
	}

	// classifications
	ass.NotEmpty(patDoc.Classifications)
	ass.Equal(4, len(patDoc.Classifications))

	for i := 0; i <= 3; i++ {
		ass.Equal(IPC, patDoc.Classifications[i].System)
		ass.Equal(1+i, patDoc.Classifications[i].Sequence)
		ass.Equal("B", patDoc.Classifications[i].SubClass)
		ass.Equal("C", patDoc.Classifications[i].Section)
		ass.Equal("04", patDoc.Classifications[i].Class)
		ass.Equal("20060101", patDoc.Classifications[i].Version)
		ass.Equal("A", patDoc.Classifications[i].ClassificationLevel)
		ass.Equal("I", patDoc.Classifications[i].ClassificationValue)
		ass.Equal("20090612", patDoc.Classifications[i].ActionDate)
		ass.Equal("B", patDoc.Classifications[i].OriginalOrReclassified)
		ass.Equal("H", patDoc.Classifications[i].Source)
		ass.Equal("EP", patDoc.Classifications[i].GeneratingOffice)
	}

	ass.Equal("C04B  26/28        20060101AFI20090612BHEP        ", patDoc.Classifications[0].Text)
	ass.Equal("26", patDoc.Classifications[0].MainGroup)
	ass.Equal("28", patDoc.Classifications[0].SubGroup)
	ass.Equal("F", patDoc.Classifications[0].FirstLater)

	ass.Equal("C04B  18/16        20060101ALI20090612BHEP        ", patDoc.Classifications[1].Text)
	ass.Equal("18", patDoc.Classifications[1].MainGroup)
	ass.Equal("16", patDoc.Classifications[1].SubGroup)
	ass.Equal("L", patDoc.Classifications[1].FirstLater)

	ass.Equal("C04B  14/36        20060101ALI20090612BHEP        ", patDoc.Classifications[2].Text)
	ass.Equal("14", patDoc.Classifications[2].MainGroup)
	ass.Equal("36", patDoc.Classifications[2].SubGroup)
	ass.Equal("L", patDoc.Classifications[2].FirstLater)

	ass.Equal("C04B  28/14        20060101ALI20090612BHEP        ", patDoc.Classifications[3].Text)
	ass.Equal("28", patDoc.Classifications[3].MainGroup)
	ass.Equal("14", patDoc.Classifications[3].SubGroup)
	ass.Equal("L", patDoc.Classifications[3].FirstLater)
}

func TestProcessXMLSimple14B2(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("test-data/grant/v1-4-B2.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("EP02776152B2", patDoc.ID)
	ass.Equal("EP02776152NWB2.xml", patDoc.File)
	ass.Equal("en", patDoc.Lang)
	ass.Equal(Country("EP"), patDoc.Country)
	ass.Equal("1451194", patDoc.DocNumber)
	ass.Equal("B2", patDoc.Kind)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20091216", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("n", patDoc.Status)
	ass.Equal("ep-patent-document-v1-4", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("HERSTELLUNG VON LEVOFLOXACIN HEMIHYDRAT", patDoc.Title[0].Text)
	ass.Equal("de", patDoc.Title[0].Language)
	ass.Equal("PREPARATION OF LEVOFLOXACIN HEMIHYDRATE", patDoc.Title[1].Text)
	ass.Equal("en", patDoc.Title[1].Language)
	ass.Equal("PREPARATION DE LEVOFLOXACINE HEMIHYDRATE", patDoc.Title[2].Text)
	ass.Equal("fr", patDoc.Title[2].Language)

	// abstract Empty
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(0, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// claims
	ass.NotEmpty(patDoc.Claims)
	ass.Equal(3, len(patDoc.Claims))
	ass.Equal(1625, len(patDoc.Claims[0].Text))
	ass.Equal("en", patDoc.Claims[0].Language)
	ass.Equal("claims01", patDoc.Claims[0].Id)
	ass.Equal(1859, len(patDoc.Claims[1].Text))
	ass.Equal("de", patDoc.Claims[1].Language)
	ass.Equal("claims02", patDoc.Claims[1].Id)
	ass.Equal(1999, len(patDoc.Claims[2].Text))
	ass.Equal("fr", patDoc.Claims[2].Language)
	ass.Equal("claims03", patDoc.Claims[2].Id)

	// description
	ass.NotEmpty(patDoc.Description)
	ass.Equal(17333, len(patDoc.Description[0].Text))
	ass.Equal("en", patDoc.Description[0].Language)

	// citations
	ass.NotEmpty(patDoc.Citations)
	ass.Equal(Country("US"), patDoc.Citations[0].Country)
	ass.Equal("4382892", patDoc.Citations[0].DocNumber)
	ass.Equal("A", patDoc.Citations[0].Kind)
	ass.Equal(Country("US"), patDoc.Citations[1].Country)
	ass.Equal("5053407", patDoc.Citations[1].DocNumber)
	ass.Equal("A", patDoc.Citations[1].Kind)
	ass.Equal(Country("US"), patDoc.Citations[2].Country)
	ass.Equal("5051505", patDoc.Citations[2].DocNumber)
	ass.Equal("A", patDoc.Citations[2].Kind)

	// Inventors
	ass.NotEmpty(patDoc.Inventors)
	ass.Equal(Country("IL"), patDoc.Inventors[0].Country)
	ass.Equal("Ein Vered 40696", patDoc.Inventors[0].City)
	ass.Equal("Pob 21,\nIsraelian", patDoc.Inventors[0].Street)
	ass.Equal("NIDDAM-HILDESHEIM, Valerie", patDoc.Inventors[0].Name)

	// representatives
	ass.NotEmpty(patDoc.Representatives)
	ass.Equal(Country("GB"), patDoc.Representatives[0].Country)
	ass.Equal("00085471", patDoc.Representatives[0].IID)
	ass.Equal("London EC1N 2DY", patDoc.Representatives[0].City)
	ass.Equal("D Young & Co \n120 Holborn", patDoc.Representatives[0].Street)
	ass.Equal("Nachshen, Neil Jacob", patDoc.Representatives[0].Name)

	// contracting states
	ass.NotEmpty(patDoc.ContractingStates)
	ass.Equal(24, len(patDoc.ContractingStates))
	for i := 0; i <= 23; i++ {
		ass.Equal(2, len(patDoc.ContractingStates[i]))
	}

	// classifications
	ass.NotEmpty(patDoc.Classifications)
	ass.Equal(5, len(patDoc.Classifications))

	for i := 0; i <= 4; i++ {
		ass.Equal(IPC, patDoc.Classifications[i].System)
		ass.Equal(i+1, patDoc.Classifications[i].Sequence)
		ass.Equal("20060101", patDoc.Classifications[i].Version)
		ass.Equal("A", patDoc.Classifications[i].ClassificationLevel)
		ass.Equal("B", patDoc.Classifications[i].OriginalOrReclassified)
		ass.Equal("H", patDoc.Classifications[i].Source)
		ass.Equal("EP", patDoc.Classifications[i].GeneratingOffice)

	}

	ass.Equal("C07D 498/06        20060101AFI20040428BHEP        ", patDoc.Classifications[0].Text)
	ass.Equal("C", patDoc.Classifications[0].Section)
	ass.Equal("07", patDoc.Classifications[0].Class)
	ass.Equal("D", patDoc.Classifications[0].SubClass)
	ass.Equal("20040428", patDoc.Classifications[0].ActionDate)
	ass.Equal("I", patDoc.Classifications[0].ClassificationValue)
	ass.Equal("498", patDoc.Classifications[0].MainGroup)
	ass.Equal("06", patDoc.Classifications[0].SubGroup)
	ass.Equal("F", patDoc.Classifications[0].FirstLater)

	ass.Equal("A61K  31/5383      20060101ALI20050425BHEP        ", patDoc.Classifications[1].Text)
	ass.Equal("A", patDoc.Classifications[1].Section)
	ass.Equal("61", patDoc.Classifications[1].Class)
	ass.Equal("K", patDoc.Classifications[1].SubClass)
	ass.Equal("20050425", patDoc.Classifications[1].ActionDate)
	ass.Equal("I", patDoc.Classifications[1].ClassificationValue)
	ass.Equal("31", patDoc.Classifications[1].MainGroup)
	ass.Equal("5383", patDoc.Classifications[1].SubGroup)
	ass.Equal("L", patDoc.Classifications[1].FirstLater)

	ass.Equal("C07D 498/06        20060101ALN20050425BHEP        ", patDoc.Classifications[2].Text)
	ass.Equal("C", patDoc.Classifications[2].Section)
	ass.Equal("07", patDoc.Classifications[2].Class)
	ass.Equal("D", patDoc.Classifications[2].SubClass)
	ass.Equal("20050425", patDoc.Classifications[2].ActionDate)
	ass.Equal("N", patDoc.Classifications[2].ClassificationValue)
	ass.Equal("498", patDoc.Classifications[2].MainGroup)
	ass.Equal("06", patDoc.Classifications[2].SubGroup)
	ass.Equal("L", patDoc.Classifications[2].FirstLater)

	ass.Equal("C07D 265/00        20060101ALN20050425BHEP        ", patDoc.Classifications[3].Text)
	ass.Equal("C", patDoc.Classifications[3].Section)
	ass.Equal("07", patDoc.Classifications[3].Class)
	ass.Equal("D", patDoc.Classifications[3].SubClass)
	ass.Equal("20050425", patDoc.Classifications[3].ActionDate)
	ass.Equal("N", patDoc.Classifications[3].ClassificationValue)
	ass.Equal("265", patDoc.Classifications[3].MainGroup)
	ass.Equal("00", patDoc.Classifications[3].SubGroup)
	ass.Equal("L", patDoc.Classifications[3].FirstLater)

	ass.Equal("C07D 221/00        20060101ALN20050425BHEP        ", patDoc.Classifications[4].Text)
	ass.Equal("C", patDoc.Classifications[4].Section)
	ass.Equal("07", patDoc.Classifications[4].Class)
	ass.Equal("D", patDoc.Classifications[4].SubClass)
	ass.Equal("20050425", patDoc.Classifications[4].ActionDate)
	ass.Equal("N", patDoc.Classifications[4].ClassificationValue)
	ass.Equal("221", patDoc.Classifications[4].MainGroup)
	ass.Equal("00", patDoc.Classifications[4].SubGroup)
	ass.Equal("L", patDoc.Classifications[4].FirstLater)
}

func TestProcessXMLSimple14B21(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("test-data/grant/v1-4-B2-1.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("EP95905359B2", patDoc.ID)
	ass.Equal("EP95905359NWB2.xml", patDoc.File)
	ass.Equal("en", patDoc.Lang)
	ass.Equal(Country("EP"), patDoc.Country)
	ass.Equal("0736085", patDoc.DocNumber)
	ass.Equal("B2", patDoc.Kind)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20091216", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("n", patDoc.Status)
	ass.Equal("ep-patent-document-v1-4", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("PERCARBONAT , AMYLASE  UND PROTEASE ENTHALTENDE WASCHMITTELZUSAMMENSETZUNGEN", patDoc.Title[0].Text)
	ass.Equal("de", patDoc.Title[0].Language)
	ass.Equal("DETERGENT COMPOSITIONS CONTAINING PERCARBONATE,  AMYLASE AND PROTEASE", patDoc.Title[1].Text)
	ass.Equal("en", patDoc.Title[1].Language)
	ass.Equal("COMPOSITION DETERGENTE CONTENANT DU PERCARBONATE, DE DE L'AMYLASE ET UNE PROTEASE", patDoc.Title[2].Text)
	ass.Equal("fr", patDoc.Title[2].Language)

	// abstract Empty
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(0, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// claims
	ass.NotEmpty(patDoc.Claims)
	ass.Equal(3, len(patDoc.Claims))
	ass.Equal(2336, len(patDoc.Claims[0].Text))
	ass.Equal("en", patDoc.Claims[0].Language)
	ass.Equal("claims01", patDoc.Claims[0].Id)
	ass.Equal(2591, len(patDoc.Claims[1].Text))
	ass.Equal("de", patDoc.Claims[1].Language)
	ass.Equal("claims02", patDoc.Claims[1].Id)
	ass.Equal(2791, len(patDoc.Claims[2].Text))
	ass.Equal("fr", patDoc.Claims[2].Language)
	ass.Equal("claims03", patDoc.Claims[2].Id)

	// description
	ass.NotEmpty(patDoc.Description)
	ass.Equal(40196, len(patDoc.Description[0].Text))
	ass.Equal("en", patDoc.Description[0].Language)

	// citations
	ass.NotEmpty(patDoc.Citations)
	ass.Equal(Country("JP"), patDoc.Citations[0].Country)
	ass.Equal("57028197", patDoc.Citations[0].DocNumber)
	ass.Equal("B", patDoc.Citations[0].Kind)
	ass.Equal(Country("WO"), patDoc.Citations[1].Country)
	ass.Equal("9401521", patDoc.Citations[1].DocNumber)
	ass.Equal("A", patDoc.Citations[1].Kind)
	ass.Equal(Country("WO"), patDoc.Citations[2].Country)
	ass.Equal("9403554", patDoc.Citations[2].DocNumber)
	ass.Equal("A", patDoc.Citations[2].Kind)

	// Inventors
	ass.NotEmpty(patDoc.Inventors)
	ass.Equal(Country("GB"), patDoc.Inventors[0].Country)
	ass.Equal("Northumberland", patDoc.Inventors[0].City)
	ass.Equal("46 Park Avenue Prudhoe", patDoc.Inventors[0].Street)
	ass.Equal("MOSS, Michael, Alan, John", patDoc.Inventors[0].Name)

	// representatives
	ass.NotEmpty(patDoc.Representatives)
	ass.Equal(Country("GB"), patDoc.Representatives[0].Country)
	ass.Equal("09397621", patDoc.Representatives[0].IID)
	ass.Equal("Newcastle upon Tyne NE12 9TS", patDoc.Representatives[0].City)
	ass.Equal("Procter & Gamble \nTechnical Centres Limited \nWhitley Road \nLongbenton", patDoc.Representatives[0].Street)
	ass.Equal("Howard, Phillip Jan", patDoc.Representatives[0].Name)

	// contracting states
	ass.NotEmpty(patDoc.ContractingStates)
	ass.Equal(6, len(patDoc.ContractingStates))
	for i := 0; i <= 5; i++ {
		ass.Equal(2, len(patDoc.ContractingStates[i]))
	}

	// classifications
	ass.NotEmpty(patDoc.Classifications)
	ass.Equal(3, len(patDoc.Classifications))

	for i := 0; i <= 2; i++ {
		ass.Equal(IPC, patDoc.Classifications[i].System)
		ass.Equal(i+1, patDoc.Classifications[i].Sequence)
		ass.Equal("C", patDoc.Classifications[i].Section)
		ass.Equal("11", patDoc.Classifications[i].Class)
		ass.Equal("20060101", patDoc.Classifications[i].Version)
		ass.Equal("A", patDoc.Classifications[i].ClassificationLevel)
		ass.Equal("I", patDoc.Classifications[i].ClassificationValue)
		ass.Equal("B", patDoc.Classifications[i].OriginalOrReclassified)
		ass.Equal("H", patDoc.Classifications[i].Source)
		ass.Equal("EP", patDoc.Classifications[i].GeneratingOffice)

	}

	ass.Equal("C11D   3/386       20060101AFI19950726BHEP        ", patDoc.Classifications[0].Text)
	ass.Equal("D", patDoc.Classifications[0].SubClass)
	ass.Equal("19950726", patDoc.Classifications[0].ActionDate)
	ass.Equal("3", patDoc.Classifications[0].MainGroup)
	ass.Equal("386", patDoc.Classifications[0].SubGroup)
	ass.Equal("F", patDoc.Classifications[0].FirstLater)

	ass.Equal("C11D   3/395       20060101ALI19970121BHEP        ", patDoc.Classifications[1].Text)
	ass.Equal("D", patDoc.Classifications[1].SubClass)
	ass.Equal("19970121", patDoc.Classifications[1].ActionDate)
	ass.Equal("3", patDoc.Classifications[1].MainGroup)
	ass.Equal("395", patDoc.Classifications[1].SubGroup)
	ass.Equal("L", patDoc.Classifications[1].FirstLater)

	ass.Equal("C11D   3/39        20060101ALI19970121BHEP        ", patDoc.Classifications[2].Text)
	ass.Equal("D", patDoc.Classifications[2].SubClass)
	ass.Equal("19970121", patDoc.Classifications[2].ActionDate)
	ass.Equal("3", patDoc.Classifications[2].MainGroup)
	ass.Equal("39", patDoc.Classifications[2].SubGroup)
	ass.Equal("L", patDoc.Classifications[2].FirstLater)
}

//v 1.5
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
	ass.NotEmpty(patDoc.Classifications)
	ass.Equal(2, len(patDoc.Classifications))

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

func TestProcessXMLSimple15A11(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("test-data/application/v1-5-A1-1.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("EP19750349A1", patDoc.ID)
	ass.Equal("19750349.3", patDoc.File)
	ass.Equal("en", patDoc.Lang)
	ass.Equal(Country("EP"), patDoc.Country)
	ass.Equal("3749080", patDoc.DocNumber)
	ass.Equal("A1", patDoc.Kind)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20201216", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("n", patDoc.Status)
	ass.Equal("ep-patent-document-v1-5", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("ADAPTIVE STEUERUNG EINES MÄHERS", patDoc.Title[0].Text)
	ass.Equal("de", patDoc.Title[0].Language)
	ass.Equal("ADAPTIVE CONTROL OF A MOWER", patDoc.Title[1].Text)
	ass.Equal("en", patDoc.Title[1].Language)
	ass.Equal("COMMANDE ADAPTATIVE DE TONDEUSE", patDoc.Title[2].Text)
	ass.Equal("fr", patDoc.Title[2].Language)

	// abstract zero
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(0, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// claims empty
	ass.Empty(patDoc.Claims)

	// description zero
	ass.NotEmpty(patDoc.Description)
	ass.Equal(0, len(patDoc.Description[0].Text))
	ass.Equal("en", patDoc.Description[0].Language)

	// citations empty
	ass.Empty(patDoc.Citations)

	// Inventors
	ass.NotEmpty(patDoc.Inventors)
	ass.Equal(Country("US"), patDoc.Inventors[0].Country)
	ass.Equal("Farmington, Utah 84025", patDoc.Inventors[0].City)
	ass.Equal("1465 North 1580 West", patDoc.Inventors[0].Street)
	ass.Equal("APOSHIAN, Steven", patDoc.Inventors[0].Name)

	ass.Equal(Country("US"), patDoc.Inventors[1].Country)
	ass.Equal("Farmington, Utah 84025", patDoc.Inventors[1].City)
	ass.Equal("1033 N. Main Street", patDoc.Inventors[1].Street)
	ass.Equal("ASTON, Eric", patDoc.Inventors[1].Name)

	ass.Equal(Country("US"), patDoc.Inventors[2].Country)
	ass.Equal("Salt Lake City, Utah 84109", patDoc.Inventors[2].City)
	ass.Equal("2412 East 2100 South", patDoc.Inventors[2].Street)
	ass.Equal("DECKER, William", patDoc.Inventors[2].Name)

	// representatives
	ass.NotEmpty(patDoc.Representatives)
	ass.Equal(Country("GB"), patDoc.Representatives[0].Country)
	ass.Equal("101020411", patDoc.Representatives[0].IID)
	ass.Equal("Leeds LS2 8PA", patDoc.Representatives[0].City)
	ass.Equal("Urquhart-Dykes & Lord LLP \nArena Point \nMerrion Way", patDoc.Representatives[0].Street)
	ass.Equal("Orr, Robert", patDoc.Representatives[0].Name)

	// contracting states
	ass.NotEmpty(patDoc.ContractingStates)
	ass.Equal(38, len(patDoc.ContractingStates))
	for i := 0; i <= 37; i++ {
		ass.Equal(2, len(patDoc.ContractingStates[i]))
	}

	// classifications
	ass.NotEmpty(patDoc.Classifications)
	ass.Equal(5, len(patDoc.Classifications))

	for i := 0; i <= 4; i++ {
		ass.Equal(IPC, patDoc.Classifications[i].System)
		ass.Equal(i+1, patDoc.Classifications[i].Sequence)
		ass.Equal("20060101", patDoc.Classifications[i].Version)
		ass.Equal("A", patDoc.Classifications[i].Section)
		ass.Equal("A", patDoc.Classifications[i].ClassificationLevel)
		ass.Equal("I", patDoc.Classifications[i].ClassificationValue)
		ass.Equal("20190816", patDoc.Classifications[i].ActionDate)
		ass.Equal("B", patDoc.Classifications[i].OriginalOrReclassified)
		ass.Equal("H", patDoc.Classifications[i].Source)
		ass.Equal("EP", patDoc.Classifications[i].GeneratingOffice)

	}

	ass.Equal("A01D  34/00        20060101AFI20190816BHEP        ", patDoc.Classifications[0].Text)
	ass.Equal("01", patDoc.Classifications[0].Class)
	ass.Equal("D", patDoc.Classifications[0].SubClass)
	ass.Equal("34", patDoc.Classifications[0].MainGroup)
	ass.Equal("00", patDoc.Classifications[0].SubGroup)
	ass.Equal("F", patDoc.Classifications[0].FirstLater)

	ass.Equal("A01D  34/63        20060101ALI20190816BHEP        ", patDoc.Classifications[1].Text)
	ass.Equal("01", patDoc.Classifications[1].Class)
	ass.Equal("D", patDoc.Classifications[1].SubClass)
	ass.Equal("34", patDoc.Classifications[1].MainGroup)
	ass.Equal("63", patDoc.Classifications[1].SubGroup)
	ass.Equal("L", patDoc.Classifications[1].FirstLater)

	ass.Equal("A01D  34/66        20060101ALI20190816BHEP        ", patDoc.Classifications[2].Text)
	ass.Equal("01", patDoc.Classifications[2].Class)
	ass.Equal("D", patDoc.Classifications[2].SubClass)
	ass.Equal("34", patDoc.Classifications[2].MainGroup)
	ass.Equal("66", patDoc.Classifications[2].SubGroup)
	ass.Equal("L", patDoc.Classifications[2].FirstLater)

	ass.Equal("A01D  34/74        20060101ALI20190816BHEP        ", patDoc.Classifications[3].Text)
	ass.Equal("01", patDoc.Classifications[3].Class)
	ass.Equal("D", patDoc.Classifications[3].SubClass)
	ass.Equal("34", patDoc.Classifications[3].MainGroup)
	ass.Equal("74", patDoc.Classifications[3].SubGroup)
	ass.Equal("L", patDoc.Classifications[3].FirstLater)

	ass.Equal("A01D  34/82        20060101ALI20190816BHEP        ", patDoc.Classifications[4].Text)
	ass.Equal("01", patDoc.Classifications[4].Class)
	ass.Equal("D", patDoc.Classifications[4].SubClass)
	ass.Equal("34", patDoc.Classifications[4].MainGroup)
	ass.Equal("82", patDoc.Classifications[4].SubGroup)
	ass.Equal("L", patDoc.Classifications[4].FirstLater)

}

func TestProcessXMLSimple15A2(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("test-data/application/v1-5-A2.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("EP19766069A2", patDoc.ID)
	ass.Equal("19766069.9", patDoc.File)
	ass.Equal("en", patDoc.Lang)
	ass.Equal(Country("EP"), patDoc.Country)
	ass.Equal("3814387", patDoc.DocNumber)
	ass.Equal("A2", patDoc.Kind)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20210505", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("n", patDoc.Status)
	ass.Equal("ep-patent-document-v1-5", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("STRAHLENUNDURCHLÄSSIGE POLYMERE", patDoc.Title[0].Text)
	ass.Equal("de", patDoc.Title[0].Language)
	ass.Equal("RADIOPAQUE POLYMERS", patDoc.Title[1].Text)
	ass.Equal("en", patDoc.Title[1].Language)
	ass.Equal("POLYMÈRES RADIO-OPAQUES", patDoc.Title[2].Text)
	ass.Equal("fr", patDoc.Title[2].Language)

	// abstract Empty!!
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(0, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// claims
	ass.Empty(patDoc.Claims)

	// description
	//	ass.Empty(patDoc.Description)

	// citations
	ass.Empty(patDoc.Citations)

	// Inventors
	ass.NotEmpty(patDoc.Inventors)
	for i := 0; i <= 5; i++ {
		ass.Equal(Country("GB"), patDoc.Inventors[i].Country)
		ass.Equal("London \nGreater London EC4M 7RD", patDoc.Inventors[i].City)
		ass.Equal("C/O BTG International Limited \n5 Fleet Place", patDoc.Inventors[i].Street)
	}
	ass.Equal("LEWIS, Andrew, Lennard", patDoc.Inventors[0].Name)
	ass.Equal("BRITTON, Hugh", patDoc.Inventors[1].Name)
	ass.Equal("TANG, Yiqing", patDoc.Inventors[2].Name)
	ass.Equal("VINCE, Jonathan", patDoc.Inventors[3].Name)
	ass.Equal("ASHRAFI, Koorosh", patDoc.Inventors[4].Name)
	ass.Equal("GUEGEN, Damien", patDoc.Inventors[5].Name)

	// representatives
	ass.NotEmpty(patDoc.Representatives)
	ass.Equal(Country("DE"), patDoc.Representatives[0].Country)
	ass.Equal("101525517", patDoc.Representatives[0].IID)
	ass.Equal("80331 München", patDoc.Representatives[0].City)
	ass.Equal("Patent- und Rechtsanwälte PartG mbB \nHermann-Sack-Strasse 3", patDoc.Representatives[0].Street)
	ass.Equal("Peterreins Schley", patDoc.Representatives[0].Name)

	// contracting states
	ass.NotEmpty(patDoc.ContractingStates)
	ass.Equal(38, len(patDoc.ContractingStates))
	for i := 0; i <= 37; i++ {
		ass.Equal(2, len(patDoc.ContractingStates[i]))
	}

	// classifications
	ass.NotEmpty(patDoc.Classifications)
	ass.Equal(21, len(patDoc.Classifications))
	for i := 0; i <= 20; i++ {
		ass.Equal(IPC, patDoc.Classifications[i].System)
		ass.Equal(i+1, patDoc.Classifications[i].Sequence)
		ass.Equal("20060101", patDoc.Classifications[i].Version)
		ass.Equal("A", patDoc.Classifications[i].ClassificationLevel)
		ass.Equal("I", patDoc.Classifications[i].ClassificationValue)
		ass.Equal("20210224", patDoc.Classifications[i].ActionDate)
		ass.Equal("B", patDoc.Classifications[i].OriginalOrReclassified)
		ass.Equal("H", patDoc.Classifications[i].Source)
		ass.Equal("EP", patDoc.Classifications[i].GeneratingOffice)

	}

	ass.Equal("C08F   8/48        20060101AFI20210224BHEP        ", patDoc.Classifications[0].Text)
	ass.Equal("C", patDoc.Classifications[0].Section)
	ass.Equal("08", patDoc.Classifications[0].Class)
	ass.Equal("F", patDoc.Classifications[0].SubClass)
	ass.Equal("8", patDoc.Classifications[0].MainGroup)
	ass.Equal("48", patDoc.Classifications[0].SubGroup)
	ass.Equal("F", patDoc.Classifications[0].FirstLater)

	ass.Equal("C08F   8/28        20060101ALI20210224BHEP        ", patDoc.Classifications[1].Text)
	ass.Equal("C", patDoc.Classifications[1].Section)
	ass.Equal("08", patDoc.Classifications[1].Class)
	ass.Equal("F", patDoc.Classifications[1].SubClass)
	ass.Equal("8", patDoc.Classifications[1].MainGroup)
	ass.Equal("28", patDoc.Classifications[1].SubGroup)
	ass.Equal("L", patDoc.Classifications[1].FirstLater)

	ass.Equal("C08F   8/30        20060101ALI20210224BHEP        ", patDoc.Classifications[2].Text)
	ass.Equal("C", patDoc.Classifications[2].Section)
	ass.Equal("08", patDoc.Classifications[2].Class)
	ass.Equal("F", patDoc.Classifications[2].SubClass)
	ass.Equal("8", patDoc.Classifications[2].MainGroup)
	ass.Equal("30", patDoc.Classifications[2].SubGroup)
	ass.Equal("L", patDoc.Classifications[2].FirstLater)

	ass.Equal("C08F  16/06        20060101ALI20210224BHEP        ", patDoc.Classifications[3].Text)
	ass.Equal("C", patDoc.Classifications[3].Section)
	ass.Equal("08", patDoc.Classifications[3].Class)
	ass.Equal("F", patDoc.Classifications[3].SubClass)
	ass.Equal("16", patDoc.Classifications[3].MainGroup)
	ass.Equal("06", patDoc.Classifications[3].SubGroup)
	ass.Equal("L", patDoc.Classifications[3].FirstLater)

	ass.Equal("C08F  20/06        20060101ALI20210224BHEP        ", patDoc.Classifications[4].Text)
	ass.Equal("C", patDoc.Classifications[4].Section)
	ass.Equal("08", patDoc.Classifications[4].Class)
	ass.Equal("F", patDoc.Classifications[4].SubClass)
	ass.Equal("20", patDoc.Classifications[4].MainGroup)
	ass.Equal("06", patDoc.Classifications[4].SubGroup)
	ass.Equal("L", patDoc.Classifications[4].FirstLater)

	ass.Equal("C08L 101/14        20060101ALI20210224BHEP        ", patDoc.Classifications[5].Text)
	ass.Equal("C", patDoc.Classifications[5].Section)
	ass.Equal("08", patDoc.Classifications[5].Class)
	ass.Equal("L", patDoc.Classifications[5].SubClass)
	ass.Equal("101", patDoc.Classifications[5].MainGroup)
	ass.Equal("14", patDoc.Classifications[5].SubGroup)
	ass.Equal("L", patDoc.Classifications[5].FirstLater)

	ass.Equal("A61K  49/04        20060101ALI20210224BHEP        ", patDoc.Classifications[6].Text)
	ass.Equal("A", patDoc.Classifications[6].Section)
	ass.Equal("61", patDoc.Classifications[6].Class)
	ass.Equal("K", patDoc.Classifications[6].SubClass)
	ass.Equal("49", patDoc.Classifications[6].MainGroup)
	ass.Equal("04", patDoc.Classifications[6].SubGroup)
	ass.Equal("L", patDoc.Classifications[6].FirstLater)

	ass.Equal("C08L  29/04        20060101ALI20210224BHEP        ", patDoc.Classifications[7].Text)
	ass.Equal("C", patDoc.Classifications[7].Section)
	ass.Equal("08", patDoc.Classifications[7].Class)
	ass.Equal("L", patDoc.Classifications[7].SubClass)
	ass.Equal("29", patDoc.Classifications[7].MainGroup)
	ass.Equal("04", patDoc.Classifications[7].SubGroup)
	ass.Equal("L", patDoc.Classifications[7].FirstLater)

	ass.Equal("A61K   9/16        20060101ALI20210224BHEP        ", patDoc.Classifications[8].Text)
	ass.Equal("A", patDoc.Classifications[8].Section)
	ass.Equal("61", patDoc.Classifications[8].Class)
	ass.Equal("K", patDoc.Classifications[8].SubClass)
	ass.Equal("9", patDoc.Classifications[8].MainGroup)
	ass.Equal("16", patDoc.Classifications[8].SubGroup)
	ass.Equal("L", patDoc.Classifications[8].FirstLater)

	ass.Equal("A61K  31/704       20060101ALI20210224BHEP        ", patDoc.Classifications[9].Text)
	ass.Equal("A", patDoc.Classifications[9].Section)
	ass.Equal("61", patDoc.Classifications[9].Class)
	ass.Equal("K", patDoc.Classifications[9].SubClass)
	ass.Equal("31", patDoc.Classifications[9].MainGroup)
	ass.Equal("704", patDoc.Classifications[9].SubGroup)
	ass.Equal("L", patDoc.Classifications[9].FirstLater)

	ass.Equal("C07C  47/575       20060101ALI20210224BHEP        ", patDoc.Classifications[10].Text)
	ass.Equal("C", patDoc.Classifications[10].Section)
	ass.Equal("07", patDoc.Classifications[10].Class)
	ass.Equal("C", patDoc.Classifications[10].SubClass)
	ass.Equal("47", patDoc.Classifications[10].MainGroup)
	ass.Equal("575", patDoc.Classifications[10].SubGroup)
	ass.Equal("L", patDoc.Classifications[10].FirstLater)

	ass.Equal("C07C  47/565       20060101ALI20210224BHEP        ", patDoc.Classifications[11].Text)
	ass.Equal("C", patDoc.Classifications[11].Section)
	ass.Equal("07", patDoc.Classifications[11].Class)
	ass.Equal("C", patDoc.Classifications[11].SubClass)
	ass.Equal("47", patDoc.Classifications[11].MainGroup)
	ass.Equal("565", patDoc.Classifications[11].SubGroup)
	ass.Equal("L", patDoc.Classifications[11].FirstLater)

	ass.Equal("C07C 229/62        20060101ALI20210224BHEP        ", patDoc.Classifications[12].Text)
	ass.Equal("C", patDoc.Classifications[12].Section)
	ass.Equal("07", patDoc.Classifications[12].Class)
	ass.Equal("C", patDoc.Classifications[12].SubClass)
	ass.Equal("229", patDoc.Classifications[12].MainGroup)
	ass.Equal("62", patDoc.Classifications[12].SubGroup)
	ass.Equal("L", patDoc.Classifications[12].FirstLater)

	ass.Equal("C07C 309/11        20060101ALI20210224BHEP        ", patDoc.Classifications[13].Text)
	ass.Equal("C", patDoc.Classifications[13].Section)
	ass.Equal("07", patDoc.Classifications[13].Class)
	ass.Equal("C", patDoc.Classifications[13].SubClass)
	ass.Equal("309", patDoc.Classifications[13].MainGroup)
	ass.Equal("11", patDoc.Classifications[13].SubGroup)
	ass.Equal("L", patDoc.Classifications[13].FirstLater)

	ass.Equal("C08F   2/14        20060101ALI20210224BHEP        ", patDoc.Classifications[14].Text)
	ass.Equal("C", patDoc.Classifications[14].Section)
	ass.Equal("08", patDoc.Classifications[14].Class)
	ass.Equal("F", patDoc.Classifications[14].SubClass)
	ass.Equal("2", patDoc.Classifications[14].MainGroup)
	ass.Equal("14", patDoc.Classifications[14].SubGroup)
	ass.Equal("L", patDoc.Classifications[14].FirstLater)

	ass.Equal("C08F 216/06        20060101ALI20210224BHEP        ", patDoc.Classifications[15].Text)
	ass.Equal("C", patDoc.Classifications[15].Section)
	ass.Equal("08", patDoc.Classifications[15].Class)
	ass.Equal("F", patDoc.Classifications[15].SubClass)
	ass.Equal("216", patDoc.Classifications[15].MainGroup)
	ass.Equal("06", patDoc.Classifications[15].SubGroup)
	ass.Equal("L", patDoc.Classifications[15].FirstLater)

	ass.Equal("C08F 290/12        20060101ALI20210224BHEP        ", patDoc.Classifications[16].Text)
	ass.Equal("C", patDoc.Classifications[16].Section)
	ass.Equal("08", patDoc.Classifications[16].Class)
	ass.Equal("F", patDoc.Classifications[16].SubClass)
	ass.Equal("290", patDoc.Classifications[16].MainGroup)
	ass.Equal("12", patDoc.Classifications[16].SubGroup)
	ass.Equal("L", patDoc.Classifications[16].FirstLater)

	ass.Equal("C08F 220/58        20060101ALI20210224BHEP        ", patDoc.Classifications[17].Text)
	ass.Equal("C", patDoc.Classifications[17].Section)
	ass.Equal("08", patDoc.Classifications[17].Class)
	ass.Equal("F", patDoc.Classifications[17].SubClass)
	ass.Equal("220", patDoc.Classifications[17].MainGroup)
	ass.Equal("58", patDoc.Classifications[17].SubGroup)
	ass.Equal("L", patDoc.Classifications[17].FirstLater)

	ass.Equal("C08L  51/06        20060101ALI20210224BHEP        ", patDoc.Classifications[18].Text)
	ass.Equal("C", patDoc.Classifications[18].Section)
	ass.Equal("08", patDoc.Classifications[18].Class)
	ass.Equal("L", patDoc.Classifications[18].SubClass)
	ass.Equal("51", patDoc.Classifications[18].MainGroup)
	ass.Equal("06", patDoc.Classifications[18].SubGroup)
	ass.Equal("L", patDoc.Classifications[18].FirstLater)

	ass.Equal("C08F 116/06        20060101ALI20210224BHEP        ", patDoc.Classifications[19].Text)
	ass.Equal("C", patDoc.Classifications[19].Section)
	ass.Equal("08", patDoc.Classifications[19].Class)
	ass.Equal("F", patDoc.Classifications[19].SubClass)
	ass.Equal("116", patDoc.Classifications[19].MainGroup)
	ass.Equal("06", patDoc.Classifications[19].SubGroup)
	ass.Equal("L", patDoc.Classifications[19].FirstLater)

	ass.Equal("C07C  45/71        20060101ALI20210224BHEP        ", patDoc.Classifications[20].Text)
	ass.Equal("C", patDoc.Classifications[20].Section)
	ass.Equal("07", patDoc.Classifications[20].Class)
	ass.Equal("C", patDoc.Classifications[20].SubClass)
	ass.Equal("45", patDoc.Classifications[20].MainGroup)
	ass.Equal("71", patDoc.Classifications[20].SubGroup)
	ass.Equal("L", patDoc.Classifications[20].FirstLater)

}

func TestProcessXMLSimple15A21(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("test-data/application/v1-5-A2-1.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("EP19751587A2", patDoc.ID)
	ass.Equal("19751587.7", patDoc.File)
	ass.Equal("en", patDoc.Lang)
	ass.Equal(Country("EP"), patDoc.Country)
	ass.Equal("3749079", patDoc.DocNumber)
	ass.Equal("A2", patDoc.Kind)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20201216", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("n", patDoc.Status)
	ass.Equal("ep-patent-document-v1-5", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("KEIMANALYSE UNTER VERWENDUNG VON RÖNTGENBILDGEBUNG", patDoc.Title[0].Text)
	ass.Equal("de", patDoc.Title[0].Language)
	ass.Equal("SEED ANALYSIS USING X-RAY IMAGING", patDoc.Title[1].Text)
	ass.Equal("en", patDoc.Title[1].Language)
	ass.Equal("ANALYSE DE GRAINES", patDoc.Title[2].Text)
	ass.Equal("fr", patDoc.Title[2].Language)

	// abstract Empty!!
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(0, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// claims
	ass.Empty(patDoc.Claims)

	// description
	//	ass.Empty(patDoc.Description)

	// citations
	ass.Empty(patDoc.Citations)

	// Inventors
	ass.NotEmpty(patDoc.Inventors)
	ass.Equal(2, len(patDoc.Inventors))

	ass.Equal(Country("US"), patDoc.Inventors[0].Country)
	ass.Equal("Saint Louis, Missouri 63167", patDoc.Inventors[0].City)
	ass.Equal("c/o Monsanto Company\n800 North Lindbergh Boulevard\nMail Zone E1NA", patDoc.Inventors[0].Street)
	ass.Equal("CHEN, Hsin-Chen", patDoc.Inventors[0].Name)
	ass.Equal(Country("US"), patDoc.Inventors[1].Country)
	ass.Equal("Saint Louis, Missouri 63167", patDoc.Inventors[1].City)
	ass.Equal("c/o Monsanto Company\n800 North Lindbergh Boulevard\nMail Zone E1NA", patDoc.Inventors[1].Street)
	ass.Equal("KOTYK, Johnny", patDoc.Inventors[1].Name)

	// representatives
	ass.NotEmpty(patDoc.Representatives)
	ass.Equal(Country("DE"), patDoc.Representatives[0].Country)
	ass.Equal("101467007", patDoc.Representatives[0].IID)
	ass.Equal("40789 Monheim am Rhein", patDoc.Representatives[0].City)
	ass.Equal("c/o Bayer Intellectual Property GmbH \nAlfred-Nobel-Straße 10", patDoc.Representatives[0].Street)
	ass.Equal("BIP Patents", patDoc.Representatives[0].Name)

	// contracting states
	ass.NotEmpty(patDoc.ContractingStates)
	ass.Equal(38, len(patDoc.ContractingStates))
	for i := 0; i <= 37; i++ {
		ass.Equal(2, len(patDoc.ContractingStates[i]))
	}

	// classifications
	ass.NotEmpty(patDoc.Classifications)
	ass.Equal(8, len(patDoc.Classifications))
	for i := 0; i <= 7; i++ {
		ass.Equal(IPC, patDoc.Classifications[i].System)
		ass.Equal(i+1, patDoc.Classifications[i].Sequence)
		ass.Equal("A", patDoc.Classifications[i].ClassificationLevel)
		ass.Equal("I", patDoc.Classifications[i].ClassificationValue)
		ass.Equal("20200807", patDoc.Classifications[i].ActionDate)
		ass.Equal("B", patDoc.Classifications[i].OriginalOrReclassified)
		ass.Equal("H", patDoc.Classifications[i].Source)
		ass.Equal("EP", patDoc.Classifications[i].GeneratingOffice)

	}

	ass.Equal("A01C   1/02        20060101AFI20200807BHEP        ", patDoc.Classifications[0].Text)
	ass.Equal("A", patDoc.Classifications[0].Section)
	ass.Equal("01", patDoc.Classifications[0].Class)
	ass.Equal("C", patDoc.Classifications[0].SubClass)
	ass.Equal("20060101", patDoc.Classifications[0].Version)
	ass.Equal("1", patDoc.Classifications[0].MainGroup)
	ass.Equal("02", patDoc.Classifications[0].SubGroup)
	ass.Equal("F", patDoc.Classifications[0].FirstLater)

	ass.Equal("A01C   1/04        20060101ALI20200807BHEP        ", patDoc.Classifications[1].Text)
	ass.Equal("A", patDoc.Classifications[1].Section)
	ass.Equal("01", patDoc.Classifications[1].Class)
	ass.Equal("C", patDoc.Classifications[1].SubClass)
	ass.Equal("20060101", patDoc.Classifications[1].Version)
	ass.Equal("1", patDoc.Classifications[1].MainGroup)
	ass.Equal("04", patDoc.Classifications[1].SubGroup)
	ass.Equal("L", patDoc.Classifications[1].FirstLater)

	ass.Equal("B07B  13/04        20060101ALI20200807BHEP        ", patDoc.Classifications[2].Text)
	ass.Equal("B", patDoc.Classifications[2].Section)
	ass.Equal("07", patDoc.Classifications[2].Class)
	ass.Equal("B", patDoc.Classifications[2].SubClass)
	ass.Equal("20060101", patDoc.Classifications[2].Version)
	ass.Equal("13", patDoc.Classifications[2].MainGroup)
	ass.Equal("04", patDoc.Classifications[2].SubGroup)
	ass.Equal("L", patDoc.Classifications[2].FirstLater)

	ass.Equal("B07B  13/08        20060101ALI20200807BHEP        ", patDoc.Classifications[3].Text)
	ass.Equal("B", patDoc.Classifications[3].Section)
	ass.Equal("07", patDoc.Classifications[3].Class)
	ass.Equal("B", patDoc.Classifications[3].SubClass)
	ass.Equal("20060101", patDoc.Classifications[3].Version)
	ass.Equal("13", patDoc.Classifications[3].MainGroup)
	ass.Equal("08", patDoc.Classifications[3].SubGroup)
	ass.Equal("L", patDoc.Classifications[3].FirstLater)

	ass.Equal("B07C   5/342       20060101ALI20200807BHEP        ", patDoc.Classifications[4].Text)
	ass.Equal("B", patDoc.Classifications[4].Section)
	ass.Equal("07", patDoc.Classifications[4].Class)
	ass.Equal("C", patDoc.Classifications[4].SubClass)
	ass.Equal("20060101", patDoc.Classifications[4].Version)
	ass.Equal("5", patDoc.Classifications[4].MainGroup)
	ass.Equal("342", patDoc.Classifications[4].SubGroup)
	ass.Equal("L", patDoc.Classifications[4].FirstLater)

	ass.Equal("G01N  23/083       20180101ALI20200807BHEP        ", patDoc.Classifications[5].Text)
	ass.Equal("G", patDoc.Classifications[5].Section)
	ass.Equal("01", patDoc.Classifications[5].Class)
	ass.Equal("N", patDoc.Classifications[5].SubClass)
	ass.Equal("20180101", patDoc.Classifications[5].Version)
	ass.Equal("23", patDoc.Classifications[5].MainGroup)
	ass.Equal("083", patDoc.Classifications[5].SubGroup)
	ass.Equal("L", patDoc.Classifications[5].FirstLater)

	ass.Equal("G06T   7/50        20170101ALI20200807BHEP        ", patDoc.Classifications[6].Text)
	ass.Equal("G", patDoc.Classifications[6].Section)
	ass.Equal("06", patDoc.Classifications[6].Class)
	ass.Equal("T", patDoc.Classifications[6].SubClass)
	ass.Equal("20170101", patDoc.Classifications[6].Version)
	ass.Equal("7", patDoc.Classifications[6].MainGroup)
	ass.Equal("50", patDoc.Classifications[6].SubGroup)
	ass.Equal("L", patDoc.Classifications[6].FirstLater)

	ass.Equal("G06T   7/60        20170101ALI20200807BHEP        ", patDoc.Classifications[7].Text)
	ass.Equal("G", patDoc.Classifications[7].Section)
	ass.Equal("06", patDoc.Classifications[7].Class)
	ass.Equal("T", patDoc.Classifications[7].SubClass)
	ass.Equal("20170101", patDoc.Classifications[7].Version)
	ass.Equal("7", patDoc.Classifications[7].MainGroup)
	ass.Equal("60", patDoc.Classifications[7].SubGroup)
	ass.Equal("L", patDoc.Classifications[7].FirstLater)

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
	ass.NotEmpty(patDoc.Classifications)
	ass.Equal(2, len(patDoc.Classifications))

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

}

func TestProcessXMLSimple15B11(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("test-data/grant/v1-5-B1-1.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("EP02256108B1", patDoc.ID)
	ass.Equal("EP02256108NWB1.xml", patDoc.File)
	ass.Equal("en", patDoc.Lang)
	ass.Equal(Country("EP"), patDoc.Country)
	ass.Equal("1291416", patDoc.DocNumber)
	ass.Equal("B1", patDoc.Kind)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20201223", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("n", patDoc.Status)
	ass.Equal("ep-patent-document-v1-5", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("Fructosyl-Peptidoxidase", patDoc.Title[0].Text)
	ass.Equal("de", patDoc.Title[0].Language)
	ass.Equal("Fructosyl peptide oxidase", patDoc.Title[1].Text)
	ass.Equal("en", patDoc.Title[1].Language)
	ass.Equal("Fructosyle peptide oxidase", patDoc.Title[2].Text)
	ass.Equal("fr", patDoc.Title[2].Language)

	// abstract !! Empty
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(0, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// claims
	ass.NotEmpty(patDoc.Claims)
	ass.Equal(3, len(patDoc.Claims))
	ass.Equal(4736, len(patDoc.Claims[0].Text))
	ass.Equal("en", patDoc.Claims[0].Language)
	ass.Equal("claims01", patDoc.Claims[0].Id)
	ass.Equal(4787, len(patDoc.Claims[1].Text))
	ass.Equal("de", patDoc.Claims[1].Language)
	ass.Equal("claims02", patDoc.Claims[1].Id)
	ass.Equal(5562, len(patDoc.Claims[2].Text))
	ass.Equal("fr", patDoc.Claims[2].Language)
	ass.Equal("claims03", patDoc.Claims[2].Id)

	// description
	ass.NotEmpty(patDoc.Description)
	ass.Equal(108637, len(patDoc.Description[0].Text))
	ass.Equal("en", patDoc.Description[0].Language)

	// citations
	ass.NotEmpty(patDoc.Citations)
	ass.Equal(Country("JP"), patDoc.Citations[0].Country)
	ass.Equal("5033997", patDoc.Citations[0].DocNumber)
	ass.Equal("A", patDoc.Citations[0].Kind)
	ass.Equal(Country("JP"), patDoc.Citations[1].Country)
	ass.Equal("11127895", patDoc.Citations[1].DocNumber)
	ass.Equal("A", patDoc.Citations[1].Kind)
	ass.Equal(Country("WO"), patDoc.Citations[2].Country)
	ass.Equal("9713872", patDoc.Citations[2].DocNumber)
	ass.Equal("A", patDoc.Citations[2].Kind)

	// Inventors
	ass.NotEmpty(patDoc.Inventors)
	ass.Equal(Country("JP"), patDoc.Inventors[0].Country)
	ass.Equal("Noda-shi,\nChiba 278-8601", patDoc.Inventors[0].City)
	ass.Equal("250, Noda", patDoc.Inventors[0].Street)
	ass.Equal("Kurosawa, Keiko,\nc/o Kikkoman Corporation", patDoc.Inventors[0].Name)

	ass.Equal(Country("JP"), patDoc.Inventors[1].Country)
	ass.Equal("Noda-shi,\nChiba 278-8601", patDoc.Inventors[1].City)
	ass.Equal("250, Noda", patDoc.Inventors[1].Street)
	ass.Equal("Hirokawa, Kozo,\nc/o Kikkoman Corporation", patDoc.Inventors[1].Name)

	// representatives
	ass.NotEmpty(patDoc.Representatives)
	ass.Equal(Country("GB"), patDoc.Representatives[0].Country)
	ass.Equal("101783151", patDoc.Representatives[0].IID)
	ass.Equal("Bristol BS1 6BX", patDoc.Representatives[0].City)
	ass.Equal("Aurora Building \nCounterslip", patDoc.Representatives[0].Street)
	ass.Equal("Mewburn Ellis LLP", patDoc.Representatives[0].Name)

	// contracting states
	ass.NotEmpty(patDoc.ContractingStates)
	ass.Equal(3, len(patDoc.ContractingStates))
	for i := 0; i <= 2; i++ {
		ass.Equal(2, len(patDoc.ContractingStates[i]))
	}

	// classifications
	ass.NotEmpty(patDoc.Classifications)
	ass.Equal(5, len(patDoc.Classifications))

	for i := 0; i <= 4; i++ {
		ass.Equal(IPC, patDoc.Classifications[i].System)
		ass.Equal(i+1, patDoc.Classifications[i].Sequence)
		ass.Equal("20060101", patDoc.Classifications[i].Version)
		ass.Equal("A", patDoc.Classifications[i].ClassificationLevel)
		ass.Equal("I", patDoc.Classifications[i].ClassificationValue)
		ass.Equal("20030117", patDoc.Classifications[i].ActionDate)
		ass.Equal("B", patDoc.Classifications[i].OriginalOrReclassified)
		ass.Equal("H", patDoc.Classifications[i].Source)
		ass.Equal("EP", patDoc.Classifications[i].GeneratingOffice)

	}

	ass.Equal("C12N   9/06        20060101AFI20030117BHEP        ", patDoc.Classifications[0].Text)
	ass.Equal("C", patDoc.Classifications[0].Section)
	ass.Equal("12", patDoc.Classifications[0].Class)
	ass.Equal("N", patDoc.Classifications[0].SubClass)
	ass.Equal("9", patDoc.Classifications[0].MainGroup)
	ass.Equal("06", patDoc.Classifications[0].SubGroup)
	ass.Equal("F", patDoc.Classifications[0].FirstLater)

	ass.Equal("C12N  15/53        20060101ALI20030117BHEP        ", patDoc.Classifications[1].Text)
	ass.Equal("C", patDoc.Classifications[1].Section)
	ass.Equal("12", patDoc.Classifications[1].Class)
	ass.Equal("N", patDoc.Classifications[1].SubClass)
	ass.Equal("15", patDoc.Classifications[1].MainGroup)
	ass.Equal("53", patDoc.Classifications[1].SubGroup)
	ass.Equal("L", patDoc.Classifications[1].FirstLater)

	ass.Equal("C12N  15/63        20060101ALI20030117BHEP        ", patDoc.Classifications[2].Text)
	ass.Equal("C", patDoc.Classifications[2].Section)
	ass.Equal("12", patDoc.Classifications[2].Class)
	ass.Equal("N", patDoc.Classifications[2].SubClass)
	ass.Equal("15", patDoc.Classifications[2].MainGroup)
	ass.Equal("63", patDoc.Classifications[2].SubGroup)
	ass.Equal("L", patDoc.Classifications[2].FirstLater)

	ass.Equal("C12N   1/15        20060101ALI20030117BHEP        ", patDoc.Classifications[3].Text)
	ass.Equal("C", patDoc.Classifications[3].Section)
	ass.Equal("12", patDoc.Classifications[3].Class)
	ass.Equal("N", patDoc.Classifications[3].SubClass)
	ass.Equal("1", patDoc.Classifications[3].MainGroup)
	ass.Equal("15", patDoc.Classifications[3].SubGroup)
	ass.Equal("L", patDoc.Classifications[3].FirstLater)

	ass.Equal("C12P  21/02        20060101ALI20030117BHEP        ", patDoc.Classifications[4].Text)
	ass.Equal("C", patDoc.Classifications[4].Section)
	ass.Equal("12", patDoc.Classifications[4].Class)
	ass.Equal("P", patDoc.Classifications[4].SubClass)
	ass.Equal("21", patDoc.Classifications[4].MainGroup)
	ass.Equal("02", patDoc.Classifications[4].SubGroup)
	ass.Equal("L", patDoc.Classifications[4].FirstLater)

}

func TestProcessXMLSimple15B2(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("test-data/grant/v1-5-B2.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("EP13189031B2", patDoc.ID)
	ass.Equal("EP13189031NWB2.xml", patDoc.File)
	ass.Equal("de", patDoc.Lang)
	ass.Equal(Country("EP"), patDoc.Country)
	ass.Equal("2743087", patDoc.DocNumber)
	ass.Equal("B2", patDoc.Kind)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20210505", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("n", patDoc.Status)
	ass.Equal("ep-patent-document-v1-5", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("Druckmaschine mit Verkleidung", patDoc.Title[0].Text)
	ass.Equal("de", patDoc.Title[0].Language)
	ass.Equal("Printing machine with cladding", patDoc.Title[1].Text)
	ass.Equal("en", patDoc.Title[1].Language)
	ass.Equal("Imprimante avec habillage", patDoc.Title[2].Text)
	ass.Equal("fr", patDoc.Title[2].Language)

	// abstract !! Empty
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(0, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// claims
	ass.NotEmpty(patDoc.Claims)
	ass.Equal(3, len(patDoc.Claims))
	ass.Equal(4910, len(patDoc.Claims[0].Text))
	ass.Equal("de", patDoc.Claims[0].Language)
	ass.Equal("claims01", patDoc.Claims[0].Id)
	ass.Equal(4748, len(patDoc.Claims[1].Text))
	ass.Equal("en", patDoc.Claims[1].Language)
	ass.Equal("claims02", patDoc.Claims[1].Id)
	ass.Equal(5622, len(patDoc.Claims[2].Text))
	ass.Equal("fr", patDoc.Claims[2].Language)
	ass.Equal("claims03", patDoc.Claims[2].Id)

	// description
	ass.NotEmpty(patDoc.Description)
	ass.Equal(29873, len(patDoc.Description[0].Text))
	ass.Equal("en", patDoc.Description[0].Language)

	// citations
	ass.NotEmpty(patDoc.Citations)
	ass.Equal(Country("DE"), patDoc.Citations[0].Country)
	ass.Equal("102009041527", patDoc.Citations[0].DocNumber)
	ass.Equal("A1", patDoc.Citations[0].Kind)

	// Inventors
	ass.NotEmpty(patDoc.Inventors)
	ass.Equal(Country("DE"), patDoc.Inventors[0].Country)
	ass.Equal("93073 Neutraubling", patDoc.Inventors[0].City)
	ass.Equal("Böhmerwaldstrasse 5", patDoc.Inventors[0].Street)
	ass.Equal("Gertlowski, Georg", patDoc.Inventors[0].Name)

	ass.Equal(Country("DE"), patDoc.Inventors[1].Country)
	ass.Equal("93073 Neutraubling", patDoc.Inventors[1].City)
	ass.Equal("Böhmerwaldstrasse 5", patDoc.Inventors[1].Street)
	ass.Equal("Winzinger, Frank", patDoc.Inventors[1].Name)

	// representatives
	ass.NotEmpty(patDoc.Representatives)
	ass.Equal(Country("DE"), patDoc.Representatives[0].Country)
	ass.Equal("100060488", patDoc.Representatives[0].IID)
	ass.Equal("80802 München", patDoc.Representatives[0].City)
	ass.Equal("Leopoldstraße 4", patDoc.Representatives[0].Street)
	ass.Equal("Grünecker Patent- und Rechtsanwälte \nPartG mbB", patDoc.Representatives[0].Name)

	// contracting states
	ass.NotEmpty(patDoc.ContractingStates)
	ass.Equal(38, len(patDoc.ContractingStates))
	for i := 0; i <= 37; i++ {
		ass.Equal(2, len(patDoc.ContractingStates[i]))
	}

	// classifications

	ass.NotEmpty(patDoc.Classifications)
	ass.Equal(3, len(patDoc.Classifications))
	for i := 0; i <= 2; i++ {
		ass.Equal(IPC, patDoc.Classifications[i].System)
		ass.Equal(i+1, patDoc.Classifications[i].Sequence)
		ass.Equal("B", patDoc.Classifications[i].Section)
		ass.Equal("20060101", patDoc.Classifications[i].Version)
		ass.Equal("A", patDoc.Classifications[i].ClassificationLevel)
		ass.Equal("I", patDoc.Classifications[i].ClassificationValue)
		ass.Equal("20140331", patDoc.Classifications[i].ActionDate)
		ass.Equal("B", patDoc.Classifications[i].OriginalOrReclassified)
		ass.Equal("H", patDoc.Classifications[i].Source)
		ass.Equal("EP", patDoc.Classifications[i].GeneratingOffice)
	}

	ass.Equal("B41J   3/407       20060101AFI20140331BHEP        ", patDoc.Classifications[0].Text)
	ass.Equal("41", patDoc.Classifications[0].Class)
	ass.Equal("J", patDoc.Classifications[0].SubClass)
	ass.Equal("3", patDoc.Classifications[0].MainGroup)
	ass.Equal("407", patDoc.Classifications[0].SubGroup)
	ass.Equal("F", patDoc.Classifications[0].FirstLater)

	ass.Equal("B41J  29/13        20060101ALI20140331BHEP        ", patDoc.Classifications[1].Text)
	ass.Equal("41", patDoc.Classifications[1].Class)
	ass.Equal("J", patDoc.Classifications[1].SubClass)
	ass.Equal("29", patDoc.Classifications[1].MainGroup)
	ass.Equal("13", patDoc.Classifications[1].SubGroup)
	ass.Equal("L", patDoc.Classifications[1].FirstLater)

	ass.Equal("B41J  29/377       20060101ALI20140331BHEP        ", patDoc.Classifications[2].Text)
	ass.Equal("41", patDoc.Classifications[2].Class)
	ass.Equal("J", patDoc.Classifications[2].SubClass)
	ass.Equal("29", patDoc.Classifications[2].MainGroup)
	ass.Equal("377", patDoc.Classifications[2].SubGroup)
	ass.Equal("L", patDoc.Classifications[2].FirstLater)

}

func TestProcessXMLSimple15B21(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("test-data/grant/v1-5-B2-1.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("EP98942751B2", patDoc.ID)
	ass.Equal("EP98942751NWB2.xml", patDoc.File)
	ass.Equal("fr", patDoc.Lang)
	ass.Equal(Country("EP"), patDoc.Country)
	ass.Equal("1011858", patDoc.DocNumber)
	ass.Equal("B2", patDoc.Kind)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20201216", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("n", patDoc.Status)
	ass.Equal("ep-patent-document-v1-5", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("AGGLOMERIERTE ZEOLITISCHE ADSORPTIONSMITTEL,  VERFAHREN ZU IHRER HERSTELLUNG UND IHRE VERWENDUNG BEI DER ADSORPTION VON PARA-XYLOL AUS C8 AROMATISCHEN KOHLENWASSERSTOFFSCHNITTEN", patDoc.Title[0].Text)
	ass.Equal("de", patDoc.Title[0].Language)
	ass.Equal("AGGLOMERATED ZEOLITE ADSORBENTS, METHOD FOR OBTAINING THEM AND USE FOR ADSORBING PARAXYLENE FROM AROMATIC C8 CUT FRACTIONS", patDoc.Title[1].Text)
	ass.Equal("en", patDoc.Title[1].Language)
	ass.Equal("ADSORBANTS ZEOLITIQUES AGGLOMERES LEUR PROCEDE D'OBTENTION ET LEUR UTILISATION POUR L'ADSORPTION DE PARAXYLENE A PARTIR DE COUPES DE C8 AROMATIQUES", patDoc.Title[2].Text)
	ass.Equal("fr", patDoc.Title[2].Language)

	// abstract !! Empty
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(0, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// claims
	ass.NotEmpty(patDoc.Claims)
	ass.Equal(3, len(patDoc.Claims))
	ass.Equal(2592, len(patDoc.Claims[0].Text))
	ass.Equal("fr", patDoc.Claims[0].Language)
	ass.Equal("claims01", patDoc.Claims[0].Id)
	ass.Equal(2707, len(patDoc.Claims[1].Text))
	ass.Equal("de", patDoc.Claims[1].Language)
	ass.Equal("claims02", patDoc.Claims[1].Id)
	ass.Equal(2572, len(patDoc.Claims[2].Text))
	ass.Equal("en", patDoc.Claims[2].Language)
	ass.Equal("claims03", patDoc.Claims[2].Id)

	// description
	ass.NotEmpty(patDoc.Description)
	ass.Equal(22392, len(patDoc.Description[0].Text))
	ass.Equal("en", patDoc.Description[0].Language)

	// citations
	ass.NotEmpty(patDoc.Citations)
	ass.Equal(Country("US"), patDoc.Citations[0].Country)
	ass.Equal("3558730", patDoc.Citations[0].DocNumber)
	ass.Equal("A", patDoc.Citations[0].Kind)

	// Inventors
	ass.NotEmpty(patDoc.Inventors)
	ass.Equal(Country("FR"), patDoc.Inventors[0].Country)
	ass.Equal("F-64140 Lons", patDoc.Inventors[0].City)
	ass.Equal("3, allée du Martin-Pêcheur", patDoc.Inventors[0].Street)
	ass.Equal("PLEE, Dominique", patDoc.Inventors[0].Name)

	ass.Equal(Country("FR"), patDoc.Inventors[1].Country)
	ass.Equal("F-92500 Rueil-Malmaison", patDoc.Inventors[1].City)
	ass.Equal("25, boulevard Coteaux", patDoc.Inventors[1].Street)
	ass.Equal("METHIVIER, Alain", patDoc.Inventors[1].Name)

	// representatives
	ass.NotEmpty(patDoc.Representatives)
	ass.Equal(Country("FR"), patDoc.Representatives[0].Country)
	ass.Equal("101457544", patDoc.Representatives[0].IID)
	ass.Equal("75002 Paris", patDoc.Representatives[0].City)
	ass.Equal("30, rue Notre-Dame des Victoires", patDoc.Representatives[0].Street)
	ass.Equal("Bandpay & Greuter", patDoc.Representatives[0].Name)

	// contracting states
	ass.NotEmpty(patDoc.ContractingStates)
	ass.Equal(6, len(patDoc.ContractingStates))
	for i := 0; i <= 5; i++ {
		ass.Equal(2, len(patDoc.ContractingStates[i]))
	}

	// classifications

	ass.NotEmpty(patDoc.Classifications)
	ass.Equal(3, len(patDoc.Classifications))
	for i := 0; i <= 2; i++ {
		ass.Equal(IPC, patDoc.Classifications[i].System)
		ass.Equal(i+1, patDoc.Classifications[i].Sequence)
		ass.Equal("20060101", patDoc.Classifications[i].Version)
		ass.Equal("A", patDoc.Classifications[i].ClassificationLevel)
		ass.Equal("I", patDoc.Classifications[i].ClassificationValue)
		ass.Equal("19990331", patDoc.Classifications[i].ActionDate)
		ass.Equal("B", patDoc.Classifications[i].OriginalOrReclassified)
		ass.Equal("H", patDoc.Classifications[i].Source)
		ass.Equal("EP", patDoc.Classifications[i].GeneratingOffice)
	}

	ass.Equal("B01J  20/18        20060101AFI19990331BHEP        ", patDoc.Classifications[0].Text)
	ass.Equal("01", patDoc.Classifications[0].Class)
	ass.Equal("J", patDoc.Classifications[0].SubClass)
	ass.Equal("B", patDoc.Classifications[0].Section)
	ass.Equal("20", patDoc.Classifications[0].MainGroup)
	ass.Equal("18", patDoc.Classifications[0].SubGroup)
	ass.Equal("F", patDoc.Classifications[0].FirstLater)

	ass.Equal("C07C  15/08        20060101ALI19990331BHEP        ", patDoc.Classifications[1].Text)
	ass.Equal("07", patDoc.Classifications[1].Class)
	ass.Equal("C", patDoc.Classifications[1].SubClass)
	ass.Equal("C", patDoc.Classifications[1].Section)
	ass.Equal("15", patDoc.Classifications[1].MainGroup)
	ass.Equal("08", patDoc.Classifications[1].SubGroup)
	ass.Equal("L", patDoc.Classifications[1].FirstLater)

	ass.Equal("C07C   7/13        20060101ALI19990331BHEP        ", patDoc.Classifications[2].Text)
	ass.Equal("07", patDoc.Classifications[2].Class)
	ass.Equal("C", patDoc.Classifications[2].SubClass)
	ass.Equal("C", patDoc.Classifications[2].Section)
	ass.Equal("7", patDoc.Classifications[2].MainGroup)
	ass.Equal("13", patDoc.Classifications[2].SubGroup)
	ass.Equal("L", patDoc.Classifications[2].FirstLater)
}

//v 1.5.1
func TestProcessXMLSimple151A1(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("test-data/application/v1-5-1-A1.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("EP19883405A1", patDoc.ID)
	ass.Equal("EP19883405NWA1.xml", patDoc.File)
	ass.Equal("en", patDoc.Lang)
	ass.Equal(Country("EP"), patDoc.Country)
	ass.Equal("3882239", patDoc.DocNumber)
	ass.Equal("A1", patDoc.Kind)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20210922", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("n", patDoc.Status)
	ass.Equal("ep-patent-document-v1-5-1", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("1,3,4-OXADIAZOLON-VERBINDUNGEN UND MEDIZIN", patDoc.Title[0].Text)
	ass.Equal("de", patDoc.Title[0].Language)
	ass.Equal("1,3,4-OXADIAZOLONE COMPOUND AND MEDICINE", patDoc.Title[1].Text)
	ass.Equal("en", patDoc.Title[1].Language)
	ass.Equal("COMPOSÉ DE 1,3,4-OXADIAZOLONE ET MÉDICAMENT", patDoc.Title[2].Text)
	ass.Equal("fr", patDoc.Title[2].Language)

	// abstract
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(553, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// claims
	ass.NotEmpty(patDoc.Claims)
	ass.Equal(1, len(patDoc.Claims))
	ass.Equal(37091, len(patDoc.Claims[0].Text))
	ass.Equal("en", patDoc.Claims[0].Language)
	ass.Equal("claims01", patDoc.Claims[0].Id)

	// description
	ass.NotEmpty(patDoc.Description)
	ass.Equal(416344, len(patDoc.Description[0].Text))
	ass.Equal("en", patDoc.Description[0].Language)

	// Inventors
	ass.NotEmpty(patDoc.Inventors)

	ass.Equal("KAMITANI Hirotaka", patDoc.Inventors[0].Name)
	ass.Equal("Kamakura-shi, Kanagawa 247-0056", patDoc.Inventors[0].City)
	ass.Equal("101 La Vie Kamakura Ofuna, 6-5-63, Ofuna", patDoc.Inventors[0].Street)

	ass.Equal("ZAIMOKU Hisaaki", patDoc.Inventors[1].Name)
	ass.Equal("Kyoto-shi, Kyoto 601-8550", patDoc.Inventors[1].City)
	ass.Equal("c/o NIPPON SHINYAKU CO., LTD., 14, Kisshoin  \nNishinosho Monguchicho, Minami-ku", patDoc.Inventors[1].Street)

	ass.Equal("HARUTA Yoshinari", patDoc.Inventors[2].Name)
	ass.Equal("Kyoto-shi, Kyoto 601-8550", patDoc.Inventors[2].City)
	ass.Equal("c/o NIPPON SHINYAKU CO., LTD., 14, Kisshoin  \nNishinosho Monguchicho, Minami-ku", patDoc.Inventors[2].Street)

	ass.Equal("KIKUCHI Takeo", patDoc.Inventors[3].Name)
	ass.Equal("Kyoto-shi, Kyoto 601-8550", patDoc.Inventors[3].City)
	ass.Equal("c/o NIPPON SHINYAKU CO., LTD., 14, Kisshoin  \nNishinosho Monguchicho, Minami-ku", patDoc.Inventors[3].Street)

	// representatives
	ass.NotEmpty(patDoc.Representatives)
	ass.Equal(Country("DE"), patDoc.Representatives[0].Country)
	ass.Equal("101361216", patDoc.Representatives[0].IID)
	ass.Equal("81675 München", patDoc.Representatives[0].City)
	ass.Equal("Siebertstraße 3", patDoc.Representatives[0].Street)
	ass.Equal("Vossius & Partner \nPatentanwälte Rechtsanwälte mbB", patDoc.Representatives[0].Name)

	// contracting states
	ass.NotEmpty(patDoc.ContractingStates)
	ass.Equal(38, len(patDoc.ContractingStates))
	for i := 0; i <= 37; i++ {
		ass.Equal(2, len(patDoc.ContractingStates[i]))
	}

	// classifications
	ass.NotEmpty(patDoc.Classifications)
	ass.Equal(14, len(patDoc.Classifications))
	for i := 0; i <= 13; i++ {
		ass.Equal(IPC, patDoc.Classifications[i].System)
		ass.Equal(i+1, patDoc.Classifications[i].Sequence)
		ass.Equal("20060101", patDoc.Classifications[i].Version)
		ass.Equal("A", patDoc.Classifications[i].ClassificationLevel)
		ass.Equal("I", patDoc.Classifications[i].ClassificationValue)
		ass.Equal("20200522", patDoc.Classifications[i].ActionDate)
		ass.Equal("B", patDoc.Classifications[i].OriginalOrReclassified)
		ass.Equal("H", patDoc.Classifications[i].Source)
		ass.Equal("EP", patDoc.Classifications[i].GeneratingOffice)

	}

	ass.Equal("C07D 271/113       20060101AFI20200522BHEP        ", patDoc.Classifications[0].Text)
	ass.Equal("C", patDoc.Classifications[0].Section)
	ass.Equal("07", patDoc.Classifications[0].Class)
	ass.Equal("D", patDoc.Classifications[0].SubClass)
	ass.Equal("271", patDoc.Classifications[0].MainGroup)
	ass.Equal("113", patDoc.Classifications[0].SubGroup)
	ass.Equal("F", patDoc.Classifications[0].FirstLater)

	ass.Equal("A61K  31/4245      20060101ALI20200522BHEP        ", patDoc.Classifications[1].Text)
	ass.Equal("A", patDoc.Classifications[1].Section)
	ass.Equal("61", patDoc.Classifications[1].Class)
	ass.Equal("K", patDoc.Classifications[1].SubClass)
	ass.Equal("31", patDoc.Classifications[1].MainGroup)
	ass.Equal("4245", patDoc.Classifications[1].SubGroup)
	ass.Equal("L", patDoc.Classifications[1].FirstLater)

	ass.Equal("A61K  31/454       20060101ALI20200522BHEP        ", patDoc.Classifications[2].Text)
	ass.Equal("A", patDoc.Classifications[2].Section)
	ass.Equal("61", patDoc.Classifications[2].Class)
	ass.Equal("K", patDoc.Classifications[2].SubClass)
	ass.Equal("31", patDoc.Classifications[2].MainGroup)
	ass.Equal("454", patDoc.Classifications[2].SubGroup)
	ass.Equal("L", patDoc.Classifications[2].FirstLater)

	ass.Equal("A61K  31/496       20060101ALI20200522BHEP        ", patDoc.Classifications[3].Text)
	ass.Equal("A", patDoc.Classifications[3].Section)
	ass.Equal("61", patDoc.Classifications[3].Class)
	ass.Equal("K", patDoc.Classifications[3].SubClass)
	ass.Equal("31", patDoc.Classifications[3].MainGroup)
	ass.Equal("496", patDoc.Classifications[3].SubGroup)
	ass.Equal("L", patDoc.Classifications[3].FirstLater)

	ass.Equal("A61K  31/538       20060101ALI20200522BHEP        ", patDoc.Classifications[4].Text)
	ass.Equal("A", patDoc.Classifications[4].Section)
	ass.Equal("61", patDoc.Classifications[4].Class)
	ass.Equal("K", patDoc.Classifications[4].SubClass)
	ass.Equal("31", patDoc.Classifications[4].MainGroup)
	ass.Equal("538", patDoc.Classifications[4].SubGroup)
	ass.Equal("L", patDoc.Classifications[4].FirstLater)

	ass.Equal("A61K  31/55        20060101ALI20200522BHEP        ", patDoc.Classifications[5].Text)
	ass.Equal("A", patDoc.Classifications[5].Section)
	ass.Equal("61", patDoc.Classifications[5].Class)
	ass.Equal("K", patDoc.Classifications[5].SubClass)
	ass.Equal("31", patDoc.Classifications[5].MainGroup)
	ass.Equal("55", patDoc.Classifications[5].SubGroup)
	ass.Equal("L", patDoc.Classifications[5].FirstLater)

	ass.Equal("A61K  31/553       20060101ALI20200522BHEP        ", patDoc.Classifications[6].Text)
	ass.Equal("A", patDoc.Classifications[6].Section)
	ass.Equal("61", patDoc.Classifications[6].Class)
	ass.Equal("K", patDoc.Classifications[6].SubClass)
	ass.Equal("31", patDoc.Classifications[6].MainGroup)
	ass.Equal("553", patDoc.Classifications[6].SubGroup)
	ass.Equal("L", patDoc.Classifications[6].FirstLater)

	ass.Equal("A61P  37/00        20060101ALI20200522BHEP        ", patDoc.Classifications[7].Text)
	ass.Equal("A", patDoc.Classifications[7].Section)
	ass.Equal("61", patDoc.Classifications[7].Class)
	ass.Equal("P", patDoc.Classifications[7].SubClass)
	ass.Equal("37", patDoc.Classifications[7].MainGroup)
	ass.Equal("00", patDoc.Classifications[7].SubGroup)
	ass.Equal("L", patDoc.Classifications[7].FirstLater)

	ass.Equal("A61P  43/00        20060101ALI20200522BHEP        ", patDoc.Classifications[8].Text)
	ass.Equal("A", patDoc.Classifications[8].Section)
	ass.Equal("61", patDoc.Classifications[8].Class)
	ass.Equal("P", patDoc.Classifications[8].SubClass)
	ass.Equal("43", patDoc.Classifications[8].MainGroup)
	ass.Equal("00", patDoc.Classifications[8].SubGroup)
	ass.Equal("L", patDoc.Classifications[8].FirstLater)

	ass.Equal("C07D 413/04        20060101ALI20200522BHEP        ", patDoc.Classifications[9].Text)
	ass.Equal("C", patDoc.Classifications[9].Section)
	ass.Equal("07", patDoc.Classifications[9].Class)
	ass.Equal("D", patDoc.Classifications[9].SubClass)
	ass.Equal("413", patDoc.Classifications[9].MainGroup)
	ass.Equal("04", patDoc.Classifications[9].SubGroup)
	ass.Equal("L", patDoc.Classifications[9].FirstLater)

	ass.Equal("C07D 413/10        20060101ALI20200522BHEP        ", patDoc.Classifications[10].Text)
	ass.Equal("C", patDoc.Classifications[10].Section)
	ass.Equal("07", patDoc.Classifications[10].Class)
	ass.Equal("D", patDoc.Classifications[10].SubClass)
	ass.Equal("413", patDoc.Classifications[10].MainGroup)
	ass.Equal("10", patDoc.Classifications[10].SubGroup)
	ass.Equal("L", patDoc.Classifications[10].FirstLater)

	ass.Equal("C07D 413/12        20060101ALI20200522BHEP        ", patDoc.Classifications[11].Text)
	ass.Equal("C", patDoc.Classifications[11].Section)
	ass.Equal("07", patDoc.Classifications[11].Class)
	ass.Equal("D", patDoc.Classifications[11].SubClass)
	ass.Equal("413", patDoc.Classifications[11].MainGroup)
	ass.Equal("12", patDoc.Classifications[11].SubGroup)
	ass.Equal("L", patDoc.Classifications[11].FirstLater)

	ass.Equal("C07D 413/14        20060101ALI20200522BHEP        ", patDoc.Classifications[12].Text)
	ass.Equal("C", patDoc.Classifications[12].Section)
	ass.Equal("07", patDoc.Classifications[12].Class)
	ass.Equal("D", patDoc.Classifications[12].SubClass)
	ass.Equal("413", patDoc.Classifications[12].MainGroup)
	ass.Equal("14", patDoc.Classifications[12].SubGroup)
	ass.Equal("L", patDoc.Classifications[12].FirstLater)

	ass.Equal("C07D 487/10        20060101ALI20200522BHEP        ", patDoc.Classifications[13].Text)
	ass.Equal("C", patDoc.Classifications[13].Section)
	ass.Equal("07", patDoc.Classifications[13].Class)
	ass.Equal("D", patDoc.Classifications[13].SubClass)
	ass.Equal("487", patDoc.Classifications[13].MainGroup)
	ass.Equal("10", patDoc.Classifications[13].SubGroup)
	ass.Equal("L", patDoc.Classifications[13].FirstLater)

}

func TestProcessXMLSimple151A13(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("test-data/application/v1-5-1-A1-3.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("EP21182450A1", patDoc.ID)
	ass.Equal("EP21182450NWA1.xml", patDoc.File)
	ass.Equal("de", patDoc.Lang)
	ass.Equal(Country("EP"), patDoc.Country)
	ass.Equal("4114164", patDoc.DocNumber)
	ass.Equal("A1", patDoc.Kind)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20230104", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("n", patDoc.Status)
	ass.Equal("ep-patent-document-v1-5-1", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("ERMITTELN EINES QUALITÄTSWERTES BEIM BESTÜCKEN EINES ELEKTRISCHEN BAUELEMENTS", patDoc.Title[0].Text)
	ass.Equal("de", patDoc.Title[0].Language)
	ass.Equal("DETERMINING A QUALITY VALUE WHEN FITTING AN ELECTRICAL COMPONENT", patDoc.Title[1].Text)
	ass.Equal("en", patDoc.Title[1].Language)
	ass.Equal("DÉTERMINATION D'UNE VALEUR DE QUALITÉ LORS DU MONTAGE D'UN ÉLÉMENT ÉLECTRIQUE", patDoc.Title[2].Text)
	ass.Equal("fr", patDoc.Title[2].Language)

	// abstract
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(749, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// claims
	ass.NotEmpty(patDoc.Claims)
	ass.Equal(1, len(patDoc.Claims))
	ass.Equal(3679, len(patDoc.Claims[0].Text))
	ass.Equal("de", patDoc.Claims[0].Language)
	ass.Equal("claims01", patDoc.Claims[0].Id)

	// description
	ass.NotEmpty(patDoc.Description)
	ass.Equal(23122, len(patDoc.Description[0].Text))
	ass.Equal("en", patDoc.Description[0].Language)

	// Inventors
	ass.NotEmpty(patDoc.Inventors)
	ass.Equal(9, len(patDoc.Inventors))
	for i := 0; i >= 8; i++ {
		ass.Equal("", patDoc.Inventors[i].Street)
	}

	ass.Equal("Blank, Rene", patDoc.Inventors[0].Name)
	ass.Equal("12249 Berlin", patDoc.Inventors[0].City)
	ass.Equal("Borwieck, Carsten", patDoc.Inventors[1].Name)
	ass.Equal("14052 Berlin", patDoc.Inventors[1].City)
	ass.Equal("Franke, Martin", patDoc.Inventors[2].Name)
	ass.Equal("14089 Berlin", patDoc.Inventors[2].City)
	ass.Equal("Frühauf, Peter", patDoc.Inventors[3].Name)
	ass.Equal("14612 Falkensee", patDoc.Inventors[3].City)
	ass.Equal("Heimann, Matthias", patDoc.Inventors[4].Name)
	ass.Equal("14469 Potsdam", patDoc.Inventors[4].City)
	ass.Equal("Knofe, Rüdiger", patDoc.Inventors[5].Name)
	ass.Equal("14513 Teltow", patDoc.Inventors[5].City)
	ass.Equal("Müller, Bernd", patDoc.Inventors[6].Name)
	ass.Equal("16259 Falkenberg", patDoc.Inventors[6].City)
	ass.Equal("Nerreter, Stefan", patDoc.Inventors[7].Name)
	ass.Equal("15754 Heidesee OT Blossin", patDoc.Inventors[7].City)
	ass.Equal("Wittreich, Ulrich", patDoc.Inventors[8].Name)
	ass.Equal("16727 Velten", patDoc.Inventors[8].City)

	// representatives
	ass.Empty(patDoc.Representatives)

	// contracting states
	ass.NotEmpty(patDoc.ContractingStates)
	ass.Equal(38, len(patDoc.ContractingStates))
	for i := 0; i <= 37; i++ {
		ass.Equal(2, len(patDoc.ContractingStates[i]))
	}

	// classifications
	ass.NotEmpty(patDoc.Classifications)
	ass.Equal(1, len(patDoc.Classifications))
	ass.Equal("H05K  13/08        20060101AFI20211202BHEP        ", patDoc.Classifications[0].Text)
	ass.Equal(IPC, patDoc.Classifications[0].System)
	ass.Equal(1, patDoc.Classifications[0].Sequence)
	ass.Equal("H", patDoc.Classifications[0].Section)
	ass.Equal("05", patDoc.Classifications[0].Class)
	ass.Equal("K", patDoc.Classifications[0].SubClass)
	ass.Equal("13", patDoc.Classifications[0].MainGroup)
	ass.Equal("08", patDoc.Classifications[0].SubGroup)
	ass.Equal("20060101", patDoc.Classifications[0].Version)
	ass.Equal("A", patDoc.Classifications[0].ClassificationLevel)
	ass.Equal("F", patDoc.Classifications[0].FirstLater)
	ass.Equal("I", patDoc.Classifications[0].ClassificationValue)
	ass.Equal("20211202", patDoc.Classifications[0].ActionDate)
	ass.Equal("B", patDoc.Classifications[0].OriginalOrReclassified)
	ass.Equal("H", patDoc.Classifications[0].Source)
	ass.Equal("EP", patDoc.Classifications[0].GeneratingOffice)

}

func TestProcessXMLSimple151A2(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("test-data/application/v1-5-1-A2.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("EP21165049A2", patDoc.ID)
	ass.Equal("EP21165049NWA2.xml", patDoc.File)
	ass.Equal("en", patDoc.Lang)
	ass.Equal(Country("EP"), patDoc.Country)
	ass.Equal("3882354", patDoc.DocNumber)
	ass.Equal("A2", patDoc.Kind)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20210922", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("n", patDoc.Status)
	ass.Equal("ep-patent-document-v1-5-1", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("MANIPULIERTE GLYCOSYLTRANSFERASEN UND VERFAHREN ZUR STEVIOL-GLYCOSID-GLUCOSYLIERUNG", patDoc.Title[0].Text)
	ass.Equal("de", patDoc.Title[0].Language)
	ass.Equal("ENGINEERED GLYCOSYLTRANSFERASES AND STEVIOL GLYCOSIDE GLUCOSYLATION METHODS", patDoc.Title[1].Text)
	ass.Equal("en", patDoc.Title[1].Language)
	ass.Equal("GLYCOSYLTRANSFÉRASES MODIFIÉES ET PROCÉDÉS DE GLUCOSYLATION DE GLYCOSIDES DE STÉVIOL", patDoc.Title[2].Text)
	ass.Equal("fr", patDoc.Title[2].Language)

	// abstract
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(978, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// claims
	ass.NotEmpty(patDoc.Claims)
	ass.Equal(1, len(patDoc.Claims))
	ass.Equal(21458, len(patDoc.Claims[0].Text))
	ass.Equal("en", patDoc.Claims[0].Language)
	ass.Equal("claims01", patDoc.Claims[0].Id)

	// description
	ass.NotEmpty(patDoc.Description)
	ass.Equal(1295579, len(patDoc.Description[0].Text))
	ass.Equal("en", patDoc.Description[0].Language)

	// Inventors
	ass.NotEmpty(patDoc.Inventors)
	ass.Equal("VROOM, Jonathan", patDoc.Inventors[0].Name)
	ass.Equal("Redwood City,  California 94063", patDoc.Inventors[0].City)
	ass.Equal("200 Penobscot Drive", patDoc.Inventors[0].Street)

	ass.Equal("GALANIE, Stephanie, Sue", patDoc.Inventors[1].Name)
	ass.Equal("Redwood City, California 94063", patDoc.Inventors[1].City)
	ass.Equal("200 Penobscot Drive", patDoc.Inventors[1].Street)

	ass.Equal("DELLAS, Nikki", patDoc.Inventors[2].Name)
	ass.Equal("Redwood City,  California 94063", patDoc.Inventors[2].City)
	ass.Equal("200 Penobscot Drive", patDoc.Inventors[2].Street)

	ass.Equal("LIANG, Jack", patDoc.Inventors[3].Name)
	ass.Equal("Redwood City,  California 94063", patDoc.Inventors[3].City)
	ass.Equal("200 Penobscot Drive", patDoc.Inventors[3].Street)

	// representatives
	ass.NotEmpty(patDoc.Representatives)
	ass.Equal(Country("GB"), patDoc.Representatives[0].Country)
	ass.Equal("101299776", patDoc.Representatives[0].IID)
	ass.Equal("London WC1B 5HA", patDoc.Representatives[0].City)
	ass.Equal("One Southampton Row", patDoc.Representatives[0].Street)
	ass.Equal("Carpmaels & Ransford LLP", patDoc.Representatives[0].Name)

	// contracting states
	ass.NotEmpty(patDoc.ContractingStates)
	ass.Equal(38, len(patDoc.ContractingStates))
	for i := 0; i <= 37; i++ {
		ass.Equal(2, len(patDoc.ContractingStates[i]))
	}

	// classifications

	ass.NotEmpty(patDoc.Classifications)
	ass.Equal(2, len(patDoc.Classifications))
	for i := 0; i <= 1; i++ {
		ass.Equal(IPC, patDoc.Classifications[i].System)
		ass.Equal(i+1, patDoc.Classifications[i].Sequence)
		ass.Equal("20060101", patDoc.Classifications[i].Version)
		ass.Equal("A", patDoc.Classifications[i].ClassificationLevel)
		ass.Equal("I", patDoc.Classifications[i].ClassificationValue)
		ass.Equal("20210804", patDoc.Classifications[i].ActionDate)
		ass.Equal("B", patDoc.Classifications[i].OriginalOrReclassified)
		ass.Equal("H", patDoc.Classifications[i].Source)
		ass.Equal("EP", patDoc.Classifications[i].GeneratingOffice)
	}

	ass.Equal("C12P  19/56        20060101AFI20210804BHEP        ", patDoc.Classifications[0].Text)
	ass.Equal("C", patDoc.Classifications[0].Section)
	ass.Equal("12", patDoc.Classifications[0].Class)
	ass.Equal("P", patDoc.Classifications[0].SubClass)
	ass.Equal("19", patDoc.Classifications[0].MainGroup)
	ass.Equal("56", patDoc.Classifications[0].SubGroup)
	ass.Equal("F", patDoc.Classifications[0].FirstLater)

	ass.Equal("C12N   9/10        20060101ALI20210804BHEP        ", patDoc.Classifications[1].Text)
	ass.Equal("C", patDoc.Classifications[1].Section)
	ass.Equal("12", patDoc.Classifications[1].Class)
	ass.Equal("N", patDoc.Classifications[1].SubClass)
	ass.Equal("9", patDoc.Classifications[1].MainGroup)
	ass.Equal("10", patDoc.Classifications[1].SubGroup)
	ass.Equal("L", patDoc.Classifications[1].FirstLater)

}

func TestProcessXMLSimple151A21(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("test-data/application/v1-5-1-A2-1.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("EP22166692A2", patDoc.ID)
	ass.Equal("EP22166692NWA2.xml", patDoc.File)
	ass.Equal("de", patDoc.Lang)
	ass.Equal(Country("EP"), patDoc.Country)
	ass.Equal("4102743", patDoc.DocNumber)
	ass.Equal("A2", patDoc.Kind)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20221214", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("n", patDoc.Status)
	ass.Equal("ep-patent-document-v1-5-1", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("ELEKTRONISCHES SYSTEM EINES KRAFTFAHRZEUGES MIT MINDESTENS ZWEI STEUERGERÄTEN, DIE JEWEILS EINEN EIGENEN TAKTGEBER AUFWEISEN UND VERFAHREN ZUR STEUERUNG EINES SOLCHEN SYSTEMS", patDoc.Title[0].Text)
	ass.Equal("de", patDoc.Title[0].Language)
	ass.Equal("ELECTRONIC SYSTEM OF A MOTOR VEHICLE WITH AT LEAST TWO CONTROL UNITS EACH HAVING ITS OWN CLOCK AND METHOD FOR CONTROLLING SUCH A SYSTEM", patDoc.Title[1].Text)
	ass.Equal("en", patDoc.Title[1].Language)
	ass.Equal("SYSTÈME ÉLECTRONIQUE D'UN VÉHICULE AUTOMOBILE POURVU D'AU MOINS DEUX APPAREILS DE COMMANDE, COMPRENANT RESPECTIVEMENT UNE HORLOGE PROPRE, ET PROCÉDÉ DE COMMANDE D'UN TEL SYSTÈME", patDoc.Title[2].Text)
	ass.Equal("fr", patDoc.Title[2].Language)

	// abstract
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(833, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// claims
	ass.NotEmpty(patDoc.Claims)
	ass.Equal(1, len(patDoc.Claims))
	ass.Equal(4806, len(patDoc.Claims[0].Text))
	ass.Equal("de", patDoc.Claims[0].Language)
	ass.Equal("claims01", patDoc.Claims[0].Id)

	// description
	ass.NotEmpty(patDoc.Description)
	ass.Equal(26403, len(patDoc.Description[0].Text))
	ass.Equal("en", patDoc.Description[0].Language)

	// Inventors
	ass.NotEmpty(patDoc.Inventors)
	ass.Equal("Kreis, Christopher", patDoc.Inventors[0].Name)
	ass.Equal("38126 Braunschweig", patDoc.Inventors[0].City)
	ass.Equal("", patDoc.Inventors[0].Street) //empty

	ass.Equal("Rüger, Tobias", patDoc.Inventors[1].Name)
	ass.Equal("38108 Braunschweig", patDoc.Inventors[1].City)
	ass.Equal("", patDoc.Inventors[1].Street) //empty

	// representatives empty
	ass.Empty(patDoc.Representatives)

	// contracting states
	ass.NotEmpty(patDoc.ContractingStates)
	ass.Equal(38, len(patDoc.ContractingStates))
	for i := 0; i <= 37; i++ {
		ass.Equal(2, len(patDoc.ContractingStates[i]))
	}

	// classifications
	ass.NotEmpty(patDoc.Classifications)
	ass.Equal(4, len(patDoc.Classifications))
	for i := 0; i <= 3; i++ {
		ass.Equal(IPC, patDoc.Classifications[i].System)
		ass.Equal(i+1, patDoc.Classifications[i].Sequence)
		ass.Equal("20060101", patDoc.Classifications[i].Version)
		ass.Equal("A", patDoc.Classifications[i].ClassificationLevel)
		ass.Equal("20220830", patDoc.Classifications[i].ActionDate)
		ass.Equal("B", patDoc.Classifications[i].OriginalOrReclassified)
		ass.Equal("H", patDoc.Classifications[i].Source)
		ass.Equal("EP", patDoc.Classifications[i].GeneratingOffice)
	}

	ass.Equal("H04J   3/06        20060101AFI20220830BHEP        ", patDoc.Classifications[0].Text)
	ass.Equal("H", patDoc.Classifications[0].Section)
	ass.Equal("I", patDoc.Classifications[0].ClassificationValue)
	ass.Equal("04", patDoc.Classifications[0].Class)
	ass.Equal("J", patDoc.Classifications[0].SubClass)
	ass.Equal("3", patDoc.Classifications[0].MainGroup)
	ass.Equal("06", patDoc.Classifications[0].SubGroup)
	ass.Equal("F", patDoc.Classifications[0].FirstLater)

	ass.Equal("H04Q   9/00        20060101ALI20220830BHEP        ", patDoc.Classifications[1].Text)
	ass.Equal("H", patDoc.Classifications[1].Section)
	ass.Equal("I", patDoc.Classifications[1].ClassificationValue)
	ass.Equal("04", patDoc.Classifications[1].Class)
	ass.Equal("Q", patDoc.Classifications[1].SubClass)
	ass.Equal("9", patDoc.Classifications[1].MainGroup)
	ass.Equal("00", patDoc.Classifications[1].SubGroup)
	ass.Equal("L", patDoc.Classifications[1].FirstLater)

	ass.Equal("B62D   6/00        20060101ALI20220830BHEP        ", patDoc.Classifications[2].Text)
	ass.Equal("B", patDoc.Classifications[2].Section)
	ass.Equal("I", patDoc.Classifications[2].ClassificationValue)
	ass.Equal("62", patDoc.Classifications[2].Class)
	ass.Equal("D", patDoc.Classifications[2].SubClass)
	ass.Equal("6", patDoc.Classifications[2].MainGroup)
	ass.Equal("00", patDoc.Classifications[2].SubGroup)
	ass.Equal("L", patDoc.Classifications[2].FirstLater)

	ass.Equal("H04L  12/40        20060101ALN20220830BHEP        ", patDoc.Classifications[3].Text)
	ass.Equal("H", patDoc.Classifications[3].Section)
	ass.Equal("N", patDoc.Classifications[3].ClassificationValue)
	ass.Equal("04", patDoc.Classifications[3].Class)
	ass.Equal("L", patDoc.Classifications[3].SubClass)
	ass.Equal("12", patDoc.Classifications[3].MainGroup)
	ass.Equal("40", patDoc.Classifications[3].SubGroup)
	ass.Equal("L", patDoc.Classifications[3].FirstLater)
}

func TestProcessXMLSimple151A22(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("test-data/application/v1-5-1-A2-2.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("EP22191906A2", patDoc.ID)
	ass.Equal("EP22191906NWA2.xml", patDoc.File)
	ass.Equal("en", patDoc.Lang)
	ass.Equal(Country("EP"), patDoc.Country)
	ass.Equal("4113994", patDoc.DocNumber)
	ass.Equal("A2", patDoc.Kind)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20230104", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("n", patDoc.Status)
	ass.Equal("ep-patent-document-v1-5-1", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("MV/MODUS-VORHERSAGE, ROI-BASIERTE ÜBERTRAGUNG, METADATENERFASSUNG UND FORMATERKENNUNG FÜR 360-VIDEO", patDoc.Title[0].Text)
	ass.Equal("de", patDoc.Title[0].Language)
	ass.Equal("MV/MODE PREDICTION, ROI-BASED TRANSMIT, METADATA CAPTURE, AND FORMAT DETECTION FOR 360 VIDEO", patDoc.Title[1].Text)
	ass.Equal("en", patDoc.Title[1].Language)
	ass.Equal("PRÉDICTION DE MODE/MV, ÉMISSION BASÉE SUR ROI, CAPTURE DE MÉTADONNÉES ET DÉTECTION DE FORMATS POUR VIDÉO 360", patDoc.Title[2].Text)
	ass.Equal("fr", patDoc.Title[2].Language)

	// abstract
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(395, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// claims
	ass.NotEmpty(patDoc.Claims)
	ass.Equal(1, len(patDoc.Claims))
	ass.Equal(3019, len(patDoc.Claims[0].Text))
	ass.Equal("en", patDoc.Claims[0].Language)
	ass.Equal("claims01", patDoc.Claims[0].Id)

	// description
	ass.NotEmpty(patDoc.Description)
	ass.Equal(261867, len(patDoc.Description[0].Text))
	ass.Equal("en", patDoc.Description[0].Language)

	// Inventors
	ass.NotEmpty(patDoc.Inventors)
	ass.Equal("BOYCE, Jill M.", patDoc.Inventors[0].Name)
	ass.Equal("Portland, OR, 97225", patDoc.Inventors[0].City)
	ass.Equal("", patDoc.Inventors[0].Street) //empty

	ass.Equal("LEE, Sang-Hee", patDoc.Inventors[1].Name)
	ass.Equal("San Jose, CA, 95129", patDoc.Inventors[1].City)
	ass.Equal("", patDoc.Inventors[1].Street) //empty

	// representatives empty
	ass.NotEmpty(patDoc.Representatives)
	ass.Equal(Country("DE"), patDoc.Representatives[0].Country)
	ass.Equal("100002032", patDoc.Representatives[0].IID)
	ass.Equal("80336 München", patDoc.Representatives[0].City)
	ass.Equal("Boehmert & Boehmert \nAnwaltspartnerschaft mbB \nPettenkoferstrasse 22", patDoc.Representatives[0].Street)
	ass.Equal("Goddar, Heinz J.", patDoc.Representatives[0].Name)

	// contracting states
	ass.NotEmpty(patDoc.ContractingStates)
	ass.Equal(38, len(patDoc.ContractingStates))
	for i := 0; i <= 37; i++ {
		ass.Equal(2, len(patDoc.ContractingStates[i]))
	}

	// classifications
	ass.NotEmpty(patDoc.Classifications)
	ass.Equal(1, len(patDoc.Classifications))
	ass.Equal("H04N  19/105       20140101AFI20221123BHEP        ", patDoc.Classifications[0].Text)
	ass.Equal(IPC, patDoc.Classifications[0].System)
	ass.Equal(1, patDoc.Classifications[0].Sequence)
	ass.Equal("H", patDoc.Classifications[0].Section)
	ass.Equal("04", patDoc.Classifications[0].Class)
	ass.Equal("N", patDoc.Classifications[0].SubClass)
	ass.Equal("19", patDoc.Classifications[0].MainGroup)
	ass.Equal("105", patDoc.Classifications[0].SubGroup)
	ass.Equal("20140101", patDoc.Classifications[0].Version)
	ass.Equal("A", patDoc.Classifications[0].ClassificationLevel)
	ass.Equal("F", patDoc.Classifications[0].FirstLater)
	ass.Equal("I", patDoc.Classifications[0].ClassificationValue)
	ass.Equal("20221123", patDoc.Classifications[0].ActionDate)
	ass.Equal("B", patDoc.Classifications[0].OriginalOrReclassified)
	ass.Equal("H", patDoc.Classifications[0].Source)
	ass.Equal("EP", patDoc.Classifications[0].GeneratingOffice)
}

func TestProcessXMLSimple151B1(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("test-data/grant/v1-5-1-B1.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("EP16849316B1", patDoc.ID)
	ass.Equal("EP16849316NWB1.xml", patDoc.File)
	ass.Equal("en", patDoc.Lang)
	ass.Equal(Country("EP"), patDoc.Country)
	ass.Equal("3383757", patDoc.DocNumber)
	ass.Equal("B1", patDoc.Kind)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20210707", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("n", patDoc.Status)
	ass.Equal("ep-patent-document-v1-5-1", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("SCHRAUBVERSCHLUSSSYSTEME MIT MAGNETISCHEM MERKMAL", patDoc.Title[0].Text)
	ass.Equal("de", patDoc.Title[0].Language)
	ass.Equal("SCREW-TYPE CLOSURE SYSTEMS WITH MAGNETIC FEATURE", patDoc.Title[1].Text)
	ass.Equal("en", patDoc.Title[1].Language)
	ass.Equal("SYSTÈMES DE FERMETURE DE TYPE À VIS COMPRENANT DES ÉLÉMENTS MAGNÉTIQUES", patDoc.Title[2].Text)
	ass.Equal("fr", patDoc.Title[2].Language)

	// abstract zero
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(0, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// claims
	ass.NotEmpty(patDoc.Claims)
	ass.Equal(3, len(patDoc.Claims))
	ass.Equal(3763, len(patDoc.Claims[0].Text))
	ass.Equal("en", patDoc.Claims[0].Language)
	ass.Equal("claims01", patDoc.Claims[0].Id)
	ass.Equal(4420, len(patDoc.Claims[1].Text))
	ass.Equal("de", patDoc.Claims[1].Language)
	ass.Equal("claims02", patDoc.Claims[1].Id)
	ass.Equal(4451, len(patDoc.Claims[2].Text))
	ass.Equal("fr", patDoc.Claims[2].Language)
	ass.Equal("claims03", patDoc.Claims[2].Id)

	// description
	ass.NotEmpty(patDoc.Description)
	ass.Equal(18652, len(patDoc.Description[0].Text))
	ass.Equal("en", patDoc.Description[0].Language)

	// citations
	ass.NotEmpty(patDoc.Citations)
	ass.Equal(Country("US"), patDoc.Citations[0].Country)
	ass.Equal("2007204872", patDoc.Citations[0].DocNumber)
	ass.Equal("A1", patDoc.Citations[0].Kind)
	ass.Equal(Country("FR"), patDoc.Citations[1].Country)
	ass.Equal("2973005", patDoc.Citations[1].DocNumber)
	ass.Equal("A1", patDoc.Citations[1].Kind)

	// Inventors
	ass.NotEmpty(patDoc.Inventors)
	ass.Equal(Country("FR"), patDoc.Inventors[0].Country)
	ass.Equal("42100 Saint-Etienne", patDoc.Inventors[0].City)
	ass.Equal("50, rue Desire Claude", patDoc.Inventors[0].Street)
	ass.Equal("JACOB, Christophe", patDoc.Inventors[0].Name)

	ass.Equal(Country("US"), patDoc.Inventors[1].Country)
	ass.Equal("New York, New York 10022", patDoc.Inventors[1].City)
	ass.Equal("300 East 55th Street, Apt. 11A", patDoc.Inventors[1].Street)
	ass.Equal("BOUIX, Herve", patDoc.Inventors[1].Name)

	// owners
	ass.NotEmpty(patDoc.Owners)
	ass.Equal(Country("US"), patDoc.Owners[0].Country)
	ass.Equal("101324505", patDoc.Owners[0].IID)
	ass.Equal("97.70.137658", patDoc.Owners[0].IRF)
	ass.Equal("Melville, NY 11747", patDoc.Owners[0].City)
	ass.Equal("155 Pinelawn Road, Suite 345 South", patDoc.Owners[0].Street)
	ass.Equal("ELC Management LLC", patDoc.Owners[0].Name)

	// representatives
	ass.NotEmpty(patDoc.Representatives)
	ass.Equal(Country("GB"), patDoc.Representatives[0].Country)
	ass.Equal("101728904", patDoc.Representatives[0].IID)
	ass.Equal("London EC4Y 8JD", patDoc.Representatives[0].City)
	ass.Equal("St. Bride's House \n10 Salisbury Square", patDoc.Representatives[0].Street)
	ass.Equal("Dehns", patDoc.Representatives[0].Name)

	// contracting states
	ass.NotEmpty(patDoc.ContractingStates)
	ass.Equal(38, len(patDoc.ContractingStates))
	for i := 0; i <= 37; i++ {
		ass.Equal(2, len(patDoc.ContractingStates[i]))
	}

	// classifications

	ass.NotEmpty(patDoc.Classifications)
	ass.Equal(5, len(patDoc.Classifications))

	ass.Equal("B65D  51/18        20060101AFI20191003BHEP        ", patDoc.Classifications[0].Text)
	ass.Equal(IPC, patDoc.Classifications[0].System)
	ass.Equal(1, patDoc.Classifications[0].Sequence)
	ass.Equal("B", patDoc.Classifications[0].Section)
	ass.Equal("65", patDoc.Classifications[0].Class)
	ass.Equal("D", patDoc.Classifications[0].SubClass)
	ass.Equal("51", patDoc.Classifications[0].MainGroup)
	ass.Equal("18", patDoc.Classifications[0].SubGroup)
	ass.Equal("20060101", patDoc.Classifications[0].Version)
	ass.Equal("A", patDoc.Classifications[0].ClassificationLevel)
	ass.Equal("F", patDoc.Classifications[0].FirstLater)
	ass.Equal("I", patDoc.Classifications[0].ClassificationValue)
	ass.Equal("20191003", patDoc.Classifications[0].ActionDate)
	ass.Equal("B", patDoc.Classifications[0].OriginalOrReclassified)
	ass.Equal("H", patDoc.Classifications[0].Source)
	ass.Equal("EP", patDoc.Classifications[0].GeneratingOffice)

	ass.Equal("B65D  43/02        20060101ALI20191003BHEP        ", patDoc.Classifications[1].Text)
	ass.Equal(IPC, patDoc.Classifications[1].System)
	ass.Equal(2, patDoc.Classifications[1].Sequence)
	ass.Equal("B", patDoc.Classifications[1].Section)
	ass.Equal("65", patDoc.Classifications[1].Class)
	ass.Equal("D", patDoc.Classifications[1].SubClass)
	ass.Equal("43", patDoc.Classifications[1].MainGroup)
	ass.Equal("02", patDoc.Classifications[1].SubGroup)
	ass.Equal("20060101", patDoc.Classifications[1].Version)
	ass.Equal("A", patDoc.Classifications[1].ClassificationLevel)
	ass.Equal("L", patDoc.Classifications[1].FirstLater)
	ass.Equal("I", patDoc.Classifications[1].ClassificationValue)
	ass.Equal("20191003", patDoc.Classifications[1].ActionDate)
	ass.Equal("B", patDoc.Classifications[1].OriginalOrReclassified)
	ass.Equal("H", patDoc.Classifications[1].Source)
	ass.Equal("EP", patDoc.Classifications[1].GeneratingOffice)

	ass.Equal("B65D  51/32        20060101ALI20191003BHEP        ", patDoc.Classifications[2].Text)
	ass.Equal(IPC, patDoc.Classifications[2].System)
	ass.Equal(3, patDoc.Classifications[2].Sequence)
	ass.Equal("B", patDoc.Classifications[2].Section)
	ass.Equal("65", patDoc.Classifications[2].Class)
	ass.Equal("D", patDoc.Classifications[2].SubClass)
	ass.Equal("51", patDoc.Classifications[2].MainGroup)
	ass.Equal("32", patDoc.Classifications[2].SubGroup)
	ass.Equal("20060101", patDoc.Classifications[2].Version)
	ass.Equal("A", patDoc.Classifications[2].ClassificationLevel)
	ass.Equal("L", patDoc.Classifications[2].FirstLater)
	ass.Equal("I", patDoc.Classifications[2].ClassificationValue)
	ass.Equal("20191003", patDoc.Classifications[2].ActionDate)
	ass.Equal("B", patDoc.Classifications[2].OriginalOrReclassified)
	ass.Equal("H", patDoc.Classifications[2].Source)
	ass.Equal("EP", patDoc.Classifications[2].GeneratingOffice)

	ass.Equal("A45D  34/00        20060101ALI20191003BHEP        ", patDoc.Classifications[3].Text)
	ass.Equal(IPC, patDoc.Classifications[3].System)
	ass.Equal(4, patDoc.Classifications[3].Sequence)
	ass.Equal("A", patDoc.Classifications[3].Section)
	ass.Equal("45", patDoc.Classifications[3].Class)
	ass.Equal("D", patDoc.Classifications[3].SubClass)
	ass.Equal("34", patDoc.Classifications[3].MainGroup)
	ass.Equal("00", patDoc.Classifications[3].SubGroup)
	ass.Equal("20060101", patDoc.Classifications[3].Version)
	ass.Equal("A", patDoc.Classifications[3].ClassificationLevel)
	ass.Equal("L", patDoc.Classifications[3].FirstLater)
	ass.Equal("I", patDoc.Classifications[3].ClassificationValue)
	ass.Equal("20191003", patDoc.Classifications[3].ActionDate)
	ass.Equal("B", patDoc.Classifications[3].OriginalOrReclassified)
	ass.Equal("H", patDoc.Classifications[3].Source)
	ass.Equal("EP", patDoc.Classifications[3].GeneratingOffice)

	ass.Equal("A45D  40/00        20060101ALI20191003BHEP        ", patDoc.Classifications[4].Text)
	ass.Equal(IPC, patDoc.Classifications[4].System)
	ass.Equal(5, patDoc.Classifications[4].Sequence)
	ass.Equal("A", patDoc.Classifications[4].Section)
	ass.Equal("45", patDoc.Classifications[4].Class)
	ass.Equal("D", patDoc.Classifications[4].SubClass)
	ass.Equal("40", patDoc.Classifications[4].MainGroup)
	ass.Equal("00", patDoc.Classifications[4].SubGroup)
	ass.Equal("20060101", patDoc.Classifications[4].Version)
	ass.Equal("A", patDoc.Classifications[4].ClassificationLevel)
	ass.Equal("L", patDoc.Classifications[4].FirstLater)
	ass.Equal("I", patDoc.Classifications[4].ClassificationValue)
	ass.Equal("20191003", patDoc.Classifications[4].ActionDate)
	ass.Equal("B", patDoc.Classifications[4].OriginalOrReclassified)
	ass.Equal("H", patDoc.Classifications[4].Source)
	ass.Equal("EP", patDoc.Classifications[4].GeneratingOffice)

	/*fmt.Println(patDoc.Title[0])
	fmt.Println(patDoc.Claims[0])
	fmt.Println(patDoc.Description[0])
	fmt.Println(patDoc.Citations[0])
	*/
}

func TestProcessXMLSimple151B11(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("test-data/grant/v1-5-1-B1-1.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("EP01976162B1", patDoc.ID)
	ass.Equal("EP01976162NWB1.xml", patDoc.File)
	ass.Equal("en", patDoc.Lang)
	ass.Equal(Country("EP"), patDoc.Country)
	ass.Equal("1316073", patDoc.DocNumber)
	ass.Equal("B1", patDoc.Kind)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20221214", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("n", patDoc.Status)
	ass.Equal("ep-patent-document-v1-5-1", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("VORRICHTUNG ZUR PRÜFUNG VON BEGLAUBIGTEN DOKUMENTEN", patDoc.Title[0].Text)
	ass.Equal("de", patDoc.Title[0].Language)
	ass.Equal("A CERTIFIED PAPER DISCRIMINATING APPARATUS", patDoc.Title[1].Text)
	ass.Equal("en", patDoc.Title[1].Language)
	ass.Equal("APPAREIL DE DISCRIMINATION DE PAPIERS CERTIFIES", patDoc.Title[2].Text)
	ass.Equal("fr", patDoc.Title[2].Language)

	// abstract empty
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(0, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// claims
	ass.NotEmpty(patDoc.Claims)
	ass.Equal(3, len(patDoc.Claims))
	ass.Equal(5069, len(patDoc.Claims[0].Text))
	ass.Equal("en", patDoc.Claims[0].Language)
	ass.Equal("claims01", patDoc.Claims[0].Id)
	ass.Equal(5869, len(patDoc.Claims[1].Text))
	ass.Equal("de", patDoc.Claims[1].Language)
	ass.Equal("claims02", patDoc.Claims[1].Id)
	ass.Equal(6202, len(patDoc.Claims[2].Text))
	ass.Equal("fr", patDoc.Claims[2].Language)
	ass.Equal("claims03", patDoc.Claims[2].Id)

	// description
	ass.NotEmpty(patDoc.Description)
	ass.Equal(75241, len(patDoc.Description[0].Text))
	ass.Equal("en", patDoc.Description[0].Language)

	// citations
	ass.NotEmpty(patDoc.Citations)
	ass.Equal(2, len(patDoc.Citations))
	ass.Equal(Country("DE"), patDoc.Citations[0].Country)
	ass.Equal("19708543", patDoc.Citations[0].DocNumber)
	ass.Equal("A1", patDoc.Citations[0].Kind)
	ass.Equal(Country("US"), patDoc.Citations[1].Country)
	ass.Equal("5719984", patDoc.Citations[1].DocNumber)
	ass.Equal("A", patDoc.Citations[1].Kind)

	// Inventors
	ass.NotEmpty(patDoc.Inventors)
	ass.Equal(Country("JP"), patDoc.Inventors[0].Country)
	ass.Equal("Osaka-shi,\nOsaka 547-00325", patDoc.Inventors[0].City)
	ass.Equal("3-15, Nishiwaki 2-chome,\nHirano-ku", patDoc.Inventors[0].Street)
	ass.Equal("MATSUI, Toru,\nJapan Cash Machine Co., Ltd", patDoc.Inventors[0].Name)

	ass.Equal(Country("DE"), patDoc.Inventors[1].Country)
	ass.Equal("10997 Berlin", patDoc.Inventors[1].City)
	ass.Equal("Schlesische Strasse 29", patDoc.Inventors[1].Street)
	ass.Equal("AHLERS, Benedikt", patDoc.Inventors[1].Name)

	// owners
	ass.NotEmpty(patDoc.Owners)
	ass.Equal(Country("DE"), patDoc.Owners[0].Country)
	ass.Equal("101510543", patDoc.Owners[0].IID)
	ass.Equal("54307", patDoc.Owners[0].IRF)
	ass.Equal("10969 Berlin", patDoc.Owners[0].City)
	ass.Equal("Kommandantenstraße 18", patDoc.Owners[0].Street)
	ass.Equal("Bundesdruckerei GmbH", patDoc.Owners[0].Name)

	// representatives
	ass.NotEmpty(patDoc.Representatives)
	ass.Equal(Country("DE"), patDoc.Representatives[0].Country)
	ass.Equal("100039344", patDoc.Representatives[0].IID)
	ass.Equal("71065 Sindelfingen", patDoc.Representatives[0].City)
	ass.Equal("Patentanwälte \nMammel & Maser \nTilsiter Strasse 3", patDoc.Representatives[0].Street)
	ass.Equal("Maser, Jochen", patDoc.Representatives[0].Name)

	// contracting states
	ass.NotEmpty(patDoc.ContractingStates)
	ass.Equal(20, len(patDoc.ContractingStates))
	for i := 0; i <= 19; i++ {
		ass.Equal(2, len(patDoc.ContractingStates[i]))
	}

	// classifications
	ass.NotEmpty(patDoc.Classifications)
	ass.Equal(1, len(patDoc.Classifications))
	ass.Equal("G07D   7/00        20160101AFI20020314BHEP        ", patDoc.Classifications[0].Text)
	ass.Equal(IPC, patDoc.Classifications[0].System)
	ass.Equal(1, patDoc.Classifications[0].Sequence)
	ass.Equal("G", patDoc.Classifications[0].Section)
	ass.Equal("07", patDoc.Classifications[0].Class)
	ass.Equal("D", patDoc.Classifications[0].SubClass)
	ass.Equal("7", patDoc.Classifications[0].MainGroup)
	ass.Equal("00", patDoc.Classifications[0].SubGroup)
	ass.Equal("20160101", patDoc.Classifications[0].Version)
	ass.Equal("A", patDoc.Classifications[0].ClassificationLevel)
	ass.Equal("F", patDoc.Classifications[0].FirstLater)
	ass.Equal("I", patDoc.Classifications[0].ClassificationValue)
	ass.Equal("20020314", patDoc.Classifications[0].ActionDate)
	ass.Equal("B", patDoc.Classifications[0].OriginalOrReclassified)
	ass.Equal("H", patDoc.Classifications[0].Source)
	ass.Equal("EP", patDoc.Classifications[0].GeneratingOffice)
}

func TestProcessXMLSimple151B2(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("test-data/grant/v1-5-1-B2.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("EP13196195B2", patDoc.ID)
	ass.Equal("EP13196195NWB2.xml", patDoc.File)
	ass.Equal("de", patDoc.Lang)
	ass.Equal(Country("EP"), patDoc.Country)
	ass.Equal("2716170", patDoc.DocNumber)
	ass.Equal("B2", patDoc.Kind)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20210922", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("n", patDoc.Status)
	ass.Equal("ep-patent-document-v1-5-1", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("Vorrichtung zum Transport von band- oder streifenförmigem Material", patDoc.Title[0].Text)
	ass.Equal("de", patDoc.Title[0].Language)
	ass.Equal("Device for transporting material in the form of strips or tape", patDoc.Title[1].Text)
	ass.Equal("en", patDoc.Title[1].Language)
	ass.Equal("Dispositif de transport de matériau en forme de bande ou de ruban", patDoc.Title[2].Text)
	ass.Equal("fr", patDoc.Title[2].Language)

	// abstract !! Empty
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(0, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// claims
	ass.NotEmpty(patDoc.Claims)
	ass.Equal(3, len(patDoc.Claims))
	ass.Equal(1793, len(patDoc.Claims[0].Text))
	ass.Equal("de", patDoc.Claims[0].Language)
	ass.Equal("claims01", patDoc.Claims[0].Id)
	ass.Equal(1649, len(patDoc.Claims[1].Text))
	ass.Equal("en", patDoc.Claims[1].Language)
	ass.Equal("claims02", patDoc.Claims[1].Id)
	ass.Equal(1878, len(patDoc.Claims[2].Text))
	ass.Equal("fr", patDoc.Claims[2].Language)
	ass.Equal("claims03", patDoc.Claims[2].Id)

	// description
	ass.NotEmpty(patDoc.Description)
	ass.Equal(29409, len(patDoc.Description[0].Text))
	ass.Equal("en", patDoc.Description[0].Language)

	// citations
	ass.NotEmpty(patDoc.Citations)
	ass.Equal(Country("DE"), patDoc.Citations[0].Country)
	ass.Equal("19751691", patDoc.Citations[0].DocNumber)
	ass.Equal("A", patDoc.Citations[0].Kind)

	// Inventors
	ass.NotEmpty(patDoc.Inventors)
	ass.Equal(Country("DE"), patDoc.Inventors[0].Country)
	ass.Equal("26871 Papenburg", patDoc.Inventors[0].City)
	ass.Equal("Richardstr. 4", patDoc.Inventors[0].Street)
	ass.Equal("Grothaus, Frank", patDoc.Inventors[0].Name)

	ass.Equal(Country("DE"), patDoc.Inventors[1].Country)
	ass.Equal("21039 Börnsen", patDoc.Inventors[1].City)
	ass.Equal("Zum Alten Elbufer 45", patDoc.Inventors[1].Street)
	ass.Equal("Schafforz, Mathias", patDoc.Inventors[1].Name)

	// representatives
	ass.NotEmpty(patDoc.Representatives)
	ass.Equal(Country("DE"), patDoc.Representatives[0].Country)
	ass.Equal("101242078", patDoc.Representatives[0].IID)
	ass.Equal("22763 Hamburg", patDoc.Representatives[0].City)
	ass.Equal("Patentanwälte \nFriedensallee 290", patDoc.Representatives[0].Street)
	ass.Equal("Müller Verweyen", patDoc.Representatives[0].Name)

	// contracting states
	ass.NotEmpty(patDoc.ContractingStates)
	ass.Equal(38, len(patDoc.ContractingStates))
	for i := 0; i <= 37; i++ {
		ass.Equal(2, len(patDoc.ContractingStates[i]))
	}

	// classifications
	ass.NotEmpty(patDoc.Classifications)
	ass.Equal(1, len(patDoc.Classifications))

	ass.Equal("A24C   5/20        20060101AFI20150420BHEP        ", patDoc.Classifications[0].Text)
	ass.Equal(IPC, patDoc.Classifications[0].System)
	ass.Equal(1, patDoc.Classifications[0].Sequence)
	ass.Equal("A", patDoc.Classifications[0].Section)
	ass.Equal("24", patDoc.Classifications[0].Class)
	ass.Equal("C", patDoc.Classifications[0].SubClass)
	ass.Equal("5", patDoc.Classifications[0].MainGroup)
	ass.Equal("20", patDoc.Classifications[0].SubGroup)
	ass.Equal("20060101", patDoc.Classifications[0].Version)
	ass.Equal("A", patDoc.Classifications[0].ClassificationLevel)
	ass.Equal("F", patDoc.Classifications[0].FirstLater)
	ass.Equal("I", patDoc.Classifications[0].ClassificationValue)
	ass.Equal("20150420", patDoc.Classifications[0].ActionDate)
	ass.Equal("B", patDoc.Classifications[0].OriginalOrReclassified)
	ass.Equal("H", patDoc.Classifications[0].Source)
	ass.Equal("EP", patDoc.Classifications[0].GeneratingOffice)

	/*fmt.Println(patDoc.Title[0])
	fmt.Println(patDoc.Claims[0])
	fmt.Println(patDoc.Description[0])
	fmt.Println(patDoc.Citations[0])
	*/
}

func TestProcessXMLSimple151B21(t *testing.T) {
	ass := assert.New(t)
	data, err := ioutil.ReadFile("test-data/grant/v1-5-1-B2-1.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.NoError(err)

	ass.Equal("EP10809927B2", patDoc.ID)
	ass.Equal("EP10809927NWB2.xml", patDoc.File)
	ass.Equal("en", patDoc.Lang)
	ass.Equal(Country("EP"), patDoc.Country)
	ass.Equal("2361995", patDoc.DocNumber)
	ass.Equal("B2", patDoc.Kind)
	ass.False(patDoc.DatePubl.IsZero())
	ass.Equal("20221214", patDoc.DatePubl.Format(layoutDatePubl))
	ass.Equal("n", patDoc.Status)
	ass.Equal("ep-patent-document-v1-5-1", patDoc.DtdVersion)

	// title
	ass.NotEmpty(patDoc.Title)
	ass.Equal("PERLITSCHIENE", patDoc.Title[0].Text)
	ass.Equal("de", patDoc.Title[0].Language)
	ass.Equal("PEARLITE RAIL", patDoc.Title[1].Text)
	ass.Equal("en", patDoc.Title[1].Language)
	ass.Equal("RAIL EN PERLITE", patDoc.Title[2].Text)
	ass.Equal("fr", patDoc.Title[2].Language)

	// abstract !! Empty
	ass.NotEmpty(patDoc.Abstract)
	ass.Equal(0, len(patDoc.Abstract[0].Text))
	ass.Equal("en", patDoc.Abstract[0].Language)

	// claims
	ass.NotEmpty(patDoc.Claims)
	ass.Equal(3, len(patDoc.Claims))
	ass.Equal(3920, len(patDoc.Claims[0].Text))
	ass.Equal("en", patDoc.Claims[0].Language)
	ass.Equal("claims01", patDoc.Claims[0].Id)
	ass.Equal(4001, len(patDoc.Claims[1].Text))
	ass.Equal("de", patDoc.Claims[1].Language)
	ass.Equal("claims02", patDoc.Claims[1].Id)
	ass.Equal(4377, len(patDoc.Claims[2].Text))
	ass.Equal("fr", patDoc.Claims[2].Language)
	ass.Equal("claims03", patDoc.Claims[2].Id)

	// description
	ass.NotEmpty(patDoc.Description)
	ass.Equal(90628, len(patDoc.Description[0].Text))
	ass.Equal("en", patDoc.Description[0].Language)

	// citations
	ass.NotEmpty(patDoc.Citations)
	ass.Equal(Country("JP"), patDoc.Citations[0].Country)
	ass.Equal("S51002616", patDoc.Citations[0].DocNumber)
	ass.Equal("B", patDoc.Citations[0].Kind)

	// Inventors
	ass.NotEmpty(patDoc.Inventors)
	ass.Equal(Country("JP"), patDoc.Inventors[0].Country)
	ass.Equal("Tokyo 100-8071", patDoc.Inventors[0].City)
	ass.Equal("c/o NIPPON STEEL CORPORATION\n6-1, Marunouchi 2-chome\nChiyoda-ku", patDoc.Inventors[0].Street)
	ass.Equal("UEDA Masaharu", patDoc.Inventors[0].Name)

	ass.Equal(Country("JP"), patDoc.Inventors[1].Country)
	ass.Equal("Tokyo 100-8071", patDoc.Inventors[1].City)
	ass.Equal("c/o NIPPON STEEL CORPORATION\n6-1, Marunouchi 2-chome\nChiyoda-ku", patDoc.Inventors[1].Street)
	ass.Equal("SONOYAMA Kyohei", patDoc.Inventors[1].Name)

	// representatives
	ass.NotEmpty(patDoc.Representatives)
	ass.Equal(1, len(patDoc.Representatives))
	ass.Equal(Country("DE"), patDoc.Representatives[0].Country)
	ass.Equal("100751388", patDoc.Representatives[0].IID)
	ass.Equal("81675 München", patDoc.Representatives[0].City)
	ass.Equal("Siebertstrasse 3", patDoc.Representatives[0].Street)
	ass.Equal("Vossius & Partner \nPatentanwälte Rechtsanwälte mbB", patDoc.Representatives[0].Name)

	// contracting states
	ass.NotEmpty(patDoc.ContractingStates)
	ass.Equal(37, len(patDoc.ContractingStates))
	for i := 0; i <= 36; i++ {
		ass.Equal(2, len(patDoc.ContractingStates[i]))
	}

	// classifications
	ass.NotEmpty(patDoc.Classifications)
	ass.Equal(5, len(patDoc.Classifications))

	for i := 0; i <= 4; i++ {
		ass.Equal(IPC, patDoc.Classifications[i].System)
		ass.Equal(i+1, patDoc.Classifications[i].Sequence)
		ass.Equal("20060101", patDoc.Classifications[i].Version)
		ass.Equal("A", patDoc.Classifications[i].ClassificationLevel)
		ass.Equal("I", patDoc.Classifications[i].ClassificationValue)
		ass.Equal("20170607", patDoc.Classifications[i].ActionDate)
		ass.Equal("B", patDoc.Classifications[i].OriginalOrReclassified)
		ass.Equal("H", patDoc.Classifications[i].Source)
		ass.Equal("EP", patDoc.Classifications[i].GeneratingOffice)

	}

	ass.Equal("C22C  38/00        20060101AFI20170607BHEP        ", patDoc.Classifications[0].Text)
	ass.Equal("C", patDoc.Classifications[0].Section)
	ass.Equal("22", patDoc.Classifications[0].Class)
	ass.Equal("C", patDoc.Classifications[0].SubClass)
	ass.Equal("38", patDoc.Classifications[0].MainGroup)
	ass.Equal("00", patDoc.Classifications[0].SubGroup)
	ass.Equal("F", patDoc.Classifications[0].FirstLater)

	ass.Equal("C22C  38/04        20060101ALI20170607BHEP        ", patDoc.Classifications[1].Text)
	ass.Equal("C", patDoc.Classifications[1].Section)
	ass.Equal("22", patDoc.Classifications[1].Class)
	ass.Equal("C", patDoc.Classifications[1].SubClass)
	ass.Equal("38", patDoc.Classifications[1].MainGroup)
	ass.Equal("04", patDoc.Classifications[1].SubGroup)
	ass.Equal("L", patDoc.Classifications[1].FirstLater)

	ass.Equal("C22C  38/58        20060101ALI20170607BHEP        ", patDoc.Classifications[2].Text)
	ass.Equal("C", patDoc.Classifications[2].Section)
	ass.Equal("22", patDoc.Classifications[2].Class)
	ass.Equal("C", patDoc.Classifications[2].SubClass)
	ass.Equal("38", patDoc.Classifications[2].MainGroup)
	ass.Equal("58", patDoc.Classifications[2].SubGroup)
	ass.Equal("L", patDoc.Classifications[2].FirstLater)

	ass.Equal("C21D   9/04        20060101ALI20170607BHEP        ", patDoc.Classifications[3].Text)
	ass.Equal("C", patDoc.Classifications[3].Section)
	ass.Equal("21", patDoc.Classifications[3].Class)
	ass.Equal("D", patDoc.Classifications[3].SubClass)
	ass.Equal("9", patDoc.Classifications[3].MainGroup)
	ass.Equal("04", patDoc.Classifications[3].SubGroup)
	ass.Equal("L", patDoc.Classifications[3].FirstLater)

	ass.Equal("C22C  38/02        20060101ALI20170607BHEP        ", patDoc.Classifications[4].Text)
	ass.Equal("C", patDoc.Classifications[4].Section)
	ass.Equal("22", patDoc.Classifications[4].Class)
	ass.Equal("C", patDoc.Classifications[4].SubClass)
	ass.Equal("38", patDoc.Classifications[4].MainGroup)
	ass.Equal("02", patDoc.Classifications[4].SubGroup)
	ass.Equal("L", patDoc.Classifications[4].FirstLater)

	/*fmt.Println(patDoc.Title[0])
	fmt.Println(patDoc.Claims[0])
	fmt.Println(patDoc.Description[0])
	fmt.Println(patDoc.Citations[0])
	*/
}
