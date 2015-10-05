package overlay

import (
	"bytes"
	"image"
	"image/draw"
	"image/png"
	"math"
	"strings"

	"github.com/alienchow/CAGHack/dto"
	"github.com/alienchow/CAGHack/forms"

	"github.com/golang/freetype"
)

const (
	fontDPI       = 72
	fontSize      = 24
	fontSize2     = 18
	genderBoxSize = 14
)

func newForm() *image.RGBA {
	bounds := form.Bounds()
	dest := image.NewRGBA(bounds)
	draw.Draw(dest, bounds, form, bounds.Min, draw.Src)
	return dest
}

func Process(data *dto.EmbarkationCardRequest) *bytes.Buffer {
	output := newForm()
	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(f)
	c.SetFontSize(fontSize)
	c.SetClip(output.Bounds())
	c.SetSrc(image.Black)
	c.SetDst(output)

	typeString(data.Fullname, forms.Config.FullName, c)
	typeString(data.PassportNumber, forms.Config.PassportNumber, c)
	typeString(data.FlightCode, forms.Config.FlightNumber, c)
	typeString(data.CountryOfBirth, forms.Config.CountryOfBirth, c)
	typeString(strings.Replace(data.DateOfBirth, "-", "", -1), forms.Config.DateOfBirth, c)
	typeString(data.Nationality, forms.Config.Nationality, c)
	typeString(data.EmbarkationLocation, forms.Config.LastCity, c)
	typeString(data.MalaysianIC, forms.Config.MalaysianIC, c)
	typeString(data.MalaysianIC, forms.Config.MalaysianIC2, c)

	c.SetFontSize(fontSize2)
	typeString2(data.Fullname, forms.Config.FullName2, c)
	typeString2(data.Nationality, forms.Config.Nationality2, c)

	var ptX, ptY int
	switch data.Gender {
	case "M":
		ptX = forms.Config.Gender[0][0]
		ptY = forms.Config.Gender[0][1]
	default:
		ptX = forms.Config.Gender[1][0]
		ptY = forms.Config.Gender[1][1]
	}
	bounds := image.Rectangle{
		Min: image.Point{ptX, ptY},
		Max: image.Point{ptX + genderBoxSize, ptY + genderBoxSize},
	}
	draw.Draw(output, bounds, image.Black, bounds.Min, draw.Src)

	outWriter := bytes.NewBuffer([]byte{})
	_ = png.Encode(outWriter, output)
	return outWriter
}

func typeString(field string, positions [][]int, c *freetype.Context) {
	stringSlice := strings.Split(field, "")
	length := int(math.Min(float64(len(stringSlice)), float64(len(positions))))
	for i := 0; i < length; i++ {
		pt := freetype.Pt(positions[i][0]+5+alphabetOffSet(stringSlice[i]), positions[i][1]+26)
		_, _ = c.DrawString(stringSlice[i], pt)
	}
}

func alphabetOffSet(alphabet string) int {
	if len(alphabet) != 1 {
		return 0
	}
	switch alphabet {
	case "I":
		return 5
	case "M", "W":
		return -2
	case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
		return 2
	}
	return 0
}

func typeString2(field string, position []int, c *freetype.Context) {
	pt := freetype.Pt(position[0], position[1])
	_, _ = c.DrawString(field, pt)
}
