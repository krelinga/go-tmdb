package tmdb_test

import (
	"context"
	"testing"

	"github.com/krelinga/go-tmdb"
)

func TestGetCollection(t *testing.T) {
	client := testClientOptions{useApiReadAccessToken: true}.newClient(t)
	collection, err := tmdb.GetCollection(context.Background(), client, 8091)
	if err != nil {
		t.Fatalf("Failed to get collection: %v", err)
	}

	checkField(t, int32(8091), collection, tmdb.Collection.ID)
	checkField(t, "Alien Collection", collection, tmdb.Collection.Name)
	checkField(t, "A science fiction horror film franchise, focusing on Lieutenant Ellen Ripley (Sigourney Weaver) and her battle with an extraterrestrial life-form, commonly referred to as \"the Alien\". Produced by 20th Century Fox, the series started with the 1979 film Alien, then Aliens in 1986, Alien³ in 1992, and Alien: Resurrection in 1997.", collection, tmdb.Collection.Overview)
	checkField(t, "/gWFHIY77cRVoBRGERwMHqpD27gc.jpg", collection, tmdb.Collection.PosterPath)
	checkField(t, "/6X42JnSMdo3dPAswOHUuvebdTq7.jpg", collection, tmdb.Collection.BackdropPath)
}