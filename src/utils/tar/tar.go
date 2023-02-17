package tar

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"path"

	"github.com/tricorder/src/utils/file"
)

// GZExtract function which extract *.tar.gz compress file to destDir
func GZExtract(tarFilePath string, destDir string) error {
	reader, closer, err := file.Reader(tarFilePath)
	if err != nil {
		return fmt.Errorf("could not reader %s error: %v", tarFilePath, err)
	}

	defer closer.Close()

	uncompressedStream, err := gzip.NewReader(reader)
	if err != nil {
		return fmt.Errorf("could not craete gzip reader %s error: %v", tarFilePath, err)
	}

	tarReader := tar.NewReader(uncompressedStream)
	for {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			return fmt.Errorf("iterate tar %s error: %v", tarFilePath, err)
		}

		switch header.Typeflag {
		case tar.TypeDir:
			continue
		case tar.TypeReg:
			expectedFile := path.Join(destDir, header.Name)
			err := file.Create(expectedFile)
			if err != nil {
				return fmt.Errorf("could not create file %s error: %v", expectedFile, err)
			}
			expectedWriter, expectedCloser, err := file.Writer(expectedFile)
			if err != nil {
				expectedCloser.Close()
				return fmt.Errorf("could not writer %s error: %v", expectedFile, err)
			}
			if _, err := io.Copy(expectedWriter, tarReader); err != nil {
				expectedCloser.Close()
				return fmt.Errorf("could not copy content from %s to %s error: %v", header.Name, expectedFile, err)
			}
			expectedCloser.Close()

		default:
			continue
		}
	}

	return nil
}
