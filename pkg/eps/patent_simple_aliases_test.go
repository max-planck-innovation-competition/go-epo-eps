package eps

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestEpPatentDocumentSimple_GenerateAliases(t *testing.T) {
	ass := assert.New(t)
	data, err := os.ReadFile("../../dtds/1-5-example.xml")
	ass.NoError(err)
	patDoc, err := ProcessXMLSimple(data)
	ass.Equal("EP3530531B1", patDoc.Aliases[0])
}
