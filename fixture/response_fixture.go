package fixture

import (
	"bytes"
	"encoding/json"
)

type ResponseFixture struct {
	Animal string `json:"animal"`
}

func RawResponseFixture() []byte {
	testStruct := ResponseFixture{Animal: "Scooby"}
	respBodyBytes := new(bytes.Buffer)
	json.NewEncoder(respBodyBytes).Encode(testStruct)

	return respBodyBytes.Bytes()
}
