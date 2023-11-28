package todo

import (
	"fmt"

	"github.com/alexeyco/simpletable"
)

func (tasks *Tasks) Print() {

	t := simpletable.New()

	t.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignLeft, Text: "#"},
			{Align: simpletable.AlignLeft, Text: "Text"},
			{Align: simpletable.AlignLeft, Text: "Status"},
			{Align: simpletable.AlignLeft, Text: "CreatedAt"},
			{Align: simpletable.AlignLeft, Text: "FinishedAt"},
		},
	}

	var cells [][]*simpletable.Cell

	for idx, row := range *tasks {
		idx++
		done := Red("no")
		if row.IsDone {
			done = Green("yes")
		}
		fmt.Println(row)
		cells = append(cells, *&[]*simpletable.Cell{
			{Text: fmt.Sprintf("%d", idx)},
			{Text: row.Text},
			{Text: done},
			{Text: row.CreatedTime.Local().String()},
			{Text: row.CompleteTime.Local().String()},
		})
	}
	t.Body = &simpletable.Body{Cells: cells}
	t.SetStyle(simpletable.StyleUnicode)
	fmt.Println(t)
}
