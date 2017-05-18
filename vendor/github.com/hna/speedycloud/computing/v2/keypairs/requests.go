package keypairs

import (
    "github.com/hna/speedycloud"
    "bytes"
    "fmt"
    "net/url"
)


// CreateOpts specifies keypair creation or import parameters.
type CreateOpts struct {
	// Name [required] is a friendly name to refer to this KeyPair in other services.
	DisplayName string

	// PublicKey [optional] is a pregenerated OpenSSH-formatted public key. If provided, this key
	// will be imported and no new key will be created.
	PublicKey string
}

// ToKeyPairCreateMap constructs a request body from CreateOpts.
func (opts CreateOpts) ToKeyPairCreateUrlEncode() (*bytes.Buffer, error) {

    server := url.Values{}

    server.Set("display_name", opts.DisplayName)
    server.Set("public_key", opts.DisplayName)

    return  bytes.NewBufferString(server.Encode()), nil
}

// Create requests the creation of a new keypair on the server, or to import a pre-existing
// keypair.
func Create(client *speedycloud.ServiceClient, opts CreateOpts) CreateResult {
	var res CreateResult

	reqBody, err := opts.ToKeyPairCreateUrlEncode()
	if err != nil {
		res.Err = err
		return res
	}

	_, res.Err = client.Post(createURL(client), reqBody, &res.Body, nil)
	return res
}

// Get returns public data about a previously uploaded KeyPair.
func Get(client *speedycloud.ServiceClient, id string) GetResult {
	var res GetResult
	_, res.Err = client.Post(getURL(client),
        bytes.NewBufferString(fmt.Sprintf("id=%s", id)),
        &res.Body,
        nil)

	return res
}

// Get returns public data about a previously uploaded KeyPair.
func GetAll(client *speedycloud.ServiceClient) GetResult {
    var res GetResult
    _, res.Err = client.Post(listURL(client), bytes.NewBufferString(""), &res.Body, nil)

    return res
}

// Delete requests the deletion of a previous stored KeyPair from the server.
func Delete(client *speedycloud.ServiceClient, id string) DeleteResult {
	var res DeleteResult
	_, res.Err = client.Post(deleteURL(client),
        bytes.NewBufferString(fmt.Sprintf("id=%s", id)),
        &res.Body,
        nil)
	return res
}
