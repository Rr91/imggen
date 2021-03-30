// server.go
package server

func imgHandler(w http.ResponseWriter, msg string){
	rend(w, "favicon")
}


func faviconHandler(w http.ResponseWriter, r *http.Request){
	rend(w, "img")

}
func robotsHandler(w http.ResponseWriter, r *http.Request){
	rend(w, "img")

}
func pingHandler(w http.ResponseWriter, r *http.Request){
	rend(w, "PONG")

}

func rend(w http.ResponseWriter, r *http.Request)
{
	_, err := w.Write([]byte("img"))
	if err != nil{
		log.Println(err)
	}
}


func Run(){
	http.HandleFunc("/", imgHandler)
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/robots", robotsHandler)
	http.ListenAndServe(":8080", nil)

	log.Println("server starting ...")
	if err := http.ListenAndServe(":" + conf.getPost(), nil); err != nil{
		log.Fatalln(err)
	}
}