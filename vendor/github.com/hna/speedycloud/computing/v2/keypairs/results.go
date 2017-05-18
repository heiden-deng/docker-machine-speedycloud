package keypairs

import (
	"github.com/mitchellh/mapstructure"
    "github.com/hna/speedycloud"
    //"github.com/hna/speedycloud/pagination"
    "fmt"
)

// KeyPair is an SSH key known to the OpenStack cluster that is available to be injected into
// servers.
type KeyPair struct {
    // Name is used to refer to this keypair from other services within this region.
    Name string `mapstructure:"name"`

    // Fingerprint is a short sequence of bytes that can be used to authenticate or validate a longer
    // public key.
    Fingerprint string `mapstructure:"fingerprint"`

    // PublicKey is the public key from this pair, in OpenSSH format. "ssh-rsa AAAAB3Nz..."
    PublicKey string `mapstructure:"ssh_public_key_content"`

    // PrivateKey is the private key from this pair, in PEM format.
    // "-----BEGIN RSA PRIVATE KEY-----\nMIICXA..." It is only present if this keypair was just
    // returned from a Create call
    DisplayName string `mapstructure:"display_name"`

    created string `mapstructure:"created_at"`

    // UserID is the user who owns this keypair.
    ID string `mapstructure:"id"`

}

// KeyPairPage stores a single, only page of KeyPair results from a List call.
//type KeyPairPage struct {
//	pagination.SinglePageBase
//}

//// IsEmpty determines whether or not a KeyPairPage is empty.
//func (page KeyPairPage) IsEmpty() (bool, error) {
//	ks, err := ExtractKeyPairs(page)
//	return len(ks) == 0, err
//}

//// ExtractKeyPairs interprets a page of results as a slice of KeyPairs.
//func ExtractKeyPairs(page pagination.Page) ([]KeyPair, error) {
//	type pair struct {
//		KeyPair KeyPair `mapstructure:"keypair"`
//	}
//
//	var resp struct {
//		KeyPairs []pair `mapstructure:"keypairs"`
//	}
//
//	err := mapstructure.Decode(page.(KeyPairPage).Body, &resp)
//	results := make([]KeyPair, len(resp.KeyPairs))
//	for i, pair := range resp.KeyPairs {
//		results[i] = pair.KeyPair
//	}
//	return results, err
//}

type keyPairResult struct {
    speedycloud.Result
}

// Extract is a method that attempts to interpret any KeyPair resource response as a KeyPair struct.
func (r keyPairResult) Extract() (*KeyPair, error) {
	if r.Err != nil {
		return nil, r.Err
	}

	var res KeyPair

    cfg := &mapstructure.DecoderConfig{
        WeaklyTypedInput: true,
        Result:     &res,
    }
    decoder, err := mapstructure.NewDecoder(cfg)
    if err != nil {
        return nil, err
    }
    decoder.Decode(r.Body)
    if err != nil{
        return nil, err
    }
	return &res, nil
}

func (r keyPairResult) ExtractByDisplayName(name string) (*KeyPair, error) {
	if r.Err != nil {
		return nil, r.Err
	}

	var res []KeyPair

    cfg := &mapstructure.DecoderConfig{
        WeaklyTypedInput: true,
        Result:     &res,
    }
    decoder, err := mapstructure.NewDecoder(cfg)
    if err != nil {
        return nil, err
    }
    decoder.Decode(r.Body)
    if err != nil{
        return nil, err
    }

    for _, val := range res {
        if val.DisplayName == name {
            return &val, nil
        }
    }
    return nil, fmt.Errorf("no keypair %s", name)

}
// CreateResult is the response from a Create operation. Call its Extract method to interpret it
// as a KeyPair.
type CreateResult struct {
	keyPairResult
}

// GetResult is the response from a Get operation. Call its Extract method to interpret it
// as a KeyPair.
type GetResult struct {
	keyPairResult
}

// DeleteResult is the response from a Delete operation. Call its Extract method to determine if
// the call succeeded or failed.
type DeleteResult struct {
    speedycloud.ErrResult
}
