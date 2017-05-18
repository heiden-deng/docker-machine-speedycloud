package keypairs

import "github.com/hna/speedycloud"

const resourcePath = "sshkey"

func resourceURL(c *speedycloud.ServiceClient) string {
	return c.ServiceURL(resourcePath)
}

func listURL(c *speedycloud.ServiceClient) string {
	return resourceURL(c)
}

func createURL(c *speedycloud.ServiceClient) string {
	return c.ServiceURL(resourcePath, "create")
}

func getURL(c *speedycloud.ServiceClient) string {
	return c.ServiceURL(resourcePath, "info")
}

func deleteURL(c *speedycloud.ServiceClient) string {
	return c.ServiceURL(resourcePath, "delete")
}
