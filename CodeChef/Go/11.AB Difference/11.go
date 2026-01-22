package main

import "fmt"

func main() {
    var a, b int

    fmt.Scanf("%d %d", &a, &b)
    
    correctAnsswer := a+b
    
    chefsAnswer := a*b
    
    differance := correctAnsswer - chefsAnswer
    
    if (differance < 0){
        differance = -differance
        
    }
    
    fmt.Println(differance)
}