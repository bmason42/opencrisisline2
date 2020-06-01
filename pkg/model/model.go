/*
 * Copyright (c) 2020.  This software is made for the Black Rock City Aid group and is provided AS IS with no support or liability under the Apache 2 license.
 */

package model

import (
	"github.com/bmason42/opencrisisline2/pkg/errors"
	v1 "github.com/bmason42/opencrisisline2/pkg/generated/v1"
	"math/rand"
	"time"
)

type SupportRequest struct {
	RequestID string
	AuthPin   string
	Data      v1.HelpRequest
}
type PersistenceLayer interface {
	SaveSupportRequest(req *SupportRequest) error
	FetchSupportRequestByPhone(phone string) (*SupportRequest, error)
}

var persistence PersistenceLayer

func InitModel() error {
	rand.Seed(time.Now().Unix())
	err := LoadConfig()
	if err != nil {
		return err
	}
	persistence = NewInmemoryLayer()
	return nil
}

func GetPersistenceLayer() (PersistenceLayer, error) {
	var err error
	if persistence == nil {
		err = errors.NewInernalError(errors.OCERROR_ERROR, errors.ERROR_CODE_SYSTEM_NOT_INITIALIZED, nil)
	} else {
		err = nil
	}
	return persistence, err
}
