/*
 * Copyright (c) 2020.  This software is made for the Black Rock City Aid group and is provided AS IS with no support or liability under the Apache 2 license.
 */

package model

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	url2 "net/url"

	"net/http"
	"strings"
)

type MessageSystem interface {
	SendText(from, name, msg string) error
}

type TwllioConfig struct {
	TwilloAccountSid string
	TwilloToken      string
}

func NewMessageSystem() MessageSystem {
	t := TwllioConfig{TwilloAccountSid: Config.TwilloAccountSid, TwilloToken: Config.TwilloToken}
	return &t
}
func (t *TwllioConfig) SendText(from string, name, msg string) error {
	msgData := url2.Values{}
	msgData.Add("To", Config.DutyNumber)
	msgData.Add("From", "12058318644")
	//msgData.Add("From","17653352431")

	bodyMsg := fmt.Sprintf("%s %s %s", from, name, msg)
	msgData.Add("Body", bodyMsg)
	rawMsg := msgData.Encode()
	msgDataReader := strings.NewReader(rawMsg)
	client := &http.Client{}
	url := fmt.Sprintf("https://api.twilio.com/2010-04-01/Accounts/%s/Messages.json", Config.TwilloAccountSid)
	//url:=fmt.Sprintf("http://localhost:8080/opencrisisline2/v1/support-request")

	req, err := http.NewRequest("POST", url, msgDataReader)
	req.SetBasicAuth(Config.TwilloAccountSid, Config.TwilloToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return err
	}
	if resp.StatusCode != 201 {
		bits, err2 := ioutil.ReadAll(resp.Body)
		var err error
		if err2 == nil {
			errorMsg := string(bits)
			err = errors.New("error from twillo " + errorMsg)
			log.Infof("error msg: %s", errorMsg)
		} else {
			err = errors.New("error from twillo " + resp.Status)
		}
		return err
	}
	return nil
}
