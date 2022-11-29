package eps

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"os"
	"strconv"
	"testing"
)

const testID = "EP2921808NWB1"

func TestGetPatentHTML(t *testing.T) {
	ass := assert.New(t)
	res, err := GetPatentHTML(testID)
	ass.NoError(err)
	ass.NotNil(res)
	ass.NotEmpty(res)
	err = SaveFile(res, "./test-data/", string(testID)+".html")
	ass.NoError(err)
}

func TestGetPatentZIP(t *testing.T) {
	ass := assert.New(t)
	res, err := GetPatentZIP(testID)
	ass.NoError(err)
	ass.NotNil(res)
	ass.NotEmpty(res)
	err = SaveFile(res, "./test-data/", string(testID)+".zip")
	ass.NoError(err)
}

func TestGetPatentXML(t *testing.T) {
	ass := assert.New(t)
	res, err := GetPatentXML(testID)
	ass.NoError(err)
	ass.NotNil(res)
	ass.NotEmpty(res)
	err = SaveFile(res, "./test-data/", string(testID)+".xml")
	ass.NoError(err)
}

func TestGetPatentPDF(t *testing.T) {
	ass := assert.New(t)
	res, err := GetPatentPDF(testID)
	ass.NoError(err)
	ass.NotNil(res)
	ass.NotEmpty(res)
	err = SaveFile(res, "./test-data/", string(testID)+".pdf")
	ass.NoError(err)
}

func doReq() (res []byte, err error) {
	// init http client
	client := NewHttpClient()
	// build req
	url := "https://data.epo.org/publication-server/forbidden.html"
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

func TestCheckIfBlackListed(t *testing.T) {
	ass := assert.New(t)
	res, err := doReq()
	ass.NoError(err)
	err = CheckIfBlackListed(res)
	ass.Error(err)
}

func TestCheckIfBlackListedLocal(t *testing.T) {
	ass := assert.New(t)
	file, err := os.ReadFile("./test-data/fail.xml")
	ass.NoError(err)
	err = CheckIfBlackListed(file)
	ass.Error(err)
}
