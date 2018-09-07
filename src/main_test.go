package main

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	// a := Api{Address: "localhost:5000"}
	// a.Run()

	resp, err := http.Get("http://localhost:5000/status")

	assert.Nil(t, err)
	assert.Equal(t, resp.StatusCode, 200)
}

// body, err := ioutil.ReadAll(resp.Body)
// if err != nil {
// 	assert.FailNow(t, "Failed to read status", +err.Error())
// }

// log.Println(string(body))

// req := test.NewRequest("localhost:5000")

// if body != nil {
// 	bs, err := json.Marshal(body)
// 	if err != nil {
// 		assert.FailNow(t, "failed to serialize request body: "+err.Error())
// 	}

// 	req = test.NewRequest(method, path, bytes.NewBuffer(bs))
// 	req.Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// }

// rsp := test.NewResponseRecorder()
// api.echo.ServeHTTP(req, rsp)
// return rsp.Status(), rsp.Body.String()

// func sendRequest() {
// 	url := "http://restapi3.apiary.io/notes"
// 	fmt.Println("URL:>", url)

// 	var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)
// 	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
// 	req.Header.Set("X-Custom-Header", "myvalue")
// 	req.Header.Set("Content-Type", "application/json")

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer resp.Body.Close()

// 	fmt.Println("response Status:", resp.Status)
// 	fmt.Println("response Headers:", resp.Header)
// 	body, _ := ioutil.ReadAll(resp.Body)
// 	fmt.Println("response Body:", string(body))
// }
