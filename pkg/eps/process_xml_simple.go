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
	// title
	/*
		<B540>
			<B541>de</B541>
			<B542>VERFAHREN UND...</B542>
			<B541>en</B541>
			<B542>PROCEDURE AND ...</B542>
		</B540>
	*/
	titles := all.Find("B540")
	titles.Children().Each(func(i int, c *goquery.Selection) {
		if c.Is("B541") && c.Next() != nil && c.Next().Is("B542") {
			patentDoc.Title = append(patentDoc.Title, Title{
				Language: c.Text(),
				Text:     c.Next().Text(),
			})
		}
	})
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
	// citations
	/*
		<patcit id="ref-pcit0001" dnum="US20120281566A">
			<document-id>
				<country>US</country>
				<doc-number>20120281566</doc-number>
				<kind>A</kind>
			</document-id>
		</patcit>
		<crossref idref="pcit0001">[0006]</crossref>
	*/
	citations := all.Find("patcit")
	citations.Each(func(i int, c *goquery.Selection) {
		docId := c.Find("document-id")
		if docId.Size() > 0 {
			patentDoc.Citations = append(patentDoc.Citations, Citation{
				Country:   docId.Find("country").Text(),
				DocNumber: docId.Find("doc-number").Text(),
				Kind:      docId.Find("kind").Text(),
			})
		}
	})
	// inventors
	/*
		<B720>
			<B721>
				<snm>MARTIN, Brian Alexander</snm>
				<adr>
					<str>c/o Sony Europe IP Europe Jays Close, Viables</str>
					<city>Basingstoke, Hampshire RG22 4SB</city>
					<ctry>GB</ctry>
				</adr>
			</B721>


	*/
	inventors := all.Find("B721")
	inventors.Each(func(i int, c *goquery.Selection) {
		patentDoc.Inventors = append(patentDoc.Inventors, Inventor{
			Country: c.Find("adr ctry").Text(),
			City:    c.Find("adr city").Text(),
			Street:  c.Find("adr str").Text(),
			Name:    c.Find("snm").Text(),
		})
	})
	// owners
	/*
		<B720>
			<B721>
				<snm>MARTIN, Brian Alexander</snm>
				<adr>
					<str>c/o Sony Europe IP Europe Jays Close, Viables</str>
					<city>Basingstoke, Hampshire RG22 4SB</city>
					<ctry>GB</ctry>
				</adr>
			</B721>


	*/
	owners := all.Find("B731")
	owners.Each(func(i int, c *goquery.Selection) {
		patentDoc.Owners = append(patentDoc.Owners, Owner{
			IID:     c.Find("iid").Text(),
			IRF:     c.Find("irf").Text(),
			Country: c.Find("adr ctry").Text(),
			City:    c.Find("adr city").Text(),
			Street:  c.Find("adr str").Text(),
			Name:    c.Find("snm").Text(),
		})
	})

	return
}
