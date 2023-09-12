package helpers

import (
	"bytes"
	"compress/zlib"
	"io"
)

func Compress(data []byte) ([]byte, error) {
	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	_, err := w.Write(data)
	if err != nil {
		return []byte{}, err
	}
	err = w.Close()
	if err != nil {
		return []byte{}, nil
	}

	return b.Bytes(), nil
}

func Decompress(data []byte) ([]byte, error) {
	b := bytes.NewReader(data)

	r, err := zlib.NewReader(b)
	defer r.Close()
	if err != nil {
		return []byte{}, err
	}

	res, err := io.ReadAll(r)
	if err != nil {
		return []byte{}, err
	}
	return res, nil
}
