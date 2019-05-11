package main

import (
	"encoding/json"
	"testing"
)

var reykjavik = ThemeMap{
	DarkBG:   true,
	Fg1:      "#b1b1b1",
	Fg2:      "#a3a3a3",
	Bg1:      "#112328",
	Bg2:      "#243549",
	Bg3:      "#37464a",
	Bg4:      "#4a585c",
	Builtin:  "#c4cbee",
	Keyword:  "#a3d4e8",
	Constant: "#a3d6cc",
	Comment:  "#5d5d5d",
	Func:     "#f1c1bd",
	String:   "#e6c2db",
	Warning:  "#e81050",
	Warning2: "#e86310",
}

var whiteSand = ThemeMap{
	DarkBG:   false,
	Fg1:      "#585858",
	Fg2:      "#656565",
	Bg1:      "#f5ebe1",
	Bg01:     "#f6ede4",
	Bg2:      "#dfd6cd",
	Bg3:      "#c9c1b9",
	Bg4:      "#b4ada6",
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

var thursday = ThemeMap{
	Fg1:      "#2f2f2f",
	Fg2:      "#404040",
	Bg1:      "#f9fbfd",
	Bg2:      "#e5e7e9",
	Bg3:      "#d1d3d5",
	Bg4:      "#bdbfc0",
	Builtin:  "#636792",
	Keyword:  "#28728f",
	Constant: "#28766e",
	Comment:  "#949494",
	Func:     "#935c54",
	String:   "#8c5c79",
	Type:     "#56724b",
	Warning:  "#fa0c0c",
	Warning2: "#fa7b0c",
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
	if theme.Bg2 != whiteSand.Bg2 {
		t.Errorf("Expected bg2 to be %s, got %s",
			whiteSand.Bg2, theme.Bg2)
	}
	if theme.Bg3 != whiteSand.Bg3 {
		t.Errorf("expected bg3 to be %s, got %s",
			whiteSand.Bg3, theme.Bg3)
	}
	if theme.Bg4 != whiteSand.Bg4 {
		t.Errorf("expected Bg4 to be %s, got %s",
			whiteSand.Bg4, theme.Bg4)
	}
	if theme.Keyword != whiteSand.Keyword {
		t.Errorf("expected keyword to be %s, got %s",
			whiteSand.Keyword, theme.Keyword)
	}
	if theme.Builtin != whiteSand.Builtin {
		t.Errorf("Expected builtin to be %s, got %s",
			whiteSand.Builtin, theme.Builtin)
	}
	if theme.String != whiteSand.String {
		t.Errorf("Expected string to be %s, got %s",
			whiteSand.String, theme.String)
	}
	if theme.Type != whiteSand.Type {
		t.Errorf("Expected type to be %s, got: %s",
			whiteSand.Type, theme.Type)
	}
	if theme.Warning != whiteSand.Warning {
		t.Errorf("Expected warning to be %s, got: %s",
			whiteSand.Warning, theme.Warning)
	}
	if theme.Warning2 != whiteSand.Warning2 {
		t.Errorf("Expected warning2 to be %s, got: %s",
			whiteSand.Warning2, theme.Warning2)
	}

}
