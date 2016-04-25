package main


import  (
"fmt"
"strings"
)


func main() {
url:="http://gitlab.com/morphyesh1/asd.git"

  i := strings.IndexAny(url, "/")
  repo := strings.TrimRight(url[i+1:], ".git")
  s := strings.Split(repo, "/")
fmt.Println(s)
  fmt.Println(s[2] + "/" + s[3])

}
