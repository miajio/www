package lib

import (
	"bytes"
	"errors"
	"image"
	"image/jpeg"
	"image/png"

	"github.com/nfnt/resize"
)

type imageCompressImpl struct{}

type imageCompress interface {
	Compress(buf []byte, w, h uint) ([]byte, error)
}

var ImageCompress imageCompress = (*imageCompressImpl)(nil)

func (*imageCompressImpl) Compress(buf []byte, w, h uint) ([]byte, error) {
	decodeBuf, layout, err := image.Decode(bytes.NewReader(buf))
	if err != nil {
		return nil, err
	}

	set := resize.Resize(w, h, decodeBuf, resize.Lanczos3)
	newBuf := bytes.Buffer{}

	switch layout {
	case "png":
		err = png.Encode(&newBuf, set)
	case "jpeg", "jpg":
		err = jpeg.Encode(&newBuf, set, &jpeg.Options{Quality: 80})
	default:
		return nil, errors.New("the image format currently does't support compression")
	}
	if err != nil {
		return nil, err
	}
	if newBuf.Len() < len(buf) {
		buf = newBuf.Bytes()
	}
	return buf, nil
}
