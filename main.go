//This is the package comment
package main

import (
  "log"
  "github.com/gotk3/gotk3/gtk"
  "fmt"
  "os"
  "encoding/json"
)

//The word object steuct
type word struct{
  Imi string
  Yomi string
  Kanji string
}
//This struct contains all words from chosen level
type words struct{
  List []word
}

var kotoba []word

var count int = 0

var conv1 *gtk.Label = nil
var conv2 *gtk.Label = nil

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

  but, _ := b.GetObject("button1")
  entryBut := but.(*gtk.Button)

  lab1, _ := b.GetObject("label1")
  label1 := lab1.(*gtk.Label)
  conv1 = label1

  lab2, _ := b.GetObject("label2")
  label2 := lab2.(*gtk.Label)
  conv2 = label2


  label1.SetText("Imi")
  label2.SetText("Yomi")

  entryBut.Connect("clicked", func(){
    if (count >= len(kotoba)){
      var itog string
      for index, elem := range kotoba {
        itog += elem.Kanji
        fmt.Print(elem.Kanji)
        if (index+1) % 30 == 0 {
          itog += "\n"
        }
      }
      label1.SetText(itog)
      label2.SetText("")
    } else {
      label1.SetText(kotoba[count].Imi)
      label2.SetText(kotoba[count].Yomi)
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
  conf := words{}
  err = decoder.Decode(&conf)
  if err != nil {
    fmt.Println(err)
  }
  return conf.List
}
