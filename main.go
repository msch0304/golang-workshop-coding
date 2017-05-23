package main

import (

       "golang-workshop-coding/server"
       "log"
       "io/ioutil"
       "fmt"
       "strings"
       "net/http"
)



func main(){
       server.Handle()
//       start()
}

func start() {

       client := &http.Client{}

       req, err := http.NewRequest(http.MethodPut, "http://localhost:8090/keyStore", strings.NewReader(""))
       req.Form.Add("key", "test2")
       req.Form.Add("value", "12234")

       response , err := client.Do(req)

       if err != nil {
              log.Fatal(err)
       } else {
              defer response.Body.Close()
              contents, err := ioutil.ReadAll(response.Body)
              if err != nil {
                     log.Fatal(err)
              }
              fmt.Println("The calculated length is:", len(string(contents)))
              fmt.Println("   ", response.StatusCode)

              fmt.Println(contents)
       }

       response, err = http.Get("http://localhost:8090/keyStore/list")
       if err != nil {
              panic(err)
       }

       if err != nil {
              log.Fatal(err)
       } else {
              defer response.Body.Close()
              contents, err := ioutil.ReadAll(response.Body)
              if err != nil {
                     log.Fatal(err)
              }
              fmt.Println("The calculated length is:", len(string(contents)))
              fmt.Println("   ", response.StatusCode)

              fmt.Println(contents)
       }

}