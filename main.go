package main

import	(
	"net/http"
	"log"
	"encoding/json"
	_"strings"
	_"io/ioutil"
	"fmt"
)

type Task struct {
	Username      string `json:username`
	Password      string `json:password`
	FirstElement  int    `json:first`
	SecondElement int    `json:second`
}

var serverAddress = "192.168.1.63"
var serverPort = ":8091"

func main() {
	resp, err := http.Get("http://" + serverAddress + serverPort +"/api/task") //Tanya ip

	if err != nil {
		log.Println("%s", err.Error())
	}

	log.Println("%v", resp.StatusCode)
	log.Println("%s", resp.Body)
    
    var task Task
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)

	err = decoder.Decode(&task)
	if err != nil {
		log.Println("JSON error %v", err)
	}

	fmt.Println("our resp struct - %+v\n", &task)
}
