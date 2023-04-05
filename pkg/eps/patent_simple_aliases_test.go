package eps

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestEpPatentDocumentSimple_GenerateAliases(t *testing.T) {
	ass := assert.New(t)
	data, err := os.ReadFile("./test-data/application/v1-5-A1.xml")
	ass.NoError(err)
	if err != nil {
		t.Error(err)
		return
	}
	patDoc, err := ProcessXMLSimple(data)
	ass.Equal("EP3782854A1", patDoc.Aliases[0])
}
