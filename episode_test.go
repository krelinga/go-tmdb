package tmdb_test

import (
	"context"
	"testing"

	"github.com/krelinga/go-tmdb"
)

func TestGetEpisode(t *testing.T) {
	client := testClientOptions{useApiReadAccessToken: true}.newClient(t)
	episode, err := tmdb.GetEpisode(context.Background(), client, 1399, 1, 1, tmdb.WithAppendToResponse("credits", "external_ids"))
	if err != nil {
		t.Fatalf("failed to get episode: %v", err)
	}

	checkField(t, "2011-04-17", episode, tmdb.Episode.AirDate)
	if crew, err := episode.Crew(); err != nil {
		t.Fatalf("failed to get crew: %v", err)
	} else if writer, err := findCredit(crew, "David Benioff"); err != nil {
		t.Fatalf("failed to find writer: %v", err)
	} else {
		checkField(t, "Writer", writer, tmdb.Credit.Job)
		checkField(t, "Writing", writer, tmdb.Credit.Department)
		checkField(t, "5256c8a019c2956ff6046e2b", writer, tmdb.Credit.CreditID)
		checkField(t, false, writer, tmdb.Credit.Adult)
		checkField(t, int32(2), writer, tmdb.Credit.Gender)
		checkField(t, int32(9813), writer, tmdb.Credit.ID)
		checkField(t, "Writing", writer, tmdb.Credit.KnownForDepartment)
		checkField(t, "David Benioff", writer, tmdb.Credit.Name)
		checkField(t, "David Benioff", writer, tmdb.Credit.OriginalName)
		checkField(t, 0.9627, writer, tmdb.Credit.Popularity)
		checkField(t, "/xvNN5huL0X8yJ7h3IZfGG4O2zBD.jpg", writer, tmdb.Credit.ProfilePath)
	}
	checkField(t, int32(1), episode, tmdb.Episode.EpisodeNumber)
	checkField(t, "standard", episode, tmdb.Episode.EpisodeType)
	if guestStarts, err := episode.GuestStars(); err != nil {
		t.Fatalf("failed to get guest stars: %v", err)
	} else if len(guestStarts) == 0 {
		t.Fatal("expected guest stars, got none")
	} else if guestStar, err := findCredit(guestStarts, "Joseph Mawle"); err != nil {
		t.Fatalf("failed to find guest star: %v", err)
	} else {
		checkField(t, "Benjen Stark", guestStar, tmdb.Credit.Character)
		checkField(t, "5256c8b919c2956ff604836a", guestStar, tmdb.Credit.CreditID)
		checkField(t, int32(61), guestStar, tmdb.Credit.Order)
		checkField(t, false, guestStar, tmdb.Credit.Adult)
		checkField(t, int32(2), guestStar, tmdb.Credit.Gender)
		checkField(t, int32(119783), guestStar, tmdb.Credit.ID)
		checkField(t, "Acting", guestStar, tmdb.Credit.KnownForDepartment)
		checkField(t, "Joseph Mawle", guestStar, tmdb.Credit.Name)
		checkField(t, "Joseph Mawle", guestStar, tmdb.Credit.OriginalName)
		checkField(t, 2.7396, guestStar, tmdb.Credit.Popularity)
		checkField(t, "/1Ocb9v3h54beGVoJMm4w50UQhLf.jpg", guestStar, tmdb.Credit.ProfilePath)
	}
	checkField(t, "Winter Is Coming", episode, tmdb.Episode.Name)
	checkField(t, "Jon Arryn, the Hand of the King, is dead. King Robert Baratheon plans to ask his oldest friend, Eddard Stark, to take Jon's place. Across the sea, Viserys Targaryen plans to wed his sister to a nomadic warlord in exchange for an army.", episode, tmdb.Episode.Overview)
	checkField(t, int32(63056), episode, tmdb.Episode.ID)
	checkField(t, "101", episode, tmdb.Episode.ProductionCode)
	checkField(t, int32(62), episode, tmdb.Episode.Runtime)
	checkField(t, int32(1), episode, tmdb.Episode.SeasonNumber)
	checkField(t, "/9hGF3WUkBf7cSjMg0cdMDHJkByd.jpg", episode, tmdb.Episode.StillPath)
	checkField(t, 8.097, episode, tmdb.Episode.VoteAverage)
	checkField(t, int32(385), episode, tmdb.Episode.VoteCount)
}
