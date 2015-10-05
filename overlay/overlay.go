/*
Package overlay is in charge of overlaying the text onto the give image profile
*/
package overlay

import (
	"bytes"
	"image"
	"image/png"
	"io/ioutil"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

const (
	fontFile = `fonts/HelveticaNeue-Bold.ttf`
	formFile = `forms/form.png`
)

var (
	f    *truetype.Font
	form *image.RGBA
)

func init() {
	fontBytes, err := ioutil.ReadFile(fontFile)
	if err != nil {
		panic("Font file read error: " + err.Error())
		return
	}
	f, err = freetype.ParseFont(fontBytes)
	if err != nil {
		panic("Error parsing freetype font: " + err.Error())
		return
	}
	formBytes, err := ioutil.ReadFile(formFile)
	if err != nil {
		panic("Form file read error: " + err.Error())
		return
	}
	reader := bytes.NewBuffer(formBytes)
	formImage, err := png.Decode(reader)
	if err != nil {
		panic("Failed to decode PNG file: " + err.Error())
		return
	}
	form, _ = formImage.(*image.RGBA)
}
