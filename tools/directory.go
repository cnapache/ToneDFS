package tools

import (
	"io/ioutil"
	"strings"
)

type Directory struct {
	Dir string
}

func NewDirectory(dir string) *Directory {
	return &Directory{Dir: dir}
}

func (me *Directory) GetFilesFilterExt(extFilter string) (files []string, err error) {
	dir, err := ioutil.ReadDir(me.Dir)
	if err != nil {
		return nil, err
	}

	suffix := strings.ToLower(extFilter)

	files = make([]string, 0, 100)

	for _, fi := range dir {
		if fi.IsDir() {
			continue
		}

		if strings.HasSuffix(strings.ToLower(fi.Name()), suffix) {
			files = append(files, fi.Name())
		}
	}

	return files, nil
}
