package main

import(
	"net/http"
	"log"
	"fmt"
	"io/ioutil"
	)

type Post struct {
	Title string
	Body []byte 
}


func (p *Post) save() error{
	filename := p.Title +".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}
func loadPage(title string) *Post{
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return &Post{Title:"Hata", Body:[]byte("Böyle bir sayfa yok!")}
	}
	return &Post{Title: title, Body: body}
}

func main(){
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServeTLS(":8000","cert.pem","key.pem",nil))
}


func handler(w http.ResponseWriter, r *http.Request) {
	baslik := r.URL.Path[len("/view/"):]
	govde := loadPage(baslik)
	fmt.Fprintf(w, "<h1>%s</h1><p>%s</p>",baslik, govde.Body)
	}
