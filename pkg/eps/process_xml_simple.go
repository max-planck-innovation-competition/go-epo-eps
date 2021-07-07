package eps

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
)

// ProcessXMLSimple transforms the raw response of the xml data into a simple patent
func ProcessXMLSimple(raw []byte) (patentDoc EpPatentDocumentSimple, err error) {
	// parse doc
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(raw))
	if err != nil {
		return
	}
	all := doc.Find("ep-patent-document")
	if all == nil {
		return
	}
	patentDoc.ID, _ = all.Attr("id")
	patentDoc.DocNumber, _ = all.Attr("doc-number")
	patentDoc.Lang, _ = all.Attr("lang")
	patentDoc.Country, _ = all.Attr("country")
	// abstract
	abstract := all.Find("abstract")
	abstract.Each(func(i int, a *goquery.Selection) {
		lang, _ := a.Attr("lang")
		patentDoc.Abstract = append(patentDoc.Abstract, Abstract{Text: a.Text(), Language: lang})
	})
	// description
	description := all.Find("description")
	description.Each(func(i int, d *goquery.Selection) {
		lang, _ := d.Attr("lang")
		patentDoc.Description = append(patentDoc.Description, Description{Text: d.Text(), Language: lang})
	})
	// claims
	claims := all.Find("claims")
	claims.Each(func(i int, c *goquery.Selection) {
		lang, _ := c.Attr("lang")
		id, _ := c.Attr("id")
		num, _ := c.Attr("num")
		patentDoc.Claims = append(patentDoc.Claims, Claim{
			Text:     c.Text(),
			Language: lang,
			Id:       id,
			Num:      num,
		})
	})

	return
}
