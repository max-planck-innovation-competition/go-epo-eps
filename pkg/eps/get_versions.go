package eps

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
	"strconv"
)

// GetVersions retrieves the REST API version from the endpoint
func GetVersions() (res []string, err error) {
	// init http client
	client := NewHttpClient()
	// make request
	url := ENDPOINT_HOST + ENDPOINT_ROOT
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

	// Find the links
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the link and the name
		link, _ := s.Attr("href")
		name := s.Text()
		log.Info("name: ", name, " link: ", link)
		res = append(res, link)
	})

	return

}
