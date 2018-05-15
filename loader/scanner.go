package loader

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

type ScannerExt interface {
	Scan() ([]string, error)
	ScanNotUpdated(lastUpdate time.Time) ([]string, error)
}

type scanner struct {
	Path string
}

func NewScanner(path string) ScannerExt {
	return &scanner{
		Path: path,
	}
}

func (s *scanner) ls(path string) ([]os.FileInfo, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	return files, nil
}

func enqueue(list []string, item string) []string {
	return append(list, item)
}

func dequeue(list []string) (string, []string) {
	if len(list) > 0 {
		return list[0], list[1:]
	} else {
		panic("ERROR in scanner: list have no element")
	}
}

func (s *scanner) Scan() ([]string, error) {
	dirList := []string{}
	songList := []string{}

	dirList = enqueue(dirList, s.Path)
	var path string
	for len(dirList) != 0 {
		path, dirList = dequeue(dirList)

		files, _ := ioutil.ReadDir(string(path))
		for _, item := range files {
			if item.IsDir() == true {
				dirList = enqueue(dirList, filepath.Join(path, item.Name()))
			} else {
				songList = enqueue(songList, filepath.Join(path, item.Name()))
			}
		}
	}

	return songList, nil
}

func (s *scanner) ScanNotUpdated(lastUpdate time.Time) ([]string, error) {
	songList, err := s.Scan()
	notUpdatedList := []string{}
	if err != nil {
		return nil, err
	}

	for _, song := range songList {
		info, err := os.Stat(song)
		if err != nil {
			return nil, err
		}

		if !info.ModTime().Before(lastUpdate) {
			notUpdatedList = append(notUpdatedList, song)
		}
	}

	return notUpdatedList, nil
}
