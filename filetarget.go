package rpl

import "os"

// FileTarget is Target for os.File.
// Used for these types and variables:
//
//  - os.File
//  - os.Stdout
//  - os.Stderr
type FileTarget struct {
	c    chan *Log
	file *os.File
}

func NewFileTarget(file *os.File) *FileTarget {
	fileTarget := FileTarget{
		c:    make(chan *Log),
		file: file,
	}

	go func(ft *FileTarget) {
		for {
			log := <-ft.c
			if log == nil {
				break
			}

			_, _ = ft.file.WriteString(log.Value + "\n")
		}
	}(&fileTarget)

	return &fileTarget
}

func (fileTarget *FileTarget) Writer() chan<- *Log {
	return fileTarget.c
}
