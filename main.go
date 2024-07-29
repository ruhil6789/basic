package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/julienschmidt/httprouter"
)

// type CustomServerMux struct {
// }

// func (p *CustomServerMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {

// 	if r.URL.Path == "/" {
// 		giveRandom(w, r)
// 	}
// }

// func giveRandom(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "random number:%f", rand.Float64())
// }

// func main() {

// 	mux := &CustomServerMux{}

// 	http.HandleFunc("/hello", MyServer)
// 	log.Fatal(http.ListenAndServe(":8080", mux))

// 	newMux := http.NewServeMux()

// 	newMux.HandleFunc("/randomFloat", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "random number:%f", rand.Float64())
// 	})
// 	newMux.HandleFunc("/randomInt", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintln(w, rand.Int())
// 	})

// 	// log.Fatal(http.ListenAndServe(":8080", newMux))

// 	// http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
// 	// 	fmt.Fprintf(w, "<h1>Welcome to the Simple API Server</h1>")
// 	// 	urlPathElement:= strings.Split(r.URL.Path,"/")

// 	// 	if urlPathElement[1]=="roman_number"{
// 	// 		number ,_:=strconv.Atoi(strings.TrimSpace(urlPathElement[2]))
// 	// 		// fmt.Fprintf(w, "<h2>The Roman Number for %d is: %s</h2>",number,ToRoman(number))

// 	// 		if number ==0 || number >10{
// 	// 			w.WriteHeader(http.StatusNotFound)
// 	// 			w.Write([]byte("404-Not Found"))
// 	// 		}else{
// 	// 			// fmt.Fprintf(w,"%q",html.EscapeString(roman))

// 	// 		}
// 	// 	}else{
// 	// 		w.WriteHeader(http.StatusNotFound)
// 	//         w.Write([]byte("404-Not Found"))
// 	// 	}
// 	// })

// 	// s:=&http.Server{
// 	// 	Addr:":8080",
// 	//     // Handler: nil,
// 	// 	ReadTimeout: 10*time.Second,
// 	// 	WriteTimeout: 10*time.Second,
// 	// 	MaxHeaderBytes: 1<<20,
// 	// }
// 	// s.ListenAndServe()
// }

// func MyServer(w http.ResponseWriter, r *http.Request) {
// 	io.WriteString(w, "world")

// }

//  this is the second

func getCommandOutput(command string, arguments ...string) string {

	cmd := exec.Command(command, arguments...)

	var out bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Start()
	if err != nil {
		log.Fatal(err, ":", stderr.String())
	}

	err = cmd.Wait()

	if err != nil {
		log.Fatal(err, ":", stderr.String())
	}
	return out.String()
}

func goVersion(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	fmt.Fprintf(w, getCommandOutput("/usr/local/bin/go", "version"))

}

func getFileContent(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprintf(w, getCommandOutput("/bin/cat", params.ByName("name")))

}

func main() {
	router := httprouter.New()
	router.GET("/api/v1/go-version", goVersion)
	router.GET("/api/v1/show-file/:name", getFileContent)
	log.Fatal(http.ListenAndServe(":8080", router))
}
