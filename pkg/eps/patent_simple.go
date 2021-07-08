package eps

type EpPatentDocumentSimple struct {
	ID          string `xml:"id,attr"`
	File        string `xml:"file,attr"`
	Lang        string `xml:"lang,attr"`
	Country     string `xml:"country,attr"`
	DocNumber   string `xml:"doc-number,attr"`
	Kind        string `xml:"kind,attr"`
	DatePubl    string `xml:"date-publ,attr"`
	Status      string `xml:"status,attr"`
	DtdVersion  string `xml:"dtd-version,attr"`
	Title       []Title
	Abstract    []Abstract
	Claims      []Claim
	Description []Description
}

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
	Num      string
}

type Description struct {
	Text     string
	Language string
}

type Citation struct {
	Country   string
	DocNumber string
	Kind      string
}
