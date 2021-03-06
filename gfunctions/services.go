// Copyright 2018 Google LLC. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

// [START functions_helloworld_get]

// Package helloworld provides a set of Cloud Function samples.
package helloworld

import (
        "log"
        "fmt"
        "net/http"
        "encoding/json"
)

type Service map[string]interface{}

// HelloGet is an HTTP Cloud Function.
// CORSEnabledFunction is an example of setting CORS headers.
// For more information about CORS and CORS preflight requests, see
// https://developer.mozilla.org/en-US/docs/Glossary/Preflight_request.
// https://cloud.google.com/functions/docs/writing/http#writing_http_helloworld-go
func ServicesGet(w http.ResponseWriter, r *http.Request) {
  // Set CORS headers for the preflight request
  if r.Method == http.MethodOptions {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
    w.Header().Set("Access-Control-Max-Age", "3600")
    w.WriteHeader(http.StatusNoContent)
      return
  }

  var myMapSlice []Service
  //https://golang.org/pkg/encoding/json/
  myMapSlice = append(myMapSlice,
    Service{"id": 1, "Name": "Wedding"},
    Service{"id": 2, "Name": "Birthday"},
   	Service{"id": 3, "Name": "Conference"})

  b, err := json.Marshal(myMapSlice)
  if err != nil {
		log.Fatal(err)
  }
  w.Header().Set("Access-Control-Allow-Origin", "*")
  fmt.Fprint(w, string(b))
}

// [END functions_helloworld_get]
