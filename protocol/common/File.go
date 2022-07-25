package common

import (
	"bytes"
	"compress/gzip"
	"io/fs"
	"os"
)

type File struct {
	path string
	info *fs.FileInfo
}

func FileRef(path string) (File, error) {
	info, err := os.Lstat(path)
	return File{
		path,
		&info,
	}, err
}

func (f File) GetSize() uint64 {
	return uint64((*f.info).Size())
}

func (f File) GetName() string {
	return (*f.info).Name()
}

func (f File) Read() ([]byte, error) {
	buf, err := os.ReadFile(f.path)
	return buf, err
}

func (f File) Compressed() ([]byte, error) {
	var buf bytes.Buffer
	data, err := f.Read()
	if err != nil {
		return nil, err
	}
	w := gzip.NewWriter(&buf)
	w.Write(data)
	w.Close()
	return buf.Bytes(), nil
}
