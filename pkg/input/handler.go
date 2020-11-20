package input 

import (
  "os"
  "fmt"
  "strconv"
  "io/ioutil"
  "strings"
  "github.com/neonify/lessgo/pkg/lessgo"
  )

func (Obj *Data)Handle(){
  
  url := Obj.Url
  file := Obj.File
  ConcReqs := Obj.ConcRequests
  FR := Obj.FollowRedirects
  GL := Obj.Grep
  HF := Obj.HeaderFyl
  PD := Obj.PostData
  
    lessgo.Color("blue","____________________________________________\n")
  
  UrlPrinter(url)
  
  if !strings.HasPrefix(url,"http"){
    lessgo.ErrorRaiser(lessgo.UrlDoesNotStartWithHttp)
  }
  
  if !strings.Contains(url,"FUZZ") &&
   !strings. Contains(PD,"FUZZ"){
    lessgo.ErrorRaiser(lessgo.FuzzableParamNotFound)
  }

  if file == ""{
    lessgo.ErrorRaiser(lessgo.FileNotSelected)
  }
  
  if FileExists(file) {
    lessgo.Color("blue","[*] Wordlist : "+file)
  }else{
    lessgo.ErrorRaiser(lessgo.FileNotFound)
  }
  
  FileLengthPrinter(file)
  lessgo.Color("blue","[*] Concurrent Threads : "+strconv.Itoa(ConcReqs))
  
  if FR == false{
  lessgo.Color("blue","[*] Follow Redirects : false")
  }else{
    lessgo.Color("blue","[*] Follow Redirects : true")
  }
  
  if PD == ""{
    Obj.Method = "GET"
  }else{
    Obj.Method = "POST"
  }
  
  lessgo.Color("blue","[*] Method : "+Obj .Method)
  
  StrGL := IntSlice2Str(GL)
  
  lessgo.Color("blue","[*] Grep List : "+StrGL)
  
  if FileExists(HF) == true{
    
    lessgo.Color("blue","[*] Header File : "+HF)
    
    
    lessgo.Color("blue","\nHEADERS\n")
    HeaderMap := HeaderParse(HF)
    
    for param,value := range(HeaderMap){
      fmt.Print(string(param)," : ")
      lessgo.Color("blue",string(value))
    }
  }else{
    fmt.Println(HF)
    lessgo.ErrorRaiser(lessgo.FileNotFound)
  }
  
  lessgo.Color("blue","\nPOST DATA\n")
  DataMap := DataParse(PD)
  
  fmt.Println(DataMap)
  

  lessgo.Color("blue","____________________________________________")
}
  
// Checks whether the file exists
func FileExists(file string) bool{
    FileExists := false
    _, err := os.Stat(file)
    
    if err == nil{
      FileExists = true
    }
    return FileExists
}

// Checks if any url is selected
func UrlPrinter(url string){
  if url == ""{
    lessgo.ErrorRaiser(lessgo.UrlNotSelected)
  }else{
    lessgo.Color("blue","[*] Target : "+url)
  }
}

// Prints the number of lines in the file
func FileLengthPrinter(file string){
  resp,_ := ioutil.ReadFile(file)
  cont := strings.Split(string(resp),"\n")
  
  lessgo.Color("blue","[*] Wordlist Length : "+strconv.Itoa(len(cont)))
}