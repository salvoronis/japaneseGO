package main

import (
  "log"
  "github.com/gotk3/gotk3/gtk"
  "fmt"
  "os"
  "encoding/json"
)

type word struct{
  Mean string `json:"meaning"`
  Read string `json:"reading"`
  Kanji string `json:"kanji"`
}

var (
  kanjiArr []word
  count int = 0
  labelMean *gtk.Label
  labelRead *gtk.Label
)


func main(){
  gtk.Init(nil)
  b, err := gtk.BuilderNew()
  if err != nil {
    log.Fatal("Ошибка", err)
  }
  err = b.AddFromFile("/home/salvoroni/myshit/myshit.glade")
  if err != nil {
    log.Fatal("Ошибка", err)
  }

  b.ConnectSignals(signals)

  obj, err := b.GetObject("window_main")
  if err != nil {
    log.Fatal("Ошибка", err)
  }

  but, err := b.GetObject("button1")
  if err != nil {
    log.Fatal("Can not get button", err)
  }
  entryBut := but.(*gtk.Button)

  lab1, err := b.GetObject("label1")
  if err != nil {
    log.Fatal("Can not get label first", err)
  }
  labelMean = lab1.(*gtk.Label)

  lab2, err := b.GetObject("label2")
  if err != nil {
    log.Fatal("Can not get label second", err)
  }
  labelRead = lab2.(*gtk.Label)

  labelMean.SetText("Imi")
  labelRead.SetText("Yomi")

  entryBut.Connect("clicked", func(){
    if (count >= len(kanjiArr)){
      var itog string
      for index, elem := range kanjiArr {
        itog += elem.Kanji
        fmt.Print(elem.Kanji)
        if (index+1) % 30 == 0 {
          itog += "\n"
        }
      }
      labelMean.SetText(itog)
      labelRead.SetText("")
    } else {
      labelMean.SetText(kanjiArr[count].Mean)
      labelRead.SetText(kanjiArr[count].Read)
    }
    count++
  })

  win := obj.(*gtk.Window)
  win.Connect("destroy", func(){
    gtk.MainQuit()
  })
  win.ShowAll()
  gtk.Main()
}

//japanese function will read the "path" file
//Decode json and return array of words
func japanese(path string)(elements []word){
  file, err := os.Open(path)
  if err != nil {
    log.Printf("Can not open %s\n",path)
  }
  defer file.Close()

  decoder := json.NewDecoder(file)
  conf := []word{}
  err = decoder.Decode(&conf)
  if err != nil {
    fmt.Println(err)
  }
  return conf
}
