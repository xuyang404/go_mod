package main

import (
  "cpustatus/sys"
  "github.com/leaanthony/mewn"
  "github.com/wailsapp/wails"
)

func basic() string {
  return "Hello World!"
}

func main() {

  js := mewn.String("./frontend/dist/app.js")
  css := mewn.String("./frontend/dist/app.css")

  stats := &sys.Stats{}
  app := wails.CreateApp(&wails.AppConfig{
    Width:  512,
    Height: 512,
    Title:  "cpu status",
    JS:     js,
    CSS:    css,
    Colour: "#131313",
  })
  app.Bind(stats)
  app.Run()
}
