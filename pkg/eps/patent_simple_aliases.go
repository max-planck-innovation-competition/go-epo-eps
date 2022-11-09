package eps

import "strings"

// GenerateAliases generates the aliases for the document
func (p *EpPatentDocumentSimple) GenerateAliases() {
	p.Aliases = []string{
		"EP" + strings.ToUpper(p.DocNumber) + strings.ToUpper(p.Kind), // e.g. EP1234567A1
	}
}
