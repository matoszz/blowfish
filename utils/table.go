package utils

import (
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/tidwall/gjson"
)

func tableWriter() table.Writer {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleLight)

	t.Style().Box = table.BoxStyle{PaddingRight: "\t"}
	t.Style().Options.SeparateHeader = false
	t.Style().Options.SeparateRows = false
	t.Style().Options.DrawBorder = false

	return t
}

// RenderTableList will print out a table based on the json and json selectors passed in
func RenderTableList(rawJSON []byte, columns []string, sortColumn string) {
	t := tableWriter()

	keys := make([]interface{}, 0, len(columns))
	for _, k := range columns {
		keys = append(keys, strings.SplitN(k, ":", 2)[0]) //nolint:gomnd
	}

	t.AppendHeader(keys)

	result := gjson.GetBytes(rawJSON, "@this")
	result.ForEach(func(key, value gjson.Result) bool {
		values := make([]interface{}, 0, len(columns))
		for _, k := range columns {
			var selector string
			spl := strings.SplitN(k, ":", 2) //nolint:gomnd
			if len(spl) == 1 {
				selector = spl[0]
			} else {
				selector = spl[1]
			}

			v := gjson.Get(value.String(), selector) //nolint:gomnd
			values = append(values, formatJSONValue(v))
		}
		t.AppendRow(values)
		return true // keep iterating
	})

	t.SortBy([]table.SortBy{
		{Name: sortColumn, Mode: table.Asc},
	})

	t.Render()
}

// RenderTable will print out data from the json passed in and the rows that we want printed
func RenderTable(rawJSON []byte, rows [][]string) {
	t := tableWriter()

	parsedJSON := gjson.ParseBytes(rawJSON)

	for _, row := range rows {
		switch row[0] {
		case "-":
			t.AppendSeparator()
		default:
			l := label(row[0])
			v := parsedJSON.Get(row[1])

			t.AppendRow([]interface{}{l, formatJSONValue(v)})
		}
	}

	t.Render()
}

func label(l string) string {
	return text.Bold.Sprint(l + ":")
}
