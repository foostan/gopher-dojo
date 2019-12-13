package imgconverter

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

type Converter struct {
	srcPath string
	dstFormat string
}

func (c Converter) exec() error {
	file, err := os.Open(c.srcPath)
	if err != nil {
		return err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	out, err := os.Create(c.DstPath())
	if err != nil {
		return err
	}
	defer out.Close()

	switch c.dstFormat {
	case "jpg", "jpeg" : jpeg.Encode(out, img, nil)
	case "png" : png.Encode(out, img)
	case "gif" : gif.Encode(out, img, nil)
	}

	return nil
}

func (c Converter) DstPath() string {
	name := filepath.Base(c.srcPath[:len(c.srcPath)-len(filepath.Ext(c.srcPath))])
	return filepath.Dir(c.srcPath) + "/" + name + "." + c.dstFormat
}

