package jsons

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/rs/zerolog/log"
)

// encoding json with Jsoniter (https://jsoniter.com/) is much faster compare to standard lib

type Data struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

var jsoner = jsoniter.ConfigCompatibleWithStandardLibrary

func JsoniterMarshal(d any) {
	_, err := jsoner.Marshal(&d)
	if err != nil {
		log.Err(err)
	}
	//fmt.Printf("data bytes: %v", d)
}

func JsoniterUnMarshal(dataBytes []byte, d any) {
	err := jsoner.Unmarshal(dataBytes, &d)
	if err != nil {
		log.Err(err)
	}
	//fmt.Printf("data: %+v", d)
}
