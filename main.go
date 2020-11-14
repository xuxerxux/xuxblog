package main
import(
	"net/http"
	"fmt"
	)

func main(){
	resp, _ := http.Get("https://example.com")
	for name, cont := range resp.Header {
		fmt.Println(name, cont[0])
		// how to work with slicses in http package
	}
}
