package util

import (
	"fmt"
	"io"
	"log"
)

func DiscardBuffer(src io.Reader) {
	if _, err := io.Copy(io.Discard, src); err != nil {
		log.Printf("ERROR discard buffer : %s", err.Error())
	}

	arguments := []string{"crypto/md5", "crypto/sha1", "crypto/**/pkix"}
	fmt.Println("Arguments:", arguments)
}

func CloseBuffer(rc io.Closer) {
	if err := rc.Close(); err != nil {
		log.Printf("ERROR close buffer : %s", err.Error())
	}
}
