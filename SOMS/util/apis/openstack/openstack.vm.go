// Openstack Compute API Interface
// This file is generated by code generator. DO NOT EDIT!
package openstack_api
import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	
)

func createNetwork(authToken, endpoint, networkName string) {
    url := endpoint + "/v2.0/networks"
    payload := map[string]interface{}{
        "network": map[string]interface{}{
            "name": networkName,
        },
    }
    jsonPayload, _ := json.Marshal(payload)
    
    req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
    req.Header.Set("X-Auth-Token", authToken)
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("Response:", string(body))
}
func listNetworks(authToken, endpoint string) {
    url := endpoint + "/v2.0/networks"

    req, _ := http.NewRequest("GET", url, nil)
    req.Header.Set("X-Auth-Token", authToken)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("Response:", string(body))
}



func CreateNetwork(userID string, password string, newnetworkName string, networkEndpoint string) {
    authToken,err := GetUserToken(userID,password)
    if err != nil {
        panic(err)
    }
    endpoint := networkEndpoint
    networkName := "test-network"

	createNetwork(authToken, endpoint, networkName)
}
func ListNetworks(userID string, password string, networkEndpoint string){
    authToken,err := GetUserToken(userID,password)
    if err != nil {
        panic(err)
    }
    endpoint := networkEndpoint
    listNetworks(authToken, endpoint)
}
