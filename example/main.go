package main

import (
	"context"

	"github.io/cylonchau/kube-haproxy/rest"
)

func main() {

	req := rest.FastRequest(
		"http://10.0.0.3:5555/v2/services/haproxy/configuration/version",
	).BasicAuth(
		"admin", "1fc917c7ad66487470e466c0ad40ddd45b9f7730a4b43e1b2542627f0596bbdc",
	).AddHeader("Content-Type", "application/json")
	req.Get().Do(context.TODO())

	// with client
	client := rest.NewDefaultRESTClient()
	req = rest.New
}
