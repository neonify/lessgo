package fuzz

import(
  "os"
  "net/http"
  "github.com/neonify/lessgo/pkg/lessgo"
  )


func Format(resp *http.Response, payload string,url string){
  
      // success
      if resp.StatusCode <= 299 && resp.StatusCode >= 200{
      Printer(lessgo.Green,"+",resp,payload)
      
      fyl,err := os.OpenFile("succ_urls.txt",os.O_CREATE|os.O_APPEND|os.O_WRONLY,0644)
      lessgo.Checkerr(err)
      
      defer fyl.Close()
      
      _,err = fyl.WriteString(url+"\n")
      lessgo.Checkerr(err)
      
    // page not found
    }else if resp.StatusCode == 404{
     Printer(lessgo.Red,"!",resp,payload)
    
    // 400 codes
    }else if resp.StatusCode <= 499 && resp.StatusCode >= 400{
     Printer(lessgo.Blue,"*",resp,payload)
     
    // redirects
    }else if resp.StatusCode <= 399 && resp.StatusCode >= 300{
   
      Printer(lessgo.Cyan,"-",resp,payload)
      
    // server errors
    }else{
      Printer(lessgo.Yellow,"~",resp,payload)
    }
    
}
