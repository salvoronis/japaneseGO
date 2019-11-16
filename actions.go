package main

import (
  "github.com/gotk3/gotk3/gtk"
  "math/rand"
  "sort"
  "time"
)

var signals = map[string]interface{}{
  "govno": govno5,
	"n4standart": govno4,
  "n3standart": govno3,
  "n2standart": govno2,
  "damnit": damn,
}


func damn()  {
  gtk.MainQuit()
}

func govno4(){
  schet = 0
  conv1.SetText("start level 4")
  conv2.SetText("im tired")
  kotoba = japanese("n4.json")
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
func govno3(){
  schet = 0
  conv1.SetText("start level 3")
  conv2.SetText("im tired")
  kotoba = japanese("n3.json")
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
func govno2(){
  schet = 0
  conv1.SetText("start level 2")
  conv2.SetText("im tired")
  kotoba = japanese("n2.json")
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
func govno5(){
  schet = 0
  conv1.SetText("start level 5")
  conv2.SetText("im tired")
  kotoba = japanese("n5.json")
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
