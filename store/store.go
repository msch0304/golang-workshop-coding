package store

import (
        "encoding/json"
        "io/ioutil"
        "os"
)

type FileStore struct{
        FileName string
        StoreMap map[string]string
}


func NewFileStore(fname string) *FileStore {


        fstore :=  &FileStore{
                FileName: fname,
                StoreMap: make(map[string]string),
        }
        err := fstore.load()
        if (err != nil) {
                println(err.Error())

        }
        return fstore
}


func (store *FileStore) Set(key string, value string){
        store.StoreMap[key] = value;
        err := store.save();
        if (err != nil){
                println(err.Error())
        }
}

func (store *FileStore) Get(key string) (string, bool){
        if v, e := store.StoreMap[key]; e{
                return v, e
        } else {
            return v,e
        }

}

func (store *FileStore) Keys() ([]string) {
        keys := make([]string, len(store.StoreMap))
        for k := range  store.StoreMap {
           keys = append(keys,k)
        }
        return keys
}

func (store *FileStore) load() error{
        store.StoreMap = make(map[string]string)

        if _, err := os.Stat(store.FileName); os.IsNotExist(err) {
                return nil
        }

        bytes, err := ioutil.ReadFile(store.FileName)
        if err == nil {
                err = json.Unmarshal(bytes, &store.StoreMap)
        }
        return err
}

func (store *FileStore) save() error {
        bytes, err := json.Marshal(&store.StoreMap)
        if err == nil {
                err = ioutil.WriteFile(store.FileName, bytes, 0644)
        }
        return err
}