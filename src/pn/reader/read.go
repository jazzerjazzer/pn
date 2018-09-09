package reader

import (
	"bufio"
	"io"
	"os"
)

type Interface interface {
	Read() ([]string, error)
}

type FileReader struct {
	path  string
	batch int
	left  int64
}

func New(path string, batch int) *FileReader {
	return &FileReader{
		path:  path,
		batch: batch,
	}
}

func (f *FileReader) Read() ([]string, error) {
	file, err := os.Open(f.path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	if f.left >= fileInfo.Size() {
		return nil, io.EOF
	}

	if _, err := file.Seek(f.left, 0); err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)

	scanLines := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		advance, token, err = bufio.ScanLines(data, atEOF)
		f.left += int64(advance)
		return
	}
	scanner.Split(scanLines)

	list := []string{}
	for scanner.Scan() {
		text := scanner.Text()
		list = append(list, text)
		if len(list) >= f.batch {
			return list, nil
		}
	}

	return list, scanner.Err()
}
