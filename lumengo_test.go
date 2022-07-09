package lumengo

import (
	"testing"
)

var identifier = "383848002224"

func TestLumengo(t *testing.T) {
	t.Log("TestLumengo")
}

func TestIdentify(t *testing.T) {
	instance := NewLumengo("Zd3SV9PoCx1d4ui7m8JVvkAasC8ADfczu8")

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
	instance := NewLumengo("Zd3SV9PoCx1d4ui7m8JVvkAasC8ADfczu8")

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
