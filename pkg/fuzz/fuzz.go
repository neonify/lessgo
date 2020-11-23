package fuzz

import (
  //"fmt"
	"strings"
	//"github.com/neonify/lessgo/pkg/lessgo"
	"github.com/neonify/lessgo/pkg/input"
)

func contains(s []int, str int) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
 

func Handle(Obj input.Data,word string,Output map[int]int)(map[int]int){

  word1 := strings.TrimSpace(word)
  word2 := strings.Replace(word1,"\n","",-1)
  
  if !(strings.HasPrefix(word2,"#")){
	  Obj.Url = strings.Replace(Obj.Url,"FUZZ",word2,-1)
	  Obj.PostData = strings.Replace(Obj.PostData,"FUZZ",word2,-1)
	  
	 HeaderMap := input.HeaderParse(Obj.HeaderFyl)
	 
	 for i,_ := range(HeaderMap){
	   if HeaderMap[i] == "FUZZ"{
	     HeaderMap[i] = word2
	   }
	 }
	  
    Output = Request(Obj,word2,Output,HeaderMap)
  }
  
  return Output
}


func Request(Obj input.Data,payload string,Output map[int]int,HeaderMap map[string]string)(map[int]int){
  
  url := Obj.Url
  FR := Obj.FollowRedirects
  GrepList := Obj.Grep
  ExcList := Obj.Exclude
  Tmout := Obj.Timeout
  
  req,err := MakeReq(Obj,HeaderMap)
  if err == nil{
  
    client := MakeClient(FR,Tmout)
    resp,err := client.Do(req)
    
    if err != nil{
      ErrorPrint(err,payload)
      Output = OutPrint(Output,000)
    }else{
      
      if len(GrepList) == 0 && len(ExcList) == 0{
        Format(resp,payload,url)
      }else if len(GrepList) != 0{
        if contains(GrepList,resp.StatusCode) == true{
          Format(resp,payload,url)
        }
      
      }else if len(ExcList) != 0{
        if contains(ExcList,resp.StatusCode) == false{
          Format(resp,payload,url)
        }
        
      }
      
      
      Output = OutPrint(Output,resp.StatusCode)
    
      
    }
  }
  return Output
  
}

func Exists(Map map[int]int,val int)(bool){
  
  for i,_ := range(Map){
    if i == val{
      return true
    }
  }
    
  return false
  
}

func OutPrint(Output map[int]int,code int)(map[int]int){
    if Exists(Output,code)==true{
          Output[code] += 1
    }else{
          Output[code] = 1
    }
    
    return Output
}