package imgconverter

import (
	"flag"
)

var path = flag.String("path", ".", "path")
var fromFormat = flag.String("from_format", "jpg", "convert format from")
var toFormat = flag.String("to_format", "png", "convert format to")

func Run() error {
	flag.Parse()
	return (Runner {
		path: *path,
		fromFormat: *fromFormat,
		toFormat: *toFormat,
	}).exec()
}
