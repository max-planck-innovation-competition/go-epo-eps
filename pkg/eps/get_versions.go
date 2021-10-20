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
	reqUrl := ENDPOINT_HOST + ENDPOINT_ROOT
	log.Debug("GET: ", reqUrl)
	resp, err := client.Get(reqUrl)
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
	// Find the links
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the link and the name
		link, _ := s.Attr("href")
		name := s.Text()
		log.Info("name: ", name, " link: ", link)
		res = append(res, link)
	})
	// close body
	err = resp.Body.Close()
	if err != nil {
		log.Error(err)
		return
	}
	return
}
