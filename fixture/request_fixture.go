package fixture

import (
	"bytes"
	"encoding/json"
)

type RequestFixture struct {
	Name string `json:"name"`
}

func RawRequestFixture() json.RawMessage {
	testStruct := RequestFixture{Name: "Test"}
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(testStruct)

	return reqBodyBytes.Bytes()
}
