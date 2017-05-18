package startstop

import (
	"github.com/hna/speedycloud"
	"bytes"
)

func actionURL(client *speedycloud.ServiceClient, id string, action string) string {
	return client.ServiceURL("cloud_servers", id, action)
}

// Start is the operation responsible for starting a Compute server.
func Start(client *speedycloud.ServiceClient, id string) speedycloud.ErrResult {
	var res speedycloud.ErrResult
	_, res.Err = client.Post(actionURL(client, id, "start"),
		bytes.NewBufferString(""),
		&res.Body,
		nil)
	return res
}

// Stop is the operation responsible for stopping a Compute server.
func Stop(client *speedycloud.ServiceClient, id string) speedycloud.ErrResult {
	var res speedycloud.ErrResult
	_, res.Err = client.Post(actionURL(client, id, "stop"),
        bytes.NewBufferString(""),
        &res.Body,
        nil)
	return res
}
