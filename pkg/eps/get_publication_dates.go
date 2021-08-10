package eps

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

type PublicationDate struct {
	Date *time.Time
	Name string
	Link string
}

const (
	layoutParsingDate = "2006/01/02"
)

// GetPublicationDates retrieves the publication dates of patents from the endpoint
func GetPublicationDates() (res []PublicationDate, err error) {
	// init http client
	client := NewHttpClient()
	// make request
	url := ENDPOINT_HOST + ENDPOINT_ROOT + "/" + VERSION + "/publication-dates"
	log.Debug("GET: ", url)
	resp, err := client.Get(url)
	if err != nil {
		log.Error(err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		err = errors.New("No 200 status code: " + strconv.Itoa(resp.StatusCode))
		log.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
		return
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the dates and links
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the link and the name
		link, _ := s.Attr("href")
		dateString := s.Text()
		// parse the date form the string
		parsedDate, errDate := time.Parse(layoutParsingDate, dateString)
		if errDate != nil {
			log.Warn("Can not parse dateString: ", dateString, " to layout ", layoutParsingDate, " err:", errDate)
		}
		log.Debug("name: ", dateString, " link: ", link, " date: ", parsedDate)
		// append the date object to the result set
		d := PublicationDate{
			Date: &parsedDate,
			Name: dateString,
			Link: link,
		}
		res = append(res, d)
	})
	return
}
