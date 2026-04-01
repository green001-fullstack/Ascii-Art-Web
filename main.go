package main



import (
	"html/template"
	"net/http"
	"log"
	"ascii-art-web/ascii"
	"strings"
)


var tmpl = template.Must(template.ParseFiles("templates/index.html", "templates/error.html"))

type PageData struct{
	Output string;
	Message string;
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{Output: ""}
	if r.Method != http.MethodGet {
		http.Error(w, "Only a GET request is allowed here", http.StatusMethodNotAllowed)
		return
	}
	

	if r.URL.Path != "/"{
		w.WriteHeader(http.StatusNotFound)
		tmpl.ExecuteTemplate(w, "error.html", nil)
		return
	}
	tmpl.ExecuteTemplate(w, "index.html", data)
}

func asciiHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
		http.Error(w, "Only a POST request is allowed here", http.StatusMethodNotAllowed)
		return
	}


	// allowedBanners := map[string]bool{
	// 	"standard": true,
	// 	"shadow": true,
	// 	"thinkertoy": true,
	// }
	// if !allowedBanners[banner]{
	// 	banner = r.FormValue("standard")
	// }
	
	text := r.FormValue("text")
	banner := r.FormValue("banner")

	if strings.TrimSpace(text) == ""{
		data := PageData{Output: "You have to provide a text"}
		tmpl.Execute(w, data)
		return
	}

	text = strings.ReplaceAll(text, "\\n", "\n")
	text1 := strings.Split(text, "\n")
	var asciiArt strings.Builder
	for _, line := range text1{
		if strings.TrimSpace(line) == ""{
			// asciiArt.WriteString("\n")
			continue
		}
		asciiArt.WriteString(ascii.GenerateAscii(line, banner))
	}
	
	data := PageData{Output: asciiArt.String()}
	tmpl.Execute(w, data)
}

func downloadHandler( w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost{
		http.Error(w, "Only a POST request is allowed here", http.StatusMethodNotAllowed)
		return
	}

	ascii := r.FormValue("ascii")

	if strings.TrimSpace(ascii) == ""{
		
	}
}

func main() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii", asciiHandler)
	http.HandleFunc("/download", downloadHandler)

	log.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}