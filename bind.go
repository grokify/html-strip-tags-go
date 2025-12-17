package strip

import (
	"errors"
	"reflect"
)

func Bind(input interface{}) error {

	// Get the value of the input
	v := reflect.ValueOf(input)

	// Ensure the input is a pointer
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return errors.New("input is a nil pointer")
		}
		v = v.Elem()
	}

	// Get the type of the value
	t := v.Type()

	// Ensure the dereferenced value is a struct
	if t.Kind() != reflect.Struct {
		return errors.New("input is not a struct or pointer to a struct")
	}

	// Iterate through the fields
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)               // Get the field
		value := v.Field(i)               // Get the field's value
		tag := field.Tag.Get("strip_tag") // Get the 'validate' tag

		// Check if the tag is set to "true" and the field is a string
		if tag == "true" && value.Kind() == reflect.String {
			// Strip the HTML tags
			clean := StripTags(value.String())
			value.SetString(clean)
		}
	}

	return nil
}
