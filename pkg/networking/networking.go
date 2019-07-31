package networking

import (
	//"encoding/json"
	//"fmt"
	//"io/ioutil"
	//"net/http"
	"github.com/skycoin/skycoin/src/api"
)

//func GetDefaultConnections(NodeAddress string)  {
//	endpoint := NodeAddress + "/api/v1/network/connections"
//	req, err := http.Get(endpoint)
//	if err != nil {
//
//	}
//
//	defer req.Body.Close()
//
//	body, err := ioutil.ReadAll(req.Body)
//
//	if err != nil {
//		fmt.Println("Uff, caboom!")
//	} else {
//		p := json.Unmarshal(body)
//		fmt.Println(string(body))
//		fmt.Println("Todo esta rico!")
//	}
//
//}


func init(){
	
}

//-----------------TODO
func nodeAddress() string{
	return""
}

//-----------------TODO
func nodeUsername() string{
	return ""
}

//----------------TODO
func nodePassord() string{
	return ""
}

func newClient() *api.Client{
	c := api.NewClient(nodeAddress())
	c.SetAuth(nodeUsername(), nodePassord())
	return c
}