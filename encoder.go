package webpbin

import (
	"image"
	"io"
)

type Encoder struct {
	//Specify the compression factor for RGB channels between 0 and 100. The default is 75.
	//
	//A small factor produces a smaller file with lower quality. Best quality is achieved by using a value of 100.
	Quality uint
}

//Writes the Image m to w in WebP format. Any Image may be encoded
func (e *Encoder) Encode(w io.Writer, m image.Image) error {
	return NewCWebP().
		Quality(e.Quality).
		InputImage(m).
		Output(w).
		Run()

}

//Writes the Image m to w in WebP format. Any Image may be encoded
func Encode(w io.Writer, m image.Image) error {
	e := &Encoder{Quality: 75}
	return e.Encode(w, m)
}
