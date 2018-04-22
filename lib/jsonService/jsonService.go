package jsonService

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func ReadByts(url string) ([]byte, error) {

	response, error := http.Get(url)
	fmt.Println(response.Body)
	if error != nil {
		panic(error.Error())
	}

	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}

func GetJson(url string, jsonObject interface{}) error {

	response, error := http.Get(url)
	if error != nil {
		panic(error.Error())
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err.Error())
	}

	err = json.Unmarshal(body, &jsonObject)

	if err != nil {
		panic(err.Error())
	}
	return err
}
