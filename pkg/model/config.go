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
}

var Config ConfigData

//loads the config from the environment
func LoadConfig() error {
	Config.TwilloToken = os.Getenv("TWILLO_TOKEN")
	if len(Config.TwilloToken) == 0 {
		return errors.New("no twillo token")
	}
	Config.TwilloAccountSid = os.Getenv("TWILLO_SID")
	if len(Config.TwilloAccountSid) == 0 {
		return errors.New("no twillo sid")
	}
	Config.DutyNumber = os.Getenv("DUTY_NUMBER")
	if len(Config.TwilloAccountSid) == 0 {
		return errors.New("no duty number")
	}

	return nil
}
