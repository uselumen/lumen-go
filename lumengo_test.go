package lumengo_test

import (
	"context"
	"testing"

	lumengo "github.com/uselumen/lumen-go"
)

var instance *lumengo.Lumengo

var identifier = "12203315"

func TestLumengo(t *testing.T) {
	instance = lumengo.NewLumengo("Zd3SV9PoCx1d4ui7m8JVvkAasC8ADfczu8")
	t.Log("TestLumengo")
}

func TestIdentify(t *testing.T) {
	params := lumengo.IdentifyParams{
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

func TestIdentifyCtx(t *testing.T) {
	params := lumengo.IdentifyParams{
		Email:     "test@tes.co",
		FirstName: "Gopher",
		LastName:  "Go",
	}

	err := instance.IdentifyCtx(context.Background(), identifier, params)

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

func TestTrackCtx(t *testing.T) {

	params := map[string]interface{}{
		"productId": 100023449,
	}
	err := instance.TrackCtx(context.Background(), identifier, "Product Clicked", params)

	if err != nil {
		t.Error(err)
		return
	}

	// instance.GetApiKey()
	t.Log("TestIdentify Successs", err)
}
