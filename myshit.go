 package main

 import (
   "log"
   "github.com/gotk3/gotk3/gtk"
   "fmt"
   "os"
   "encoding/json"
   "math/rand"
   "sort"
   "time"
 )

type word struct{
  Imi string
  Yomi string
  Kanji string
}
type words struct{
  Fuck []word
}


func main(){
  kanjj := japanese("n4.json")
  //fmt.Println(kanjj[1].Imi)
  sort.Slice(kanjj, func(i, j int)bool{
    rand.Seed(time.Now().UnixNano())
    ra := rand.Intn(100)
    if ra <= 49{
      return true
    } else {
      return false
    }
  })


  gtk.Init(nil)
  b, err := gtk.BuilderNew()
  if err != nil {
    log.Fatal("Ошибка", err)
  }
  err = b.AddFromFile("myshit.glade")
  if err != nil {
    log.Fatal("Ошибка", err)
  }

  obj, err := b.GetObject("window_main")
  if err != nil {
    log.Fatal("Ошибка", err)
  }

  but, _ := b.GetObject("button1")
  entryBut := but.(*gtk.Button)

  lab1, _ := b.GetObject("label1")
  label1 := lab1.(*gtk.Label)

  lab2, _ := b.GetObject("label2")
  label2 := lab2.(*gtk.Label)

  var schet int = 0

  label1.SetText(kanjj[0].Imi)
  label2.SetText(kanjj[0].Yomi)

  entryBut.Connect("clicked", func(){
    schet++
    if (schet >= len(kanjj)){
      var itog string
      for _, elem := range kanjj {
        itog += elem.Kanji
        fmt.Println(itog)
      }
      label1.SetText(itog)
      label2.SetText("You was good\nAnd did good")
    } else {
      label1.SetText(kanjj[schet].Imi)
      label2.SetText(kanjj[schet].Yomi)
    }
  })

  win := obj.(*gtk.Window)
  win.Connect("destroy", func(){
    gtk.MainQuit()
  })
  win.ShowAll()
  gtk.Main()
}

func japanese(path string)(elements []word){
  file, _ := os.Open(path)
  defer file.Close()

  decoder := json.NewDecoder(file)
  conf := words{}
  err := decoder.Decode(&conf)
  if err != nil {
    fmt.Println(err)
  }
  return conf.Fuck
}
