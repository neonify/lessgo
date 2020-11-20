package fuzz

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "github.com/neonify/lessgo/pkg/lessgo"
  )

func Printer(clr lessgo.Hue,symb string,resp *http.Response,payload string){
  
  code := resp.Status
  loc := ""
  loc = resp.Header.Get("Location")
  
  
  body,_ := ioutil.ReadAll(resp.Body)
  defer resp.Body.Close()
  
  length := len(string(body))
  
  fmt.Printf(clr.Code)
  
  fmt.Printf("["+symb+"] ")
  fmt.Printf("%-032s",code)
  fmt.Printf("%7d",length)
  fmt.Printf("%-15s","  "+payload)
  fmt.Println(loc)
  fmt.Printf(lessgo.Reset.Code)
  
}


func Heading(){
  fmt.Printf("\n")
  fmt.Printf("%-030s","  Status")
  fmt.Printf("%13s","Length")
  fmt.Printf("%-10s","  Payload")
  fmt.Printf("%-10s","  Remarks")
  fmt.Printf("\n")
  
  fmt.Printf(lessgo.Gray.Code)
  fmt.Printf("%-035s","-----------")
  fmt.Printf("---------------------------------\n")
  fmt.Printf(lessgo.Reset.Code)
}

func ErrorPrint(err error,payload string){
  fmt.Printf(lessgo.Red.Code)
  fmt.Printf("[!] ")
  fmt.Printf("%-032s","000 Error")
  fmt.Printf("%07d",000)
  fmt.Printf("%-15s","  "+payload)
  fmt.Println(err)
  fmt.Printf(lessgo.Reset.Code)
  
}