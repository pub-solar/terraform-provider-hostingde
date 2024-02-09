package hostingde

import (
	"errors"
	"fmt"
	"net/http"
)

// https://www.hosting.de/api/?json#listing-zones
func (c *Client) listZones(findRequest ZonesFindRequest) (*ZonesFindResponse, error) {
	uri := c.baseURL + "/zonesFind"

	findResponse := &ZonesFindResponse{}

	rawResp, err := c.doRequest(http.MethodPost, uri, findRequest, findResponse)
	if err != nil {
		return nil, err
	}

	if len(findResponse.Response.Data) == 0 {
		return nil, fmt.Errorf("%v: uri: %s %s", err, uri, rawResp)
	}

	if findResponse.Status != "success" && findResponse.Status != "pending" {
		return findResponse, errors.New(toErrorWithNewlines(uri, rawResp))
	}

	return findResponse, nil
}

// https://www.hosting.de/api/?json#creating-new-zones
func (c *Client) createZone(createRequest ZoneCreateRequest) (*ZoneCreateResponse, error) {
	uri := c.baseURL + "/zoneCreate"

	createResponse := &ZoneCreateResponse{}

	rawResp, err := c.doRequest(http.MethodPost, uri, createRequest, createResponse)
	if err != nil {
		return nil, err
	}

	if createResponse.Status != "success" && createResponse.Status != "pending" {
		return nil, errors.New(toErrorWithNewlines(uri, rawResp))
	}

	return createResponse, nil
}

// https://www.hosting.de/api/?json#updating-zones
func (c *Client) updateZone(updateRequest ZoneUpdateRequest) (*ZoneUpdateResponse, error) {
	uri := c.baseURL + "/zoneUpdate"

	updateResponse := &ZoneUpdateResponse{}

	rawResp, err := c.doRequest(http.MethodPost, uri, updateRequest, updateResponse)
	if err != nil {
		return nil, err
	}

	if updateResponse.Status != "success" && updateResponse.Status != "pending" {
		return nil, errors.New(toErrorWithNewlines(uri, rawResp))
	}

	return updateResponse, nil
}

// https://www.hosting.de/api/?json#deleting-zones
func (c *Client) deleteZone(deleteRequest ZoneDeleteRequest) (*ZoneDeleteResponse, error) {
	uri := c.baseURL + "/zoneDelete"

	deleteResponse := &ZoneDeleteResponse{}

	rawResp, err := c.doRequest(http.MethodPost, uri, deleteRequest, deleteResponse)
	if err != nil {
		return nil, err
	}

	if deleteResponse.Status != "success" && deleteResponse.Status != "pending" {
		return nil, errors.New(toErrorWithNewlines(uri, rawResp))
	}

	return deleteResponse, nil
}

// https://www.hosting.de/api/?json#purging-zones
func (c *Client) purgeZone(purgeRequest ZoneDeleteRequest) (*ZoneDeleteResponse, error) {
	uri := c.baseURL + "/zonePurgeRestorable"

	purgeResponse := &ZoneDeleteResponse{}

	rawResp, err := c.doRequest(http.MethodPost, uri, purgeRequest, purgeResponse)
	if err != nil {
		return nil, err
	}

	if purgeResponse.Status != "success" && purgeResponse.Status != "pending" {
		return nil, errors.New(toErrorWithNewlines(uri, rawResp))
	}

	return purgeResponse, nil
}
