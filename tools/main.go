package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
        "io/ioutil"
)

func main() {

	client := &http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/repos", "https://api.github.com/users", "sfawcett191"), nil)
	if err != nil {
		log.Fatal(err)
	}

	r, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer r.Body.Close()

        repos := make([]map[string]interface{}, 0)
        body, err := ioutil.ReadAll(r.Body)

        err = json.Unmarshal( body, &repos) 
	if err != nil {
	        fmt.Printf("%s\n", r.Body)
		log.Fatal(err)
	}

        flatten( &repos )
}

func flatten( repos *[]map[string]interface{} ) []interface{} {

  if repos != nil {
     ois := make([]interface{}, len(*repos), len(*repos))

     for i, m := range *repos {
        oi := make(map[string]interface{})

        oi["id"] = m["id"]
        oi["name"] = m["name"]

        fmt.Println( m["id"] )
        fmt.Println( m["name"] )

        ois[i] = oi 
     }
     return ois
  }
   
  return make( []interface{} , 0 )
}
