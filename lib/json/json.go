package json

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ReadBytes(url string) ([]byte, error) {

	response, error := http.Get(url)
	if error != nil {
		panic(error.Error())
	}

	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}

func Get(url string, jsonObject interface{}) error {

	body, err := ReadBytes(url)
	if err != nil {
		panic(err.Error())
	}
	return Deserialize(body, &jsonObject)
}

func Deserialize(jsonText []byte, jsonObject interface{}) error {
	err := json.Unmarshal(jsonText, &jsonObject)
	return err
}
