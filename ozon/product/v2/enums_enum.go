// Code generated by go-enum DO NOT EDIT.
// Version:
// Revision:
// Build Date:
// Built By:

package v2

import (
	"fmt"
	"strings"
)

const (
	// ImportRequestItemCurrencyCodeRUB is a ImportRequestItemCurrencyCode of type RUB.
	ImportRequestItemCurrencyCodeRUB ImportRequestItemCurrencyCode = iota
	// ImportRequestItemCurrencyCodeBYN is a ImportRequestItemCurrencyCode of type BYN.
	ImportRequestItemCurrencyCodeBYN
	// ImportRequestItemCurrencyCodeKZT is a ImportRequestItemCurrencyCode of type KZT.
	ImportRequestItemCurrencyCodeKZT
	// ImportRequestItemCurrencyCodeEUR is a ImportRequestItemCurrencyCode of type EUR.
	ImportRequestItemCurrencyCodeEUR
	// ImportRequestItemCurrencyCodeUSD is a ImportRequestItemCurrencyCode of type USD.
	ImportRequestItemCurrencyCodeUSD
	// ImportRequestItemCurrencyCodeCNY is a ImportRequestItemCurrencyCode of type CNY.
	ImportRequestItemCurrencyCodeCNY
)

var ErrInvalidImportRequestItemCurrencyCode = fmt.Errorf("not a valid ImportRequestItemCurrencyCode, try [%s]", strings.Join(_ImportRequestItemCurrencyCodeNames, ", "))

const _ImportRequestItemCurrencyCodeName = "RUBBYNKZTEURUSDCNY"

var _ImportRequestItemCurrencyCodeNames = []string{
	_ImportRequestItemCurrencyCodeName[0:3],
	_ImportRequestItemCurrencyCodeName[3:6],
	_ImportRequestItemCurrencyCodeName[6:9],
	_ImportRequestItemCurrencyCodeName[9:12],
	_ImportRequestItemCurrencyCodeName[12:15],
	_ImportRequestItemCurrencyCodeName[15:18],
}

// ImportRequestItemCurrencyCodeNames returns a list of possible string values of ImportRequestItemCurrencyCode.
func ImportRequestItemCurrencyCodeNames() []string {
	tmp := make([]string, len(_ImportRequestItemCurrencyCodeNames))
	copy(tmp, _ImportRequestItemCurrencyCodeNames)
	return tmp
}

var _ImportRequestItemCurrencyCodeMap = map[ImportRequestItemCurrencyCode]string{
	ImportRequestItemCurrencyCodeRUB: _ImportRequestItemCurrencyCodeName[0:3],
	ImportRequestItemCurrencyCodeBYN: _ImportRequestItemCurrencyCodeName[3:6],
	ImportRequestItemCurrencyCodeKZT: _ImportRequestItemCurrencyCodeName[6:9],
	ImportRequestItemCurrencyCodeEUR: _ImportRequestItemCurrencyCodeName[9:12],
	ImportRequestItemCurrencyCodeUSD: _ImportRequestItemCurrencyCodeName[12:15],
	ImportRequestItemCurrencyCodeCNY: _ImportRequestItemCurrencyCodeName[15:18],
}

