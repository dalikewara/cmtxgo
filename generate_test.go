package cmtxgo_test

import (
	"fmt"
	"github.com/dalikewara/cmtxgo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAttribute_Order(t *testing.T) {
	cmtx := cmtxgo.NewCmtx()
	cmtx.SetHeader(&cmtxgo.Field{
		"first": &cmtxgo.Attribute{
			Order:  1,
			Type:   "lps",
			Length: 5,
			Value:  "First",
		},
		"third": &cmtxgo.Attribute{
			Order:  3,
			Type:   "lps",
			Length: 5,
			Value:  "Third",
		},
		"second": &cmtxgo.Attribute{
			Order:  2,
			Type:   "lps",
			Length: 6,
			Value:  "Second",
		},
	})
	cmtxStr := cmtx.Generate()
	assert.Equal(t, "FirstSecondThird", cmtxStr)
}

func TestAttribute_Type(t *testing.T) {
	cmtx := cmtxgo.NewCmtx()
	t.Run("Left Padding", func(t *testing.T) {
		cmtx.SetHeader(&cmtxgo.Field{
			"name": &cmtxgo.Attribute{
				Order:  1,
				Type:   "lps",
				Length: 10,
				Value:  "Company",
			},
		})
		cmtxStr := cmtx.Generate()
		assert.Equal(t, "   Company", cmtxStr)

		cmtx.SetHeader(&cmtxgo.Field{
			"record": &cmtxgo.Attribute{
				Order:  1,
				Type:   "lpz",
				Length: 10,
				Value:  "1",
			},
		})
		cmtxStr = cmtx.Generate()
		assert.Equal(t, "0000000001", cmtxStr)

		cmtx.SetHeader(&cmtxgo.Field{
			"name": &cmtxgo.Attribute{
				Order:  1,
				Type:   "lps",
				Length: 10,
				Value:  "Company",
			},
			"record": &cmtxgo.Attribute{
				Order:  2,
				Type:   "lpz",
				Length: 10,
				Value:  "1",
			},
		})
		cmtxStr = cmtx.Generate()
		assert.Equal(t, "   Company0000000001", cmtxStr)

		cmtx.SetHeader(&cmtxgo.Field{
			"name": &cmtxgo.Attribute{
				Order:  1,
				Type:   "lps",
				Length: 4,
				Value:  "Company",
			},
		})
		cmtxStr = cmtx.Generate()
		assert.Equal(t, "Comp", cmtxStr)

		cmtx.SetHeader(&cmtxgo.Field{
			"record": &cmtxgo.Attribute{
				Order:  1,
				Type:   "lpz",
				Length: 4,
				Value:  "12345",
			},
		})
		cmtxStr = cmtx.Generate()
		assert.Equal(t, "1234", cmtxStr)

		cmtx.SetHeader(&cmtxgo.Field{
			"name": &cmtxgo.Attribute{
				Order:  1,
				Type:   "lps",
				Length: 7,
				Value:  "Company",
			},
		})
		cmtxStr = cmtx.Generate()
		assert.Equal(t, "Company", cmtxStr)

		cmtx.SetHeader(&cmtxgo.Field{
			"record": &cmtxgo.Attribute{
				Order:  1,
				Type:   "lpz",
				Length: 5,
				Value:  "12345",
			},
		})
		cmtxStr = cmtx.Generate()
		assert.Equal(t, "12345", cmtxStr)
	})

	t.Run("Right Padding", func(t *testing.T) {
		cmtx.SetHeader(&cmtxgo.Field{
			"name": &cmtxgo.Attribute{
				Order:  1,
				Type:   "rps",
				Length: 10,
				Value:  "Company",
			},
		})
		cmtxStr := cmtx.Generate()
		assert.Equal(t, "Company   ", cmtxStr)

		cmtx.SetHeader(&cmtxgo.Field{
			"record": &cmtxgo.Attribute{
				Order:  1,
				Type:   "rpz",
				Length: 10,
				Value:  "1",
			},
		})
		cmtxStr = cmtx.Generate()
		assert.Equal(t, "1000000000", cmtxStr)

		cmtx.SetHeader(&cmtxgo.Field{
			"name": &cmtxgo.Attribute{
				Order:  1,
				Type:   "rps",
				Length: 10,
				Value:  "Company",
			},
			"record": &cmtxgo.Attribute{
				Order:  2,
				Type:   "rpz",
				Length: 10,
				Value:  "1",
			},
		})
		cmtxStr = cmtx.Generate()
		assert.Equal(t, "Company   1000000000", cmtxStr)

		cmtx.SetHeader(&cmtxgo.Field{
			"name": &cmtxgo.Attribute{
				Order:  1,
				Type:   "rps",
				Length: 4,
				Value:  "Company",
			},
		})
		cmtxStr = cmtx.Generate()
		assert.Equal(t, "Comp", cmtxStr)

		cmtx.SetHeader(&cmtxgo.Field{
			"record": &cmtxgo.Attribute{
				Order:  1,
				Type:   "rpz",
				Length: 4,
				Value:  "12345",
			},
		})
		cmtxStr = cmtx.Generate()
		assert.Equal(t, "1234", cmtxStr)

		cmtx.SetHeader(&cmtxgo.Field{
			"name": &cmtxgo.Attribute{
				Order:  1,
				Type:   "rps",
				Length: 7,
				Value:  "Company",
			},
		})
		cmtxStr = cmtx.Generate()
		assert.Equal(t, "Company", cmtxStr)

		cmtx.SetHeader(&cmtxgo.Field{
			"record": &cmtxgo.Attribute{
				Order:  1,
				Type:   "rpz",
				Length: 5,
				Value:  "12345",
			},
		})
		cmtxStr = cmtx.Generate()
		assert.Equal(t, "12345", cmtxStr)
	})
}

