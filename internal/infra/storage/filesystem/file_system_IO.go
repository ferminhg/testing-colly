package filesystem

import (
	"bytes"
	"os"
)

const OutputDir = "output/"

func WriteBuffer2File(outputFile string, buf *bytes.Buffer) error {
	f, err := os.OpenFile(outputFile, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.Write(buf.Bytes()); err != nil {
		return err
	}
	return nil
}
