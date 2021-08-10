package eps

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

type PatentItem struct {
	Name string
	Link string
}

const (
	layoutRetrievingDate = "20060102"
)

// GetPublicationDatePatents retrieves the patent list of a given date
func GetPublicationDatePatents(date time.Time) (res []PatentItem, err error) {
	// generate the date string for the url param from the time object
	urlDateString := date.Format(layoutRetrievingDate)
	// init http client
	client := NewHttpClient()
	// make request
	url := ENDPOINT_HOST + ENDPOINT_ROOT + "/" + VERSION + "/publication-dates/" + urlDateString + "/patents"
	log.Debug("GET: ", url)
	resp, err := client.Get(url)
	if err != nil {
		log.Error(err)
		return
	}
	if resp.StatusCode != 200 {
		err = errors.New("No 200 status code: " + strconv.Itoa(resp.StatusCode))
		log.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
		return
	}
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Error(err)
		return
	}
	// close body
	err = resp.Body.Close()
	if err != nil {
		log.Error(err)
		return
	}
	// Find the dates and links
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the link and the name
		link, _ := s.Attr("href")
		dateString := s.Text()
		log.Debug("name: ", dateString, " link: ", link)
		// append the date object to the result set
		d := PatentItem{
			Name: dateString,
			Link: link,
		}
		res = append(res, d)
	})
	return

}
