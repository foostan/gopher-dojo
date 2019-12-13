package imgconverter

import (
	"os"
	"path/filepath"
)

type Runner struct {
	path string
	fromFormat string
	toFormat string
}

func (r Runner) GetTargetFilePathList() ([]string, error) {
	var pathList []string

	if err := filepath.Walk(r.path, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == "." + r.fromFormat {
			pathList = append(pathList, path)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return pathList, nil
}

func (r Runner) exec() error {

	pathList, err := r.GetTargetFilePathList()
	if err != nil {
		return err
	}

	for _, path := range pathList {
		if err := (Converter {
			srcPath: path,
			dstFormat: r.toFormat,
		}).exec(); err != nil {
			return err
		}
	}

	return nil
}
