package render

import (
	"encoding/json"
	"net/http"
)

// JSON contains the given interface object.
type JSON struct {
	Data interface{}
}

var (
	jsonContentType = []string{"application/json; charset=utf-8"}
	// jsonpContentType     = []string{"application/javascript; charset=utf-8"}
	// jsonASCIIContentType = []string{"application/json"}
)

// Render (JSON) writes data with custom ContentType.
func (r JSON) Render(w http.ResponseWriter) (err error) {
	if err = WriteJSON(w, r.Data); err != nil {
		panic(err)
	}
	return
}

// WriteContentType (JSON) writes JSON ContentType.
func (r JSON) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, jsonContentType)
}

// WriteJSON marshals the given interface object and writes it with custom ContentType.
func WriteJSON(w http.ResponseWriter, obj interface{}) error {
	writeContentType(w, jsonContentType)
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	_, err = w.Write(jsonBytes)
	return err
}
