package fuzz

import (
  "io/ioutil"
  "fmt"
  "os"
  "os/signal"
  "strings"
  "github.com/zenthangplus/goccm"
  "github.com/neonify/lessgo/pkg/lessgo"
  "github.com/neonify/lessgo/pkg/input"
  )


func ConcManager(Obj input.Data) {

  Output := make(map[int]int)
              
  D := make(chan os.Signal, 1) 
  signal.Notify(D, os.Interrupt) 

  Heading()
  
  list,err:= ioutil.ReadFile(Obj.File)
  
  lessgo.Checkerr(err)
  
  ArrList := strings.Split(string(list),"\n")
  
  nbJobs := len(ArrList)
  
  // Number of concurent go requests
  c := goccm.New(Obj.ConcRequests)
    
    for i := 0; i < nbJobs; i++ {
        c.Wait()
        go func(i int) {
            Output = Handle(Obj,ArrList[i],Output)
            c.Done()
            
            go func(){ 
              for _ = range D { 
                WindUp(Output)
                os.Exit(2)
              } 
            }()
            
            
            
        }(i)
    }
    
    c.WaitAllDone()
    
    WindUp(Output)
}


func WindUp(Output map[int]int){
  
    lessgo.Color("blue","____________________________________________")
    lessgo.Color("blue","\n\nPROGRESS\n")
    lessgo.Color("cyan"," Status    Requests")
    for stat,nos := range(Output){
      fmt.Printf("[*] ")
      fmt.Printf("%03d",stat)
      fmt.Printf(" : ")
      fmt.Println(nos)
    }
    
    lessgo.Color("green","\n[+] All the successful urls are stored in the succ_urls.txt file.\n")
}
	