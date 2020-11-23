
package input

import (
  "flag"
  "os"
  "path/filepath"
  //"github.com/neonify/lessgo/pkg/lessgo"
  )
  
type Data struct{
  Url string
  File string
  Method string
  ConcRequests int
  FollowRedirects bool
  Grep []int
  Exclude []int
  HeaderFyl string
  PostData string
  Timeout int
}

var what string = ""

func Input()(Data){
  
  path := FindGoPath()

  ListPath := filepath.Join(path,"/src/github.com/neonify/lessgo/")
  
  DefaultHeaderFyl := filepath.Join(ListPath+"/lists/default_headers")
  
  url := flag.String("u","","website to be fuzzed.(Required)")
  subd := flag.Bool("subd",false,"default wordlist for fuzzing subdomains")
  dirs := flag.Bool("dirs",false,"default wordlist for fuzzing directories")
  lfi := flag.Bool("lfi",false,"default wordlist of local file inclusion payloads")
  file := flag.String("f","","wordlist name")
  ConcReqs := flag.Int("c",50,"No of concurrent requests")
  FollowRedirects := flag.Bool("R",false,"To follow redirects")
  GrepStr := flag.String("G","","To grep the given list of status codes")
  Exclude := flag.String("E","","To exclude the given statis codes")
  HeaderFile := flag.String("H",DefaultHeaderFyl,"To specify the file name containing headers")
  PostData := flag.String("D","","To specify POST data")
  TimeOut := flag.Int("T",0,"To specify timeout")
  
  flag.Usage = NewUsage
  flag.Parse()
  
  
  if len(os.Args) == 1{
    flag.Usage()
    os.Exit(213)
  }
  

  
  if *file != ""{
    what = *file
  }else if *subd != false{
    what = filepath.Join(ListPath,"lists/subd_list")
  }else if *dirs != false{
    what = filepath.Join(ListPath,"lists/dirs_list")
  }else if *lfi != false{
    what = filepath.Join(ListPath,"lists/local_file_inclusion_list")
  }
  
  GrepList := Str2IntSlice(*GrepStr)
  ExcList := Str2IntSlice(*Exclude)
  
  
  Obj := Data{*url,what,"",*ConcReqs,*FollowRedirects,GrepList,ExcList,*HeaderFile,*PostData,*TimeOut}
  
  return Obj
}
