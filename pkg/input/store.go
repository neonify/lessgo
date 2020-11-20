package input

import(
  "os"
  "net/url"
  "strings"
  "strconv"
  "io/ioutil"
  "go/build"
)

var IntSlice []int

func FindGoPath() string{
gopath := os.Getenv("GOPATH")
if gopath == "" {
        gopath = build.Default.GOPATH
    }
return gopath

}

func Str2IntSlice(str string)([]int){
  
  var IntSlice []int
  
  if strings.Contains(str,",") == true{
    StrSlice := strings.Split(str,",")
  
    for i := range(StrSlice){
      j,_ := strconv.Atoi(StrSlice[i])
      IntSlice = append(IntSlice,j)
    }
    
  }else if str != ""{
    j,_ := strconv.Atoi(str)
    IntSlice = append(IntSlice,j)
  }
  
  return IntSlice
}

func IntSlice2Str(IntS []int)string{
  
  Str := ""
  
  if len(IntS) != 0{

    for i := range(IntS){
      j := strconv.Itoa(IntS[i])
      Str += j+" "
    }
    
  }
    
  return Str
}


func HeaderParse(FylName string)(map[string]string){
  
  HeaderMap := make(map[string]string)
  
  body,_:=ioutil.ReadFile(FylName)
  
  ComboSlice := strings.Split(string(body),"\n")
  
  for i := range(ComboSlice){
    lyn := ComboSlice[i]
    if strings.Contains(lyn,":"){
      
      ActSlice := strings.Split(lyn,":")
      
      val := strings.Join(ActSlice[1:],":")
      
      Param := strings.TrimSpace(ActSlice[0])
      Value := strings.TrimSpace(val)
      
      HeaderMap[Param] = Value
    }
  }
  
  return HeaderMap
}

func DataParse(DataStr string)(*url.Values){
  
  data := url.Values{}
  
  if strings.Contains(DataStr,"&"){
    
    ComboSplit := strings.Split(DataStr,"&")
    for i := range(ComboSplit){
      lyn := ComboSplit[i]
      ActSplit := strings.Split(lyn,"=")
      
      param := ActSplit[0]
      value := ActSplit[1]
      
      data.Set(param,value)
    }
    
  }else if DataStr != ""{

      ActSplit := strings.Split(DataStr,"=")
      
      param := ActSplit[0]
      value := ActSplit[1]
      
      data.Set(param,value)
  }
  
  return &data
  
}