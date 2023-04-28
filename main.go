package main

import (
	"encoding/json"
	"fmt"
	"jmap"
	"os"
)

const (
	URL   = "https://api.fastmail.com/.well-known/jmap"
	TOKEN = "fmb1-cb4140a6-c8a9bb1770306f2085bb1af7c0b70742-1657609200-dbf05362d5b594cca52d35d86b344cc4"
)

func main() {

	client, err := jmap.CreateClient(URL, TOKEN)

	if err != nil {
		panic(err)
	}

	client.GetContact()
	resp, err := client.Do()
	if err != nil {
		panic(err)
	}

	//api.GetProfileData()

	for _, response := range resp.MethodResponses {

		b, err := json.Marshal(response.Arguments.(map[string]interface{})["list"])
		if err != nil {
			panic(err)
		}
		list := make([]jmap.Contact, 0)
		err = json.Unmarshal(b, &list)

		if err != nil {
			panic(err)
		}

		//f1, err := os.Create("people.csv")
		//if err != nil {
		//	panic(err)
		//}
		//w := csv.NewWriter(f1)
		for _, l := range list {
			//fmt.Printf("%+v\n", l)
			if l.Avatar != nil {
				//fmt.Printf("%+v\n", l.Avatar)
				name := l.Avatar.Name
				if name == "" {
					name = l.FirstName + l.LastName
				}
				fmt.Println(name)
				err = client.DownloadFile(l.Avatar.BlobId, name, l.Avatar.Type)
				if err != nil {
					panic(err)
				}
			}
			//email := ""
			//if len(l.Emails) > 0 {
			//	email = l.Emails[0].Value
			//}
			//if err := w.Write([]string{l.FirstName, l.LastName, string(l.Id), email}); err != nil {
			//	log.Fatalln("error writing record to csv:", err)
			//}
		}
		//w.Flush()

		//if err := w.Error(); err != nil {
		//	log.Fatal(err)
		//}
		//err = f1.Close()
		//if err != nil {
		//	panic(err)
		//}

		f, err := os.Create("people.json")
		if err != nil {
			panic(err)
		}
		encoder := json.NewEncoder(f)
		encoder.SetIndent("", "\t")
		err = encoder.Encode(&list)
		if err != nil {
			panic(err)
		}
		err = f.Close()
		if err != nil {
			panic(err)
		}

	}

}
