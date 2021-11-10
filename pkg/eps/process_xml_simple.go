package eps

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"time"
)

const (
	layoutDatePubl = "20060102"
)

// ProcessXMLSimple transforms the raw response of the xml data into a simple patent
func ProcessXMLSimple(raw []byte) (patentDoc EpPatentDocumentSimple, err error) {
	// parse doc
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(raw))
	if err != nil {
		return
	}
	root := doc.Find("ep-patent-document")
	if root == nil {
		return
	}

	patentDoc.ID, _ = root.Attr("id")
	patentDoc.DocNumber, _ = root.Attr("doc-number")
	patentDoc.Kind, _ = root.Attr("kind")
	patentDoc.Status, _ = root.Attr("status")
	patentDoc.DtdVersion, _ = root.Attr("dtd-version")

	// parse the date form the string
	dateString, _ := root.Attr("date-publ")
	parsedDate, errDate := time.Parse(layoutDatePubl, dateString)
	if errDate != nil {
		log.Warn("can not parse date", dateString)
	} else {
		patentDoc.DatePubl = parsedDate
	}

	patentDoc.Lang, _ = root.Attr("lang")
	patentDoc.File, _ = root.Attr("file")
	country, _ := root.Attr("country")
	patentDoc.Country = Country(strings.TrimSpace(country))
	// title
	/*
		<B540>
			<B541>de</B541>
			<B542>VERFAHREN UND...</B542>
			<B541>en</B541>
			<B542>PROCEDURE AND ...</B542>
		</B540>
	*/
	titles := root.Find("B540")
	titles.Children().Each(func(i int, c *goquery.Selection) {
		if c.Is("B541") && c.Next() != nil && c.Next().Is("B542") {
			patentDoc.Title = append(patentDoc.Title, Title{
				Language: strings.ToLower(strings.TrimSpace(c.Text())),
				Text:     strings.TrimSpace(c.Next().Text()),
			})
		}
	})
	// abstract
	abstract := root.Find("abstract")
	patentDoc.Abstract = append(
		patentDoc.Abstract,
		Abstract{
			Text:     strings.TrimSpace(abstract.Text()),
			Language: strings.ToLower(strings.TrimSpace("en")),
		},
	)
	// description
	description := root.Find("description")
	patentDoc.Description = append(
		patentDoc.Description,
		Description{
			Text:     strings.TrimSpace(description.Text()),
			Language: strings.ToLower(strings.TrimSpace("en")),
		})
	// claims
	claims := root.Find("claims")
	claims.Each(func(i int, c *goquery.Selection) {
		lang, _ := c.Attr("lang")
		id, _ := c.Attr("id")
		patentDoc.Claims = append(patentDoc.Claims, Claim{
			Text:     strings.TrimSpace(c.Text()),
			Language: strings.TrimSpace(strings.ToLower(lang)),
			Id:       id,
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
	citations := root.Find("patcit")
	citations.Each(func(i int, c *goquery.Selection) {
		docId := c.Find("document-id")
		if docId.Size() > 0 {
			patentDoc.Citations = append(patentDoc.Citations, Citation{
				Country:   Country(strings.TrimSpace(docId.Find("country").Text())),
				DocNumber: strings.TrimSpace(docId.Find("doc-number").Text()),
				Kind:      strings.TrimSpace(docId.Find("kind").Text()),
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
	inventors := root.Find("B721")
	inventors.Each(func(i int, c *goquery.Selection) {
		patentDoc.Inventors = append(patentDoc.Inventors, Inventor{
			Country: Country(strings.TrimSpace(c.Find("adr ctry").Text())),
			City:    strings.TrimSpace(c.Find("adr city").Text()),
			Street:  strings.TrimSpace(c.Find("adr str").Text()),
			Name:    strings.TrimSpace(c.Find("snm").Text()),
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
	owners := root.Find("B731")
	owners.Each(func(i int, c *goquery.Selection) {
		patentDoc.Owners = append(patentDoc.Owners, Owner{
			Country: Country(strings.TrimSpace(c.Find("adr ctry").Text())),
			IID:     strings.TrimSpace(c.Find("iid").Text()),
			IRF:     strings.TrimSpace(c.Find("irf").Text()),
			City:    strings.TrimSpace(c.Find("adr city").Text()),
			Street:  strings.TrimSpace(c.Find("adr str").Text()),
			Name:    strings.TrimSpace(c.Find("snm").Text()),
		})
	})
	// representatives
	/*
		<B740>
			<B741>
				<snm>D Young & Co LLP</snm>
				<iid>101533551</iid>
				<adr>
					<str>120 Holborn</str>
					<city>London EC1N 2DY</city>
					<ctry>GB</ctry>
				</adr>
			</B741>
		</B740>
	*/
	representatives := root.Find("B741")
	representatives.Each(func(i int, c *goquery.Selection) {
		patentDoc.Representatives = append(patentDoc.Representatives, Representative{
			Country: Country(strings.TrimSpace(c.Find("adr ctry").Text())),
			IID:     strings.TrimSpace(c.Find("iid").Text()),
			City:    strings.TrimSpace(c.Find("adr city").Text()),
			Street:  strings.TrimSpace(c.Find("adr str").Text()),
			Name:    strings.TrimSpace(c.Find("snm").Text()),
		})
	})
	// ContractingStates
	/*
		<B800>
			<B840>
				<ctry>AL</ctry>
				<ctry>AT</ctry>

	*/
	countries := root.Find("B840 ctry")
	countries.Each(func(i int, c *goquery.Selection) {
		patentDoc.ContractingStates = append(patentDoc.ContractingStates, Country(strings.TrimSpace(c.Text())))
	})
	// Classifications
	/*
		<B510EP>
			<classification-ipcr sequence="1">
				<text>B60T 17/22 20060101AFI20200403BHEP</text>
			</classification-ipcr>
		</B510EP>
	*/
	classes := root.Find("B510EP classification-ipcr")
	classes.Each(func(i int, c *goquery.Selection) {
		seq, ex := c.Attr("sequence")
		if !ex {
			log.Warn("classification ipcr: seq does not exist")
		}
		seqInt, warn := strconv.Atoi(seq)
		if warn != nil {
			log.Warn("classification ipcr: can not parse seq string", warn)
		}
		// do not use trim here
		item := NewClassificationItemFromString(c.Find("text").Text(), seqInt)
		patentDoc.Classifications = append(patentDoc.Classifications, item)
	})
	return
}
