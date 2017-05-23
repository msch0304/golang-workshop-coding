package store

import (
        "github.com/golang/mock/gomock"
        "github.com/stretchr/testify/assert"
        "testing"
)

//go:generate go get github.com/golang/mock/mockgen
//go:generate $GOPATH/bin/mockgen -source $GOPATH/src/golang-workshop-coding/store/interfaces.go -package golang-workshop-coding/store -destination $GOPATH/src/golang-workshop-coding/store/mocks_test.go
func TestFileStore_Get(t *testing.T) {
        store := newFileStore("test.json")
        assert.Equal(t, "test.json", store.FileName)
        assert.Equal(t, 0, len(store.StoreMap))
}

