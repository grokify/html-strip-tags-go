package strip

import "testing"

func TestBindStruct(t *testing.T) {
	// Sample struct
	type user struct {
		Name  string `strip_tag:"true"`
		Age   int
		Email string `strip_tag:"true"`
		Photo string
	}

	users := user{
		Name:  "<b>John Doe</b>",
		Age:   30,
		Email: "<script>alert('Hello')</script>",
		Photo: "<h1>aa</h1>profile.jpg",
	}

	if err := Bind(&users); err != nil {
		t.Error(err)
	}

	t.Log(users)
}
