// generated by jsonenums -type=ArticleAccessStatus; DO NOT EDIT

package types

import (
	"encoding/json"
	"fmt"
)

var (
	_ArticleAccessStatusNameToValue = map[string]ArticleAccessStatus{
		"ArticleDraft":   ArticleDraft,
		"ArticlePublic":  ArticlePublic,
		"ArticlePrivate": ArticlePrivate,
	}

	_ArticleAccessStatusValueToName = map[ArticleAccessStatus]string{
		ArticleDraft:   "ArticleDraft",
		ArticlePublic:  "ArticlePublic",
		ArticlePrivate: "ArticlePrivate",
	}
)

func init() {
	var v ArticleAccessStatus
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_ArticleAccessStatusNameToValue = map[string]ArticleAccessStatus{
			interface{}(ArticleDraft).(fmt.Stringer).String():   ArticleDraft,
			interface{}(ArticlePublic).(fmt.Stringer).String():  ArticlePublic,
			interface{}(ArticlePrivate).(fmt.Stringer).String(): ArticlePrivate,
		}
	}
}

// MarshalJSON is generated so ArticleAccessStatus satisfies json.Marshaler.
func (r ArticleAccessStatus) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _ArticleAccessStatusValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid ArticleAccessStatus: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so ArticleAccessStatus satisfies json.Unmarshaler.
func (r *ArticleAccessStatus) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("ArticleAccessStatus should be a string, got %s", data)
	}
	v, ok := _ArticleAccessStatusNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid ArticleAccessStatus %q", s)
	}
	*r = v
	return nil
}
