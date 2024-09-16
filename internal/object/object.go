package object

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

const (
	objectsDirName = "objects"
)

type Type string

const (
	Blob Type = "blob"
)

type Key string

func (k Key) Dir(tigDir string) string {
	return filepath.Join(tigDir, objectsDirName, string(k[:2]))
}

func (k Key) Path(tigDir string) string {
	return filepath.Join(k.Dir(tigDir), string(k[2:]))
}

func newKey(t Type, data []byte) Key {
	// [type, length, \0, data] -> SHA-1 Hash
	str := newContent(t, data)

	h := sha1.New()
	h.Write(str)
	return Key(fmt.Sprintf("%x", h.Sum(nil)))

}

func newContent(t Type, data []byte) []byte {
	// [Type] [Length of Type]\0[Data]
	return []byte(fmt.Sprintf("%s %d\000%s", t, len(data), data))
}

func zlibCompress(data []byte) ([]byte, error) {
	b := new(bytes.Buffer)

	// zlib write, write to Buffer
	w := zlib.NewWriter(b)
	if _, err := w.Write(data); err != nil {
		return nil, err
	}

	err := w.Close()
	return b.Bytes(), err
}

type Object struct {
	Type   Type
	Length int
	Data   []byte
}

func parseObject(tigDir string, objectHash string) (Object, error) {
	path := Key(objectHash).Path(tigDir)
	data, err := os.ReadFile(path)
	if err != nil {
		return Object{}, err
	}

	var obj Object
	if err := UnmarshalObject(data, &obj); err != nil {
		return Object{}, err

	}

	return obj, nil

}

func UnmarshalObject(data []byte, o *Object) error {
	r, err := zlib.NewReader(bytes.NewReader(data))
	if err != nil {
		return err
	}

	decryption := new(bytes.Buffer)
	if _, err := decryption.ReadFrom(r); err != nil {
		return err
	}

	// HEADER\000BODY
	headerAndBody := bytes.Split(decryption.Bytes(), []byte{'\x00'})
	if len(headerAndBody) != 2 {
		return errors.New("invalid object format")
	}

	var (
		objectType   Type
		objectLength int
	)

	if _, err := fmt.Sscanf(string(headerAndBody[0]), "%s %d", &objectType, &objectLength); err != nil {
		return err
	}

	o.Type = objectType
	o.Length = objectLength
	o.Data = headerAndBody[1]

	return nil
}
