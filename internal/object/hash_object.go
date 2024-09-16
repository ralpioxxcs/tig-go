package object

import (
	"os"
)

type HashObjectParam struct {
	DryRun bool
	TigDir string
	Type   Type
	Data   []byte
}

func HashObject(param HashObjectParam) (string, error) {
	// Creating Key
	key := newKey(param.Type, param.Data)

	// DryRun -> quick return
	if param.DryRun {
		return string(key), nil
	}

	// Creating contents
	val, err := zlibCompress(newContent(param.Type, param.Data))
	if err != nil {
		return "", err
	}

	// Creating required directory
	if err := os.MkdirAll(key.Dir(param.TigDir), os.ModePerm); err != nil {
		return "", err

	}

	// Creating object
	if err := os.WriteFile(key.Path(param.TigDir), val, os.ModePerm); err != nil {
		return "", err
	}

	return string(key), nil

}
