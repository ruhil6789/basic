// package rpcserver

// // go 2 :	"net/rpc" &net/rpc/jsonrpc
// // rpc server and client based upon two thing one is argument based and second is value based
// import (
//     "fmt"
//     "log"
//     "net"
// 	"net/rpc"
// 	"net/http"
// )

// // Define a type for our RPC service
// type Args struct
// type TimeServer int64

// func (t *TimeServer) Now(args *Args, reply *int64) error {
//     *reply = time.Now().UnixNano() / 1e6
//     return nil
// }

// func main(){
// // create a new rpc server
// timeserver:=new(TimeServer)
// // register the server
// rpc.Register(timeserver)
// rpc.HandleHttp()

// // listen for request on port 1234
// l,e:=net.Listen("tcp",":1234")

// if e !=nil{
// 	log.Fatal("listen error:",e)
// }
// http.Serve(l,nil)

// }

package main

import (
	"io/ioutil"
	"net/http"
	"net/rpc"
	"os"
jsonparse "encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
)

// Args hold argument passed to json rpc service

type Args struct {
}

type Book struct {
	Id     string "json:string,omitempty"
	Name   string "json:string,omitempty"
	Author string "json:string,omitempty"
}

type JSONServer struct{}


func (s *JSONServer) GetBook(r *rpc.Request, args *Args, reply *Book) error {


	var books []Book

	// read json file and load data
	 raw ,readerr:=ioutil.ReadFile("./books.json")
	 if readerr!=nil{
      
		return readerr
      os.Exit() 
	}

	// unmarshal json raw data into books array
	marshalerr:=jsonparse.Unmarshal(raw,&books)
	if marshalerr!=nil{
		return marshalerr
	}
	for _,book:=range books{
		if book.Id==args.Id{
            *reply=book
            break
        }
	}
}
func main() {

	r := mux.NewRouter()
	s := rpc.NewServer()
	s.RegisterCodec(json.NewCodec(), "application/json")
	// register service by creating a new json server
	s.RegisterService(new(JSONServer), "")
	r.Handle("/rpc", s)
	http.ListenAndServe(":1234", r)
	// http.Handle("/", r)
	// l,e:=net.Listen("tcp",":8080")
	// if e!=nil{
	//     log.Fatal("listen error:",e)
	// }
	// http.Serve(l,nil)
}
