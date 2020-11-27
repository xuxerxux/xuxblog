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
		return &Post{Title:"Hata", Body:[]byte("there is not such a file!")}
	}
	return &Post{Title: title, Body: body}
}

func main(){
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8000",nil))
}


func handler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	body := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><p>%s</p>",title, body.Body)
	}
