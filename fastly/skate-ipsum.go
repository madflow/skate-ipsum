package main

import (
	"context"
	"fmt"
	"os"

	"github.com/fastly/compute-sdk-go/fsthttp"

	"github.com/madflow/skate-ipsum/generator"
)

// The entry point for your application.
//
// Use this function to define your main request handling logic. It could be
// used to route based on the request properties (such as method or path), send
// the request to a backend, make completely new requests, and/or generate
// synthetic responses.

func main() {
	// Log service version
	fmt.Println("FASTLY_SERVICE_VERSION:", os.Getenv("FASTLY_SERVICE_VERSION"))

	fsthttp.ServeFunc(func(ctx context.Context, w fsthttp.ResponseWriter, r *fsthttp.Request) {
		// Filter requests that have unexpected methods.
		if r.Method == "POST" || r.Method == "PUT" || r.Method == "PATCH" || r.Method == "DELETE" {
			w.WriteHeader(fsthttp.StatusMethodNotAllowed)
			fmt.Fprintf(w, "This method is not allowed\n")
			return
		}

		// If request is to the `/` path...
		if r.URL.Path == "/" {

			paragraphs := generator.IpsumJson(10)
			// Below are some common patterns for Compute@Edge services using TinyGo.
			// Head to https://developer.fastly.com/learning/compute/go/ to discover more.

			// Create a new request.
			// req, err := fsthttp.NewRequest("GET", "https://example.com", nil)
			// if err != nil {
			//   // Handle Error
			// }

			// Add request headers.
			// req.Header.Set("Custom-Header", "Welcome to Compute@Edge!")
			// req.Header.Set(
			//   "Another-Custom-Header",
			//   "Recommended reading: https://developer.fastly.com/learning/compute"
			// )

			// Override cache TTL.
			// req.CacheOptions.TTL = 60

			// Forward the request to a backend named "TheOrigin".
			// resp, err := req.Send(ctx, "TheOrigin")
			// if err != nil {
			//	 w.WriteHeader(fsthttp.StatusBadGateway)
			//	 fmt.Fprintln(w, err)
			//	 return
			// }

			// Remove response headers.
			// resp.Header.Del("Yet-Another-Custom-Header")

			// Copy all headers from the response.
			// w.Header().Reset(resp.Header.Clone())

			// Log to a Fastly endpoint.
			// NOTE: You will need to import "github.com/fastly/compute-sdk-go/rtlog"
			// for this to work
			// endpoint := rtlog.Open("my_endpoint")
			// fmt.Fprintln(endpoint, "Hello from the edge!")

			// Send a default synthetic response.
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(fsthttp.StatusOK)
			w.Write([]byte(paragraphs))
			fmt.Fprintln(w)
			return
		}

		// Catch all other requests and return a 404.
		w.WriteHeader(fsthttp.StatusNotFound)
	})
}
