       _
      | |___ ___ ___ ___ ___
      | | -_|_ -|_ -| . | . |
      |_|___|___|___|_  |___|
                    |___|
<img src="https://img.shields.io/badge/Language-Golang-blue">    <img src="https://img.shields.io/badge/License-GNU General Public License v3.0-green">
<img src="https://img.shields.io/badge/Author-Neonify-red">      <img src="https://img.shields.io/badge/Credits-NeGo-yellow">


# lessgo
A __fast__ web fuzzer in golang

## SPEED 
This tool is written golang which is very fast. 
It can fuzz more than 500 words within a minute.
It by default has 50 threads which you can alter
by passing the flag "-c" . \
Example : 
```
lessgo -u https://www.example.com/FUZZ -dirT -c 100
```
<img src="https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/1200px-Go_Logo_Blue.svg.png">

## REQUIREMENTS
* GoLang

## INSTALLATION
* `go get github.com/neonify/lessgo`
* `cd /usr/bin`
* `go build github.com/neonify/lessgo`

__UPDATE__
* `go get -u github.com/neonify/lessgo`
* `cd /usr/bin`
* `go build github.com/neonify/lessgo`


## USAGE 

* Enter the URL by passing the flag "-u" replace the value to be fuzzed with the word "FUZZ"

* Enter the wordlist by passing the flag "-f"

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
* Replace the value tp be fuzzed with the word "FUZZ" 

__EXAMPLE__ \
`https://www.example.com?id=123&pwd=FUZZ`

### FUZZING POST DATA 
__STEPS__
* Enter the parameters and values by passing the flag "-D"
    * The parameters and their values should be differenciated by a "=" sign
    * A combination of a parameter and its value should be differnciated
from the other by a "&" sign.
* Replace the value to be fuzzed with the word "FUZZ" 

__EXAMPLE__ \
`-D "id=123&pwd=FUZZ"`

## EXAMPLES
For Specifying Headers 
* `lessgo -u https://www.example.com -H file_containing_headers.txt`

For Fuzzing __GET__ data
* `lessgo -u https://www.example.com?id=FUZZ -f list.txt`

For Fuzzing __POST__ data
* `lessgo -u https://www.example.com -D "uid=1001&pwd=FUZZ" -f list.txt`

For Grepping Status Codes
* `lessgo -u https://FUZZ.example.com -subd -G 403,404,400`

For subdomain fuzzing
* `lessgo -u https://FUZZ.example.com -subd`

Attacking for Directories/Path Traversal/Local File Inclusion
* `lessgo -u https://www.example.com/FUZZ -dirs/-dirT/-lfi`


### VERSION
<strong>v 2.0.2</strong>

### SUPPORT 
INSTAGRAM : https://m.instagram.com/kryshh_na
TWITTER : https://mobile.twitter.com/neonify4
