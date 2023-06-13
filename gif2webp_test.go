package webpbin

import (
	"image/gif"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeGif(t *testing.T) {
	g := NewGif2WebP()
	f, err := os.Open("test2.gif")
	assert.Nil(t, err)
	gifImg, err := gif.DecodeAll(f)
	assert.Nil(t, err)
	g.InputGif(gifImg)
	g.Quality(100)
	g.OutputFile("test2.webp")
	err = g.Run()
	assert.Nil(t, err)
}

func TestVersionGif2WebP(t *testing.T) {
	g := NewGif2WebP()
	r, err := g.Version()
	assert.Nil(t, err)

	if _, ok := os.LookupEnv("DOCKER_ARM_TEST"); !ok {
		assert.Equal(t, "WebP Encoder version: 1.2.0WebP Mux version: 1.2.0", r)
	}
}
