package main
import(
	"net/http"
	)

func main(){
	resp, _ := http.Get("https://example.com")
	// got the response!
}
