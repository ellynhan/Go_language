package main

import (
    "code.google.com/p/go-tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {
    m:=make(map[string]int)
    var arr []string
    arr=strings.Fields(s)
    for i:=0; i<len(arr); i++ {
        _, ok:=m[arr[i]]
        if(ok==false){
        	m[arr[i]]=1
        }else{
            m[arr[i]]++
        }
    }
    return m
    
}

func main() {
    wc.Test(WordCount)
}
