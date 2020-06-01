/*
 * Copyright (c) 2020.  This software is made for the Black Rock City Aid group and is provided AS IS with no support or liability under the Apache 2 license.
 */

package model

import (
	"crypto/tls"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	url2 "net/url"

	"net/http"
	"strings"
)

type MessageSystem interface {
	AsyncSendText(toNumber string, msg string)
	SendText(toNumber string, msg string) error
}

type TwllioConfig struct {
	TwilloAccountSid string
	TwilloToken      string
}

func NewMessageSystem() MessageSystem {
	t := TwllioConfig{TwilloAccountSid: config.TwilloAccountSid, TwilloToken: config.TwilloToken}
	return &t
}
func (t *TwllioConfig) AsyncSendText(toNumber string, msg string) {
	err := t.SendText(toNumber, msg)
	if err != nil {
		log.Error("Unable to send text")
	}
}
func (t *TwllioConfig) SendText(toNumber string, msg string) error {

	msgData := url2.Values{}
	msgData.Add("To", toNumber)
	msgData.Add("From", GetConfig().TwilloNumber)

	msgData.Add("Body", msg)
	rawMsg := msgData.Encode()
	msgDataReader := strings.NewReader(rawMsg)
	url := fmt.Sprintf("https://api.twilio.com/2010-04-01/Accounts/%s/Messages.json", config.TwilloAccountSid)
	//url:=fmt.Sprintf("http://localhost:8080/opencrisisline2/v1/support-request")
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	req, err := http.NewRequest("POST", url, msgDataReader)
	req.SetBasicAuth(config.TwilloAccountSid, config.TwilloToken)
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
