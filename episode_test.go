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
	if credits, err := episode.Credits(); err != nil {
		t.Fatalf("failed to get credits: %v", err)
	} else {
		if cast, err := credits.Cast(); err != nil {
			t.Fatalf("failed to get cast from credits: %v", err)
		} else if len(cast) == 0 {
			t.Fatal("expected cast, got none")
		} else if actor, err := findCredit(cast, "Peter Dinklage"); err != nil {
			t.Fatalf("failed to find actor: %v", err)
		} else {
			checkField(t, false, actor, tmdb.Credit.Adult)
			checkField(t, int32(2), actor, tmdb.Credit.Gender)
			checkField(t, int32(22970), actor, tmdb.Credit.ID)
			checkField(t, "Acting", actor, tmdb.Credit.KnownForDepartment)
			checkField(t, "Peter Dinklage", actor, tmdb.Credit.Name)
			checkField(t, "Peter Dinklage", actor, tmdb.Credit.OriginalName)
			checkField(t, 1.9543, actor, tmdb.Credit.Popularity)
			checkField(t, "/9CAd7wr8QZyIN0E7nm8v1B6WkGn.jpg", actor, tmdb.Credit.ProfilePath)
			checkField(t, "Tyrion 'The Halfman' Lannister", actor, tmdb.Credit.Character)
			checkField(t, "5256c8b219c2956ff6047cd8", actor, tmdb.Credit.CreditID)
			checkField(t, int32(0), actor, tmdb.Credit.Order)

		}
		if crew, err := credits.Crew(); err != nil {
			t.Fatalf("failed to get crew from credits: %v", err)
		} else if len(crew) == 0 {
			t.Fatal("expected crew, got none")
		} else if director, err := findCredit(crew, "Tim Van Patten"); err != nil {
			t.Fatalf("failed to find director: %v", err)
		} else {
			checkField(t, "Directing", director, tmdb.Credit.Department)
			checkField(t, "Director", director, tmdb.Credit.Job)
			checkField(t, "5256c8a219c2956ff6046e77", director, tmdb.Credit.CreditID)
			checkField(t, false, director, tmdb.Credit.Adult)
			checkField(t, int32(2), director, tmdb.Credit.Gender)
			checkField(t, int32(44797), director, tmdb.Credit.ID)
			checkField(t, "Directing", director, tmdb.Credit.KnownForDepartment)
			checkField(t, "Tim Van Patten", director, tmdb.Credit.Name)
			checkField(t, "Tim Van Patten", director, tmdb.Credit.OriginalName)
			checkField(t, 0.7112, director, tmdb.Credit.Popularity)
			checkField(t, "/vwcARZBg4PEzOwnPsXdjRWeUVrZ.jpg", director, tmdb.Credit.ProfilePath)
		}
		if guestStars, err := credits.GuestStars(); err != nil {
			t.Fatalf("failed to get guest stars from credits: %v", err)
		} else if len(guestStars) == 0 {
			t.Fatal("expected guest stars, got none")
		} else if hodor, err := findCredit(guestStars, "Kristian Nairn"); err != nil {
			t.Fatalf("failed to find guest star: %v", err)
		} else {
			checkField(t, "Hodor", hodor, tmdb.Credit.Character)
			checkField(t, "5256c8be19c2956ff6048446", hodor, tmdb.Credit.CreditID)
			checkField(t, int32(81), hodor, tmdb.Credit.Order)
			checkField(t, false, hodor, tmdb.Credit.Adult)
			checkField(t, int32(2), hodor, tmdb.Credit.Gender)
			checkField(t, int32(1223792), hodor, tmdb.Credit.ID)
			checkField(t, "Acting", hodor, tmdb.Credit.KnownForDepartment)
			checkField(t, "Kristian Nairn", hodor, tmdb.Credit.Name)
			checkField(t, "Kristian Nairn", hodor, tmdb.Credit.OriginalName)
			checkField(t, 0.4995, hodor, tmdb.Credit.Popularity)
			checkField(t, "/dlbq6cCW0xdpFY15q6flP6lDXWV.jpg", hodor, tmdb.Credit.ProfilePath)
		}
	}
	checkField(t, "tt1480055", episode, tmdb.Episode.ExternalIDs, tmdb.ExternalIDs.IMDBID)
	checkField(t, "/m/0gmc6ph", episode, tmdb.Episode.ExternalIDs, tmdb.ExternalIDs.FreebaseMID)
	checkField(t, "/en/winter_is_coming", episode, tmdb.Episode.ExternalIDs, tmdb.ExternalIDs.FreebaseID)
	checkField(t, int32(3254641), episode, tmdb.Episode.ExternalIDs, tmdb.ExternalIDs.TVDBID)
	checkField(t, int32(1065008299), episode, tmdb.Episode.ExternalIDs, tmdb.ExternalIDs.TVRageID)
	checkField(t, "Q2614622", episode, tmdb.Episode.ExternalIDs, tmdb.ExternalIDs.WikidataID)
}
