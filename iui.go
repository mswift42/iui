package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/mswift42/iui/cl"
	"github.com/mswift42/iui/ui"
	"io/ioutil"
	"os"
	"text/template"
)

func loadFile(fp string) ([]byte, error) {
	file, err := os.Open(fp)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(file)
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
	var td ui.ThemeFile
	bytes, _ := loadFile("white-sand.xml")
	xml.Unmarshal(bytes, &td)

	tm, err := ui.NewThemeMap(&td)
	if err != nil {
		panic(err)
	}
	var jt ui.JsonThemeFile
	bytes, err = loadFile("ThemeColors.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(bytes, &jt)
	if err != nil {
		panic(err)
	}
	fmt.Println(jt)
	fmt.Println(tm)
	ntfj, err := ui.NewThemeMapFromJson(&jt)
	if err != nil {
		panic(err)
	}
	fmt.Println(ntfj)
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
