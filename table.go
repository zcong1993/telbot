package main

import (
	"bytes"
	"github.com/olekukonko/tablewriter"
)

// CreateTableText format data to text table
func CreateTableText(data [][]string) bytes.Buffer {
	var bf bytes.Buffer
	table := tablewriter.NewWriter(&bf)
	//table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.AppendBulk(data) // Add Bulk Data
	table.Render()         // Send output
	return bf
}
