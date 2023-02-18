package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"web_test/csvdecoder"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/main", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, 200)
}

func TestGet(t *testing.T) {
	var newAns csvdecoder.Person
	persons := csvdecoder.Convert("ueba.csv")

	ans, _ := get(persons, []string{"2163 1835"})
	json.Unmarshal([]byte(ans[0]), &newAns)
	
	assert.Equal(t, 1835, newAns.Id)
	assert.Equal(t, "Цара Селезнёва", newAns.Cn)
}