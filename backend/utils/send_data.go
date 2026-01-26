package utils

import (
	"encoding/json"
	"net/http"
)

// interface{} is datatype which can represent every datatype. In this case, not matter which datatype int/float/slice/array it can handle it.
// sendData writes an HTTP response with the given status code and encodes the provided data as JSON into the response body.
func SendData(w http.ResponseWriter, data interface{}, statusCode int) {
	//set the status code and send it as response
	w.WriteHeader(statusCode)
	// Creating Encoder object
	encoder := json.NewEncoder(w)
	// Converting text into JSON.
	encoder.Encode(data)
}
