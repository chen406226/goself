package tutorials

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func windowScreen(_ fyne.Window) fyne.CanvasObject {
	selectV := widget.NewSelect([]string{"Option 1", "Option 2", "Option 3"},
		func(s string) {
			fmt.Println("selected", s)
		},
	)
	selectV.Selected = "Option 1"
	return container.NewVBox(
		selectV,
		widget.NewButton("Run",
			func() {
				fmt.Println("tapped text button ")
				fmt.Println(selectV.Selected)
			},
		),
	)
}
