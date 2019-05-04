package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/lucasb-eyer/go-colorful"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
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

type ThemeFile struct {
	Colors     []ColorOptions `xml:"colors>option"`
	ThemeAttrs []AttrOption   `xml:"attributes>option"`
}

type AttrOption struct {
	Option string      `xml:"name,attr"`
	Values []AttrValue `xml:"value>option"`
}
type AttrValue struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

type ColorOptions struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

func attrMap(attros []AttrOption) map[string]string {
	am := make(map[string]string)
	for _, i := range attros {
		for _, j := range i.Values {
			lower := strings.ToLower(j.Name)
			if lower == "foreground" {
				am[i.Option] = j.Value
			}
		}
	}
	return am
}

func colMap(cm []ColorOptions) map[string]string {
	cmap := make(map[string]string)
	for _, i := range cm {
		cmap[i.Name] = i.Value
	}
	return cmap
}

func addColors(colors map[string]string) map[string]string {
	bg, _ := colorful.Hex(colors["bg1"])
	var bg01 string
	if hasDarkBG(&bg) {
		bg01 = darken(&bg, 0.1)
	} else {
		bg01 = lighten(&bg, 0.1)
	}
	colors["bg01"] = bg01
	builtin, _ := colorful.Hex(colors["builtin"])
	keyw, _ := colorful.Hex(colors["keyword"])
	typ, _ := colorful.Hex(colors["type"])
	fnc, _ := colorful.Hex(colors["func"])
	warn1, _ := colorful.Hex(colors["warning"])
	warn2, _ := colorful.Hex(colors["warning2"])
	str, _ := colorful.Hex(colors["string"])

	colors["invbuiltin"] = invertColor(&bg, &builtin)
	colors["invkeyword"] = invertColor(&bg, &keyw)
	colors["invtype"] = invertColor(&bg, &typ)
	colors["invfunc"] = invertColor(&bg, &fnc)
	colors["invwarning"] = invertColor(&bg, &warn1)
	colors["invwarning2"] = invertColor(&bg, &warn2)
	colors["invstring"] = invertColor(&bg, &str)

	return colors
}

func main() {
	var res bytes.Buffer
	tmpl, err := template.ParseFiles("templ.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(tmpl)
	fmt.Println(err)
	fmt.Println(res.String())
	file, err := os.Open("white-sand.xml")
	if err != nil {
		panic(err)
	}
	var td ThemeFile
	bytes, _ := ioutil.ReadAll(file)
	xml.Unmarshal(bytes, &td)
	fmt.Println(td.ThemeAttrs)
	fmt.Println(attrMap(td.ThemeAttrs))
	fmt.Println(colMap(td.Colors))
}
