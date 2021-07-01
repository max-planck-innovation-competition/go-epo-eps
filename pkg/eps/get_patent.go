package eps

import (
	"errors"
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
	client := &http.Client{}
	// build req
	url := ENDPOINT_HOST + ENDPOINT_ROOT + "/" + VERSION + "/patents/" + patentID + "/document." + strings.ToLower(string(format))
	log.Debug("GET: ", url)
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
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		err = errors.New("No 200 status code: " + strconv.Itoa(resp.StatusCode))
		log.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
		return
	}
	res, err = io.ReadAll(resp.Body)
	return
}

// GetPatentXML returns the patent in the xml format
func GetPatentXML(patentID string) (res []byte, err error) {
	return getPatent(patentID, XML)
}

// GetPatentHTML returns the patent in the html format
func GetPatentHTML(patentID string) (res []byte, err error) {
	return getPatent(patentID, HTML)
}

// GetPatentZIP returns the patent in the zip format
func GetPatentZIP(patentID string) (res []byte, err error) {
	return getPatent(patentID, ZIP)
}

// GetPatentPDF returns the patent in the pdf format
func GetPatentPDF(patentID string) (res []byte, err error) {
	return getPatent(patentID, PDF)
}
