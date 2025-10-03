package okx_connector

import (
	"context"
	"encoding/json"
	"net/http"
)

type AnnouncementTypeService struct {
	c *Client
}

type AnnouncementType struct {
	AnnType     string `json:"annType"`
	AnnTypeDesc string `json:"annTypeDesc"`
}

// Send the request
func (s *AnnouncementTypeService) Do(ctx context.Context, opts ...RequestOption) (res []*AnnouncementType, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v5/support/announcement-types",
		secType:  secTypeNone,
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	result := new([]*AnnouncementType)
	err = json.Unmarshal(data, result)
	if err != nil {
		return nil, err
	}
	return *result, nil
}

type AnnouncementService struct {
	c       *Client
	annType *string
	page    *int
}

func (s *AnnouncementService) AnnType(annType string) *AnnouncementService {
	s.annType = &annType
	return s
}

func (s *AnnouncementService) Page(page int) *AnnouncementService {
	s.page = &page
	return s
}

type AnnouncementResponse struct {
	TotalPage string          `json:"totalPage"`
	Details   []*Announcement `json:"details"`
}

type Announcement struct {
	Title   string `json:"title"`
	AnnType string `json:"annType"`
	PTime   string `json:"pTime"`
	Url     string `json:"url"`
}

func (s *AnnouncementService) Do(ctx context.Context, opts ...RequestOption) (res []*AnnouncementResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v5/support/announcements",
		secType:  secTypeSigned,
	}
	if s.annType != nil {
		r.setParam("annType", *s.annType)
	}
	if s.page != nil {
		r.setParam("page", *s.page)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	result := new([]*AnnouncementResponse)
	err = json.Unmarshal(data, result)
	if err != nil {
		return nil, err
	}
	return *result, nil
}
