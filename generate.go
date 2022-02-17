package cmtxgo

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"sync"
)

type Attribute struct {
	// Order sorts the generation of the fields. Example:
	//
	//cmtx.SetHeader(&cmtxgo.Field{
	//	"first": &cmtxgo.Attribute{
	//		Order:  2,
	//		Type:   "lps",
	//		Length: 6,
	//		Value:  "Second",
	//	},
	//	"second": &cmtxgo.Attribute{
	//		Order:  1,
	//		Type:   "lps",
	//		Length: 5,
	//		Value:  "First",
	//	},
	//})
	//
	// Output: "FirstSecond"
	Order int `json:"order"`

	// Type formats the field value based on padding type. Available padding type: `lps`, `lpz`, `rps`, `rpz`.
	// Default is `rps`.
	//
	// `lps`: Left Padding Space (ex: "  john doe")
	//
	// `lpz`: Left Padding Zero (ex: "0000000123")
	//
	// `rps`: Right Padding Space (ex: "john doe   ")
	//
	// `rpz`: Right Padding Zero (ex: "1230000000")
	Type string `json:"type"`

	// Length sets max length to the field value.
	Length int `json:"length"`

	// Value sets field value. This only effect on `header` & `footer` section.
	// For `detail` section, the field value got from the matched field name in detail data.
	// If Value is empty, DefaultValue will be used.
	Value string `json:"value"`

	// DefaultValue sets the default value to the field if the Value is empty. Default is "".
	DefaultValue string `json:"default_value"`

	// RemoveAllChars removes specified characters from the field value. Example:
	//
	//cmtx.SetFooter(&cmtxgo.Field{
	//	"total_amount": &cmtxgo.Attribute{
	//		Order:          1,
	//		Type:           "lpz",
	//		Length:         10,
	//		Value:          "1.000.000.000",
	//		RemoveAllChars: ".",
	//	},
	//})
	//
	// Output: "1000000000"
	RemoveAllChars string `json:"remove_all_chars"`

	// ReplaceAllChars replaces specified characters with the new one from the field value.
	// Example:
	//
	//cmtx.SetFooter(&cmtxgo.Field{
	//	"total_amount": &cmtxgo.Attribute{
	//		Order:           1,
	//		Type:            "lpz",
	//		Length:          13,
	//		Value:           "1.000.000.000",
	//		ReplaceAllChars: [2]string{".", ","},
	//	},
	//})
	//
	// Output: "1,000,000,000"
	ReplaceAllChars [2]string `json:"replace_all_chars"`
}

type Option struct {
	// MaxLengthPerSection sets max length on every section value.
	MaxLengthPerSection int `json:"max_length_per_section"`

	// AddCharPerSection adds specified character at the end of the value on every section.
	// Keep in mind that this will add new character, so it will break
	// the value length setting.
	AddCharPerSection string `json:"add_char_per_section"`
}

type Field map[string]*Attribute

type cmtx struct {
	header     *Field
	footer     *Field
	detail     *Field
	detailData *[]map[string]interface{}
	option     *Option
}

var orderSep = " __{{[[((0RD3R_53P))]]}}__ "

// NewCmtx generates new cmtx object.
func NewCmtx() *cmtx {
	return &cmtx{}
}

// SetHeader sets the cemtext header.
func (c *cmtx) SetHeader(field *Field) {
	c.header = field
}

// SetDetail sets the cemtext detail.
func (c *cmtx) SetDetail(field *Field, detailData *[]map[string]interface{}) {
	c.detail = field
	c.detailData = detailData
}

// SetFooter sets the cemtext footer.
func (c *cmtx) SetFooter(field *Field) {
	c.footer = field
}

// SetOption sets the generation option of each section.
func (c *cmtx) SetOption(option *Option) {
	c.option = option
}

