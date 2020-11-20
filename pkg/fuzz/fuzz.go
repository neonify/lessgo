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
    Output = Request(Obj,word2,Output)
  }
  
  return Output
}


func Request(Obj input.Data,payload string,Output map[int]int)(map[int]int){
  
  url := Obj.Url
  FR := Obj.FollowRedirects
  GrepList := Obj.Grep
  
  req,err := MakeReq(Obj)
  if err == nil{
  
    client := MakeClient(FR)
    resp,err := client.Do(req)
    
    if err != nil{
      ErrorPrint(err,payload)
    }else{
      if len(GrepList) == 0{
        Format(resp,payload,url)
      }else{
        if contains(GrepList,resp.StatusCode) == true{
          Format(resp,payload,url)
        }
        
      }
      
      
        if Exists(Output,resp.StatusCode)==true{
          Output[resp.StatusCode] += 1
        }else{
          Output[resp.StatusCode] = 1
        }
    
      
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