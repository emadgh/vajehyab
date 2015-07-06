// vajeh
package vajehyab

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type VajehYab struct {
	Developer string
	Improve   bool
	Ads       bool
}
type Result struct {
	Search Search `json:"search"`
	Data   Data   `json:"data"`
	Error  Error  `json:"error"`
	Ads    Ads    `json:"ads"`
}

type Search struct {
	Q    string `json:"q"`
	Code int    `json:"code"`
}

type Data struct {
	Title         string `json:"title"`
	Pronunciation string `json:"pronunciation"`
	Text          string `json:"text"`
	Source        string `json:"source"`
	Permalink     string `json:"permalink"`
}

type Error struct {
	Message string `json:"message"`
	Reason  string `json:"reason"`
}
type Ads struct {
	Text string `json:"text"`
	Url  string `json:"url"`
}

func (vy *VajehYab) Search(word string) (*Result, error) {
	v := url.Values{}
	v.Add("q", word)
	v.Add("developer", vy.Developer)
	v.Add("improve", btos(vy.Improve))
	v.Add("ads", btos(vy.Ads))
	query_url := v.Encode()
	res, err := http.Get("http://api.vajehyab.com/v2/public/?" + query_url)
	if err != nil {
		return nil, err
	}

	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, err
	}

	vajeh := Result{}
	err = json.Unmarshal(robots[4:], &vajeh)
	if err != nil {
		return nil, err
	}

	return &vajeh, nil
}
func btos(b bool) string {
	r := "0"
	if b {
		r = "1"
	}
	return r
}
