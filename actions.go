package main

import (
  "github.com/gotk3/gotk3/gtk"
  "math/rand"
  "sort"
  "time"
)

var signals = map[string]interface{}{
  "govno": standart5,
	"n4standart": standart4,
  "n3standart": standart3,
  "n2standart": standart2,
  "damnit": exit,
}


func exit()  {
  gtk.MainQuit()
}

func standart4(){
  count = 0
  conv1.SetText("start level 4")
  conv2.SetText("")
  kotoba = japanese("/home/salvoroni/myshit/n4.json")
  sort.Slice(kotoba, func(i, j int)bool{
    rand.Seed(time.Now().UnixNano())
    ra := rand.Intn(100)
    if ra <= 49{
      return true
    } else {
      return false
    }
  })
}
func standart3(){
  count = 0
  conv1.SetText("start level 3")
  conv2.SetText("")
  kotoba = japanese("/home/salvoroni/myshit/n3.json")
  sort.Slice(kotoba, func(i, j int)bool{
    rand.Seed(time.Now().UnixNano())
    ra := rand.Intn(100)
    if ra <= 49{
      return true
    } else {
      return false
    }
  })
}
func standart2(){
  count = 0
  conv1.SetText("start level 2")
  conv2.SetText("")
  kotoba = japanese("/home/salvoroni/myshit/n2.json")
  sort.Slice(kotoba, func(i, j int)bool{
    rand.Seed(time.Now().UnixNano())
    ra := rand.Intn(100)
    if ra <= 49{
      return true
    } else {
      return false
    }
  })
}
func standart5(){
  count = 0
  conv1.SetText("start level 5")
  conv2.SetText("")
  kotoba = japanese("/home/salvoroni/myshit/n5.json")
  sort.Slice(kotoba, func(i, j int)bool{
    rand.Seed(time.Now().UnixNano())
    ra := rand.Intn(100)
    if ra <= 49{
      return true
    } else {
      return false
    }
  })
}
