package main

import (
  "github.com/gotk3/gotk3/gtk"
  "math/rand"
  "sort"
  "time"
)

var signals = map[string]interface{}{
  "n5standart": standart5,
	"n4standart": standart4,
  "n3standart": standart3,
  "n2standart": standart2,
  "exit": exit,
  "edstandart": editstandartmenu,
}

var signalsedit = map[string]interface{}{
  "lvlchanger": setlevel,
}


func exit()  {
  gtk.MainQuit()
}

func standart4(){
  standart(currentDir+"/standart/n4.json", "n4")
}
func standart3(){
  standart(currentDir+"/standart/n3.json", "n3")
}
func standart2(){
  standart(currentDir+"/standart/n2.json", "n2")
}
func standart5(){
  standart(currentDir+"/standart/n5.json", "n5")
}

func standart(filepath, label string){
  count = 0
  labelMean.SetText("start "+label)
  labelRead.SetText("")
  kanjiArr = japanese(filepath)
  sort.Slice(kanjiArr, func(i, j int)bool{
    rand.Seed(time.Now().UnixNano())
    ra := rand.Intn(100)
    if ra <= 49{
      return true
    } else {
      return false
    }
  })
}

var editwindow *gtk.Builder
var target *gtk.ComboBoxText

func editstandartmenu() {
  editwindow, err := gtk.BuilderNew()
  if err != nil {
    logger.Fatal("Innitial gtk error ", err)
  }
  err = editwindow.AddFromFile(currentDir+"/glade/lvleditor.glade")
  if err != nil {
    logger.Fatal("Glade file error ", err)
  }
  box, err := editwindow.GetObject("combo_box")
  if err != nil {
    logger.Fatal("Can not get lvl select box ", err)
  }
  target = box.(*gtk.ComboBoxText)
  editwindow.ConnectSignals(signalsedit)
  obj, err := editwindow.GetObject("lvleditor")
  if err != nil {
    logger.Fatal("Can not get main window ", err)
  }
  win := obj.(*gtk.Window)
  win.ShowAll()
}

func setlevel() {
  logger.Println(target.GetActiveText())
}