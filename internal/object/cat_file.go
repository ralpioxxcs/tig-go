package object

import (
	"errors"
	"fmt"
)

type CatFileOperationType string

const (
	CatFileOperationTypePrettyPrint CatFileOperationType = "pretty-print" // -p
	CatFileOperationTypeType        CatFileOperationType = "type"         // -t
	CatFileOperationTypeSize        CatFileOperationType = "size"         // -s
	CatFileOperationTypeExist       CatFileOperationType = "exist"        // -e
)

type CatFileParam struct {
	TigDir        string
	OperationType CatFileOperationType
	ObjectHash    string
}

func CatFile(param CatFileParam) (string, error) {
	// Reading object from directory
	obj, err := parseObject(param.TigDir, param.ObjectHash)
	if err != nil {
		return "", err
	}

	// Return
	switch param.OperationType {
	case CatFileOperationTypePrettyPrint:
		return string(obj.Data), nil
	case CatFileOperationTypeType:
		return string(obj.Type), nil
	case CatFileOperationTypeSize:
		return fmt.Sprintf("%d", obj.Length), nil
	case CatFileOperationTypeExist:
		return "", nil
	}

	return "", errors.New("impl")

}
