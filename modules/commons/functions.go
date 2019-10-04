package commons

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
)

// Reflect if an interface is either a struct or a pointer to a struct
// and has the defined member method. If error is nil, it means
// the MethodName is accessible with reflect.
func ReflectStructMethod(Iface interface{}, MethodName string) error {
	ValueIface := reflect.ValueOf(Iface)

	// Check if the passed interface is a pointer
	if ValueIface.Type().Kind() != reflect.Ptr {
		// Create a new type of Iface, so we have a pointer to work with
		ValueIface = reflect.New(reflect.TypeOf(Iface))
	}

	// Get the method by name
	Method := ValueIface.MethodByName(MethodName)
	if !Method.IsValid() {
		return fmt.Errorf("Couldn't find method `%s` in interface `%s`, is it Exported?", MethodName, ValueIface.Type())
	}
	return nil
}

// Check if struct has field
// Reflect if an interface is either a struct or a pointer to a struct
// and has the defined member field, if error is nil, the given
// FieldName exists and is accessible with reflect.
func HasField(Iface interface{}, FieldName string) bool {
	ValueIface := reflect.ValueOf(Iface)

	// Check if the passed interface is a pointer
	if ValueIface.Type().Kind() != reflect.Ptr {
		// Create a new type of Iface's Type, so we have a pointer to work with
		ValueIface = reflect.New(reflect.TypeOf(Iface))
	}

	// 'dereference' with Elem() and get the field by name
	Field := ValueIface.Elem().FieldByName(FieldName)
	if !Field.IsValid() {
		// return fmt.Errorf("Interface `%s` does not have the field `%s`", ValueIface.Type(), FieldName)
		return false

	}
	return true
}

// respondWithError if the request fail i send the error
func RespondWithError(w http.ResponseWriter, code int, msg string) {
	RespondWithJSON(w, code, map[string]string{"error": msg})
}

//respondWithJSON generate the response with the items from db
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// respondWithError if the request fail i send the error
// func RespondWithError(w http.ResponseWriter, code int, msg string) {
// 	respondWithJSON(w, code, map[string]string{"error": msg})
// }
