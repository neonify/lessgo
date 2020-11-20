package lessgo

import (
  "fmt"
)
  
  
func Banner(){
  
  ver := Version()
  
  fmt.Println(`

         _                     
        | |___ ___ ___ ___ ___ 
        | | -_|_ -|_ -| . | . |
        |_|___|___|___|_  |___|
                      |___|    
                       
  ____________________________________
  
            lessgo `+ver+`
  
           Author : Neonify 
      A Fast URL Fuzzer In Golang
  ____________________________________
  
`)
Cautions()
}

func Cautions(){
    
  Color("red","\a\n[!] Make sure you have good internet connection\a")
  Color("red","\a[!] Please do not use this tool for illegal purposes\a\n")
  
}