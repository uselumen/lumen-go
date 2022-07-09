package lumengo

import (
	"testing"
)

var identifier = "383848002224"

var instance *Lumengo

func TestLumengo(t *testing.T) {

	instance = NewLumengo("<< api-key-here >>")
	t.Log("TestLumengo")
}

func TestIdentify(t *testing.T) {

	params := IdentifyParams{
		Email:     "test@tes.co",
		FirstName: "Gopher",
		LastName:  "Go",
	}

	err := instance.Identify(identifier, params)

	if err != nil {
		t.Error(err)
		return
	}

	// instance.GetApiKey()
	t.Log("TestIdentify Successs", err)
}

func TestTrack(t *testing.T) {

	params := map[string]interface{}{
		"productId": 100023449,
	}

	err := instance.Track(identifier, "Product Clicked", params)

	if err != nil {
		t.Error(err)
		return
	}

	// instance.GetApiKey()
	t.Log("TestIdentify Successs", err)
}