func TestAttribute_Length(t *testing.T) {
	cmtx := cmtxgo.NewCmtx()
	cmtx.SetHeader(&cmtxgo.Field{
		"name": &cmtxgo.Attribute{
			Order:  1,
			Type:   "lps",
			Length: 4,
			Value:  "Company",
		},
	})
	cmtxStr := cmtx.Generate()
	assert.Equal(t, 4, len(cmtxStr))
	assert.Equal(t, "Comp", cmtxStr)

	cmtx.SetHeader(&cmtxgo.Field{
		"name": &cmtxgo.Attribute{
			Order:  1,
			Type:   "lpz",
			Length: 6,
			Value:  "1234567",
		},
	})
	cmtxStr = cmtx.Generate()
	assert.Equal(t, 6, len(cmtxStr))
	assert.Equal(t, "123456", cmtxStr)

	cmtx.SetHeader(&cmtxgo.Field{
		"name": &cmtxgo.Attribute{
			Order:  1,
			Type:   "lps",
			Length: 7,
			Value:  "Company",
		},
	})
	cmtxStr = cmtx.Generate()
	assert.Equal(t, 7, len(cmtxStr))
	assert.Equal(t, "Company", cmtxStr)

	cmtx.SetHeader(&cmtxgo.Field{
		"name": &cmtxgo.Attribute{
			Order:  1,
			Type:   "lpz",
			Length: 7,
			Value:  "1234567",
		},
	})
	cmtxStr = cmtx.Generate()
	assert.Equal(t, 7, len(cmtxStr))
	assert.Equal(t, "1234567", cmtxStr)
}

func TestAttribute_Value(t *testing.T) {
	cmtx := cmtxgo.NewCmtx()
	cmtx.SetHeader(&cmtxgo.Field{
		"name": &cmtxgo.Attribute{
			Order:  1,
			Type:   "lps",
			Length: 7,
			Value:  "Company",
		},
		"record": &cmtxgo.Attribute{
			Order:  2,
			Type:   "lpz",
			Length: 10,
			Value:  "1",
		},
	})
	cmtxStr := cmtx.Generate()
	assert.Equal(t, "Company0000000001", cmtxStr)
}

func TestAttribute_DefaultValue(t *testing.T) {
	cmtx := cmtxgo.NewCmtx()
	cmtx.SetHeader(&cmtxgo.Field{
		"name": &cmtxgo.Attribute{
			Order:        1,
			Type:         "lps",
			Length:       7,
			DefaultValue: "Company",
		},
		"record": &cmtxgo.Attribute{
			Order:        2,
			Type:         "lpz",
			Length:       10,
			Value:        "2",
			DefaultValue: "1",
		},
		"branch": &cmtxgo.Attribute{
			Order:        3,
			Type:         "lpz",
			Length:       10,
			Value:        "",
			DefaultValue: "3",
		},
	})
	cmtxStr := cmtx.Generate()
	assert.Equal(t, "Company00000000020000000003", cmtxStr)
}

