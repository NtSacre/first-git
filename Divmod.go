package main
import "fmt"

func Divmod(a,b int, div *int,mod *int){
*div,*mod=a/b,a%b
}
func main (){
 a:=13
 b:=2
 var(
 div int
 mod int
 )
 Divmod(a,b,&div,&mod)
 fmt.Println(div)
 fmt.Println(mod)
}