package main

import (
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

func invertColor(bgcol, col *colorful.Color) string {
	_, _, l1 := bgcol.Hsl()
	h2, s2, _ := col.Hsl()
	return colorful.Hsl(h2, s2, l1).Hex()
}

func luminance(col *colorful.Color) float64 {
	_, _, l := col.Hsl()
	return l
}

func addColors(colors map[string]string) map[string]string {
	bg, _ := colorful.Hex(colors["bg"])
	var bg01 string
	if hasDarkBG(&bg) {
		bg01 = darken(&bg, 0.12)
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
}
