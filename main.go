package main

import (
  "log"
  "github.com/gotk3/gotk3/gtk"
  "fmt"
  "os"
  "encoding/json"
  "io/ioutil"
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
  logger *log.Logger
  currentDir string
)


func main(){
  logfile, _ := os.Create("./log.txt")
  logger = log.New(logfile, "japaneseGo ", log.LstdFlags|log.Lshortfile)
  currentDir, _ = os.Getwd()

  gtk.Init(nil)
  application, err := gtk.BuilderNew()
  if err != nil {
    logger.Fatal("Innitial gtk error", err)
  }
  err = application.AddFromFile("/home/salvoroni/myshit/myshit.glade")
  if err != nil {
    logger.Fatal("Glade file error", err)
  }

  addOwnKanji(application)

  application.ConnectSignals(signals)

  obj, err := application.GetObject("window_main")
  if err != nil {
    logger.Fatal("Can not get main window", err)
  }

  but, err := application.GetObject("button1")
  if err != nil {
    logger.Fatal("Can not get button", err)
  }
  entryBut := but.(*gtk.Button)

  lab1, err := application.GetObject("label1")
  if err != nil {
    logger.Fatal("Can not get label first", err)
  }
  labelMean = lab1.(*gtk.Label)

  lab2, err := application.GetObject("label2")
  if err != nil {
    logger.Fatal("Can not get label second", err)
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
    logger.Printf("Can not open %s\n",path)
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

func addOwnKanji(application *gtk.Builder){
  ownKanjiBar, _ := application.GetObject("subownlib")
  ownKanji := ownKanjiBar.(*gtk.Menu)
  files, err := ioutil.ReadDir("./user")
  if err != nil {
    logger.Println("Can not get files from user data")
  }
  for _, file := range files {
    var filename string = file.Name()
    item, err := gtk.MenuItemNewWithLabel(file.Name())
    if err != nil {
      logger.Println("Can not get file")
    }
    item.Connect("activate", func(){
      standart(currentDir+"/user/"+filename, filename[:len(filename)-5])
    })
    ownKanji.Append(item)
  }
}
