package eps

import (
	"bytes"
	"errors"
	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strconv"
	"strings"
)

// PatentExportFormat is a predefined export format
type PatentExportFormat string

// Current forms of export
const (
	PDF  PatentExportFormat = "PDF"
	HTML PatentExportFormat = "HTML"
	ZIP  PatentExportFormat = "ZIP"
	XML  PatentExportFormat = "XML"
)

// getPatent executes the http request using the id and the export format
func getPatent(patentID string, format PatentExportFormat) (res []byte, err error) {
	// init http client
	client := NewHttpClient()
	// build req
	reqUrl := ENDPOINT_HOST + ENDPOINT_ROOT + "/" + VERSION + "/patents/" + patentID + "/document." + strings.ToLower(string(format))
	log.Debug("GET: ", reqUrl)
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		log.Error(err)
		return
	}
	// add header
	req.Header.Add("user-agent", "raw")
	// make request
	resp, err := client.Do(req)
	if err != nil {
		log.Error(err)
		return
	}
	if resp.StatusCode != 200 {
		err = errors.New("No 200 status code: " + strconv.Itoa(resp.StatusCode))
		log.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
		return
	}
	res, err = io.ReadAll(resp.Body)
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
	// check if blacklisted
	err = CheckIfBlackListed(res)
	if err != nil {
		log.Error(err)
		return
	}
	return
}

// GetPatentXML returns the patent in the xml format
func GetPatentXML(patentID string) (res []byte, err error) {
	return getPatent(patentID, XML)
}

// GetPatentHTML returns the patent in the html format
func GetPatentHTML(patentID string) (res []byte, err error) {
	initialResponse, err := getPatent(patentID, HTML)
	if err != nil {
		log.Error("GetPatentXML: can not get initial response", err)
		return
	}
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(initialResponse))
	if err != nil {
		log.Error("GetPatentXML: can not read document", err)
		return
	}
	// find the iframe on the website
	iframes := doc.Find("#documentCenter")
	url, exists := iframes.First().Attr("src")
	if !exists {
		err = errors.New("can not find iframe")
		log.Error("GetPatentXML:", err)
		return
	}
	// now perform the second request
	// init http client
	client := NewHttpClient()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Error(err)
		return
	}
	// add header
	req.Header.Add("user-agent", "raw")
	// make request
	resp, err := client.Do(req)
	if err != nil {
		log.Error(err)
		return
	}
	if resp.StatusCode != 200 {
		err = errors.New("No 200 status code: " + strconv.Itoa(resp.StatusCode))
		log.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
		return
	}
	res, err = io.ReadAll(resp.Body)
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
	return
}

// GetPatentZIP returns the patent in the zip format
func GetPatentZIP(patentID string) (res []byte, err error) {
	return getPatent(patentID, ZIP)
}

// GetPatentPDF returns the patent in the pdf format
func GetPatentPDF(patentID string) (res []byte, err error) {
	initialResponse, err := getPatent(patentID, PDF)
	if err != nil {
		log.Error("GetPatentXML: can not get initial response", err)
		return
	}
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(initialResponse))
	if err != nil {
		log.Error("GetPatentXML: can not read document", err)
		return
	}
	// find the iframe on the website
	iframes := doc.Find("#body > div.epoToolBar.document > ul > li > a")
	url, exists := iframes.First().Attr("href")
	if !exists {
		err = errors.New("can not find iframe")
		log.Error("GetPatentXML:", err)
		return
	}
	// now perform the second request
	// init http client
	client := NewHttpClient()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Error(err)
		return
	}
	// add header
	req.Header.Add("user-agent", "raw")
	// make request
	resp, err := client.Do(req)
	if err != nil {
		log.Error(err)
		return
	}
	if resp.StatusCode != 200 {
		err = errors.New("No 200 status code: " + strconv.Itoa(resp.StatusCode))
		log.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
		return
	}
	res, err = io.ReadAll(resp.Body)
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
	// check if blacklisted
	err = CheckIfBlackListed(res)
	if err != nil {
		log.Error(err)
		return
	}
	return
}

func CheckIfBlackListed(res []byte) (err error) {
	bodyString := string(res)
	if strings.Contains(bodyString, "The European publication server has detected a very high level of data flow to your IP address. Such traffic could potentially disturb the access to the service for other users.") {
		err = errors.New("client and IP blacklisted")
	}
	return
}
