package webpbin

import (
	"errors"
	"fmt"
	"image/gif"
	"io"

	"github.com/nickalie/go-binwrapper"
)

// Gif2WebP converts a GIF image to a WebP image.
// https://developers.google.com/speed/webp/docs/gif2webp
type Gif2WebP struct {
	*binwrapper.BinWrapper
	inputFile  string
	inputGif   *gif.GIF
	input      io.Reader
	outputFile string
	output     io.Writer
	quality    int
}

// NewGif2WebP creates new Gif2WebP instance.
func NewGif2WebP(optionFuncs ...OptionFunc) *Gif2WebP {
	bin := &Gif2WebP{
		BinWrapper: createBinWrapper(optionFuncs...),
		quality:    -1,
	}
	bin.ExecPath("gif2webp")

	return bin
}

// Version returns cwebp version.
func (g *Gif2WebP) Version() (string, error) {
	return version(g.BinWrapper)
}

// InputFile sets image file to convert.
// Input or InputImage called before will be ignored.
func (g *Gif2WebP) InputFile(file string) *Gif2WebP {
	g.input = nil
	g.inputGif = nil
	g.inputFile = file
	return g
}

// Input sets reader to convert.
// InputFile or InputImage called before will be ignored.
func (g *Gif2WebP) Input(reader io.Reader) *Gif2WebP {
	g.inputFile = ""
	g.inputGif = nil
	g.input = reader
	return g
}

// InputGif sets gif to convert.
// InputFile or Input called before will be ignored.
func (g *Gif2WebP) InputGif(gif *gif.GIF) *Gif2WebP {
	g.inputFile = ""
	g.input = nil
	g.inputGif = gif
	return g
}

// OutputFile specify the name of the output WebP file.
// Output called before will be ignored.
func (g *Gif2WebP) OutputFile(file string) *Gif2WebP {
	g.output = nil
	g.outputFile = file
	return g
}

// Output specify writer to write webp file content.
// OutputFile called before will be ignored.
func (g *Gif2WebP) Output(writer io.Writer) *Gif2WebP {
	g.outputFile = ""
	g.output = writer
	return g
}

// Quality specify the compression factor for RGB channels between 0 and 100. The default is 75.
//
// A small factor produces a smaller file with lower quality. Best quality is achieved by using a value of 100.
func (g *Gif2WebP) Quality(quality uint) *Gif2WebP {
	if quality > 100 {
		quality = 100
	}

	g.quality = int(quality)
	return g
}

// Run starts gif2webp with specified parameters.
func (g *Gif2WebP) Run() error {
	defer g.BinWrapper.Reset()

	g.Arg("-min_size")
	g.Arg("-mt")

	if g.quality > -1 {
		g.Arg("-q", fmt.Sprintf("%d", g.quality))
	}

	output, err := g.getOutput()

	if err != nil {
		return err
	}

	g.Arg("-o", output)

	err = g.setInput()

	if err != nil {
		return err
	}

	if g.output != nil {
		g.SetStdOut(g.output)
	}

	err = g.BinWrapper.Run()

	if err != nil {
		return errors.New(err.Error() + ". " + string(g.StdErr()))
	}

	return nil
}

// Reset all parameters to default values
func (g *Gif2WebP) Reset() *Gif2WebP {
	g.quality = -1
	return g
}

func (g *Gif2WebP) setInput() error {
	if g.input != nil {
		g.Arg("--").Arg("-")
		g.StdIn(g.input)
	} else if g.inputGif != nil {
		r, err := createReaderFromGif(g.inputGif)

		if err != nil {
			return err
		}

		g.Arg("--").Arg("-")
		g.StdIn(r)
	} else if g.inputFile != "" {
		g.Arg(g.inputFile)
	} else {
		return errors.New("Undefined input")
	}

	return nil
}

func (g *Gif2WebP) getOutput() (string, error) {
	if g.output != nil {
		return "-", nil
	} else if g.outputFile != "" {
		return g.outputFile, nil
	} else {
		return "", errors.New("Undefined output")
	}
}
