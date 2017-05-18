package servers

import "github.com/hna/speedycloud"

const resourcePath = "cloud_servers"

func createURL(client *speedycloud.ServiceClient) string {
	return client.ServiceURL(resourcePath, "provision")
}


func listURL(client *speedycloud.ServiceClient) string {
	return client.ServiceURL(resourcePath)
}

func getURL(client *speedycloud.ServiceClient, id string) string {
	return client.ServiceURL(resourcePath, id)
}

func actionURL(client *speedycloud.ServiceClient, id string, action string) string {
	return client.ServiceURL(resourcePath, id, action)
}

