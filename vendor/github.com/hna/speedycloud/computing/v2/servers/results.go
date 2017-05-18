package servers

import (
	"reflect"

	"github.com/mitchellh/mapstructure"
    "github.com/hna/speedycloud"
    //"github.com/hna/speedycloud/pagination"
)

type serverResult struct {
    speedycloud.Result
}

// Extract interprets any serverResult as a Server, if possible.
func (r serverResult) Extract() (*Server, error) {
	if r.Err != nil {
		return nil, r.Err
	}

	var response Server

	config := &mapstructure.DecoderConfig{
        WeaklyTypedInput: true,
		DecodeHook: toMapFromString,
		Result:     &response,
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return nil, err
	}

	err = decoder.Decode(r.Body)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// CreateResult temporarily contains the response from a Create call.
type CreateResult struct {
	serverResult
}

// GetResult temporarily contains the response from a Get call.
type GetResult struct {
	serverResult
}

// UpdateResult temporarily contains the response from an Update call.
type UpdateResult struct {
	serverResult
}

// DeleteResult temporarily contains the response from a Delete call.
type DeleteResult struct {
    speedycloud.ErrResult
}

// RebuildResult temporarily contains the response from a Rebuild call.
type RebuildResult struct {
	serverResult
}

// ActionResult represents the result of server action operations, like reboot
type ActionResult struct {
    speedycloud.ErrResult
}

// RescueResult represents the result of a server rescue operation
type RescueResult struct {
	ActionResult
}

// Extract interprets any RescueResult as an AdminPass, if possible.
func (r RescueResult) Extract() (string, error) {
	if r.Err != nil {
		return "", r.Err
	}

	var response struct {
		AdminPass string `mapstructure:"adminPass"`
	}

	err := mapstructure.Decode(r.Body, &response)
	return response.AdminPass, err
}

type AZ struct {
	DisplayName string `mapstructure:"display_name"`
	Name string `mapstructure:"name"`
}

type Network struct {
    DisplayName string `mapstructure:"display_name"`
    Name string `mapstructure:"name"`
}
// Server exposes only the standard OpenStack fields corresponding to a given server on the user's account.
type Server struct {
    Az AZ `mapstructure:"availability_zone"`
    Image string `mapstructure:"image"`
    GroupName string `mapstructure:"group_name"`
    Ips []string `mapstructure:"ips"`
    Bandwidth string `mapstructure:"bandwidth"`
    DiskCapacity string `mapstructure:"disk"`
    ID string `mapstructure:"id"`
    SecurityGroups []string `mapstructure:"security_groups"`
    MonitorUrl string `mapstructure:"monitor_url"`
    Alias   string `mapstructure:"alias"`
    DisablePrivateNetwork bool `mapstructure:"disable_private_network"`
    Hostname string `mapstructure:"hostname"`
    BootScript   string `mapstructure:"bootscript"`
    Memory string `mapstructure:"memory"`
    Metadata string `mapstructure:"metadata"`
    Status string `mapstructure:"status"`
    JoinedNetworks []string `mapstructure:"joined_networks"`
    Tags string `mapstructure:"tags"`
    DiskType string `mapstructure:"disk_type"`
    HasRunningJob bool `mapstructure:"has_running_job"`
    Name string `mapstructure:"name"`
    CreateTime string `mapstructure:"created_at"`
    Network Network `mapstructure:"network"`
    Volumes []string `mapstructure:"volumes"`
    Password string `mapstructure:"default_password"`
    CpuNumber string `mapstructure:"cpu"`
}

// ServerPage abstracts the raw results of making a List() request against the API.
// As OpenStack extensions may freely alter the response bodies of structures returned to the client, you may only safely access the
// data provided through the ExtractServers call.
//type ServerPage struct {
//	pagination.LinkedPageBase
//}

// IsEmpty returns true if a page contains no Server results.
//func (page ServerPage) IsEmpty() (bool, error) {
//	servers, err := ExtractServers(page)
//	if err != nil {
//		return true, err
//	}
//	return len(servers) == 0, nil
//}

// NextPageURL uses the response's embedded link reference to navigate to the next page of results.
//func (page ServerPage) NextPageURL() (string, error) {
//	type resp struct {
//		Links []speedycloud.Link `mapstructure:"servers_links"`
//	}
//
//	var r resp
//	err := mapstructure.Decode(page.Body, &r)
//	if err != nil {
//		return "", err
//	}
//
//	return speedycloud.ExtractNextURL(r.Links)
//}
//
//// ExtractServers interprets the results of a single page from a List() call, producing a slice of Server entities.
//func ExtractServers(page pagination.Page) ([]Server, error) {
//	casted := page.(ServerPage).Body
//
//	var response struct {
//		Servers []Server `mapstructure:"servers"`
//	}
//
//	config := &mapstructure.DecoderConfig{
//		DecodeHook: toMapFromString,
//		Result:     &response,
//	}
//	decoder, err := mapstructure.NewDecoder(config)
//	if err != nil {
//		return nil, err
//	}
//
//	err = decoder.Decode(casted)
//
//	return response.Servers, err
//}

// MetadataResult contains the result of a call for (potentially) multiple key-value pairs.
type MetadataResult struct {
    speedycloud.Result
}

// GetMetadataResult temporarily contains the response from a metadata Get call.
type GetMetadataResult struct {
	MetadataResult
}

// ResetMetadataResult temporarily contains the response from a metadata Reset call.
type ResetMetadataResult struct {
	MetadataResult
}

// UpdateMetadataResult temporarily contains the response from a metadata Update call.
type UpdateMetadataResult struct {
	MetadataResult
}

// MetadatumResult contains the result of a call for individual a single key-value pair.
type MetadatumResult struct {
    speedycloud.Result
}

// GetMetadatumResult temporarily contains the response from a metadatum Get call.
type GetMetadatumResult struct {
	MetadatumResult
}

// CreateMetadatumResult temporarily contains the response from a metadatum Create call.
type CreateMetadatumResult struct {
	MetadatumResult
}

// DeleteMetadatumResult temporarily contains the response from a metadatum Delete call.
type DeleteMetadatumResult struct {
    speedycloud.ErrResult
}

// Extract interprets any MetadataResult as a Metadata, if possible.
func (r MetadataResult) Extract() (map[string]string, error) {
	if r.Err != nil {
		return nil, r.Err
	}

	var response struct {
		Metadata map[string]string `mapstructure:"metadata"`
	}

	err := mapstructure.Decode(r.Body, &response)
	return response.Metadata, err
}

// Extract interprets any MetadatumResult as a Metadatum, if possible.
func (r MetadatumResult) Extract() (map[string]string, error) {
	if r.Err != nil {
		return nil, r.Err
	}

	var response struct {
		Metadatum map[string]string `mapstructure:"meta"`
	}

	err := mapstructure.Decode(r.Body, &response)
	return response.Metadatum, err
}

func toMapFromString(from reflect.Kind, to reflect.Kind, data interface{}) (interface{}, error) {
	if (from == reflect.String) && (to == reflect.Map) {
		return map[string]interface{}{}, nil
	}
	return data, nil
}

// Address represents an IP address.
type Address struct {
	Version int    `mapstructure:"version"`
	Address string `mapstructure:"addr"`
}

// AddressPage abstracts the raw results of making a ListAddresses() request against the API.
// As OpenStack extensions may freely alter the response bodies of structures returned
// to the client, you may only safely access the data provided through the ExtractAddresses call.
//type AddressPage struct {
//	pagination.SinglePageBase
//}
//
//// IsEmpty returns true if an AddressPage contains no networks.
//func (r AddressPage) IsEmpty() (bool, error) {
//	addresses, err := ExtractAddresses(r)
//	if err != nil {
//		return true, err
//	}
//	return len(addresses) == 0, nil
//}
//
//// ExtractAddresses interprets the results of a single page from a ListAddresses() call,
//// producing a map of addresses.
//func ExtractAddresses(page pagination.Page) (map[string][]Address, error) {
//	casted := page.(AddressPage).Body
//
//	var response struct {
//		Addresses map[string][]Address `mapstructure:"addresses"`
//	}
//
//	err := mapstructure.Decode(casted, &response)
//	if err != nil {
//		return nil, err
//	}
//
//	return response.Addresses, err
//}
//
//// NetworkAddressPage abstracts the raw results of making a ListAddressesByNetwork() request against the API.
//// As OpenStack extensions may freely alter the response bodies of structures returned
//// to the client, you may only safely access the data provided through the ExtractAddresses call.
//type NetworkAddressPage struct {
//	pagination.SinglePageBase
//}
//
//// IsEmpty returns true if a NetworkAddressPage contains no addresses.
//func (r NetworkAddressPage) IsEmpty() (bool, error) {
//	addresses, err := ExtractNetworkAddresses(r)
//	if err != nil {
//		return true, err
//	}
//	return len(addresses) == 0, nil
//}
//
//// ExtractNetworkAddresses interprets the results of a single page from a ListAddressesByNetwork() call,
//// producing a slice of addresses.
//func ExtractNetworkAddresses(page pagination.Page) ([]Address, error) {
//	casted := page.(NetworkAddressPage).Body
//
//	var response map[string][]Address
//	err := mapstructure.Decode(casted, &response)
//	if err != nil {
//		return nil, err
//	}
//
//	var key string
//	for k := range response {
//		key = k
//	}
//
//	return response[key], err
//}
