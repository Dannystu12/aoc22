package day07

import (
	"fmt"
	"regexp"
)

type fileName string

func (f fileName) validate() error {
	re := regexp.MustCompile(`^[a-zA-Z\d]+(\.[a-zA-Z\d]+)?$`)
	if !re.MatchString(string(f)) {
		return fmt.Errorf("invalid file name, can only contain alpha numeric characters and . extension: %s", string(f))
	}
	return nil
}

func (f fileName) getKey() FsEntryKey {
	return FsEntryKey(f)
}

type file struct {
	name fileName
	size uint
}

func newFile(name fileName, size uint) (*file, error) {
	if err := name.validate(); err != nil {
		return nil, err
	}
	return &file{
		name: name,
		size: size,
	}, nil
}

func (f *file) IsFile() bool {
	return true
}

func (f *file) IsDir() bool {
	return false
}

func (f *file) getName() fsEntryName {
	return f.name
}

func (f *file) GetSize() uint {
	return f.size
}
