/*
 * Copyright (c) 2020.  This software is made for the Black Rock City Aid group and is provided AS IS with no support or liability under the Apache 2 license.
 */

package model

import (
	"os"
	"testing"
)

func TestTwllioConfig_SendText(t *testing.T) {
	err := LoadConfig()
	if err != nil {
		panic("No Config")
	}
	type fields struct {
		TwilloAccountSid string
		TwilloToken      string
	}
	type args struct {
		from string
		msg  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "happy test", fields: fields{TwilloToken: os.Getenv("TWILLO_TOKEN"), TwilloAccountSid: os.Getenv("TWILLO_SID")}, args: args{from: "12058318644", msg: "test"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &TwllioConfig{
				TwilloAccountSid: tt.fields.TwilloAccountSid,
				TwilloToken:      tt.fields.TwilloToken,
			}
			if err := tr.SendText(tt.args.from, "name", tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("TwllioConfig.SendText() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
