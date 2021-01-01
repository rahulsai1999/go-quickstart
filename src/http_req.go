package src

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

//structure for unmarshalling json
type obj struct {
	Positive float32
	Negative float32
}
type resultBody struct {
	Status      string
	Predictions []obj
}
type requestBody struct {
	Text []string `json:"text"`
}

//HTTPRequests -> HTTP Requests using the network library
func HTTPRequests() {
	url := "http://localhost:5000/model/predict"

	//preparing the object
	reqobj := requestBody{}
	reqobj.Text = append(reqobj.Text, "Kubernetes is a great orchestration service", "Uber is not the best taxi service")

	//json conversion
	data, _ := json.Marshal(reqobj)

	//convert to byte[]
	payload := strings.NewReader(string(data))
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	var finobj resultBody
	json.Unmarshal([]byte(body), &finobj)

	fmt.Println(finobj.Predictions)
}

//RESTapi -> starting a web server using Labstack's Echo
func RESTapi() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}

//HTTPRequestWithReturn -> http request with returns
func HTTPRequestWithReturn(x string) (pos float32) {
	url := "http://localhost:5000/model/predict"

	//preparing the object
	reqobj := requestBody{}
	reqobj.Text = append(reqobj.Text, x)

	//json conversion
	data, _ := json.Marshal(reqobj)

	//convert to byte[]
	payload := strings.NewReader(string(data))
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	var finobj resultBody
	json.Unmarshal([]byte(body), &finobj)

	fmt.Println(finobj.Predictions)

	return finobj.Predictions[0].Positive
}
