package main

import (
	"encoding/xml"
	"fmt"
	"github.com/lucasb-eyer/go-colorful"
	"io/ioutil"
	"os"
)

func lighten(col *colorful.Color, factor float64) string {
	white, _ := colorful.Hex("#ffffff")
	return col.BlendLab(white, factor).Hex()
}

func darken(col *colorful.Color, factor float64) string {
	black, _ := colorful.Hex("#000000")
	return col.BlendLab(black, factor).Hex()
}

func hasDarkBG(col *colorful.Color) bool {
	l, _, _ := col.Lab()
	return l < 0.5
}

func invertColor(bgcol, col *colorful.Color) string {
	_, _, l1 := bgcol.Hsl()
	h2, s2, _ := col.Hsl()
	return colorful.Hsl(h2, s2, l1).Hex()
}

func luminance(col *colorful.Color) float64 {
	_, _, l := col.Hsl()
	return l
}

type ThemeData struct {
	MainFG  string `xml:"scheme""`
	Attribs ThemeAttributes
}

type ThemeAttributes struct {
	Attributes []ThemeAttributeOption `xml:"attributes"`
}

type ThemeAttributeOption struct {
	option string    `xml:"option,attr"`
	AV     AttrValue `xml:"value"`
}

type AttrValue struct {
	Name  string `xml:"option>name,attr"`
	Value string `xml:"option>value,attr"`
}

type TD struct {
	Attributes []struct {
		Optname    string `xml:"name,attr"`
		FaceOption struct {
			FaceVal  string `xml:"value,attr"`
			FaceName string `xml:"name,attr"`
		} `xml:"attributes>option>value"`
	} `xml:"attributes>option"`
}

func main() {
	c1, _ := colorful.Hex("#f5ebe1")
	fmt.Println(lighten(&c1, 0.16))
	keywordcol, _ := colorful.Hex("#4a858c")
	fmt.Println("keyword face: ", invertColor(&c1, &keywordcol))
	builtincol, _ := colorful.Hex("#1a8591")
	funcname, _ := colorful.Hex("#bd745e")
	typeface, _ := colorful.Hex("#8c4a79")
	variableface, _ := colorful.Hex("#476238")
	stringface, _ := colorful.Hex("#b3534b")
	warningface, _ := colorful.Hex("#ff1276")
	warning2face, _ := colorful.Hex("#ff4d12")
	fmt.Println("builtin face: ", invertColor(&c1, &builtincol))
	fmt.Println("function name: ", invertColor(&c1, &funcname))
	fmt.Println("type face: ", invertColor(&c1, &typeface))
	fmt.Println("variable face: ", invertColor(&c1, &variableface))
	fmt.Println("string face: ", invertColor(&c1, &stringface))
	fmt.Println("warning face: ", invertColor(&c1, &warningface))
	fmt.Println("warning2 face: ", invertColor(&c1, &warning2face))
	fmt.Println(luminance(&c1))
	fmt.Println("darken: ", darken(&c1, 0.1))
	fmt.Println(hasDarkBG(&c1))
	icls, err := os.Open("white-sand.icls")
	if err != nil {
		panic(err)
	}
	var theme ThemeAttributes
	var td TD
	bytes, err := ioutil.ReadAll(icls)
	if err != nil {
		panic(err)
	}
	xml.Unmarshal(bytes, &theme)
	xml.Unmarshal(bytes, &td)
	fmt.Println(theme)
}
