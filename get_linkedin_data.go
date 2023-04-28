package main

import (
	"contact-sync/proxycurl"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"jmap"
	"log"
	"mime"
	"net/http"
	neturl "net/url"
	"os"
	"strings"
	"time"
)

const BearerToken = "3e479446-7c37-4aae-9dc6-73938ac9e4d2"
const Url = "https://api.fastmail.com/.well-known/jmap"
const Token = "fmu1-cb4140a6-594b144d15e51a4e0234278c02ce72f2-0-36ebf754ee909fbead0416d1ca22d2b3"

func getLinkedinUser(publicIdentifier string) (*proxycurl.Person, error) {
	url, err := neturl.Parse("https://nubela.co/proxycurl/api/v2/linkedin")
	if err != nil {
		return nil, err
	}

	q := url.Query()
	q.Set("url", "https://www.linkedin.com/in/"+publicIdentifier)
	url.RawQuery = q.Encode()

	req, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+BearerToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error getting linkedin user", publicIdentifier)
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	person := proxycurl.Person{}
	err = json.NewDecoder(resp.Body).Decode(&person)
	if err != nil {
		return nil, err
	}

	return &person, nil
}

func main() {
	f, err := os.Open("people.csv")
	if err != nil {
		panic(err)
	}
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		panic(err)
	}

	client, err := jmap.CreateClient(Url, Token)
	if err != nil {
		panic(err)
	}

	update := make(map[jmap.Id]jmap.PatchObject)

	for _, record := range records {
		//fname := record[0]
		//lname := record[1]
		id := jmap.Id(record[2])
		email := record[3]
		publicIdentifier := record[4]
		imageUrl := record[5]

		if publicIdentifier == "" {
			continue
		}

		err = os.MkdirAll("linkedin", 0750)
		if err != nil {
			panic(err)
		}

		filename := fmt.Sprintf("linkedin/%s.json", publicIdentifier)
		_, err := os.Stat(filename)
		if os.IsNotExist(err) {
			writeLinkedinUser(publicIdentifier, filename)
		}

		person := readLinkedinUser(filename)
		//if person.ProfilePicURL == "" && imageUrl == ""{
		//	continue
		//}
		if imageUrl == "" {
			imageUrl = person.ProfilePicURL
		}
		fmt.Println(person.FullName, imageUrl)
		if imageUrl == "" {
			continue
		}

		resp, err := http.Get(imageUrl)
		if err != nil {
			panic(err)
		}
		if resp.StatusCode != http.StatusOK {
			log.Println(resp.Status)
			continue
		}

		//fmt.Println("WE1")
		file, err := client.UploadFile("ucb4140a6", resp.Body)
		if err != nil {
			panic(err)
		}
		file.Type = resp.Header.Get("Content-Type")
		exts, err := mime.ExtensionsByType(file.Type)
		if err != nil {
			panic(err)
		}
		file.Name = person.FullName + exts[len(exts)-1]
		update[id] = map[string]interface{}{
			"avatar": file,
		}
		if !strings.HasSuffix(email, "amazon.com") {
			tp := time.Now()
			ip := 0
			for i, exp := range person.Experiences {
				if exp.EndsAt == nil {
					t, err := time.Parse("2006-1-2", fmt.Sprintf("%d-%d-%d", exp.StartsAt.Year, exp.StartsAt.Month, exp.StartsAt.Day))
					if err != nil {
						panic(err)
					}
					if t.Before(tp) {
						tp = t
						ip = i
					}
				}
			}
			update[id]["company"] = person.Experiences[ip].Company
			//update[id]["department"] =
			update[id]["jobTitle"] = person.Experiences[ip].Title
		}
	}
	client.SetContact(nil, nil, &update, nil)
	_, err = client.Do()
	if err != nil {
		panic(err)
	}

	//for _, r := range resp.MethodResponses {
	//	fmt.Printf("%+v\n", r)
	//}
}

func readLinkedinUser(filename string) *proxycurl.Person {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	person := proxycurl.Person{}
	err = json.NewDecoder(f).Decode(&person)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}

	return &person
}

func writeLinkedinUser(publicIdentifier string, filename string) {
	person, err1 := getLinkedinUser(publicIdentifier)
	if err1 != nil {
		panic(err1)
	}

	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	e := json.NewEncoder(f)
	e.SetIndent("", "\t")
	err1 = e.Encode(&person)

	if err != nil {
		panic(err)
	}
}
