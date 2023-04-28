package main

import (
	"contact-sync/proxycurl"
	"encoding/json"
	"net/http"
	neturl "net/url"
	"os"
)

// LinkedinUrl const BearerToken = "3e479446-7c37-4aae-9dc6-73938ac9e4d2"
const LinkedinUrl = "https://www.linkedin.com/in/rkkochie/"

func main() {

	url, err := neturl.Parse("https://nubela.co/proxycurl/api/v2/linkedin")
	if err != nil {
		panic(err)
	}

	q := url.Query()
	q.Set("url", LinkedinUrl)

	url.RawQuery = q.Encode()

	req, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Authorization", "Bearer "+BearerToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode != http.StatusOK {
		panic(resp.Status)
	}

	person := proxycurl.Person{}

	err = json.NewDecoder(resp.Body).Decode(&person)
	if err != nil {
		panic(err)
	}

	f, err := os.Create("person.json")
	if err != nil {
		panic(err)
	}

	enc := json.NewEncoder(f)
	enc.SetIndent("", "\t")
	err = enc.Encode(&person)
	if err != nil {
		panic(err)
	}
}
