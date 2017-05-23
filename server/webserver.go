package server


import (
        "fmt"
        "net/http"
//        "golang-workshop-coding/store"
        "encoding/json"
        "golang-workshop-coding/store"
)

func Handle() {

        mux := http.NewServeMux()

        store := store.NewFileStore("web.json")
        mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                fmt.Fprintf(w, "No function!")
        }))
        mux.Handle("/status", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                fmt.Fprintf(w, "ok!")
        }))
        mux.Handle("/keyStore/list", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                c, e := keysToJson(store.Keys())
                if (e == nil){
                        w.Write(c)
                }else{
                        fmt.Fprintf(w, "Error during request!")
                }
        }))
        mux.Handle("/keyStore", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                switch (r.Method){
                        case http.MethodGet:
                                key:= r.FormValue("key")
                                val, exists := store.Get(key)
                                if exists {
                                        fmt.Fprintf(w, "Value für %v is %v", key, val)
                                }else{
                                        fmt.Fprintf(w,"No value found for %v", key)
                                }
                        case http.MethodPut:
                                key := r.FormValue("key")
                                val := r.FormValue("value")
                                if (key  != "" && val != ""){
                                        store.Set(key, val)
                                        fmt.Fprintf(w, "Value %v für %v inserted", val,key)
                                }else{
                                        fmt.Fprintf(w, "Key and/or value was empty");
                                }
                }
        }))
        err := http.ListenAndServe(":8090", mux)
        if err != nil {
                fmt.Printf(err.Error())
                panic(err)

        }
}


func keysToJson(keys []string) ([]byte, error){
        c , e := json.Marshal(keys)
        return c, e
}