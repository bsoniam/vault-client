package client

import (
	"testing"
	"fmt"
)

func createClient(t *testing.T) Client{
	var token string = "6aeec159-592b-a807-ac54-3060f204116c"
	var path string = "127.0.0.1:8200"
	client, err := NewClient(token,path)
	if err != nil{
		t.Errorf("%s: I have %s and I should get %s", t.Name(), err, nil)
	}
	return client
}

func TestClient_CreatePolicy(t *testing.T){
	client := createClient(t)
	err := client.CreatePolicy("transit/keys/*","createkey", "createkey")
	if err != nil{
		t.Errorf("%s: I have %s and I should get %s", t.Name(), err, nil)
	}
}

func TestClient_CreateToken(t *testing.T) {
	client := createClient(t)
	token, err := client.CreateToken("createkey")
	if err != nil{
		t.Errorf("%s: I have %s and I should get %s", t.Name(), err, nil)
	}
	fmt.Println(token)

}

func TestClient_Read(t *testing.T) {
	client := createClient(t)
	secret, err := client.Read("transit/keys/key10","6aeec159-592b-a807-ac54-3060f204116c" )
	if err != nil{
		t.Errorf("%s: I have %s and I should get %s", t.Name(), err, nil)
	}
	fmt.Println(secret.Data)
}

func TestClient_Write(t *testing.T) {
	client := createClient(t)
	secret, err := client.Write("transit/keys/key10",map[string]interface{}{
		"type": "aes256-gcm96",
		"derived": false,
		"exportable": true},"6aeec159-592b-a807-ac54-3060f204116c" )
	if err != nil{
		t.Errorf("%s: I have %s and I should get %s", t.Name(), err, nil)
	}
	fmt.Println(secret.Data)
}