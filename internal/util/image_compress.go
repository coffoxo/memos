package utils

import (
    "image"
    "image/jpeg"
    "io"
    "os"

    "github.com/disintegration/imaging"
)

// CompressImage compresses the image and returns a byte slice.
func CompressImage(input io.Reader, output io.Writer, maxWidth int, maxHeight int, quality int) error {
    img, err := imaging.Decode(input)
    if err != nil {
        return err
    }

    // Resize the image to fit the max width/height
    img = imaging.Resize(img, maxWidth, maxHeight, imaging.Lanczos)

    // Create a JPEG encoder with the specified quality
    jpegOptions := jpeg.Options{Quality: quality}
    return jpeg.Encode(output, img, &jpegOptions)
}
