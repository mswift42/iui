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

type ThemeAttributes struct {
	FG string
	BG string
}

func attrMap(attros []AttrOption) map[string]ThemeAttributes {
	tamap := make(map[string]ThemeAttributes)
	for _, i := range attros {
		var ta ThemeAttributes
		for _, j := range i.Values {
			lower := strings.ToLower(j.Name)
			if lower == "foreground" {
				ta.FG = j.Value
			}
			if lower == "background" {
				ta.BG = j.Value
			}
		}
		tamap[i.Option] = ta
	}
	return tamap
}

func colMap(cm []ColorOptions) map[string]string {
	cmap := make(map[string]string)
	for _, i := range cm {
		cmap[i.Name] = i.Value
	}
	return cmap
}

type ThemeMap struct {
	ThemeName   string
	Author      string
	DarkBG      bool
	Fg1         string
	Fg2         string
	Bg1         string
	Bg01        string
	Bg2         string
	Bg3         string
	Bg4         string
	Builtin     string
	Keyword     string
	Constant    string
	Comment     string
	Func        string
	String      string
	Warning     string
	Warning2    string
	InvBuiltin  string
	InvKeyword  string
	InvType     string
	InvFunc     string
	InvString   string
	InvWarning  string
	InvWarning2 string
}

func addColors(colors map[string]string) ThemeMap {
	bg, _ := colorful.Hex(colors["bg1"])
	var bg01 string

	if hasDarkBG(&bg) {
		bg01 = darken(&bg, 0.1)
	} else {
		bg01 = lighten(&bg, 0.1)
	}
	var tm ThemeMap
	tm.Author = colors["author"]
	tm.ThemeName = colors["themename"]
	tm.Bg1 = colors["bg1"]
	tm.Bg01 = bg01
	tm.Fg1 = colors["fg1"]
	tm.Fg2 = colors["fg2"]
	tm.Bg2 = colors["bg2"]
	tm.Bg3 = colors["bg3"]
	tm.Bg4 = colors["bg4"]
	tm.DarkBG = hasDarkBG(&bg)
	builtin, _ := colorful.Hex(colors["builtin"])
	keyw, _ := colorful.Hex(colors["keyword"])
	typ, _ := colorful.Hex(colors["type"])
	fnc, _ := colorful.Hex(colors["func"])
	warn1, _ := colorful.Hex(colors["warning"])
	warn2, _ := colorful.Hex(colors["warning2"])
	str, _ := colorful.Hex(colors["string"])
	tm.Builtin = colors["builtin"]
	tm.Keyword = colors["keyword"]
	tm.Constant = colors["const"]
	tm.Comment = colors["comment"]
	tm.Func = colors["func"]
	tm.String = colors["string"]
	tm.Warning = colors["warning"]
	tm.Warning2 = colors["warning2"]
	tm.InvBuiltin = invertColor(&bg, &builtin)
	tm.InvKeyword = invertColor(&bg, &keyw)
	tm.InvType = invertColor(&bg, &typ)
	tm.InvFunc = invertColor(&bg, &fnc)
	tm.InvString = invertColor(&bg, &str)
	tm.InvWarning = invertColor(&bg, &warn1)
	tm.InvWarning2 = invertColor(&bg, &warn2)

	return tm
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
	am := attrMap(td.ThemeAttrs)
	fmt.Println(am)
	fmt.Println(am["TEXT"].BG)
	fmt.Println(am["TEXT"].FG)
}
