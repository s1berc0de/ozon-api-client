// Code generated by go-enum DO NOT EDIT.
// Version: 0.5.5
// Revision: b9e7d1ac24b2b7f6a5b451fa3d21706ffd8d79e2
// Build Date: 2023-01-30T01:49:43Z
// Built By: goreleaser

package uploaddigitalcodes

import (
	"fmt"
	"strings"
)

const (
	// InfoResponseResultStatusPending is a InfoResponseResultStatus of type Pending.
	InfoResponseResultStatusPending InfoResponseResultStatus = iota + 1
	// InfoResponseResultStatusImported is a InfoResponseResultStatus of type Imported.
	InfoResponseResultStatusImported
	// InfoResponseResultStatusFailed is a InfoResponseResultStatus of type Failed.
	InfoResponseResultStatusFailed
)

var ErrInvalidInfoResponseResultStatus = fmt.Errorf("not a valid InfoResponseResultStatus, try [%s]", strings.Join(_InfoResponseResultStatusNames, ", "))

const _InfoResponseResultStatusName = "pendingimportedfailed"

var _InfoResponseResultStatusNames = []string{
	_InfoResponseResultStatusName[0:7],
	_InfoResponseResultStatusName[7:15],
	_InfoResponseResultStatusName[15:21],
}

// InfoResponseResultStatusNames returns a list of possible string values of InfoResponseResultStatus.
func InfoResponseResultStatusNames() []string {
	tmp := make([]string, len(_InfoResponseResultStatusNames))
	copy(tmp, _InfoResponseResultStatusNames)
	return tmp
}

var _InfoResponseResultStatusMap = map[InfoResponseResultStatus]string{
	InfoResponseResultStatusPending:  _InfoResponseResultStatusName[0:7],
	InfoResponseResultStatusImported: _InfoResponseResultStatusName[7:15],
	InfoResponseResultStatusFailed:   _InfoResponseResultStatusName[15:21],
}

// String implements the Stringer interface.
func (x InfoResponseResultStatus) String() string {
	if str, ok := _InfoResponseResultStatusMap[x]; ok {
		return str
	}
	return fmt.Sprintf("InfoResponseResultStatus(%d)", x)
}

var _InfoResponseResultStatusValue = map[string]InfoResponseResultStatus{
	_InfoResponseResultStatusName[0:7]:   InfoResponseResultStatusPending,
	_InfoResponseResultStatusName[7:15]:  InfoResponseResultStatusImported,
	_InfoResponseResultStatusName[15:21]: InfoResponseResultStatusFailed,
}

// ParseInfoResponseResultStatus attempts to convert a string to a InfoResponseResultStatus.
func ParseInfoResponseResultStatus(name string) (InfoResponseResultStatus, error) {
	if x, ok := _InfoResponseResultStatusValue[name]; ok {
		return x, nil
	}
	return InfoResponseResultStatus(0), fmt.Errorf("%s is %w", name, ErrInvalidInfoResponseResultStatus)
}

// MarshalText implements the text marshaller method.
func (x InfoResponseResultStatus) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *InfoResponseResultStatus) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseInfoResponseResultStatus(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}
