package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON!")

	EncodeJSON()
	DecodeJSON()
}

func EncodeJSON() {
	myCourses := []course{
		{"Android Bootcamp", 299, "ankiitdev.in", "abc123", []string{"android-dev", "kotlin"}},
		{"MEAN Bootcamp", 399, "ankiitdev.in", "bcd123", []string{"full-stack", "js"}},
		{"MERN Bootcamp", 499, "ankiitdev.in", "gfg123", nil},
	}

	// finalJson, err := json.Marshal(lcoCourses)
	finalJson, err := json.MarshalIndent(myCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func DecodeJSON() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "Android Bootcamp",
        "price": 299,
        "website": "ankiitdev.in",
        "tags": ["android-dev","kotlin"]
	}
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON is valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("Invalid JSON")
	}

	// some cases where you just want to add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and Type is %T\n", k, v, v)
	}

}
