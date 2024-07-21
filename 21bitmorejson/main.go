package main

import (
	"encoding/json"
	"fmt"
)

//course is a data structure 
type course struct {
	Name string `json:"coursename"` //name of the course
	Price int //price of the course
	Platform string `json:"website"` //platform where course is hosted
	Password string `json:"-"` //don't show this field
	Tags []string `json:"tags,omitempty"` //if empty don't show
}

func main()  {
	fmt.Println("Welcome to bit more json");
	EncodeJson() //encode golang data to json
	DecodeJson() //decode json to golang data
}

//EncodeJson is to encode golang data to json
func EncodeJson()  {
	//create a slice of courses
	locCources := []course{
		{"ReactJS Bootcamp", 299, "learncodeonline.in", "abc123", []string{"web-dev", "js"}},
		{"Angular Bootcamp", 199, "learncodeonline.in", "abc12233", []string{"full-stack", "js"}},
		{"VueJS Bootcamp", 99, "learncodeonline.in", "abcsd123", nil},
	}

	//package this data as json data
	//MarshalIndent - to format the json with indentation
	finalJson, err := json.MarshalIndent(locCources, "", "\t")
	if err!= nil {
        panic(err)
    }
	fmt.Printf("%s\n",finalJson)
}


//DecodeJson is to decode json to golang data
func DecodeJson()  {
	//create a json data from web
	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJS Bootcamp",
		"Price": 299,
		"website": "learncodeonline.in",
		"tags": ["web-dev", "js"]
	}
	`)

	//create a variable of type course
	var result course

	//check if json is valid
	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("Json is valid")
		//unmarshal the json data to the result variable
		json.Unmarshal(jsonDataFromWeb, &result)
		fmt.Printf("%#v\n", result)
	} else {
		fmt.Println("Json is not valid")
	}

	//some cases where you just want to add data to key value pair 

    // instead of creating a struct. In such cases you can use the map
    var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	// traverse through the map
	for k,v := range myOnlineData {
		fmt.Printf("key is %v and value is %v is of type: %T\n", k, v, v)
	}
}




