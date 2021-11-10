package eps

import (
	log "github.com/sirupsen/logrus"
	"regexp"
	"time"
)

// EpPatentDocumentSimple is a simple representation of the xml data
type EpPatentDocumentSimple struct {
	ID                string
	File              string
	Lang              string
	Country           Country
	DocNumber         string
	Kind              string
	DatePubl          time.Time
	Status            string
	DtdVersion        string
	Title             []Title
	Abstract          []Abstract
	Claims            []Claim
	Description       []Description
	Citations         []Citation
	Inventors         []Inventor
	Owners            []Owner
	Representatives   []Representative
	ContractingStates []Country
	Classifications   []ClassificationItem
}

type Country string

type Title struct {
	Text     string
	Language string
}

type Abstract struct {
	Text     string
	Language string
}

type Claim struct {
	Text     string
	Language string
	Id       string
}

type Description struct {
	Text     string
	Language string
}

type Citation struct {
	Country   Country
	DocNumber string
	Kind      string
}

type Inventor struct {
	Country Country
	City    string
	Street  string
	Name    string
}

type Owner struct {
	Country Country
	IID     string
	IRF     string
	City    string
	Street  string
	Name    string
}

type Representative struct {
	Country Country
	IID     string
	City    string
	Street  string
	Name    string
}

type ClassificationItem struct {
	Text                   string
	Sequence               int
	Section                string
	Class                  string
	SubClass               string
	MainGroup              string
	SubGroup               string
	Version                string
	ClassificationLevel    string
	FirstLater             string
	ClassificationValue    string // (inventive or non-inventive)
	ActionDate             string // (inventive or non-inventive)
	OriginalOrReclassified string
	Source                 string
	GeneratingOffice       string
}

var reClassification = regexp.MustCompile(`([ABCDEFGH])([0-9]{1,2})([A-Z]) *([0-9]{1,4})\/([0-9]{1,6}) *([0-9]{8})([CAS])([FL])([IN])([0-9]{8})([BRVD])([HMG])([A-Z]{2}) *`)

func NewClassificationItemFromString(text string, sequence int) (c ClassificationItem) {
	c = ClassificationItem{
		Text:     text,
		Sequence: sequence,
	}
	/*
		e.g. H04W 76/28 20180101AFI20201221BHEP
		e.g. E04B 1/88 20060101AFI20070316BHEP
		1 Section A-H
		2 to 3 ClassificationItem 01-99
		4 Subclass A-Z
		5 to 8 Main Group (right aligned) 1-9999, blank
		9 Separating character / ("Slash")
		10 to 15 Subgroup (left aligned) 00-999999 blank
		16 to 19 For future use 4 blanks
		20 to 27 Version indicator YYYYMMDD date format
		28 Classification level C, A, S
		29 First or later position of symbol F, L
		30 Classification value (inventive or non-inventive) I, N
		31 to 38 Action date YYYYMMDD date format
		39 Original or reclassified data BRVD
		40 Source of classification data H M G
		41-42 Generating office AA ZZ (ST.3)
		43-50 For future use 8 blanks
	*/
	regexRes := reClassification.FindAllStringSubmatch(text, -1)
	if len(regexRes) == 0 {
		log.WithField("text", text).Warn("can not find IPC8 pattern")
		return
	}
	if len(regexRes[0]) != 14 {
		log.WithField("text", text).Warn("can not find IPC8 pattern")
		return
	}
	// 1 Section A-H
	c.Section = regexRes[0][1]
	// 2 to 3 ClassificationItem 01-99
	c.Class = regexRes[0][2]
	// 4 Subclass A-Z
	c.SubClass = regexRes[0][3]
	// 5 to 8 Main Group (right aligned) 1-9999, blank
	c.MainGroup = regexRes[0][4]
	// 9 Separating character / ("Slash")
	// 10 to 15 Subgroup (left aligned) 00-999999 blank
	c.SubGroup = regexRes[0][5]
	// 16 to 19 For future use 4 blanks
	// 20 to 27 Version indicator YYYYMMDD date format
	c.Version = regexRes[0][6]
	// 28 Classification level C, A, S
	c.ClassificationLevel = regexRes[0][7]
	// 29 First or later position of symbol F, L
	c.FirstLater = regexRes[0][8]
	// 30 Classification value (inventive or non-inventive) I, N
	c.ClassificationValue = regexRes[0][9]
	// 31 to 38 Action date YYYYMMDD date format
	c.ActionDate = regexRes[0][10]
	// 39 Original or reclassified data BRVD
	c.OriginalOrReclassified = regexRes[0][11]
	// 40 Source of classification data H M G
	c.Source = regexRes[0][12]
	// 41-42 Generating office AA ZZ (ST.3)
	c.GeneratingOffice = regexRes[0][13]
	// 43-50 For future use 8 blanks

	return c
}
