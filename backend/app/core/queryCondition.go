package core

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
)

// ConditionParams represent request query params
// that will be used on transport and repository level.
type ConditionParams struct {
	AccountID uuid.UUID           `json:"account_id"`
	Limit     string              `json:"limit"`
	Offset    string              `json:"offset"`
	Filter    []QuerySliceElement `json:"filter"`
	Sort      []QuerySliceElement `json:"sort"`
	Export    string              `json:"export"`
	CheckList ListValidationFilds `json:"chek_list"`
}

type QuerySliceElement struct {
	Key string
	Val string
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

	if cp.CheckList.AccountID {
		_, err := uuid.Parse(string(cp.AccountID.String()))
		if err != nil {
			return err
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
