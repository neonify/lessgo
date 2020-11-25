
```
     _
    | |___ ___ ___ ___ ___
    | | -_|_ -|_ -| . | . |
    |_|___|___|___|_  |___|
                  |___|
```

<p align="center">
<img src="https://img.shields.io/badge/build-passing-green"><br>
<img src="https://img.shields.io/badge/made with-go-orange">
<img src="https://img.shields.io/badge/go-v1.15-blue?logo=go"> 
<img src="https://img.shields.io/badge/author-neonify-blue">      
<img src="https://img.shields.io/badge/credits-nego-yellow">
<img src="https://img.shields.io/badge/license-GPL v3.0-green">
</p>

<h3 align="center">𝖑𝖊𝖘𝖘𝖌𝖔</h3>
<p align="center">
A <code>fast</code> web fuzzer in golang
</p><hr>


## Table Of Contents
* <a href="#FEATURES">`Features`</a> 
* <a href="#INSTALLATION">`Installation`</a>
* <a href="#USAGE">`Usage`</a>
* <a href="#EXAMPLES">`Examples`</a>
* <a href="#SUPPORT-FEEDBACK">`Support/Feedback`</a>

## FEATURES
* Speed
   * It is super fast as it uses golang
   * It has 50 go routines (threads) by default , you can alter them by passing
the flag `-c`
* Easy Usage
   * It is very easy to <a href="#INSTALLATION">`install`</a> & use
* Cross Platform

## REQUIREMENTS
* <a href="https://golang.org">`GoLang`</a>

## INSTALLATION
###### Same works for updating
``` bash
$ go get -u github.com/neonify/lessgo
$ cd /usr/bin
$ go build github.com/neonify/lessgo
```

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
-lfi    wordlist of local file inclusion payloads (length : 961)

OPTIONAL 
-h      for help
-c      no of threads (default : 50)
-R      to follow redirects (default : false)
-G      to grep the given status codes (default : *)
-E      to exclude/hide the given status code (default : none)
-H      to specify the file containg headers
-D      to specify data to be sent (default : none)
-T      to specify timeout (default : 0)
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
<details>
<summary>Command</summary>
<br>
<pre>
lessgo -u https://www.example.com -D "uid=1001&pwd=FUZZ" -f list.txt
</pre>
</details>


### FUZZING HEADERS
__STEPS__
* Paste the headers in a file 
   * The parameters and their values should be differenciated by a `:` sign
   * A combination of parameter and its value should be differenciated from other by a new line
* Replace the value/parameter to be fuzzed with the word `FUZZ`
* Specify the file name by passing the flag `-H`

__EXAMPLE OF FILE CONTAINING HEADERS__
```
Host: FUZZ
Accept: */*
Connection: close
User-Agent: lessgo fuzzer
Referer: www.google.com
```
<details>
<summary>Command</summary>
<br>
<pre>
lessgo -u https://www.example.com -H file_containing_headers.txt
</pre>
</details>

## EXAMPLES

* For Fuzzing `GET` data
```
lessgo -u https://www.example.com?id=FUZZ -f list.txt
```

* For Grepping Status Codes
``` 
lessgo -u https://FUZZ.example.com -subd -G 403,404,400
```

* For subdomain fuzzing
``` 
lessgo -u https://FUZZ.example.com -subd
```

* Attacking for Directories/Local File Inclusion
```
lessgo -u https://www.example.com/FUZZ -dirs/-lfi
```


### VERSION
<strong>v 2.1.0</strong>

### SUPPORT-FEEDBACK

__SUPPORT__ \
<a href="https://mobile.twitter.com/neonify4"><img src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQE5lwX-Jw06voMwUSqXccNMHPQkWMTx4Odvg&usqp=CAU" width="70px"></a>
<a href="https://reddit.com/user/n3onify"><img src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRSbvqesjWPhKrOOHcJABPo2-7uvM4iapo3Gw&usqp=CAU" width="80px"></a>

__FEEDBACK__ \
<a href="mailto:lessgofuzzer@gmail.com"><img src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTLx3fjKDNLRpmJS8MpPLMleNNALBrgoE2VTA&usqp=CAU" width="70px"></a>
<a href="https://github.com/neonify/lessgo/issues"><img src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRhgYgVFsExjMYtlmONymM58gcWsRGKlgb7FQ&usqp=CAU" width="80px" height="40px"></a>

__STARGAZERS__
[![Stargazers repo roster for @neonify/lessgo](https://reporoster.com/stars/neonify/lessgo)](https://github.com/neonify/lessgo/stargazers)
__FORKERS__
[![Forkers repo roster for @neonify/lessgo](https://reporoster.com/forks/neonify/lessgo)](https://github.com/neonify/lessgo/network/members)


<hr>
<p align="center">
<a href="#">
<img src="https://mozpk.com/wp-content/uploads/2016/11/Back-to-top-button.png" height="30px">
</a></p><hr>
