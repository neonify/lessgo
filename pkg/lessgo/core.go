package lessgo

import (
	"fmt"
)

type Hue struct{
  Code string
}

var (
 Red Hue = Hue{"\033[0;31m"}
 Green Hue = Hue{"\033[0;32m"}
 Blue Hue = Hue{"\033[0;34m"}
 Yellow Hue = Hue{"\033[0;33m"}
 Cyan Hue = Hue{"\033[0;36m"}
 Gray Hue = Hue{"\033[1;30m"}
 Reset Hue = Hue{"\033[0m"}
)

func Color(clr string,txt string){
  switch clr {
    case "red":
      fmt.Println(Red.Code+txt+Reset.Code)
    case "green":
      fmt.Println(Green.Code+txt+Reset.Code)
    case "blue":
      fmt.Println(Blue.Code+txt+Reset.Code)
    case "yellow":
      fmt.Println(Yellow.Code+txt+Reset.Code)
    case "cyan":
      fmt.Println(Cyan.Code+txt+Reset.Code)
    case "gray":
      fmt.Println(Gray.Code+txt+Reset.Code)
  }
}


func Checkerr(err error){
	if err != nil{
		Color("red","Error")
		fmt.Println(err)
	}
}
