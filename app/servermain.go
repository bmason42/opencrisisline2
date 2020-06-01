/*
 * Copyright (c) 2020.  This software is made for the Black Rock City Aid group and is provided AS IS with no support or liability under the Apache 2 license.
 */

package main

import (
	"github.com/bmason42/opencrisisline2/pkg/apiimpl"
	"github.com/bmason42/opencrisisline2/pkg/model"
)

func main() {
	model.LoadConfig()
	apiimpl.RunServer()
}
