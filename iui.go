package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/lucasb-eyer/go-colorful"
	"io/ioutil"
	"os"
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
	Colors []ColorOptions `xml:"colors>option"`
}

type ColorOptions struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
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

var reykjavik = map[string]string{
	"author":    "Martin Haesler",
	"themename": "reykjavik",
	"fg1":       "#b1b1b1",
	"fg2":       "#a3a3a3",
	"bg1":       "#112328",
	"bg2":       "#243549",
	"bg3":       "#37464a",
	"bg4":       "#4a585c",
	"builtin":   "#c4cbee",
	"keyword":   "#a3d4e8",
	"const":     "#a3d6cc",
	"comment":   "#5d5d5d",
	"func":      "#f1c1bd",
	"string":    "#e6c2db",
	"warning":   "#e81050",
	"warning2":  "#e86310",
}

var soft_charcoal = map[string]string{
	"author":    "Martin Haesler",
	"themename": "soft-charcoal",
	"fg1":       "#c2c2c2",
	"fg2":       "#b2b2b2",
	"bg1":       "#191919",
	"bg2":       "#2b2b2b",
	"bg3":       "#3e3e3e",
	"bg4":       "#505050",
	"builtin":   "#54686d",
	"keyword":   "#8aa234",
	"const":     "8aa6c1",
	"comment":   "#808080",
	"func":      "#7a8bbd",
	"string":    "#5d90cd",
	"warning":   "#ff6523",
	"warning2":  "#ff2370",
}

var madrid = map[string]string{
	"author":    "Martin Haesler",
	"themename": "madrid",
	"fg1":       "#b1b1b1",
	"fg2":       "#a3a3a3",
	"bg1":       "#1b1b1b",
	"bg2":       "#2d2d2d",
	"bg3":       "#3f3f3f",
	"bg4":       "#525252",
	"builtin":   "#a78360",
	"keyword":   "#b7797d",
	"const":     "#a27ea4",
	"comment":   "#5d5d5d",
	"func":      "#96546c",
	"string":    "#818f62",
	"warning":   "#e81050",
	"warning2":  "#e86310",
}

var silkworm = map[string]string{
	"author":    "Martin Haesler",
	"themename": "silkworm",
	"fg1":       "#585858",
	"fg2":       "#656565",
	"bg1":       "#ece3db",
	"bg2":       "#d9d1c9",
	"bg3":       "#c6bfb8",
	"bg4":       "#b3ada6",
	"builtin":   "#0073b5",
	"keyword":   "#367a7f",
	"const":     "#a27ea4",
	"comment":   "#a9a9a9",
	"func":      "#ad4271",
	"string":    "#3b4bab",
	"warning":   "#ff1276",
	"warning2":  "#ff4d12",
}

var metalheart = map[string]string{
	"author":    "Martin Haesler",
	"themename": "metalheart",
	"fg1":       "#8693ae",
	"fg2":       "#7b87a0",
	"bg1":       "#1d272a",
	"bg2":       "#2f383b",
	"bg3":       "#414a4c",
	"bg4":       "#535b5d",
	"builtin":   "#9a91b8",
	"keyword":   "#6e94ad",
	"const":     "5980a3",
	"comment":   "#727272",
	"func":      "#937aba",
	"string":    "#ba899c",
	"warning":   "#ff0000",
	"warning2":  "#ff8800",
}

var foggyNight = map[string]string{
	"author":    "Martin Haesler",
	"themename": "foggy-night",
	"fg1":       "#8f8f8f",
	"fg2":       "#848484",
	"bg1":       "#292929",
	"bg2":       "#3a3a3a",
	"bg3":       "#4b4b4b",
	"bg4":       "#5c5c5c",
	"builtin":   "#997599",
	"keyword":   "#6b83ac",
	"const":     "#3e8c9d",
	"comment":   "#626262",
	"func":      "#9e7a5a",
	"string":    "#ad7a76",
	"warning":   "#e81050",
	"warning2":  "#e86310",
}

var whiteSand = map[string]string{
	"author":    "Martin Haesler",
	"themename": "white-sand",
	"fg1":       "#585858",
	"fg2":       "#656565",
	"bg1":       "#f5ebe1",
	"bg2":       "#e1d8cf",
	"bg3":       "#cec5bd",
	"bg4":       "#bab3ab",
	"builtin":   "#1a8591",
	"keyword":   "#4a858c",
	"const":     "#697024",
	"comment":   "#a9a9a9",
	"func":      "#bd745e",
	"string":    "#b3534b",
	"type":      "#8c4a79",
	"warning":   "#ff1276",
	"warning2":  "#ff4d12",
}

var warmNight = map[string]string{
	"author":    "Martin Haesler",
	"themename": "warm-night",
	"fg1":       "#b1b1b1",
	"fg2":       "#a3a3a3",
	"bg1":       "#292424",
	"bg2":       "#3a3636",
	"bg3":       "#4b4747",
	"bg4":       "#5c5959",
	"builtin":   "#71a46c",
	"keyword":   "#96905f",
	"const":     "#bd845f",
	"comment":   "#5d5a58",
	"func":      "#b680b1",
	"string":    "#71a19f",
	"type":      "#8b8fc6",
	"warning":   "#e81050",
	"warning2":  "#e86310",
}

var thursday = map[string]string{
	"author":    "Martin Haesler",
	"themename": "thursday",
	"fg1":       "#2f2f2f",
	"fg2":       "#404040",
	"bg1":       "#f9fbfd",
	"bg2":       "#e5e7e9",
	"bg3":       "#d1d3d5",
	"bg4":       "#bdbfc0",
	"builtin":   "#636792",
	"keyword":   "#28728f",
	"const":     "#28766e",
	"comment":   "#949494",
	"func":      "#935c54",
	"string":    "#8c5c79",
	"type":      "#56724b",
	"warning":   "#fa0c0c",
	"warning2":  "#fa7b0c",
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
	var res bytes.Buffer
	tmpl, err := template.ParseFiles("templ.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(err)
	fmt.Println(res.String())
	foggycols := addColors(foggyNight)
	fmt.Println("foggyNight cols: ", foggycols)
	reykjavikcols := addColors(reykjavik)
	fmt.Println("Reykjavik colours: ", reykjavikcols)
	reykbg, _ := colorful.Hex("#112328")
	reykfg2, _ := colorful.Hex("#a3a3a3")
	reykfg1, _ := colorful.Hex("#b1b1b1")
	softcharcoalcols := addColors(soft_charcoal)
	fmt.Println(softcharcoalcols)
	thursdaycols := addColors(thursday)
	fmt.Println("inverted reykjavik fg2: ", invertColor(&reykbg, &reykfg2))
	fmt.Println("inverted reykjavik fg1: ", invertColor(&reykbg, &reykfg1))
	err = tmpl.Execute(&res, thursdaycols)
	if err := ioutil.WriteFile(thursdaycols["themename"]+".theme.json", res.Bytes(), 0644); err != nil {
		panic(err)
	}
	file, err := os.Open("white-sand.xml")
	if err != nil {
		panic(err)
	}
	var td ThemeFile
	bytes, _ := ioutil.ReadAll(file)
	xml.Unmarshal(bytes, &td)
	fmt.Println(td)
}
