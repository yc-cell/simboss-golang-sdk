package simboss

import (
	"github.com/yc-cell/simboss-golang-sdk/utils"
	"net/url"
)

type RealnameService struct {
	client *Client
}

func (r *RealnameService) Submit(params url.Values) error {
	if err := RequiredCardId(params); err != nil {
		return err
	}
	if !utils.Required(params, "name", "licenseType", "licenseCode", "phone", "pic1", "pic2") {
		return ErrRequired
	}
	_, err := r.client.Post("/realname/submitRealname", params)
	if err != nil {
		return err
	}
	return nil
}
