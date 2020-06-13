package main

import (
	"fmt"
	"io/ioutil"
	link "linkparser/lib"
	"strings"
)

func check(e error) {
  if e != nil {
      panic(e)
  }
}

func findExample() []string {
  const com string = "html"
  fileinfo, err := ioutil.ReadDir(".")
  check(err)
  var retlist [] string
  for _, file := range fileinfo {
    str := file.Name()
    prefix, flag := "", false
    for _ , c := range str{
      if flag {
        prefix += string(c)
      }
      if c == '.'{
        flag = true
      }
    }
    if prefix == com {
      retlist = append(retlist, str)
    }
  }
  return retlist
}

func main()  {

  filelist := findExample()
  for _, filepath := range filelist {
    dat, err := ioutil.ReadFile(filepath)
    check(err)
    r := strings.NewReader(string(dat))
    links, err := link.Parse(r)
    if err != nil {
      panic(err)
    }
    fmt.Printf("# %s test \n", filepath)
    fmt.Printf("%+v\n", links)
  }
}