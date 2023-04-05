package eps

import (
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetVersions(t *testing.T) {
	// skip if env is test
	if os.Getenv("ENV") == "TEST" {
		t.Skip()
	}
	log.SetLevel(log.TraceLevel)
	ass := assert.New(t)
	res, err := GetVersions()
	ass.NoError(err)
	ass.Len(res, 3)
	ass.Equal(res[0], "/publication-server/rest/v1.0")
	ass.Equal(res[1], "/publication-server/rest/v1.1")
	ass.Equal(res[2], "/publication-server/rest/v1.2")
}
