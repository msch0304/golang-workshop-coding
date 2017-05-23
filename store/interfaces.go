package store



type Store interface {
        Set(key, value string)
        Get(key string) (value string, exists bool)
        Keys() (keys []string)
        save() (err error)
        load() (err error)
}


