/*
 * Copyright (c) 2020.  This software is made for the Black Rock City Aid group and is provided AS IS with no support or liability under the Apache 2 license.
 */

package model

import (
	"errors"
	"os"
)

type ConfigData struct {
	TwilloAccountSid string
	TwilloToken      string
	DutyNumber       string
	TwilloNumber     string
}

var config ConfigData

func GetConfig() ConfigData {
	return config
}

//loads the config from the environment
func LoadConfig() error {
	config.TwilloToken = os.Getenv("TWILLO_TOKEN")
	if len(config.TwilloToken) == 0 {
		return errors.New("no twillo token")
	}
	config.TwilloAccountSid = os.Getenv("TWILLO_SID")
	if len(config.TwilloAccountSid) == 0 {
		return errors.New("no twillo sid")
	}
	config.DutyNumber = os.Getenv("DUTY_NUMBER")
	if len(config.TwilloAccountSid) == 0 {
		return errors.New("no duty number")
	}
	//todo, load this from env
	config.TwilloNumber = "12058318644"
	return nil
}
