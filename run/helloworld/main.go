// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// [START cloudrun_helloworld_service]
// [START run_helloworld_service]

// Sample run-helloworld is a minimal Cloud Run service.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"context"
	"io"
//	"google.golang.org/api/idtoken"

)

func main() {
	log.Print("starting server...")
	http.HandleFunc("/", handler)

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	// Start HTTP server.
	log.Printf("listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	if name == "" {
		name = "World2"
	}
	
	iapHeader := r.Header.Get("X-Goog-IAP-JWT-Assertion")

	fmt.Fprintf(w, "Hello %s!\n", name)
	fmt.Fprintf(w, "Header = %s", iapHeader)
}


// validateJWTFromComputeEngine validates a JWT found in the
// "x-goog-iap-jwt-assertion" header.
// func validateJWTFromComputeEngine(w io.Writer, iapJWT, projectNumber, backendServiceID string) error {
// 	// iapJWT := "YmFzZQ==.ZW5jb2RlZA==.and0" // req.Header.Get("X-Goog-IAP-JWT-Assertion")
// 	/projectNumber := "123456789"
// 	// backendServiceID := "backend-service-id"
// 	ctx := context.Background()
// 	aud := fmt.Sprintf("/projects/%s/global/backendServices/%s", projectNumber, backendServiceID)

// 	payload, err := idtoken.Validate(ctx, iapJWT, aud)
// 	if err != nil {
// 			return fmt.Errorf("idtoken.Validate: %v", err)
// 	}

// 	// payload contains the JWT claims for further inspection or validation
// 	fmt.Fprintf(w, "payload: %v", payload)

// 	return nil
// }
// [END run_helloworld_service]
// [END cloudrun_helloworld_service]
