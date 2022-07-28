package util

import (
	"bytes"
	"encoding/binary"

	"github.com/HappyGick/filebroadcastprotocol/protocol/common"
)

type ByteBuffer struct {
	ErrorStorer
	data []byte
	err  error
}

func NewByteBuffer(startval []byte) ByteBuffer {
	return ByteBuffer{
		data: startval,
		err:  nil,
	}
}

func (b ByteBuffer) GetError() error {
	return b.err
}

func (b ByteBuffer) String() string {
	if b.err != nil {
		return ""
	}
	return string(b.data)
}

func (b ByteBuffer) Bytes() []byte {
	return b.data
}

func (b ByteBuffer) LEUint() uint64 {
	if b.err != nil {
		return 0
	}
	return binary.LittleEndian.Uint64(b.data)
}

func (b ByteBuffer) BEUint() uint64 {
	if b.err != nil {
		return 0
	}
	return binary.BigEndian.Uint64(b.data)
}

func (b ByteBuffer) Append(array []byte) ByteBuffer {
	return ByteBuffer{
		data: append(b.data, array...),
		err:  b.err,
	}
}

func (b ByteBuffer) AppendMultiple(arrays ...[]byte) ByteBuffer {
	buf := b
	for _, v := range arrays {
		buf = buf.Append(v)
	}
	return buf
}

func (b ByteBuffer) AppendBuffer(buf ByteBuffer) ByteBuffer {
	err := b.err
	if buf.GetError() != nil {
		err = buf.GetError()
	}
	return ByteBuffer{
		err:  err,
		data: append(b.data, buf.Bytes()...),
	}
}

func (b ByteBuffer) AppendUint64BE(n uint64) ByteBuffer {
	numbytes := make([]byte, 8)
	binary.BigEndian.PutUint64(numbytes, n)
	return b.Append(numbytes)
}

func (b ByteBuffer) AppendUint64LE(n uint64) ByteBuffer {
	numbytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(numbytes, n)
	return b.Append(numbytes)
}

func (b ByteBuffer) AppendFile(file common.File) ByteBuffer {
	buf, err := file.Read()

	if err != nil {
		return ByteBuffer{
			data: b.data,
			err:  err,
		}
	}

	return ByteBuffer{
		data: append(b.data, buf...),
		err:  b.err,
	}
}

func (b ByteBuffer) AppendFrom(reader *bytes.Reader, amt uint64) ByteBuffer {
	buf := make([]byte, amt)
	_, err := reader.Read(buf)

	if err != nil {
		return ByteBuffer{
			data: b.data,
			err:  err,
		}
	}

	return ByteBuffer{
		data: append(b.data, buf...),
		err:  b.err,
	}
}

func (b ByteBuffer) AppendUntil(reader *bytes.Reader, marker byte) ByteBuffer {
	buf, err := ReadUntil(reader, marker)

	if err != nil {
		return ByteBuffer{
			data: b.data,
			err:  err,
		}
	}

	return ByteBuffer{
		data: append(b.data, buf...),
		err:  b.err,
	}
}

func (b ByteBuffer) AppendAll(reader *bytes.Reader) ByteBuffer {
	return b.AppendFrom(reader, uint64(reader.Len()))
}
