// Copyright 2018 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"strconv"

	"google.golang.org/appengine"
)

func main() {
	http.HandleFunc("/", handle)
	appengine.Main()
}

type risk_info_sub struct{
	Safety string
	Reliability  string
	CustomerData string
	Financial string
}

func handle(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t  risk_info_sub
	err := decoder.Decode(&t)
	if err != nil{
		panic(err)
	}

	m := map[string]int{}

	m["safety"],_ = strconv.Atoi(t.Safety);
	m["reliability"],_ = strconv.Atoi(t.Reliability);
	m["customerData"],_ = strconv.Atoi(t.CustomerData);
	m["financial"],_ = strconv.Atoi(t.Financial);

	fmt.Fprintln(w, t)
	for key,value := range(m){
		fmt.Fprintln(w, "Key:", key, "Value:" , value)
	}

	fmt.Fprintln(w, "Isn't it a wonderful world?")
}
