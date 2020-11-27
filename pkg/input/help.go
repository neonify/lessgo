package input

import(
  "fmt"
  "github.com/neonify/lessgo/pkg/lessgo"
  )

func NewUsage(){
  fmt.Println(`
lessgo `+lessgo.Version()+` : A fast web fuzzer in golang

REQUIRED

:: Url
---------
-u      to specify the url 

CHOOSE ANY ONE 

:: Wordlists
-------------
-B      to specify a range of numbers to use them as payloads
-f      to specify the wordlist
-subd   wordlist of common subdomains (length : 871)
-dirs   wordlist of common directories (length : 1273822)
-lfi    wordlist of local file inclusion payloads (length : 961)

OPTIONAL 

:: Filters
-------------
-G      to grep the given status codes (default : *)
-E      to Exclude the given status codes (default : none)
-W      to grep a given word (default : none)

:: Other Options
------------------
-c      no of threads (default : 50)
-R      to follow redirects (default : false)
-H      to specify the file containg headers
-D      to specify data to be sent (default : none)
-T      to specify timeout (default : 0)

EXAMPLE

$ lessgo -u http://FUZZ.example.com -subd
(subdomain fuzzing)
$ lessgo -u http://www.example.com/FUZZ -dirs
(directory fuzzing)
$ lessgo -u http://example.com -D "id=FUZZ&uid=admin" -f wordlist.txt
(POST data fuzzing)
$ lessgo -u http://example.com?id=FUZZ&uid=admin -f wordlist.txt
(GET data fuzzing)


NOTE 
:: For further help read README.md
    
    `)
}