       _
      | |___ ___ ___ ___ ___
      | | -_|_ -|_ -| . | . |
      |_|___|___|___|_  |___|
                    |___|

<img src="https://img.shields.io/badge/build-passing-green">   <img src="https://img.shields.io/badge/made with-go-orange">
<img src="https://img.shields.io/badge/go-v1.15-blue"> 
<img src="https://img.shields.io/badge/author-neonify-blue">      <img src="https://img.shields.io/badge/credits-nego-yellow">
<img src="https://img.shields.io/badge/license-GPL v3.0-green">

### 𝕝𝕖𝕤𝕤𝕘𝕠
A _`fast`_ web fuzzer in <img src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRuLO9M7e-FIj8ZEacXuSVv47NdBHScc3ZlMA&usqp=CAU" width="40px">

## Table Of Contents
* <a href="#FEATURES">`Features`</a> 
* <a href="#INSTALLATION">`Installation`</a>
* <a href="#USAGE">`Usage`</a>
* <a href="#EXAMPLES">`Examples`</a>
* <a href="#SUPPORT & FEEDBACK">`Support/Feedback`</a>

## FEATURES
* Speed
   * It is super fast as it uses golang
   * It has 50 go routines (threads) by default , you can alter them by passing
the flag `-c`
* Easy Usage
   * It is very easy to <a href="#INSTALLATION">`install`</a> & use

## REQUIREMENTS
* <a href="https://golang.org">`GoLang`</a>

## INSTALLATION
* `go get github.com/neonify/lessgo`
* `cd /usr/bin`
* `go build github.com/neonify/lessgo`


## USAGE 

* Enter the URL by passing the flag `-u` , replace the value to be fuzzed with the word `FUZZ`

* Enter the wordlist by passing the flag `-f`

Example : 
```
lessgo -u https://www.example.com/FUZZ -f wordlist.txt
```

#### FLAGS
```
REQUIRED
-u      to specify the url 

CHOOSE ANY ONE 
-f      to specify the wordlist
-subd   wordlist of common subdomains (length : 871)
-dirs   wordlist of common directories (length : 1273822)
-dirT   wordlist of path traversal payloads (length : 7988)
-lfi    wordlist of local file inclusion payloads (length : 961)

OPTIONAL 
-c      no of threads (default : 50)
-R      to follow redirects (default : false)
-G      to grep the given status codes (default : *)
-H      to specify the file containg headers
-D      to specify data to be sent (default : none)
```
### FUZZING GET DATA
__STEPS__
* Enter the parameters & their values as a part of the url
* Replace the value tp be fuzzed with the word `FUZZ`

__EXAMPLE__ 
```
https://www.example.com?id=123&pwd=FUZZ
```

### FUZZING POST DATA 
__STEPS__
* Enter the parameters and values by passing the flag `-D`
    * The parameters and their values should be differenciated by a `=` sign
    * A combination of a parameter and its value should be differnciated
from the other by a `&` sign.
* Replace the value to be fuzzed with the word `FUZZ`

__EXAMPLE__ 
``` 
-D "id=123&pwd=FUZZ"
 ```

## EXAMPLES
* For Specifying Headers 
```
lessgo -u https://www.example.com -H file_containing_headers.txt
```

* For Fuzzing `GET` data
```
lessgo -u https://www.example.com?id=FUZZ -f list.txt
```

* For Fuzzing `POST` data
``` 
lessgo -u https://www.example.com -D "uid=1001&pwd=FUZZ" -f list.txt
```

* For Grepping Status Codes
``` 
lessgo -u https://FUZZ.example.com -subd -G 403,404,400
```

* For subdomain fuzzing
``` 
lessgo -u https://FUZZ.example.com -subd
```

* Attacking for Directories/Path Traversal/Local File Inclusion
```
lessgo -u https://www.example.com/FUZZ -dirs/-dirT/-lfi
```


### VERSION
<strong>v 2.0.2</strong>

### SUPPORT & FEEDBACK

__SUPPORT__
* <a href="https://mobile.twitter.com/neonify4"><img src="https://img.shields.io/badge/twitter-follow-blue?logo=twitter&style=social"></a>
* <a href="https://reddit.com/user/n3onify"><img src="https://img.shields.io/badge/reddit-follow-red?logo=reddit&style=social"></a>
 
__FEEDBACK__
* <a href="mailto:lessgofuzzer@gmail.com"><img src="https://img.shields.io/badge/mail-us-red?logo=gmail&style=social"></a>
* <a href="https://github.com/neonify/lessgo/issues"><img src="https://img.shields.io/badge/raise an-issue-red?logo=github&style=social"></a>
