package main

import (

  "github.com/neonify/lessgo/pkg/lessgo"
  "github.com/neonify/lessgo/pkg/fuzz"
  "github.com/neonify/lessgo/pkg/input"
  )

var file string = ""

func main(){
  
  Obj := input.Input()
  
  lessgo.Banner()
  Obj.Handle()
  
  fuzz.ConcManager(Obj)
  
}