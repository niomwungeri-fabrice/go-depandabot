package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// function which return "geeks"
func ReturnGeeks() string {
	return "geeks"
}

const (
	apiURL      = "https://bookbeta.com/api/v1"
	releasesURL = apiURL + "/releases"
)

// Let's define a single error for the purposes of this example
var ErrFailedAPICall = errors.New("bad response from BookBeta API")

type Release struct {
	ID          int64  `json:"id"`
	BookName    string `json:"bookName"`
	AuthorName  string `json:"authorName"`
	IsAvailable bool   `json:"isAvailable"`
}

func GetAvailableReleases() ([]Release, error) {
	res, err := http.Get(releasesURL)
	if err != nil {
		return nil, fmt.Errorf("failed to get releases:%s %v", err, ErrFailedAPICall)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d - %s %v", res.StatusCode, res.Status, ErrFailedAPICall)
	}
	var releases []Release
	if err := json.NewDecoder(res.Body).Decode(&releases); err != nil {
		return nil, fmt.Errorf("failed to decode body into release slice: %w", err)
	}
	var availableReleases []Release
	for _, r := range releases {
		if r.IsAvailable {
			availableReleases = append(availableReleases, r)
		}
	}
	return availableReleases, nil
}

func main() {
	ReturnGeeks()
}
