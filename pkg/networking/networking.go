package networking

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetDefaultConnections(NodeAddress string)  {
	endpoint := NodeAddress + "/api/v1/network/connections"
	req, err := http.Get(endpoint)
	if err != nil {

	}

	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		fmt.Println("Uff, caboom!")
	} else {
		p := json.Unmarshal(body)
		fmt.Println(string(body))
		fmt.Println("Todo esta rico!")
	}

}