// String implements the Stringer interface.
func (x ImportRequestItemCurrencyCode) String() string {
	if str, ok := _ImportRequestItemCurrencyCodeMap[x]; ok {
		return str
	}
	return fmt.Sprintf("ImportRequestItemCurrencyCode(%d)", x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x ImportRequestItemCurrencyCode) IsValid() bool {
	_, ok := _ImportRequestItemCurrencyCodeMap[x]
	return ok
}

var _ImportRequestItemCurrencyCodeValue = map[string]ImportRequestItemCurrencyCode{
	_ImportRequestItemCurrencyCodeName[0:3]:   ImportRequestItemCurrencyCodeRUB,
	_ImportRequestItemCurrencyCodeName[3:6]:   ImportRequestItemCurrencyCodeBYN,
	_ImportRequestItemCurrencyCodeName[6:9]:   ImportRequestItemCurrencyCodeKZT,
	_ImportRequestItemCurrencyCodeName[9:12]:  ImportRequestItemCurrencyCodeEUR,
	_ImportRequestItemCurrencyCodeName[12:15]: ImportRequestItemCurrencyCodeUSD,
	_ImportRequestItemCurrencyCodeName[15:18]: ImportRequestItemCurrencyCodeCNY,
}

// ParseImportRequestItemCurrencyCode attempts to convert a string to a ImportRequestItemCurrencyCode.
func ParseImportRequestItemCurrencyCode(name string) (ImportRequestItemCurrencyCode, error) {
	if x, ok := _ImportRequestItemCurrencyCodeValue[name]; ok {
		return x, nil
	}
	return ImportRequestItemCurrencyCode(0), fmt.Errorf("%s is %w", name, ErrInvalidImportRequestItemCurrencyCode)
}

// MarshalText implements the text marshaller method.
func (x ImportRequestItemCurrencyCode) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *ImportRequestItemCurrencyCode) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseImportRequestItemCurrencyCode(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

const (
	// ImportRequestItemDimensionUnitMm is a ImportRequestItemDimensionUnit of type Mm.
	ImportRequestItemDimensionUnitMm ImportRequestItemDimensionUnit = iota + 1
	// ImportRequestItemDimensionUnitCm is a ImportRequestItemDimensionUnit of type Cm.
	ImportRequestItemDimensionUnitCm
	// ImportRequestItemDimensionUnitIn is a ImportRequestItemDimensionUnit of type In.
	ImportRequestItemDimensionUnitIn
)

var ErrInvalidImportRequestItemDimensionUnit = fmt.Errorf("not a valid ImportRequestItemDimensionUnit, try [%s]", strings.Join(_ImportRequestItemDimensionUnitNames, ", "))

const _ImportRequestItemDimensionUnitName = "mmcmin"

var _ImportRequestItemDimensionUnitNames = []string{
	_ImportRequestItemDimensionUnitName[0:2],
	_ImportRequestItemDimensionUnitName[2:4],
	_ImportRequestItemDimensionUnitName[4:6],
}

// ImportRequestItemDimensionUnitNames returns a list of possible string values of ImportRequestItemDimensionUnit.
func ImportRequestItemDimensionUnitNames() []string {
	tmp := make([]string, len(_ImportRequestItemDimensionUnitNames))
	copy(tmp, _ImportRequestItemDimensionUnitNames)
	return tmp
}

var _ImportRequestItemDimensionUnitMap = map[ImportRequestItemDimensionUnit]string{
	ImportRequestItemDimensionUnitMm: _ImportRequestItemDimensionUnitName[0:2],
	ImportRequestItemDimensionUnitCm: _ImportRequestItemDimensionUnitName[2:4],
	ImportRequestItemDimensionUnitIn: _ImportRequestItemDimensionUnitName[4:6],
}

// String implements the Stringer interface.
func (x ImportRequestItemDimensionUnit) String() string {
	if str, ok := _ImportRequestItemDimensionUnitMap[x]; ok {
		return str
	}
	return fmt.Sprintf("ImportRequestItemDimensionUnit(%d)", x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x ImportRequestItemDimensionUnit) IsValid() bool {
	_, ok := _ImportRequestItemDimensionUnitMap[x]
	return ok
}

var _ImportRequestItemDimensionUnitValue = map[string]ImportRequestItemDimensionUnit{
	_ImportRequestItemDimensionUnitName[0:2]: ImportRequestItemDimensionUnitMm,
	_ImportRequestItemDimensionUnitName[2:4]: ImportRequestItemDimensionUnitCm,
	_ImportRequestItemDimensionUnitName[4:6]: ImportRequestItemDimensionUnitIn,
}

// ParseImportRequestItemDimensionUnit attempts to convert a string to a ImportRequestItemDimensionUnit.
func ParseImportRequestItemDimensionUnit(name string) (ImportRequestItemDimensionUnit, error) {
	if x, ok := _ImportRequestItemDimensionUnitValue[name]; ok {
		return x, nil
	}
	return ImportRequestItemDimensionUnit(0), fmt.Errorf("%s is %w", name, ErrInvalidImportRequestItemDimensionUnit)
}

// MarshalText implements the text marshaller method.
func (x ImportRequestItemDimensionUnit) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *ImportRequestItemDimensionUnit) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseImportRequestItemDimensionUnit(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

const (
	// ImportRequestItemServiceTypeISCODESERVICE is a ImportRequestItemServiceType of type IS_CODE_SERVICE.
	ImportRequestItemServiceTypeISCODESERVICE ImportRequestItemServiceType = iota
	// ImportRequestItemServiceTypeISNOCODESERVICE is a ImportRequestItemServiceType of type IS_NO_CODE_SERVICE.
	ImportRequestItemServiceTypeISNOCODESERVICE
)

var ErrInvalidImportRequestItemServiceType = fmt.Errorf("not a valid ImportRequestItemServiceType, try [%s]", strings.Join(_ImportRequestItemServiceTypeNames, ", "))

const _ImportRequestItemServiceTypeName = "IS_CODE_SERVICEIS_NO_CODE_SERVICE"

var _ImportRequestItemServiceTypeNames = []string{
	_ImportRequestItemServiceTypeName[0:15],
	_ImportRequestItemServiceTypeName[15:33],
}

// ImportRequestItemServiceTypeNames returns a list of possible string values of ImportRequestItemServiceType.
func ImportRequestItemServiceTypeNames() []string {
	tmp := make([]string, len(_ImportRequestItemServiceTypeNames))
	copy(tmp, _ImportRequestItemServiceTypeNames)
	return tmp
}

var _ImportRequestItemServiceTypeMap = map[ImportRequestItemServiceType]string{
	ImportRequestItemServiceTypeISCODESERVICE:   _ImportRequestItemServiceTypeName[0:15],
	ImportRequestItemServiceTypeISNOCODESERVICE: _ImportRequestItemServiceTypeName[15:33],
}

// String implements the Stringer interface.
func (x ImportRequestItemServiceType) String() string {
	if str, ok := _ImportRequestItemServiceTypeMap[x]; ok {
		return str
	}
	return fmt.Sprintf("ImportRequestItemServiceType(%d)", x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x ImportRequestItemServiceType) IsValid() bool {
	_, ok := _ImportRequestItemServiceTypeMap[x]
	return ok
}

var _ImportRequestItemServiceTypeValue = map[string]ImportRequestItemServiceType{
	_ImportRequestItemServiceTypeName[0:15]:  ImportRequestItemServiceTypeISCODESERVICE,
	_ImportRequestItemServiceTypeName[15:33]: ImportRequestItemServiceTypeISNOCODESERVICE,
}

// ParseImportRequestItemServiceType attempts to convert a string to a ImportRequestItemServiceType.
func ParseImportRequestItemServiceType(name string) (ImportRequestItemServiceType, error) {
	if x, ok := _ImportRequestItemServiceTypeValue[name]; ok {
		return x, nil
	}
	return ImportRequestItemServiceType(0), fmt.Errorf("%s is %w", name, ErrInvalidImportRequestItemServiceType)
}

// MarshalText implements the text marshaller method.
func (x ImportRequestItemServiceType) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *ImportRequestItemServiceType) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseImportRequestItemServiceType(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

const (
	// ImportRequestItemWeightUnitG is a ImportRequestItemWeightUnit of type G.
	ImportRequestItemWeightUnitG ImportRequestItemWeightUnit = iota + 1
	// ImportRequestItemWeightUnitKg is a ImportRequestItemWeightUnit of type Kg.
	ImportRequestItemWeightUnitKg
	// ImportRequestItemWeightUnitLb is a ImportRequestItemWeightUnit of type Lb.
	ImportRequestItemWeightUnitLb
)

var ErrInvalidImportRequestItemWeightUnit = fmt.Errorf("not a valid ImportRequestItemWeightUnit, try [%s]", strings.Join(_ImportRequestItemWeightUnitNames, ", "))

const _ImportRequestItemWeightUnitName = "gkglb"

var _ImportRequestItemWeightUnitNames = []string{
	_ImportRequestItemWeightUnitName[0:1],
	_ImportRequestItemWeightUnitName[1:3],
	_ImportRequestItemWeightUnitName[3:5],
}

// ImportRequestItemWeightUnitNames returns a list of possible string values of ImportRequestItemWeightUnit.
func ImportRequestItemWeightUnitNames() []string {
	tmp := make([]string, len(_ImportRequestItemWeightUnitNames))
	copy(tmp, _ImportRequestItemWeightUnitNames)
	return tmp
}

var _ImportRequestItemWeightUnitMap = map[ImportRequestItemWeightUnit]string{
	ImportRequestItemWeightUnitG:  _ImportRequestItemWeightUnitName[0:1],
	ImportRequestItemWeightUnitKg: _ImportRequestItemWeightUnitName[1:3],
	ImportRequestItemWeightUnitLb: _ImportRequestItemWeightUnitName[3:5],
}

// String implements the Stringer interface.
func (x ImportRequestItemWeightUnit) String() string {
	if str, ok := _ImportRequestItemWeightUnitMap[x]; ok {
		return str
	}
	return fmt.Sprintf("ImportRequestItemWeightUnit(%d)", x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x ImportRequestItemWeightUnit) IsValid() bool {
	_, ok := _ImportRequestItemWeightUnitMap[x]
	return ok
}

var _ImportRequestItemWeightUnitValue = map[string]ImportRequestItemWeightUnit{
	_ImportRequestItemWeightUnitName[0:1]: ImportRequestItemWeightUnitG,
	_ImportRequestItemWeightUnitName[1:3]: ImportRequestItemWeightUnitKg,
	_ImportRequestItemWeightUnitName[3:5]: ImportRequestItemWeightUnitLb,
}

// ParseImportRequestItemWeightUnit attempts to convert a string to a ImportRequestItemWeightUnit.
func ParseImportRequestItemWeightUnit(name string) (ImportRequestItemWeightUnit, error) {
	if x, ok := _ImportRequestItemWeightUnitValue[name]; ok {
		return x, nil
	}
	return ImportRequestItemWeightUnit(0), fmt.Errorf("%s is %w", name, ErrInvalidImportRequestItemWeightUnit)
}

// MarshalText implements the text marshaller method.
func (x ImportRequestItemWeightUnit) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *ImportRequestItemWeightUnit) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseImportRequestItemWeightUnit(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

const (
	// InfoResponseResultCurrencyCodeRUB is a InfoResponseResultCurrencyCode of type RUB.
	InfoResponseResultCurrencyCodeRUB InfoResponseResultCurrencyCode = iota + 1
	// InfoResponseResultCurrencyCodeBYN is a InfoResponseResultCurrencyCode of type BYN.
	InfoResponseResultCurrencyCodeBYN
	// InfoResponseResultCurrencyCodeKZT is a InfoResponseResultCurrencyCode of type KZT.
	InfoResponseResultCurrencyCodeKZT
	// InfoResponseResultCurrencyCodeEUR is a InfoResponseResultCurrencyCode of type EUR.
	InfoResponseResultCurrencyCodeEUR
	// InfoResponseResultCurrencyCodeUSD is a InfoResponseResultCurrencyCode of type USD.
	InfoResponseResultCurrencyCodeUSD
	// InfoResponseResultCurrencyCodeCNY is a InfoResponseResultCurrencyCode of type CNY.
	InfoResponseResultCurrencyCodeCNY
)

var ErrInvalidInfoResponseResultCurrencyCode = fmt.Errorf("not a valid InfoResponseResultCurrencyCode, try [%s]", strings.Join(_InfoResponseResultCurrencyCodeNames, ", "))

const _InfoResponseResultCurrencyCodeName = "RUBBYNKZTEURUSDCNY"

var _InfoResponseResultCurrencyCodeNames = []string{
	_InfoResponseResultCurrencyCodeName[0:3],
	_InfoResponseResultCurrencyCodeName[3:6],
	_InfoResponseResultCurrencyCodeName[6:9],
	_InfoResponseResultCurrencyCodeName[9:12],
	_InfoResponseResultCurrencyCodeName[12:15],
	_InfoResponseResultCurrencyCodeName[15:18],
}

// InfoResponseResultCurrencyCodeNames returns a list of possible string values of InfoResponseResultCurrencyCode.
func InfoResponseResultCurrencyCodeNames() []string {
	tmp := make([]string, len(_InfoResponseResultCurrencyCodeNames))
	copy(tmp, _InfoResponseResultCurrencyCodeNames)
	return tmp
}

var _InfoResponseResultCurrencyCodeMap = map[InfoResponseResultCurrencyCode]string{
	InfoResponseResultCurrencyCodeRUB: _InfoResponseResultCurrencyCodeName[0:3],
	InfoResponseResultCurrencyCodeBYN: _InfoResponseResultCurrencyCodeName[3:6],
	InfoResponseResultCurrencyCodeKZT: _InfoResponseResultCurrencyCodeName[6:9],
	InfoResponseResultCurrencyCodeEUR: _InfoResponseResultCurrencyCodeName[9:12],
	InfoResponseResultCurrencyCodeUSD: _InfoResponseResultCurrencyCodeName[12:15],
	InfoResponseResultCurrencyCodeCNY: _InfoResponseResultCurrencyCodeName[15:18],
}

// String implements the Stringer interface.
func (x InfoResponseResultCurrencyCode) String() string {
	if str, ok := _InfoResponseResultCurrencyCodeMap[x]; ok {
		return str
	}
	return fmt.Sprintf("InfoResponseResultCurrencyCode(%d)", x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x InfoResponseResultCurrencyCode) IsValid() bool {
	_, ok := _InfoResponseResultCurrencyCodeMap[x]
	return ok
}

var _InfoResponseResultCurrencyCodeValue = map[string]InfoResponseResultCurrencyCode{
	_InfoResponseResultCurrencyCodeName[0:3]:   InfoResponseResultCurrencyCodeRUB,
	_InfoResponseResultCurrencyCodeName[3:6]:   InfoResponseResultCurrencyCodeBYN,
	_InfoResponseResultCurrencyCodeName[6:9]:   InfoResponseResultCurrencyCodeKZT,
	_InfoResponseResultCurrencyCodeName[9:12]:  InfoResponseResultCurrencyCodeEUR,
	_InfoResponseResultCurrencyCodeName[12:15]: InfoResponseResultCurrencyCodeUSD,
	_InfoResponseResultCurrencyCodeName[15:18]: InfoResponseResultCurrencyCodeCNY,
}

// ParseInfoResponseResultCurrencyCode attempts to convert a string to a InfoResponseResultCurrencyCode.
func ParseInfoResponseResultCurrencyCode(name string) (InfoResponseResultCurrencyCode, error) {
	if x, ok := _InfoResponseResultCurrencyCodeValue[name]; ok {
		return x, nil
	}
	return InfoResponseResultCurrencyCode(0), fmt.Errorf("%s is %w", name, ErrInvalidInfoResponseResultCurrencyCode)
}

// MarshalText implements the text marshaller method.
func (x InfoResponseResultCurrencyCode) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *InfoResponseResultCurrencyCode) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseInfoResponseResultCurrencyCode(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

const (
	// InfoResponseResultPriceIndexesPriceIndexWITHOUTINDEX is a InfoResponseResultPriceIndexesPriceIndex of type WITHOUT_INDEX.
	InfoResponseResultPriceIndexesPriceIndexWITHOUTINDEX InfoResponseResultPriceIndexesPriceIndex = iota + 1
	// InfoResponseResultPriceIndexesPriceIndexPROFIT is a InfoResponseResultPriceIndexesPriceIndex of type PROFIT.
	InfoResponseResultPriceIndexesPriceIndexPROFIT
	// InfoResponseResultPriceIndexesPriceIndexAVGPROFIT is a InfoResponseResultPriceIndexesPriceIndex of type AVG_PROFIT.
	InfoResponseResultPriceIndexesPriceIndexAVGPROFIT
	// InfoResponseResultPriceIndexesPriceIndexNONPROFIT is a InfoResponseResultPriceIndexesPriceIndex of type NON_PROFIT.
	InfoResponseResultPriceIndexesPriceIndexNONPROFIT
)

var ErrInvalidInfoResponseResultPriceIndexesPriceIndex = fmt.Errorf("not a valid InfoResponseResultPriceIndexesPriceIndex, try [%s]", strings.Join(_InfoResponseResultPriceIndexesPriceIndexNames, ", "))

const _InfoResponseResultPriceIndexesPriceIndexName = "WITHOUT_INDEXPROFITAVG_PROFITNON_PROFIT"

var _InfoResponseResultPriceIndexesPriceIndexNames = []string{
	_InfoResponseResultPriceIndexesPriceIndexName[0:13],
	_InfoResponseResultPriceIndexesPriceIndexName[13:19],
	_InfoResponseResultPriceIndexesPriceIndexName[19:29],
	_InfoResponseResultPriceIndexesPriceIndexName[29:39],
}

// InfoResponseResultPriceIndexesPriceIndexNames returns a list of possible string values of InfoResponseResultPriceIndexesPriceIndex.
func InfoResponseResultPriceIndexesPriceIndexNames() []string {
	tmp := make([]string, len(_InfoResponseResultPriceIndexesPriceIndexNames))
	copy(tmp, _InfoResponseResultPriceIndexesPriceIndexNames)
	return tmp
}

var _InfoResponseResultPriceIndexesPriceIndexMap = map[InfoResponseResultPriceIndexesPriceIndex]string{
	InfoResponseResultPriceIndexesPriceIndexWITHOUTINDEX: _InfoResponseResultPriceIndexesPriceIndexName[0:13],
	InfoResponseResultPriceIndexesPriceIndexPROFIT:       _InfoResponseResultPriceIndexesPriceIndexName[13:19],
	InfoResponseResultPriceIndexesPriceIndexAVGPROFIT:    _InfoResponseResultPriceIndexesPriceIndexName[19:29],
	InfoResponseResultPriceIndexesPriceIndexNONPROFIT:    _InfoResponseResultPriceIndexesPriceIndexName[29:39],
}

// String implements the Stringer interface.
func (x InfoResponseResultPriceIndexesPriceIndex) String() string {
	if str, ok := _InfoResponseResultPriceIndexesPriceIndexMap[x]; ok {
		return str
	}
	return fmt.Sprintf("InfoResponseResultPriceIndexesPriceIndex(%d)", x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x InfoResponseResultPriceIndexesPriceIndex) IsValid() bool {
	_, ok := _InfoResponseResultPriceIndexesPriceIndexMap[x]
	return ok
}

var _InfoResponseResultPriceIndexesPriceIndexValue = map[string]InfoResponseResultPriceIndexesPriceIndex{
	_InfoResponseResultPriceIndexesPriceIndexName[0:13]:  InfoResponseResultPriceIndexesPriceIndexWITHOUTINDEX,
	_InfoResponseResultPriceIndexesPriceIndexName[13:19]: InfoResponseResultPriceIndexesPriceIndexPROFIT,
	_InfoResponseResultPriceIndexesPriceIndexName[19:29]: InfoResponseResultPriceIndexesPriceIndexAVGPROFIT,
	_InfoResponseResultPriceIndexesPriceIndexName[29:39]: InfoResponseResultPriceIndexesPriceIndexNONPROFIT,
}

// ParseInfoResponseResultPriceIndexesPriceIndex attempts to convert a string to a InfoResponseResultPriceIndexesPriceIndex.
func ParseInfoResponseResultPriceIndexesPriceIndex(name string) (InfoResponseResultPriceIndexesPriceIndex, error) {
	if x, ok := _InfoResponseResultPriceIndexesPriceIndexValue[name]; ok {
		return x, nil
	}
	return InfoResponseResultPriceIndexesPriceIndex(0), fmt.Errorf("%s is %w", name, ErrInvalidInfoResponseResultPriceIndexesPriceIndex)
}

// MarshalText implements the text marshaller method.
func (x InfoResponseResultPriceIndexesPriceIndex) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *InfoResponseResultPriceIndexesPriceIndex) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseInfoResponseResultPriceIndexesPriceIndex(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

const (
	// ListRequestFilterVisibilityALL is a ListRequestFilterVisibility of type ALL.
	ListRequestFilterVisibilityALL ListRequestFilterVisibility = iota
	// ListRequestFilterVisibilityVISIBLE is a ListRequestFilterVisibility of type VISIBLE.
	ListRequestFilterVisibilityVISIBLE
	// ListRequestFilterVisibilityINVISIBLE is a ListRequestFilterVisibility of type INVISIBLE.
	ListRequestFilterVisibilityINVISIBLE
	// ListRequestFilterVisibilityEMPTYSTOCK is a ListRequestFilterVisibility of type EMPTY_STOCK.
	ListRequestFilterVisibilityEMPTYSTOCK
	// ListRequestFilterVisibilityNOTMODERATED is a ListRequestFilterVisibility of type NOT_MODERATED.
	ListRequestFilterVisibilityNOTMODERATED
	// ListRequestFilterVisibilityMODERATED is a ListRequestFilterVisibility of type MODERATED.
	ListRequestFilterVisibilityMODERATED
	// ListRequestFilterVisibilityDISABLED is a ListRequestFilterVisibility of type DISABLED.
	ListRequestFilterVisibilityDISABLED
	// ListRequestFilterVisibilitySTATEFAILED is a ListRequestFilterVisibility of type STATE_FAILED.
	ListRequestFilterVisibilitySTATEFAILED
	// ListRequestFilterVisibilityREADYTOSUPPLY is a ListRequestFilterVisibility of type READY_TO_SUPPLY.
	ListRequestFilterVisibilityREADYTOSUPPLY
	// ListRequestFilterVisibilityVALIDATIONSTATEPENDING is a ListRequestFilterVisibility of type VALIDATION_STATE_PENDING.
	ListRequestFilterVisibilityVALIDATIONSTATEPENDING
	// ListRequestFilterVisibilityVALIDATIONSTATEFAIL is a ListRequestFilterVisibility of type VALIDATION_STATE_FAIL.
	ListRequestFilterVisibilityVALIDATIONSTATEFAIL
	// ListRequestFilterVisibilityVALIDATIONSTATESUCCESS is a ListRequestFilterVisibility of type VALIDATION_STATE_SUCCESS.
	ListRequestFilterVisibilityVALIDATIONSTATESUCCESS
	// ListRequestFilterVisibilityTOSUPPLY is a ListRequestFilterVisibility of type TO_SUPPLY.
	ListRequestFilterVisibilityTOSUPPLY
	// ListRequestFilterVisibilityINSALE is a ListRequestFilterVisibility of type IN_SALE.
	ListRequestFilterVisibilityINSALE
	// ListRequestFilterVisibilityREMOVEDFROMSALE is a ListRequestFilterVisibility of type REMOVED_FROM_SALE.
	ListRequestFilterVisibilityREMOVEDFROMSALE
	// ListRequestFilterVisibilityBANNED is a ListRequestFilterVisibility of type BANNED.
	ListRequestFilterVisibilityBANNED
	// ListRequestFilterVisibilityOVERPRICED is a ListRequestFilterVisibility of type OVERPRICED.
	ListRequestFilterVisibilityOVERPRICED
	// ListRequestFilterVisibilityCRITICALLYOVERPRICED is a ListRequestFilterVisibility of type CRITICALLY_OVERPRICED.
	ListRequestFilterVisibilityCRITICALLYOVERPRICED
	// ListRequestFilterVisibilityEMPTYBARCODE is a ListRequestFilterVisibility of type EMPTY_BARCODE.
	ListRequestFilterVisibilityEMPTYBARCODE
	// ListRequestFilterVisibilityBARCODEEXISTS is a ListRequestFilterVisibility of type BARCODE_EXISTS.
	ListRequestFilterVisibilityBARCODEEXISTS
	// ListRequestFilterVisibilityQUARANTINE is a ListRequestFilterVisibility of type QUARANTINE.
	ListRequestFilterVisibilityQUARANTINE
	// ListRequestFilterVisibilityARCHIVED is a ListRequestFilterVisibility of type ARCHIVED.
	ListRequestFilterVisibilityARCHIVED
	// ListRequestFilterVisibilityOVERPRICEDWITHSTOCK is a ListRequestFilterVisibility of type OVERPRICED_WITH_STOCK.
	ListRequestFilterVisibilityOVERPRICEDWITHSTOCK
	// ListRequestFilterVisibilityPARTIALAPPROVED is a ListRequestFilterVisibility of type PARTIAL_APPROVED.
	ListRequestFilterVisibilityPARTIALAPPROVED
	// ListRequestFilterVisibilityIMAGEABSENT is a ListRequestFilterVisibility of type IMAGE_ABSENT.
	ListRequestFilterVisibilityIMAGEABSENT
	// ListRequestFilterVisibilityMODERATIONBLOCK is a ListRequestFilterVisibility of type MODERATION_BLOCK.
	ListRequestFilterVisibilityMODERATIONBLOCK
)

var ErrInvalidListRequestFilterVisibility = fmt.Errorf("not a valid ListRequestFilterVisibility, try [%s]", strings.Join(_ListRequestFilterVisibilityNames, ", "))

const _ListRequestFilterVisibilityName = "ALLVISIBLEINVISIBLEEMPTY_STOCKNOT_MODERATEDMODERATEDDISABLEDSTATE_FAILEDREADY_TO_SUPPLYVALIDATION_STATE_PENDINGVALIDATION_STATE_FAILVALIDATION_STATE_SUCCESSTO_SUPPLYIN_SALEREMOVED_FROM_SALEBANNEDOVERPRICEDCRITICALLY_OVERPRICEDEMPTY_BARCODEBARCODE_EXISTSQUARANTINEARCHIVEDOVERPRICED_WITH_STOCKPARTIAL_APPROVEDIMAGE_ABSENTMODERATION_BLOCK"

var _ListRequestFilterVisibilityNames = []string{
	_ListRequestFilterVisibilityName[0:3],
	_ListRequestFilterVisibilityName[3:10],
	_ListRequestFilterVisibilityName[10:19],
	_ListRequestFilterVisibilityName[19:30],
	_ListRequestFilterVisibilityName[30:43],
	_ListRequestFilterVisibilityName[43:52],
	_ListRequestFilterVisibilityName[52:60],
	_ListRequestFilterVisibilityName[60:72],
	_ListRequestFilterVisibilityName[72:87],
	_ListRequestFilterVisibilityName[87:111],
	_ListRequestFilterVisibilityName[111:132],
	_ListRequestFilterVisibilityName[132:156],
	_ListRequestFilterVisibilityName[156:165],
	_ListRequestFilterVisibilityName[165:172],
	_ListRequestFilterVisibilityName[172:189],
	_ListRequestFilterVisibilityName[189:195],
	_ListRequestFilterVisibilityName[195:205],
	_ListRequestFilterVisibilityName[205:226],
	_ListRequestFilterVisibilityName[226:239],
	_ListRequestFilterVisibilityName[239:253],
	_ListRequestFilterVisibilityName[253:263],
	_ListRequestFilterVisibilityName[263:271],
	_ListRequestFilterVisibilityName[271:292],
	_ListRequestFilterVisibilityName[292:308],
	_ListRequestFilterVisibilityName[308:320],
	_ListRequestFilterVisibilityName[320:336],
}

// ListRequestFilterVisibilityNames returns a list of possible string values of ListRequestFilterVisibility.
func ListRequestFilterVisibilityNames() []string {
	tmp := make([]string, len(_ListRequestFilterVisibilityNames))
	copy(tmp, _ListRequestFilterVisibilityNames)
	return tmp
}

var _ListRequestFilterVisibilityMap = map[ListRequestFilterVisibility]string{
	ListRequestFilterVisibilityALL:                    _ListRequestFilterVisibilityName[0:3],
	ListRequestFilterVisibilityVISIBLE:                _ListRequestFilterVisibilityName[3:10],
	ListRequestFilterVisibilityINVISIBLE:              _ListRequestFilterVisibilityName[10:19],
	ListRequestFilterVisibilityEMPTYSTOCK:             _ListRequestFilterVisibilityName[19:30],
	ListRequestFilterVisibilityNOTMODERATED:           _ListRequestFilterVisibilityName[30:43],
	ListRequestFilterVisibilityMODERATED:              _ListRequestFilterVisibilityName[43:52],
	ListRequestFilterVisibilityDISABLED:               _ListRequestFilterVisibilityName[52:60],
	ListRequestFilterVisibilitySTATEFAILED:            _ListRequestFilterVisibilityName[60:72],
	ListRequestFilterVisibilityREADYTOSUPPLY:          _ListRequestFilterVisibilityName[72:87],
	ListRequestFilterVisibilityVALIDATIONSTATEPENDING: _ListRequestFilterVisibilityName[87:111],
	ListRequestFilterVisibilityVALIDATIONSTATEFAIL:    _ListRequestFilterVisibilityName[111:132],
	ListRequestFilterVisibilityVALIDATIONSTATESUCCESS: _ListRequestFilterVisibilityName[132:156],
	ListRequestFilterVisibilityTOSUPPLY:               _ListRequestFilterVisibilityName[156:165],
	ListRequestFilterVisibilityINSALE:                 _ListRequestFilterVisibilityName[165:172],
	ListRequestFilterVisibilityREMOVEDFROMSALE:        _ListRequestFilterVisibilityName[172:189],
	ListRequestFilterVisibilityBANNED:                 _ListRequestFilterVisibilityName[189:195],
	ListRequestFilterVisibilityOVERPRICED:             _ListRequestFilterVisibilityName[195:205],
	ListRequestFilterVisibilityCRITICALLYOVERPRICED:   _ListRequestFilterVisibilityName[205:226],
	ListRequestFilterVisibilityEMPTYBARCODE:           _ListRequestFilterVisibilityName[226:239],
	ListRequestFilterVisibilityBARCODEEXISTS:          _ListRequestFilterVisibilityName[239:253],
	ListRequestFilterVisibilityQUARANTINE:             _ListRequestFilterVisibilityName[253:263],
	ListRequestFilterVisibilityARCHIVED:               _ListRequestFilterVisibilityName[263:271],
	ListRequestFilterVisibilityOVERPRICEDWITHSTOCK:    _ListRequestFilterVisibilityName[271:292],
	ListRequestFilterVisibilityPARTIALAPPROVED:        _ListRequestFilterVisibilityName[292:308],
	ListRequestFilterVisibilityIMAGEABSENT:            _ListRequestFilterVisibilityName[308:320],
	ListRequestFilterVisibilityMODERATIONBLOCK:        _ListRequestFilterVisibilityName[320:336],
}

// String implements the Stringer interface.
func (x ListRequestFilterVisibility) String() string {
	if str, ok := _ListRequestFilterVisibilityMap[x]; ok {
		return str
	}
	return fmt.Sprintf("ListRequestFilterVisibility(%d)", x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x ListRequestFilterVisibility) IsValid() bool {
	_, ok := _ListRequestFilterVisibilityMap[x]
	return ok
}

var _ListRequestFilterVisibilityValue = map[string]ListRequestFilterVisibility{
	_ListRequestFilterVisibilityName[0:3]:     ListRequestFilterVisibilityALL,
	_ListRequestFilterVisibilityName[3:10]:    ListRequestFilterVisibilityVISIBLE,
	_ListRequestFilterVisibilityName[10:19]:   ListRequestFilterVisibilityINVISIBLE,
	_ListRequestFilterVisibilityName[19:30]:   ListRequestFilterVisibilityEMPTYSTOCK,
	_ListRequestFilterVisibilityName[30:43]:   ListRequestFilterVisibilityNOTMODERATED,
	_ListRequestFilterVisibilityName[43:52]:   ListRequestFilterVisibilityMODERATED,
	_ListRequestFilterVisibilityName[52:60]:   ListRequestFilterVisibilityDISABLED,
	_ListRequestFilterVisibilityName[60:72]:   ListRequestFilterVisibilitySTATEFAILED,
	_ListRequestFilterVisibilityName[72:87]:   ListRequestFilterVisibilityREADYTOSUPPLY,
	_ListRequestFilterVisibilityName[87:111]:  ListRequestFilterVisibilityVALIDATIONSTATEPENDING,
	_ListRequestFilterVisibilityName[111:132]: ListRequestFilterVisibilityVALIDATIONSTATEFAIL,
	_ListRequestFilterVisibilityName[132:156]: ListRequestFilterVisibilityVALIDATIONSTATESUCCESS,
	_ListRequestFilterVisibilityName[156:165]: ListRequestFilterVisibilityTOSUPPLY,
	_ListRequestFilterVisibilityName[165:172]: ListRequestFilterVisibilityINSALE,
	_ListRequestFilterVisibilityName[172:189]: ListRequestFilterVisibilityREMOVEDFROMSALE,
	_ListRequestFilterVisibilityName[189:195]: ListRequestFilterVisibilityBANNED,
	_ListRequestFilterVisibilityName[195:205]: ListRequestFilterVisibilityOVERPRICED,
	_ListRequestFilterVisibilityName[205:226]: ListRequestFilterVisibilityCRITICALLYOVERPRICED,
	_ListRequestFilterVisibilityName[226:239]: ListRequestFilterVisibilityEMPTYBARCODE,
	_ListRequestFilterVisibilityName[239:253]: ListRequestFilterVisibilityBARCODEEXISTS,
	_ListRequestFilterVisibilityName[253:263]: ListRequestFilterVisibilityQUARANTINE,
	_ListRequestFilterVisibilityName[263:271]: ListRequestFilterVisibilityARCHIVED,
	_ListRequestFilterVisibilityName[271:292]: ListRequestFilterVisibilityOVERPRICEDWITHSTOCK,
	_ListRequestFilterVisibilityName[292:308]: ListRequestFilterVisibilityPARTIALAPPROVED,
	_ListRequestFilterVisibilityName[308:320]: ListRequestFilterVisibilityIMAGEABSENT,
	_ListRequestFilterVisibilityName[320:336]: ListRequestFilterVisibilityMODERATIONBLOCK,
}

// ParseListRequestFilterVisibility attempts to convert a string to a ListRequestFilterVisibility.
func ParseListRequestFilterVisibility(name string) (ListRequestFilterVisibility, error) {
	if x, ok := _ListRequestFilterVisibilityValue[name]; ok {
		return x, nil
	}
	return ListRequestFilterVisibility(0), fmt.Errorf("%s is %w", name, ErrInvalidListRequestFilterVisibility)
}

// MarshalText implements the text marshaller method.
func (x ListRequestFilterVisibility) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *ListRequestFilterVisibility) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseListRequestFilterVisibility(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}
