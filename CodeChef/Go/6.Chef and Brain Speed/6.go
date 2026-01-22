package main
import "fmt"

func main(){
    var x,y int
    fmt.Scan(&x)
    fmt.Scan(&y)
    
    if y > x {
        fmt.Println("Yes")
    }else{
        fmt.Println("No")
    }
}
