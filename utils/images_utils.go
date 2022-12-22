package utils

import (
	"bytes"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"

	"github.com/nfnt/resize"
)

type ImageFormat struct {
	Format string
	Width  int
	Height int
}

// ResizeImage resizes an image to the specified width and height using the
// Lanczos resampling algorithm and returns the resized image as an io.Reader.
func (c *ImageFormat) ResizeImage(src io.Reader, width, height uint) (io.Reader, error) {
	// Decode the image.
	img, _, err := image.Decode(src)
	if err != nil {
		return nil, fmt.Errorf("decode image: %w", err)
	}

	// Resize the image.
	resizedImg := resize.Resize(width, height, img, resize.Lanczos3)

	// Encode the resized image as a JPEG.
	buf := new(bytes.Buffer)
	if err := jpeg.Encode(buf, resizedImg, nil); err != nil {
		return nil, fmt.Errorf("encode image: %w", err)
	}

	return buf, nil
}

// convertImageFormat converts the image in srcFile to the desired format and
// writes it to dst.
func (c *ImageFormat) convertImageFormat(srcFile io.Reader, format string) (io.Reader, error) {
	// Decode the source image
	srcImage, _, err := image.Decode(srcFile)
	if err != nil {
		return nil, err
	}

	// Create a buffer to hold the converted image
	dst := new(bytes.Buffer)

	// Encode the image in the desired format
	switch format {
	case "jpeg":
		err = jpeg.Encode(dst, srcImage, nil)
	case "png":
		err = png.Encode(dst, srcImage)
	case "gif":
		err = gif.Encode(dst, srcImage, nil)
	default:
		err = fmt.Errorf("unsupported image format: %s", format)
	}
	if err != nil {
		return nil, err
	}

	return dst, nil
}
