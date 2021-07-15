package eps

import "time"

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
	Classes           []Class
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
	Num      string
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

type Class struct {
	Text     string
	Sequence int
}
