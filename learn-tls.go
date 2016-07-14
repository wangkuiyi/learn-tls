// This example program comes from
// https://gist.github.com/denji/12b3a568f092ab951456, together with
// the way to generate a private key and a certFile for this HTTPS
// server:
//
// Generate private key (.key)
/*
   # Key considerations for algorithm "RSA" ≥ 2048-bit
   openssl genrsa -out server.key 2048

   # (Optional) Key considerations for algorithm "ECDSA" ≥ secp384r1
   # List ECDSA the supported curves (openssl ecparam -list_curves)
   openssl ecparam -genkey -name secp384r1 -out server.key
*/
// Generation of self-signed(x509) public key (PEM-encodings
// .pem|.crt) based on the private (.key)
/*
   openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
*/
package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "hello, world!\n")
	})
	err := http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
