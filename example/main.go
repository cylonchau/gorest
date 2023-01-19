package main

import (
	"context"

	rest "github.io/cylonchau/restclient"
)

func main() {

	client := rest.NewRESTClient("http://10.0.0.3:5555/v2/services/haproxy/configuration/version")
	req := rest.NewDefaultRequest(client).BasicAuth(
		"admin", "1fc917c7ad66487470e466c0ad40ddd45b9f7730a4b43e1b2542627f0596bbdc",
	).AddHeader("Content-Type", "application/json")
	req.Verb("GET").Do(context.TODO())
}
