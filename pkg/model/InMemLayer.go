/*
 * Copyright (c) 2020.  This software is made for the Black Rock City Aid group and is provided AS IS with no support or liability under the Apache 2 license.
 */

package model

type InMemoryPersistence struct {
	IDToRequest    map[string]*SupportRequest
	PhoneToRequest map[string]*SupportRequest
}

func NewInmemoryLayer() PersistenceLayer {
	var x InMemoryPersistence
	x.IDToRequest = make(map[string]*SupportRequest, 0)
	x.PhoneToRequest = make(map[string]*SupportRequest, 0)
	return &x
}
func (t *InMemoryPersistence) SaveSupportRequest(req *SupportRequest) error {
	t.IDToRequest[req.RequestID] = req
	return nil
}
func (t *InMemoryPersistence) FetchSupportRequestByPhone(phone string) (*SupportRequest, error) {
	ret := t.PhoneToRequest[phone]
	return ret, nil
}
