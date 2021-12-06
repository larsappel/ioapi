package ioapi

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// Get makes api GET requests for JSON
func GetJson(baseURL string, endpoint string, params string, header map[string]string) ([]byte, error) {

	reqUrl := baseURL + endpoint + "?" + params
	fmt.Println(reqUrl)
	httpClient := &http.Client{}

	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		log.Println(err)
	}

	for k, v := range header {
		req.Header.Add(k, v)
	}

	res, err := httpClient.Do(req)
	if err != nil || res.StatusCode != http.StatusOK {
		log.Println("Call returned:", err)
		log.Printf("Status code is: %d\n", res.StatusCode)
	}
	defer res.Body.Close()

	// Read the content of the respons body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}

	return body, err
}

// // Post posts to the server
// func Post(endpoint string, postBody string, authorization string) string {
// 	exceptionless := GetClient()

// 	var baseURL string = GetBaseURL()
// 	if exceptionless.ServerURL != "" {
// 		baseURL = exceptionless.ServerURL
// 	}

// 	url := baseURL + endpoint
// 	var jsonStr = []byte(postBody)
// 	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
// 	req.Header.Set("Authorization", "Bearer "+authorization)
// 	req.Header.Set("Content-Type", "application/json")
// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer resp.Body.Close()
// 	// body, _ := ioutil.ReadAll(resp.Body)
// 	return string(resp.Status)
// }
