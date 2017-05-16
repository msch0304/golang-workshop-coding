package store

import (
        "github.com/golang/mock/gomock"
        "github.com/stretchr/testify/assert"
        "testing"
)

//go:generate go get github.com/golang/mock/mockgen
//go:generate $GOPATH/bin/mockgen -self_package store -package store -destination ./store/mocks_test.go
func TestFileStore_Get(t *testing.T) {
        store := newFileStore("test.json")
        assert.Equal(t, "test.json", store.FileName)
        assert.Equal(t, 0, len(store.StoreMap))
}

