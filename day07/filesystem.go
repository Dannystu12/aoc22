package day07

import (
	"fmt"
)

type FsEntryKey string

type fsEntryName interface {
	validate() error
	getKey() FsEntryKey
}

type fsEntryMap map[FsEntryKey]FsEntry

func (fs fsEntryMap) traverse(fn func(key FsEntryKey, entry FsEntry)) {
	for key, entry := range fs {
		fn(key, entry)
		if d, ok := entry.(*dir); ok {
			d.entries.traverse(fn)
		}
	}
}

type simpleFS struct {
	capacity        uint
	entries         fsEntryMap
	currentLocation []dirName
}

func newSimpleFS(capacity uint) *simpleFS {
	return &simpleFS{
		capacity:        capacity,
		entries:         make(fsEntryMap),
		currentLocation: make([]dirName, 0),
	}
}

func (fs *simpleFS) getCurrentDirectoryContents() (fsEntryMap, error) {
	currentDirectory := fs.entries
	for _, dName := range fs.currentLocation {
		entry, ok := currentDirectory[dName.getKey()]
		if !ok {
			return nil, fmt.Errorf("current directory does not exist %q not found", dName)
		}

		if entry == nil {
			return nil, fmt.Errorf("directory is nil %q", dName)
		}

		cur, ok := entry.(*dir)
		if !ok {
			return nil, fmt.Errorf("%q is not a directory", dName)
		}

		currentDirectory = cur.entries
	}

	return currentDirectory, nil
}

func (fs *simpleFS) addEntry(entry FsEntry) error {
	if entry == nil {
		return fmt.Errorf("entry is nil")
	}

	entries, err := fs.getCurrentDirectoryContents()
	if err != nil {
		return fmt.Errorf("could not get current directory: %w", err)
	}

	if _, exists := entries[entry.getName().getKey()]; exists {
		return fmt.Errorf("entry already exists: %s", entry.getName())
	}

	currentSpaceUsed := fs.getSize()
	if entry.GetSize()+currentSpaceUsed > fs.capacity {
		return fmt.Errorf("no space available for entry: %s", entry.getName())
	}

	entries[entry.getName().getKey()] = entry
	return nil
}

func (fs *simpleFS) getSize() uint {
	total := uint(0)
	for _, entry := range fs.entries {
		total += entry.GetSize()
	}
	return total
}

func (fs *simpleFS) RecommendDirectoryForDeletion(requiredSpace uint) (*dir, error) {

	if requiredSpace > fs.capacity {
		return nil, fmt.Errorf("required space is greater than capacity")
	}

	currentSpaceUsed := fs.getSize()
	freeSpace := fs.capacity - currentSpaceUsed
	if requiredSpace <= freeSpace {
		return nil, fmt.Errorf("already have enough space")
	}

	requiredSpace = requiredSpace - freeSpace

	var result *dir
	resSize := uint(0)
	fs.Traverse(func(key FsEntryKey, entry FsEntry) {
		if entry.IsDir() {
			entrySize := entry.GetSize()
			if entrySize >= requiredSpace {
				if result == nil || resSize > entrySize {
					result = entry.(*dir)
					resSize = entrySize
				}
			}
		}
	})

	if result == nil {
		return nil, fmt.Errorf("no directories found for deletion")
	}
	return result, nil

}

func (fs *simpleFS) cd(directoryName string) error {
	if directoryName == ".." {
		if len(fs.currentLocation) == 0 {
			return fmt.Errorf("cannot cd to parent of root")
		}

		fs.currentLocation = fs.currentLocation[:len(fs.currentLocation)-1]
		return nil
	}

	if directoryName == "/" {
		fs.currentLocation = []dirName{}
		return nil
	}

	currentDirectory, err := fs.getCurrentDirectoryContents()
	if err != nil {
		return fmt.Errorf("could not get current directory: %w", err)
	}

	entry, ok := currentDirectory[dirName(directoryName).getKey()]
	if !ok {
		return fmt.Errorf("directory does not exist %q", directoryName)
	}

	if !entry.IsDir() {
		return fmt.Errorf("%q is not a directory", directoryName)
	}

	fs.currentLocation = append(fs.currentLocation, dirName(directoryName))
	return nil
}

func (fs *simpleFS) Traverse(fn func(key FsEntryKey, entry FsEntry)) {
	fs.entries.traverse(fn)
}

type FsEntry interface {
	IsFile() bool
	IsDir() bool
	getName() fsEntryName
	GetSize() uint
}
