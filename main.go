package main

import(
	"html/template"
	"net/http"
	"log"
	// "fmt"
	"io/ioutil"
	"errors"
	"time"
	"encoding/hex"
	)

type Post struct {
	Title string
	Body []byte
}


func (p *Post) save() error{
	filename := p.Title +".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}
func loadPage(title string) (*Post, error){
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return &Post{}, errors.New("loading page was not possible, empty Post returned")
	}
	return &Post{Title: title, Body: body}, nil
}

var templates = template.Must(template.ParseFiles("edit.html", "view.html"))
func main(){
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8000",nil))
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &Post{Title:title, Body:[]byte(body)}
	p.save()
	http.Redirect(w,r,"/view/"+title, http.StatusFound)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	if title == "" {
		src := []byte(time.Now().Weekday().String())
		temp := make([]byte, hex.EncodedLen(len(src)))
		hex.Encode(temp, src)
		title = string(temp)
	}
	body , err:= loadPage(title)
	if err != nil {
		body = &Post{Title:title}
	}
	err = templates.ExecuteTemplate(w,"edit.html",body)
	if err != nil {
		http.Error(w, err.Error(),http.StatusInternalServerError)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	body , err:= loadPage(title)
	if err != nil {
		http.Redirect(w,r,"/edit/"+title, http.StatusFound) // this is the way to redirect!
		// editHandler(w, r) // make this a redirect with warning! not let it be seemless
		return
	}
	err = templates.ExecuteTemplate(w,"view.html",body)
	if err != nil {
		http.Error(w, err.Error(),http.StatusInternalServerError)
	}
}
