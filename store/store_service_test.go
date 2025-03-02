package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testStoreService = &StorageService{}

func init() {
	testStoreService = InitializeStore()

}

func TestStoreInit(t *testing.T) {
	assert.True(t, testStoreService.redisClient != nil)
}

func TestInsertionAndRetrieval(t *testing.T) {
	initialLink := "https://solscan.io/account/DCD53cE15Qpxaz6hcqLs462HZKSSTpGZtbBBtPsaQrY7"
	userUUId := "e0dba740-fc4b-4977-872c-d360239e6b1a"
	shortUrl := "Burner"

	SaveUrlMapping(shortUrl, initialLink, userUUId)

	res := RetrieveInitialUrl(shortUrl)

	assert.Equal(t, res, initialLink)

}
