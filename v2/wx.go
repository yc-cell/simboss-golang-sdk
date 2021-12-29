package simboss

import (
	"net/url"
	"simboss-golang-sdk/utils"
)

type WxService struct {
	client *Client
}

// 微信预支付-单卡续费
func (s *SmsService) PrePayment(params url.Values) error {
	if err := RequiredCardId(params); err != nil {
		return err
	}
	if !utils.Required(params, "ratePlanId", "month", "outTradeNo") {
		return ErrRequired
	}
	_, err := s.client.Post("/wx/device/prePayment", params)
	if err != nil {
		return err
	}
	return nil
}
