package main

import (
	"awesomeProject2/internal/car"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)




func main(){
 //log.Println("create router")
 router := httprouter.New()
 ////log.Println("register car handler")
	handler := car.NewHandler()
	handler.Register(router)
 start(router)

}
type CarData struct {

	ID        int    `json:"id"`
	Year     int `json:"year"`
	Make string `json:"make"`
	Model  string `json:"model"`
	Engine    string `json:"engine"`

}

 func start(router *httprouter.Router){

	listener, err := net.Listen("tcp", ":1234")
	if err != nil{
		panic(err)
	}


	server := &http.Server{
		Handler: router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}

	log.Println(server.Serve(listener))

	 resp, err := http.Get("http://api.carmd.com/v3.0/decode_enh?vin=1GNALDEK9FZ108495")
	 if err != nil {
		 fmt.Println("No response from request")
	 }
	 defer resp.Body.Close()
	 body, err := ioutil.ReadAll(resp.Body) // response body is []byte

	 var result CarData
	 if err := json.Unmarshal(body, &result); err != nil {  // Parse []byte to the go struct pointer
		 fmt.Println("Can not unmarshal JSON")
	 }
}