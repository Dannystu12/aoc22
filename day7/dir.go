package day7

import (
	"fmt"
	"regexp"
)

type dirName string

func (d dirName) validate() error {
	re := regexp.MustCompile(`^[a-zA-Z\d]+$`)
	if !re.MatchString(string(d)) {
		return fmt.Errorf("invalid directory name, can only contain alpha numeric characters: %s", string(d))
	}
	return nil
}

func (d dirName) getKey() FsEntryKey {
	return FsEntryKey(d)
}

type dir struct {
	name    dirName
	entries fsEntryMap
}

func newDir(name dirName) (*dir, error) {
	if err := name.validate(); err != nil {
		return nil, err
	}
	return &dir{
		name:    name,
		entries: make(fsEntryMap),
	}, nil
}

func (d *dir) IsFile() bool {
	return false
}

func (d *dir) IsDir() bool {
	return true
}

func (d *dir) getName() fsEntryName {
	return d.name
}

func (d *dir) GetSize() uint {
	totalSize := uint(0)
	for _, entry := range d.entries {
		totalSize += entry.GetSize()
	}
	return totalSize
}