func TestAttribute_RemoveAllChars(t *testing.T) {
	cmtx := cmtxgo.NewCmtx()
	cmtx.SetFooter(&cmtxgo.Field{
		"total_amount": &cmtxgo.Attribute{
			Order:          1,
			Type:           "lpz",
			Length:         11,
			Value:          "1.000.000.000",
			RemoveAllChars: ".",
		},
	})
	cmtxStr := cmtx.Generate()
	assert.Equal(t, "01000000000", cmtxStr)

	cmtx.SetFooter(&cmtxgo.Field{
		"total_amount": &cmtxgo.Attribute{
			Order:          1,
			Type:           "lpz",
			Length:         12,
			Value:          "1000000,000",
			RemoveAllChars: ".",
		},
	})
	cmtxStr = cmtx.Generate()
	assert.Equal(t, "01000000,000", cmtxStr)
}

func TestAttribute_ReplaceAllChars(t *testing.T) {
	cmtx := cmtxgo.NewCmtx()
	cmtx.SetFooter(&cmtxgo.Field{
		"total_amount": &cmtxgo.Attribute{
			Order:           1,
			Type:            "lpz",
			Length:          11,
			Value:           "1.000.000.000",
			ReplaceAllChars: [2]string{".", ","},
		},
	})
	cmtxStr := cmtx.Generate()
	assert.Equal(t, "1,000,000,0", cmtxStr)

	cmtx.SetFooter(&cmtxgo.Field{
		"total_amount": &cmtxgo.Attribute{
			Order:           1,
			Type:            "lpz",
			Length:          12,
			Value:           "1000000,000",
			ReplaceAllChars: [2]string{",", "."},
		},
	})
	cmtxStr = cmtx.Generate()
	assert.Equal(t, "01000000.000", cmtxStr)

	cmtx.SetFooter(&cmtxgo.Field{
		"total_amount": &cmtxgo.Attribute{
			Order:           1,
			Type:            "lpz",
			Length:          12,
			Value:           "1000000000",
			ReplaceAllChars: [2]string{",", "."},
		},
	})
	cmtxStr = cmtx.Generate()
	assert.Equal(t, "001000000000", cmtxStr)
}

func TestOption_MaxLengthPerSection(t *testing.T) {
	cmtx := cmtxgo.NewCmtx()
	cmtx.SetFooter(&cmtxgo.Field{
		"name": &cmtxgo.Attribute{
			Order:  1,
			Type:   "lps",
			Length: 10,
			Value:  "Company",
		},
	})
	cmtx.SetOption(&cmtxgo.Option{
		MaxLengthPerSection: 20,
	})
	cmtxStr := cmtx.Generate()
	assert.Equal(t, 20, len(cmtxStr))
	assert.Equal(t, "   Company          ", cmtxStr)

	cmtx.SetFooter(&cmtxgo.Field{
		"name": &cmtxgo.Attribute{
			Order:  1,
			Type:   "lps",
			Length: 20,
			Value:  "Company",
		},
	})
	cmtx.SetOption(&cmtxgo.Option{
		MaxLengthPerSection: 20,
	})
	cmtxStr = cmtx.Generate()
	assert.Equal(t, 20, len(cmtxStr))
	assert.Equal(t, "             Company", cmtxStr)
}

func TestOption_AddCharPerSection(t *testing.T) {
	cmtx := cmtxgo.NewCmtx()
	cmtx.SetFooter(&cmtxgo.Field{
		"name": &cmtxgo.Attribute{
			Order:  1,
			Type:   "lps",
			Length: 10,
			Value:  "Company",
		},
	})
	cmtx.SetOption(&cmtxgo.Option{
		AddCharPerSection: "\n",
	})
	cmtxStr := cmtx.Generate()
	assert.Equal(t, "   Company\n", cmtxStr)

	cmtx.SetFooter(&cmtxgo.Field{
		"name": &cmtxgo.Attribute{
			Order:  1,
			Type:   "lps",
			Length: 10,
			Value:  "Company",
		},
	})
	cmtx.SetOption(&cmtxgo.Option{
		AddCharPerSection: "ADDED",
	})
	cmtxStr = cmtx.Generate()
	assert.Equal(t, "   CompanyADDED", cmtxStr)
}