// Generate generates cemtext format as string.
func (c *cmtx) Generate() string {
	header := ""
	detail := ""
	footer := ""

	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		defer wg.Done()
		header = generateHeader(c.header, c.option)
	}()
	go func() {
		defer wg.Done()
		detail = generateDetail(c.detail, c.detailData, c.option)
	}()
	go func() {
		defer wg.Done()
		footer = generateFooter(c.footer, c.option)
	}()

	wg.Wait()

	return fmt.Sprintf("%s%s%s", header, detail, footer)
}

// GenerateToFile generates cemtext format and save it into a file.
func (c *cmtx) GenerateToFile(output string) error {
	cmtxStr := c.Generate()
	file, err := os.Create(output)
	if err != nil {
		return err
	} else {
		_, errWriteString := file.WriteString(cmtxStr)
		if errWriteString != nil {
			return errWriteString
		}
	}
	err = file.Close()
	if err != nil {
		return err
	}
	return nil
}

func generateCommon(field *Field) string {
	if field == nil {
		return ""
	}
	var orders []string
	for key, attr := range *field {
		orders = append(orders, fmt.Sprintf("%v%s%s%s", attr.Order, key, orderSep, generateValue(attr)))
	}
	sort.Strings(orders)
	str := ""
	for _, s := range orders {
		str = str + strings.Split(s, orderSep)[1]
	}
	return str
}

func generateHeader(field *Field, option *Option) string {
	str := generateCommon(field)
	if str != "" && option != nil {
		str = formatMaxLength(str, option.MaxLengthPerSection)
		if option.AddCharPerSection != "" {
			str = addCharAtTrail(str, option.AddCharPerSection)
		}
	}
	return str
}

func generateDetail(field *Field, detailData *[]map[string]interface{}, option *Option) string {
	if field == nil || detailData == nil {
		return ""
	}
	str := ""
	for _, v := range *detailData {
		dataField := interface{}(*field).(Field)
		for n, _ := range dataField {
			dataField[n].Value = ""
			if v[n] != "" && v[n] != nil {
				dataField[n].Value = fmt.Sprintf("%v", v[n])
			}
		}
		cStr := generateCommon(field)
		if cStr != "" && option != nil {
			cStr = formatMaxLength(cStr, option.MaxLengthPerSection)
			if option.AddCharPerSection != "" {
				cStr = addCharAtTrail(cStr, option.AddCharPerSection)
			}
		}
		str = str + cStr
	}
	return str
}

func generateFooter(field *Field, option *Option) string {
	return generateHeader(field, option)
}

func generateValue(attr *Attribute) string {
	val := attr.DefaultValue
	if attr.Value != "" {
		val = attr.Value
	}
	if attr.RemoveAllChars != "" {
		val = removeAllChars(val, attr.RemoveAllChars)
	}
	if attr.ReplaceAllChars[0] != "" && attr.ReplaceAllChars[1] != "" {
		val = replaceAllChars(val, attr.ReplaceAllChars[0], attr.ReplaceAllChars[1])
	}
	return formatValue(val, attr)
}

func formatMaxLength(val string, maxLength int) string {
	if maxLength != 0 {
		return rightPad(val, " ", maxLength)
	}
	return val
}

func formatValue(val string, attr *Attribute) string {
	switch attr.Type {
	case "lps":
		return leftPad(val, " ", attr.Length)
	case "lpz":
		return leftPad(val, "0", attr.Length)
	case "rpz":
		return rightPad(val, "0", attr.Length)
	default:
		return rightPad(val, " ", attr.Length)
	}
}

func leftPad(val string, pad string, length int) string {
	for i := len(val); i < length; i++ {
		val = pad + val
	}
	return val[:length]
}

func rightPad(val string, pad string, length int) string {
	for i := len(val); i < length; i++ {
		val = val + pad
	}
	return val[:length]
}

func addCharAtTrail(val string, char string) string {
	return fmt.Sprintf("%s%s", val, char)
}

func replaceAllChars(val, old, new string) string {
	return strings.ReplaceAll(val, old, new)
}

func removeAllChars(val, char string) string {
	return strings.ReplaceAll(val, char, "")
}
