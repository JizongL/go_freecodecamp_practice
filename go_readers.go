package main 
import (
	"fmt"
	"io"
	
	"os"
	"io/ioutil"
	"strings")





func main(){
	// make a reader
	//var r1 io.Reader
	
	r1,err := os.Open("./readme.md")
	if err!=nil{
		fmt.Println(err)
	}
	
	buf, err := ioutil.ReadAll(r1)
	fmt.Println(string(buf),"buf")
	


	// method of making a reader with strings
	r:= strings.NewReader("Testing String")
	arr:=make([]byte,8)
	for{
		n,err:= r.Read(arr)
		fmt.Printf("n=%v err=%v b=%v\n",n,err,arr)
		fmt.Printf("%q\n",arr[:n])
		if err == io.EOF{
			break
		}
	}
}