package fuzz

import (
  "net/http"
  "time"
  "strings"
  "github.com/neonify/lessgo/pkg/input"
  )
  
func MakeReq(Obj input.Data)(*http.Request,error){
  var req *http.Request
  var err error
  
  if Obj.Method == "GET"{
    req,err = http.NewRequest(Obj.Method,Obj.Url,nil)
    
  }else if Obj.Method == "POST"{
    
    data := input.DataParse(Obj.PostData)
    
    req,err = http.NewRequest(Obj.Method,Obj.Url,strings.NewReader(data.Encode()))
  }
  
  HeaderMap := input.HeaderParse(Obj.HeaderFyl)
  
  for param, value := range(HeaderMap){
    req.Header.Set(param,value)
  }
  
  return req,err
  
}

func MakeClient(FR bool)(http.Client){
  timeout := 10 * time.Second
  tr := &http.Transport{
      MaxIdleConns:        1000,
			MaxIdleConnsPerHost: 500,
			MaxConnsPerHost:     500,
    }
  
  client := http.Client{Timeout:timeout,Transport : tr}
  
  if FR == false{
    client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
        return http.ErrUseLastResponse
      }
  }
  
  return client
}