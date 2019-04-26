package main

import (
	"encoding/xml"
	"fmt"
	"github.com/lucasb-eyer/go-colorful"
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

type ThemeData struct {
	MainFG string `xml:scheme`
}

type ThemeAttributes struct {
	Attributes []ThemeAttribute `xml:attributes`
}

type ThemeAttribute struct {
	XMLName xml.Name  `xml:option`
	AV      AttrValue `xml:value`
}

type AttrValue struct {
	Name  string `xml:name,attr`
	Value string `xml:value,attr`
}

func main() {
	c1, _ := colorful.Hex("#f5ebe1")
	fmt.Println(lighten(&c1, 0.16))
	fmt.Println("darken: ", darken(&c1, 0.1))
	fmt.Println(hasDarkBG(&c1))
}