func TestDetail(t *testing.T) {
	cmtx := cmtxgo.NewCmtx()
	t.Run("Plain", func(t *testing.T) {
		cmtx.SetDetail(&cmtxgo.Field{
			"name": &cmtxgo.Attribute{
				Order:  1,
				Type:   "lps",
				Length: 10,
			},
			"username": &cmtxgo.Attribute{
				Order:  2,
				Type:   "lps",
				Length: 10,
			},
		}, &[]map[string]interface{}{
			{
				"name":     "John Doe",
				"username": "john_doe",
			},
			{
				"name":     "Adam Smith",
				"username": "adam_smith",
			},
		})
		cmtxStr := cmtx.Generate()
		assert.Equal(t, "  John Doe  john_doeAdam Smithadam_smith", cmtxStr)

		cmtx.SetDetail(&cmtxgo.Field{
			"name": &cmtxgo.Attribute{
				Order:  1,
				Type:   "lps",
				Length: 10,
			},
			"username": &cmtxgo.Attribute{
				Order:  2,
				Type:   "lps",
				Length: 10,
			},
		}, &[]map[string]interface{}{
			{
				"name": "John Doe",
			},
			{
				"name":     "Dali Kewara",
				"username": "dalikewara",
			},
			{
				"name": "Adam Smith",
			},
		})
		cmtxStr = cmtx.Generate()
		assert.Equal(t, "  John Doe          Dali KewardalikewaraAdam Smith          ", cmtxStr)
	})
}

func TestExample(t *testing.T) {
	t.Run("Basic Usage", func(t *testing.T) {
		cmtx := cmtxgo.NewCmtx()
		cmtx.SetHeader(&cmtxgo.Field{
			"myHeaderFieldName": &cmtxgo.Attribute{
				Order:  1,
				Type:   "lpz",
				Length: 11,
				Value:  "999",
			},
			"myHeaderFieldName2": &cmtxgo.Attribute{
				Order:  2,
				Type:   "rps",
				Length: 35,
				Value:  "My COMPANY",
			},
		})
		detailData := []map[string]interface{}{
			{
				"myDetailFieldName":  "123-456-789",
				"myDetailFieldName2": "John Doe",
				"myDetailFieldName3": "1000000",
			},
			{
				"myDetailFieldName":  "098-765-432",
				"myDetailFieldName2": "Adam Smith",
				"myDetailFieldName3": "1005500",
			},
			{
				"myDetailFieldName":  "123-098-456",
				"myDetailFieldName2": "Dali Kewara",
				"myDetailFieldName3": "204000",
			},
		}
		cmtx.SetDetail(&cmtxgo.Field{
			"myDetailFieldName": &cmtxgo.Attribute{
				Order:  1,
				Type:   "rps",
				Length: 15,
			},
			"myDetailFieldName2": &cmtxgo.Attribute{
				Order:  2,
				Type:   "rps",
				Length: 20,
			},
			"myDetailFieldName3": &cmtxgo.Attribute{
				Order:  3,
				Type:   "lpz",
				Length: 11,
			},
		}, &detailData)
		cmtx.SetFooter(&cmtxgo.Field{
			"myFooterFieldName": &cmtxgo.Attribute{
				Order:  1,
				Type:   "lpz",
				Length: 11,
				Value:  fmt.Sprintf("%v", len(detailData)),
			},
			"myFooterFieldName2": &cmtxgo.Attribute{
				Order:  2,
				Type:   "lpz",
				Length: 35,
				Value:  "2209500",
			},
		})
		cmtxStr := cmtx.Generate()
		err := cmtx.GenerateToFile("tmp/cemtext.ctx")
		assert.Nil(t, err)
		assert.Equal(t, "00000000999My COMPANY                         123-456-789    John Doe            00001000000098-765-432    Adam Smith          00001005500123-098-456    Dali Kewara         000002040000000000000300000000000000000000000000002209500", cmtxStr)
	})
}
