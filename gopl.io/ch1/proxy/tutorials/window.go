package tutorials

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"gopl.io/ch1/proxy/data"
)

func windowScreen(_ fyne.Window) fyne.CanvasObject {
	var selectList []string
	for _, v := range data.CashData.ProviderList {
		selectList = append(selectList, v.Name)
	}
	selectV := widget.NewSelect(selectList,
		func(s string) {

		},
	)
	selectV.Selected = data.RunName
	return container.NewVBox(
		selectV,
		widget.NewButton("Run",
			func() {
				data.RunName = selectV.Selected
				data.CashData.Default = selectV.Selected
				data.SaveDefault()
			},
		),
	)
}
