package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/mswift42/iui/cl"
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

type ThemeFile struct {
	Colors     []ColorOptions `xml:"colors>option"`
	ThemeAttrs []AttrOption   `xml:"attributes>option"`
}

type JsonThemeFile struct {
	Foreground   string
	Background   string
	Keyword      string
	Functionname string
	Comment      string
	Constant     string
	String       string
	Type         string
	Builtin      string
	Warning      string
	Warning2     string
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

type themeAttributes struct {
	fg string
	bg string
}

func attrMap(attros []AttrOption) map[string]themeAttributes {
	tamap := make(map[string]themeAttributes)
	for _, i := range attros {
		var ta themeAttributes
		for _, j := range i.Values {
			lower := strings.ToLower(j.Name)
			if lower == "foreground" {
				ta.fg = j.Value
			}
			if lower == "background" {
				ta.bg = j.Value
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
	Type        string
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

const HexHash = "#"

func (tm *ThemeMap) addColors() error {
	bg, err := colorful.Hex(tm.Bg1)
	if err != nil {
		return err
	}
	fg, err := colorful.Hex(tm.Fg1)
	if err != nil {
		return err
	}
	var bg01, bg2, bg3, bg4, fg2 string

	if hasDarkBG(&bg) {
		bg01 = darken(&bg, 0.1)
		bg2 = lighten(&bg, 0.08)
		bg3 = lighten(&bg, 0.16)
		bg4 = lighten(&bg, 0.24)
		fg2 = darken(&fg, 0.08)
	} else {
		bg01 = lighten(&bg, 0.1)
		bg2 = darken(&bg, 0.08)
		bg3 = darken(&bg, 0.16)
		bg4 = darken(&bg, 0.24)
		fg2 = lighten(&fg, 0.08)
	}
	tm.Bg01 = bg01
	tm.DarkBG = hasDarkBG(&bg)
	tm.Bg2 = bg2
	tm.Bg3 = bg3
	tm.Bg4 = bg4
	tm.Fg2 = fg2
	builtin, _ := colorful.Hex(tm.Builtin)
	keyw, _ := colorful.Hex(tm.Keyword)
	typ, _ := colorful.Hex(tm.Type)
	fnc, _ := colorful.Hex(tm.Func)
	warn1, _ := colorful.Hex(tm.Warning)
	warn2, _ := colorful.Hex(tm.Warning2)
	str, _ := colorful.Hex(tm.String)
	tm.InvBuiltin = invertColor(&bg, &builtin)
	tm.InvKeyword = invertColor(&bg, &keyw)
	tm.InvType = invertColor(&bg, &typ)
	tm.InvFunc = invertColor(&bg, &fnc)
	tm.InvString = invertColor(&bg, &str)
	tm.InvWarning = invertColor(&bg, &warn1)
	tm.InvWarning2 = invertColor(&bg, &warn2)
	return nil
}

func newThemeMap(td *ThemeFile) (ThemeMap, error) {
	var tm ThemeMap
	am := attrMap(td.ThemeAttrs)
	tm.Bg1 = HexHash + am["TEXT"].bg
	tm.Fg1 = HexHash + am["TEXT"].fg
	tm.Func = HexHash + am["DEFAULT_FUNCTION_DECLARATION"].fg
	tm.Comment = HexHash + am["DEFAULT_BLOCK_COMMENT"].fg
	tm.Constant = HexHash + am["DEFAULT_CONSTANT"].fg
	tm.Keyword = HexHash + am["DEFAULT_KEYWORD"].fg
	tm.String = HexHash + am["DEFAULT_STRING"].fg
	tm.Type = HexHash + am["DEFAULT_CLASS_NAME"].fg
	tm.Builtin = HexHash + am["DEFAULT_INSTANCE_FIELD"].fg
	tm.Warning = HexHash + am["LOG_ERROR_OUTPUT"].fg
	tm.Warning2 = HexHash + am["LOG_WARNING_OUTPUT"].fg
	err := tm.addColors()
	return tm, err
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
	tm, err := newThemeMap(&td)
	if err != nil {
		panic(err)
	}
	fmt.Println(tm)
	if err := tmpl.Execute(&res, tm); err != nil {
		panic(err)
	}
	if err := ioutil.WriteFile("white_sand.theme.json", res.Bytes(), 0644); err != nil {
		fmt.Println(err)
	}
	app := cl.InitCli()
	if err := app.Run(os.Args); err != nil {
		panic(err)
	}

}
