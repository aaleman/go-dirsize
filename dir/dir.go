package dir

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	ByteUnit  = 1024
	ByteCount = "KMGTPE"
)

type Entry struct {
	Name  string
	Path  string
	IsDir bool
	Size  int64
	Files []Entry
}

func (e *Entry) String() string {
	var sb strings.Builder
	sb.WriteString(e.Name)
	sb.WriteString(" ")
	sb.WriteString(HumanSize(e.Size))

	return sb.String()
}

func (e *Entry) PrintRec() {
	fmt.Println(e)
	if e.IsDir {
		for _, f := range e.Files {
			f.PrintRec()
		}
	}
}

func (e *Entry) add(newEntry Entry) {
	e.Files = append(e.Files, newEntry)
	e.Size += newEntry.Size
}

func ReadFolder(name string) *Entry {
	dirEntry := &Entry{
		Name:  name,
		Path:  name,
		IsDir: true,
		Size:  0,
		Files: []Entry{},
	}

	files, err := os.ReadDir(name)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		filePath := name + "/" + file.Name()
		isDir := file.IsDir()

		if file.IsDir() {
			subfolder := ReadFolder(filePath)
			dirEntry.add(*subfolder)
		} else {
			info, _ := file.Info()
			// if err != nil {
			// 	return 0
			// }
			infoSize := info.Size()
			fileEntry := &Entry{
				Name:  file.Name(),
				Path:  filePath,
				IsDir: isDir,
				Size:  infoSize,
				Files: nil,
			}
			dirEntry.Size += infoSize
			dirEntry.Files = append(dirEntry.Files, *fileEntry)
		}

	}

	return dirEntry
}

func HumanSize(b int64) string {
	if b < ByteUnit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(ByteUnit), 0
	for n := b / ByteUnit; n >= ByteUnit; n /= ByteUnit {
		div *= ByteUnit
		exp++
	}
	return fmt.Sprintf("%.2f %ciB", float64(b)/float64(div), ByteCount[exp])
}
