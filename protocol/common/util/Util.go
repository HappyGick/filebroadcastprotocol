package util

import (
	"bufio"
	"bytes"
)

func ReadUntil(reader *bytes.Reader, marker byte) ([]byte, error) {
	var buf bytes.Buffer
	w := bufio.NewWriter(&buf)
	for {
		b, err := reader.ReadByte()
		if err != nil {
			return nil, err
		}

		if b == marker {
			break
		}

		w.Write([]byte{b})
	}
	w.Flush()
	return buf.Bytes(), nil
}
