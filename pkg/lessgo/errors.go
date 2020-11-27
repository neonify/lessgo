package lessgo

import (
  "fmt"
  "os"
  "regexp"
  "net"
)

type NewError struct {
   title   string
   desc    string
   example string
   code    int
}


var FileNotFound NewError = NewError{"File Not Found","Check the file path","lessgo -u http://www.example.com/FUZZ -f correct_path.txt",1001}

var UrlNotSelected NewError = NewError{"Url Not Selected","Choose a url by passing the flag [ -u <url> ]","lessgo -u http://www.example.com/FUZZ -dirs",1002}

var FileNotSelected NewError = NewError{"File Not Selected","Choose a file by passing the flag [ -f <file name> ]","lessgo -u http://www.example.com/FUZZ -f wordlist.txt",1003}

var FuzzableParamNotFound NewError = NewError{"No Fuzzable Parameter Found","Replace the value that should be fuzzed with the word 'FUZZ'","lessgo -u http://FUZZ.example.com -subd",1004}

var UrlDoesNotStartWithHttp NewError = NewError{"Url Selected Does Not Start With http(s)","Make sure the url starts with http(s)","lessgo -u http://www.example.com/FUZZ -dirs",1005}


func ErrorRaiser(ErrName NewError){
  fmt.Print(ErrName.code," - ")
  fmt.Println(ErrName.title)
  fmt.Println(ErrName.desc)
  fmt.Println("Example : ",ErrName.example)
  
  os.Exit(ErrName.code)
}


func RespErrHandle(err error)(string){
  
  var err2 string = err.Error()
  
  if err2  ==""{
    
  } else if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
    err2 = ("Timeout error")
    
  } else if match, _ := regexp.MatchString(".*lookup.*", err.Error()); match {
    err2 = ("Invalid url")
    
  } else if match, _ := regexp.MatchString(".*connection refused.*", err.Error()); match {
    err2 = ("Connection refused")
  
  } else if match, _ := regexp.MatchString(".*connection reset by peer.*", err.Error()); match {
    err2 = ("Connection reset by peer")
    
  }
  
  return err2
}