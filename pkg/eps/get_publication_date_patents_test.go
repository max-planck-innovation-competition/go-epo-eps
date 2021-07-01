package eps

import (
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetPublicationDatePatents(t *testing.T) {
	log.SetLevel(log.TraceLevel)
	ass := assert.New(t)
	// time zone
	l, err := time.LoadLocation("Europe/Berlin")
	ass.NoError(err)
	d := time.Date(2021, 06, 30, 0, 0, 0, 0, l)
	// get patents of date YYYY-MM-DD -> 2021-06-30
	res, err := GetPublicationDatePatents(d)
	ass.NoError(err)
	ass.Greater(len(res), 10)
	log.Debug(res)
}
