package file

import (
	"path/filepath"
	"os"
	"errors"
	"fmt"
)

func ScanDir(dir string)(files []string, err error){
	f, err := os.Stat(dir)
	if err != nil {
		return
	}

	if !f.IsDir(){
		return nil, errors.New(fmt.Sprint("%s is not a directory\n", dir))
	}

	visit := func (filePath string, f os.FileInfo, err error) error {
		if !f.IsDir(){
			files = append(files, filePath)
		}

		return err
	}

	err = filepath.Walk(dir, visit)
	if err != nil{
		return
	}

	return
}
