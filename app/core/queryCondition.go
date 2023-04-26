package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	minOffset               = 0
	maxOffset               = 1000
	minRate                 = 0
	maxRate                 = 10
	allowedLimitVal         = []string{"20", "50", "100"}
	allowedFilterKey        = []string{"genre", "rate", "type", "account_id"}
	allowedSortKey          = []string{"rate", "release_date", "duration"}
	allowedSortValue        = []string{"asc", "desc"}
	allowedExportValue      = []string{"csv", "none"}
	ErrUnallowedOffset      = errors.New("unallowed offset")
	ErrUnallowedFilterKey   = errors.New("unallowed filter key")
	ErrUnallowedSort        = errors.New("unallowed sort")
	ErrUnallowedLimit       = errors.New("unallowed limit")
	ErrUnallowedExportValue = errors.New("unallowed export value")
	ErrUnallowedRateValue   = errors.New("unallowed rate value")
	ErrUnkownConditionKey   = errors.New("condition has unknown parameters")
)

// ConditionParams represent request query params
// that will be used on transport and repository level.
type ConditionParams struct {
	Limit     string              `json:"limit"`
	Offset    string              `json:"offset"`
	Filter    []QuerySliceElement `json:"filter"`
	Sort      []QuerySliceElement `json:"sort"`
	Export    string              `json:"export"`
	CheckList ListValidationFilds `json:"check_list"`
}

type QuerySliceElement struct {
	Key string
	Val string
}

func (cp *ConditionParams) Prepare(c *gin.Context) error {
	cp.CheckList.Export = true
	cp.CheckList.Filter = true
	cp.CheckList.Sort = true
	cp.CheckList.Offset = true
	cp.CheckList.Limit = true

	cp.Limit = c.Query("limit")
	cp.Offset = c.Query("offset")
	cp.Export = c.Query("export")

	for _, v := range c.QueryArray("f") {
		keyval := strings.Split(v, ":")

		var element QuerySliceElement

		element.Key = keyval[0]
		element.Val = keyval[1]

		cp.Filter = append(cp.Filter, element)
	}

	for _, v := range c.QueryArray("s") {
		keyval := strings.Split(v, ":")

		var element QuerySliceElement

		element.Key = keyval[0]
		element.Val = keyval[1]

		cp.Sort = append(cp.Sort, element)
	}

	cp.SetDefaultValues()

	if err := cp.Validate(); err != nil {
		return fmt.Errorf("query preparetion failed: %w", err)
	}

	return nil
}

type ListValidationFilds struct {
	AccountID bool
	Limit     bool
	Offset    bool
	Filter    bool
	Sort      bool
	Export    bool
}

type DateTime struct {
	time.Time
}

// Convert the internal date as CSV string.
func (date *DateTime) MarshalCSV() (string, error) {
	return date.Time.Format("2006-01-02"), nil
}

func (date *DateTime) UnmarshalCSV(csv string) (err error) {
	date.Time, err = time.Parse("2006-01-02", csv)

	return fmt.Errorf("custom csv unmarshal got an error: %w", err)
}

func (date DateTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", date.Time.Format("2006-01-02"))), nil
}

func (date *DateTime) UnmarshalJSON(data []byte) error {
	var rawDate string

	err := json.Unmarshal(data, &rawDate)
	if err != nil {
		return fmt.Errorf("custom unmarshal error: %w", err)
	}

	parsedDate, err := time.Parse("2006-01-02", rawDate)
	if err != nil {
		return fmt.Errorf("date parsing error: %w", err)
	}

	*date = DateTime{parsedDate}

	return nil
}

// Set the default values to the Limit and Offset fields.
func (cp *ConditionParams) SetDefaultValues() {
	if cp.Limit == "" {
		cp.Limit = "20"
	}

	if cp.Offset == "" {
		cp.Offset = "0"
	}

	if cp.Export == "" {
		cp.Export = "none"
	}
}

func (cp ConditionParams) Validate() error {
	if cp.CheckList.Limit {
		if notInSlice(cp.Limit, allowedLimitVal) {
			return ErrUnallowedLimit
		}
	}

	if cp.CheckList.Offset {
		offset, err := strconv.Atoi(cp.Offset)
		if err != nil {
			return fmt.Errorf("offset has to be an integer: %w", err)
		}

		if offset < minOffset || offset > maxOffset {
			return fmt.Errorf("offset should be in range from %v to %v: %w", minOffset, maxOffset, ErrUnallowedOffset)
		}
	}

	if cp.CheckList.Filter {
		for _, elem := range cp.Filter {
			if notInSlice(elem.Key, allowedFilterKey) {
				return ErrUnallowedFilterKey
			}

			if elem.Key == "rate" {
				i, err := strconv.Atoi(elem.Val)
				if err != nil {
					return fmt.Errorf("the value shlould be an integer: %w", err)
				}

				if i < minRate || i > maxRate {
					return fmt.Errorf("the value should be in range from 1 to 10 :%w", ErrUnallowedRateValue)
				}
			}
		}
	}

	if cp.CheckList.Sort {
		for _, elem := range cp.Sort {
			if notInSlice(elem.Key, allowedSortKey) {
				return fmt.Errorf("key %v is bad: %w", elem.Key, ErrUnallowedSort)
			}

			if notInSlice(elem.Val, allowedSortValue) {
				return fmt.Errorf("the sort value: %w", ErrUnallowedSort)
			}
		}
	}

	if cp.CheckList.Export {
		if notInSlice(cp.Export, allowedExportValue) {
			return fmt.Errorf("value %v: %w", cp.Export, ErrUnallowedExportValue)
		}
	}

	return nil
}

func notInSlice(element string, slice []string) bool {
	for _, s := range slice {
		if s == element {
			return false
		}
	}

	return true
}
