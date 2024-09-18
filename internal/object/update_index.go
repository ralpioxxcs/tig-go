package object

import (
	"errors"
	"os"
)

type Entry struct {
	Mode       os.FileMode
	File       string
	ObejctHash string
}

type UpdateIndexParam struct {
	Files  []string
	Caches []*Entry
	TigDir string
	Add    bool
}

func UpdateIndex(param UpdateIndexParam) error {
	// Open the index file

	if len(param.Caches) != 0 {
		// If it runs Caches

		// Searching in Index

		// If it exists, update to new cache

		// If Add options is off, return error

		// Add index

		// Reflect file
	}

	// If it run as file

	// Check it is actually exist by looping file paths

	// Creating blob object by reading file contents

	// Search by index

	return errors.New("implement me")

}
