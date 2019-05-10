package main

import (
	"encoding/json"
	"testing"
)

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

var whiteSand = ThemeMap{
	DarkBG:   false,
	Fg1:      "#585858",
	Fg2:      "#656565",
	Bg1:      "#f5ebe1",
	Bg01:     "#f6ede4",
	Bg2:      "#e1d8cf",
	Bg3:      "#cec5bd",
	Bg4:      "#bab3ab",
	Builtin:  "#1a8591",
	Keyword:  "#4a858c",
	Constant: "#697024",
	Comment:  "#a9a9a9",
	Func:     "#bd745e",
	String:   "#b3534b",
	Type:     "#8c4a79",
	Warning:  "#ff1276",
	Warning2: "#ff4d12",
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

func TestNewThemeMapFromJson(t *testing.T) {
	bytes, err := loadFile("white-sand.json")
	if err != nil {
		panic(err)
	}
	var jt JsonThemeFile
	if err := json.Unmarshal(bytes, &jt); err != nil {
		panic(err)
	}
	theme, err := newThemeMapFromJson(&jt)
	if err != nil {
		panic(err)
	}
	if theme.Bg1 != whiteSand.Bg1 {
		t.Errorf("Expected bg1 to be %s, got %s",
			whiteSand.Bg1, theme.Bg1)
	}
	if theme.Fg1 != whiteSand.Fg1 {
		t.Errorf("Expected fg1 to be %s, got %s",
			whiteSand.Fg1, theme.Fg1)
	}
	if theme.Bg01 != whiteSand.Bg01 {
		t.Errorf("expected bg01 to be %s, got %s",
			whiteSand.Bg01, theme.Bg01)
	}

}
