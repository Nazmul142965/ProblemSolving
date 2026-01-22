package main
import "fmt"

func main(){
    var x,y int
    fmt.Scan(&x)
    fmt.Scan(&y)
    
    if (x >= 2*y){
        fmt.Println("Yes")
    }else{
        fmt.Println("No")
    }
}
