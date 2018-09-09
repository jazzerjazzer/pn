package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	pnapi "pn/api"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	serverURL string
)

func TestE2E(t *testing.T) {
	mux := http.NewServeMux()
	api := pnapi.API{}
	mux.HandleFunc("/promotions/", api.PromotionsHandler)
	mux.HandleFunc("/upload", api.UploadHandler)
	ts := httptest.NewServer(mux)
	defer ts.Close()
	serverURL = ts.URL

	// First upload ids.csv file
	postFile(t, "../../ids.csv")

	res, err := http.Get(ts.URL + "/promotions/9befe59a-2b4c-42d0-9944-41fb27adaffe")
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, res.StatusCode)

	found := pnapi.Found{}

	err = json.NewDecoder(res.Body).Decode(&found)
	require.NoError(t, err)

	require.Equal(t, pnapi.Found{
		ID:             "9befe59a-2b4c-42d0-9944-41fb27adaffe",
		Price:          "10.193619",
		ExpirationDate: "2018-06-03 15:37:29 +0200 CEST",
	}, found)

	// Let's upload the ids2.csv file
	// This request will implicitly delete ids.csv file
	// Hence, this id should not be found
	postFile(t, "../../ids2.csv")
	res, err = http.Get(ts.URL + "/promotions/9befe59a-2b4c-42d0-9944-41fb27adaffe")
	require.NoError(t, err)
	require.Equal(t, http.StatusNotFound, res.StatusCode)

	// But this id must be found, since it is in the ids2.csv file...
	res, err = http.Get(ts.URL + "/promotions/abcdefgh-2b4c-42d0-9944-41fb27adaffe")
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, res.StatusCode)

	found = pnapi.Found{}

	err = json.NewDecoder(res.Body).Decode(&found)
	require.NoError(t, err)

	require.Equal(t, pnapi.Found{
		ID:             "abcdefgh-2b4c-42d0-9944-41fb27adaffe",
		Price:          "10.193619",
		ExpirationDate: "2018-06-03 15:37:29 +0200 CEST",
	}, found)
}

func postFile(t *testing.T, filename string) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	fileWriter, err := bodyWriter.CreateFormFile("file", filename)
	require.NoError(t, err)

	fh, err := os.Open(filename)
	require.NoError(t, err)
	defer fh.Close()

	_, err = io.Copy(fileWriter, fh)
	require.NoError(t, err)

	contentType := bodyWriter.FormDataContentType()
	require.NoError(t, bodyWriter.Close())

	resp, err := http.Post(serverURL+"/upload", contentType, bodyBuf)
	require.NoError(t, err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("BODY: %v\n\n", string(body))
	require.NoError(t, err)
	require.NotEmpty(t, body)
}
