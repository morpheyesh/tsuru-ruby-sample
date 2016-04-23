package main

import (
     "fmt"
      )

func main() {

git := "www.github.com"

   if git == "gitlab.com" || git == "www.gitlab.com" {

    fmt.Println("gitlab")

   } else if git == "github.com" || git == "www.github.com" {

     fmt.Println("github.com")
 }     
}
