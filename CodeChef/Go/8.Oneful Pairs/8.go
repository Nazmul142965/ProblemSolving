package main
import "fmt"

func main(){
    var a,b int
    fmt.Scan(&a)
    fmt.Scan(&b)
    
    if a + b + (a*b) == 111{
        fmt.Println("Yes")
    }else{
        fmt.Println("No")
    }
    
}
