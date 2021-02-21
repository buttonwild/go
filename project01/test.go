package main

import "fmt"
func main(){
    for i:= 5;i > 0 ; i --{
        for ii := i;ii > 0 ; ii--{
            fmt.Print("*")
        }
        fmt.Println("\n")
    }
}