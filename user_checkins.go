package untappd

import (
	"math"
	"net/http"
	"net/url"
	"strconv"
)

// Checkins queries for information about a User's checkins.
// The username parameter specifies the User whose checkins will be
// returned.
//
// This method returns up to 25 of the User's most recent checkins.
// For more granular control, and to page through the checkins list using ID
// parameters, use CheckinsMinMaxIDLimit instead.
func (u *UserService) Checkins(username string) ([]*Checkin, *http.Response, error) {
	// Use default parameters as specified by API.  Max ID is somewhat
	// arbitrary, but should provide plenty of headroom, just in case.
	return u.CheckinsMinMaxIDLimit(username, 0, math.MaxInt32, 25)
}

// CheckinsMinMaxIDLimit queries for information about a User's checkins,
// but also accepts minimum checkin ID, maximum checkin ID, and a limit
// parameter to enable paging through checkins. The username parameter
// specifies the User whose checkins will be returned.
//
// 50 checkins is the maximum number of checkins which may be returned by
// one call.
func (u *UserService) CheckinsMinMaxIDLimit(username string, minID int, maxID int, limit int) ([]*Checkin, *http.Response, error) {
	q := url.Values{
		"min_id": []string{strconv.Itoa(minID)},
		"max_id": []string{strconv.Itoa(maxID)},
		"limit":  []string{strconv.Itoa(limit)},
	}

	// Temporary struct to unmarshal checkin JSON
	var v struct {
		Response struct {
			Checkins struct {
				Count int           `json:"count"`
				Items []*rawCheckin `json:"items"`
			} `json:"checkins"`
		} `json:"response"`
	}

	// Perform request for user checkins by username
	res, err := u.client.request("GET", "user/checkins/"+username, q, &v)
	if err != nil {
		return nil, res, err
	}

	// Build result slice from struct
	checkins := make([]*Checkin, v.Response.Checkins.Count)
	for i := range v.Response.Checkins.Items {
		checkins[i] = v.Response.Checkins.Items[i].export()
	}

	return checkins, res, nil
}
