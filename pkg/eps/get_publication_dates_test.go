package eps

import (
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetPublicationDates(t *testing.T) {
	log.SetLevel(log.TraceLevel)
	ass := assert.New(t)
	res, err := GetPublicationDates()
	ass.NoError(err)
	ass.NotNil(res)
	ass.Greater(len(res), 10)
	log.Println(res)
}
