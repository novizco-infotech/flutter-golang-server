package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func PrintRequest(r *http.Request, bodyBytes []byte) {
	fmt.Println("****  Request Pretty Printing Starts ****")
	log.Println("****  Request Pretty Printing Starts ****")

	var err error
	fmt.Printf("Headers: %+v\n", r.Header)

	if len(bodyBytes) > 0 {
		var prettyJSON bytes.Buffer
		if err = json.Indent(&prettyJSON, bodyBytes, "", "\t"); err != nil {
			fmt.Printf("JSON parse error: %v", err)
			return
		}
		//fmt.Println(string(prettyJSON.Bytes()))
		fmt.Println(prettyJSON.String())
		log.Println(prettyJSON.String())

	} else {
		fmt.Printf("Body: No Body Supplied\n")
	}

	fmt.Println("****  Request Pretty Printing Ends !!!****")

}
