package models

import "time"

// Model for token
// Request
type TRKDTokenRequest struct {
	ServiceTokenRequest ServiceTokenRequest `json:"CreateServiceToken_Request_1"`
}

type ServiceTokenRequest struct {
	ApplicationID string `json:"ApplicationID"`
	Username      string `json:"Username"`
	Password      string `json:"Password"`
}

// Response
type TRKDTokenResponse struct {
	ServiceTokenResponse ServiceTokenResponse `json:"CreateServiceToken_Response_1"`
}

type ServiceTokenResponse struct {
	Expiration time.Time `json:"Expiration"`
	Token      string    `json:"Token"`
}

// Model for FX
// Request
type TRKDFXRequest struct {
	RetrieveItemRequest RetrieveItemRequest `json:"RetrieveItem_Request_3"`
}

type RetrieveItemRequest struct {
	ItemRequests        []ItemRequest `json:"ItemRequest"`
	TrimResponse        bool          `json:"TrimResponse"`
	IncludeChildItemQoS bool          `json:"IncludeChildItemQoS"`
}

type ItemRequest struct {
	Fields      string       `json:"Fields"`
	RequestKeys []RequestKey `json:"RequestKey"`
	Scope       string       `json:"Scope"`
}

type RequestKey struct {
	Name     string `json:"Name"`
	Service  string `json:"Service,omitempty"`
	NameType string `json:"NameType"`
}

// Response
type TRKDFXResponse struct {
	RetrieveItemResponse RetrieveItemResponse `json:"RetrieveItem_Response_3"`
}

type RetrieveItemResponse struct {
	ItemResponses []ItemResponse `json:"ItemResponse"`
}

type ItemResponse struct {
	Items []Item `json:"Item"`
}

type Item struct {
	RequestKey RequestKey  `json:"RequestKey"`
	Fields     Fields      `json:"Fields"`
	ChildItems []ChildItem `json:"ChildItem"`
}

type Fields struct {
	Field []Field `json:"Field"`
}

type Field struct {
	DataType   string  `json:"DataType"`
	Name       string  `json:"Name"`
	Utf8String string  `json:"Utf8String,omitempty"`
	Double     float64 `json:"Double,omitempty"`
}

type ChildItem struct {
	Name   string `json:"Name"`
	Fields Fields `json:"Fields"`
}
