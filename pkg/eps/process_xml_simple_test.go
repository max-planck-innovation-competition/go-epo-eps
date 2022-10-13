package eps

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

/*
tests:  A1 A2 B1 B2
v 1.0   O  O  X  X
v 1.1   O  O  O  O
v 1.2   O  O  -  -
v 1.3	-  -  -  O
v 1.4   O  -  -  -
v 1.5   O  O  O  O
v 1.5.1 O  O  O  O
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
	for i := 0; i <= 6; i++ {
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

//v 1.3
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
	//1
	ass.Equal("C08F   8/48        20060101AFI20210224BHEP        ", patDoc.Classifications[0].Text)
	ass.Equal("C", patDoc.Classifications[0].Section)
	ass.Equal("08", patDoc.Classifications[0].Class)
	ass.Equal("F", patDoc.Classifications[0].SubClass)
	ass.Equal("8", patDoc.Classifications[0].MainGroup)
	ass.Equal("48", patDoc.Classifications[0].SubGroup)
	ass.Equal("F", patDoc.Classifications[0].FirstLater)
	//2
	ass.Equal("C08F   8/28        20060101ALI20210224BHEP        ", patDoc.Classifications[1].Text)
	ass.Equal("C", patDoc.Classifications[1].Section)
	ass.Equal("08", patDoc.Classifications[1].Class)
	ass.Equal("F", patDoc.Classifications[1].SubClass)
	ass.Equal("8", patDoc.Classifications[1].MainGroup)
	ass.Equal("28", patDoc.Classifications[1].SubGroup)
	ass.Equal("L", patDoc.Classifications[1].FirstLater)
	//3
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

	// abstract !! Empty
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

	fmt.Println(patDoc.Title[0])
	fmt.Println(patDoc.Claims[0])
	fmt.Println(patDoc.Description[0])
	fmt.Println(patDoc.Citations[0])
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

	fmt.Println(patDoc.Title[0])
	fmt.Println(patDoc.Claims[0])
	fmt.Println(patDoc.Description[0])
	fmt.Println(patDoc.Citations[0])
}
