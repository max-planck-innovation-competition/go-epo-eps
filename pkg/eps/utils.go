package eps

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func SaveFile(data []byte, filepath, filename string) (err error) {
	// create file
	f, err := os.Create(filepath + filename)
	if err != nil {
		log.Error(err)
		return
	}
	// write
	_, err = f.Write(data)
	if err != nil {
		log.Error(err)
		return
	}
	// close
	err = f.Close()
	if err != nil {
		log.Error(err)
		return
	}
	return
}
