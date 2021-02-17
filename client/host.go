package client

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type Host struct {
	Id        string   `json:"id"`
	NameLabel string   `json:"name_label"`
	Tags      []string `json:"tags"`
	Pool      string   `json:"$pool"`
}

func (h Host) Compare(obj interface{}) bool {
	otherHost := obj.(Host)
	if otherHost.Id == h.Id {
		return true
	}
	if h.Pool == otherHost.Pool {
		return true
	}
	if otherHost.NameLabel != "" && h.NameLabel == otherHost.NameLabel {
		return true
	}
	return false
}

func (c *Client) GetHostByName(nameLabel string) (hosts []Host, err error) {
	obj, err := c.FindFromGetAllObjects(Host{NameLabel: nameLabel})
	if err != nil {
		return
	}
	return obj.([]Host), nil
}

func (c *Client) GetHostById(id string) (host Host, err error) {
	obj, err := c.FindFromGetAllObjects(Host{Id: id})
	if err != nil {
		return
	}
	hosts, ok := obj.([]Host)

	if !ok {
		return host, errors.New("failed to coerce response into Host slice")
	}

	if len(hosts) != 1 {
		return host, errors.New(fmt.Sprintf("expected a single host to be returned, instead received: %d in the response: %v", len(hosts), obj))
	}

	return hosts[0], nil
}

func FindHostForTests(hostId string, host *Host) {
	c, err := NewClient(GetConfigFromEnv())
	if err != nil {
		fmt.Printf("failed to create client with error: %v", err)
		os.Exit(-1)
	}

	queriedHost, err := c.GetHostById(hostId)

	if err != nil {
		fmt.Printf("failed to find a host with id: %v with error: %v\n", hostId, err)
		os.Exit(-1)
	}

	*host = queriedHost
}

func (c *Client) GetHostsByPoolName(pool string) (hosts map[string]string, err error) {
	hosts = make(map[string]string, 0)
	obj, err := c.FindFromGetAllObjects(Host{Pool: pool})

	if err != nil {
		return
	}
	slice := obj.([]Host)
	for _, v := range slice {
		hosts[v.NameLabel[:strings.Index(v.NameLabel, ".")]] = v.Id
	}

	return hosts, nil
}
