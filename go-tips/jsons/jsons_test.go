package jsons

import (
	"encoding/json"
	"testing"
)

type Payload struct {
	A int    `json:"a"`
	b string `json:"-"`
	C int    `json:"-"`
	D string `json:"d"`
	E string `json:"e"`
}

var testPayload = []byte(`{"a":123456,"b":"this is a short string","c":343242343,"d":"a new string as well","e":"last"}`)

func BenchmarkDecodeStd(b *testing.B) {
	b.ReportAllocs()
	var data Payload
	for i := 0; i < b.N; i++ {
		json.Unmarshal(testPayload, &data)
	}
}

func BenchmarkDecodeJsoniter(b *testing.B) {
	b.ReportAllocs()
	var data Payload
	for i := 0; i < b.N; i++ {
		JsoniterUnMarshal(testPayload, &data)
	}
}

func BenchmarkEncodeStd(b *testing.B) {
	var data Payload
	json.Unmarshal(testPayload, &data)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		json.Marshal(data)
	}
}

func BenchmarkEncodeJsoniter(b *testing.B) {
	var data Payload
	JsoniterUnMarshal(testPayload, &data)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		JsoniterMarshal(data)
	}
}
