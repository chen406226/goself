package tutorials

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"gopl.io/ch1/proxy/data"
)

func windowScreen(win fyne.Window) fyne.CanvasObject {
	var ho string
	var selectList []string
	for _, v := range data.CashData.ProviderList {
		selectList = append(selectList, v.Name)
		if v.Name == data.RunName {
			ho = v.Host
		}
	}
	text := canvas.NewText("代理地址:" + ho, nil)
	text.Alignment = fyne.TextAlignLeading
	text.TextStyle = fyne.TextStyle{Italic: true}

	selectV := widget.NewSelect(selectList,
		func(s string) {
			for _, v := range data.CashData.ProviderList {
				if v.Name == s {
					ho = v.Host
				}
			}
			text.Text = "代理地址:" + ho
			text.Refresh()
		},
	)
	selectV.Selected = data.RunName
	button := widget.NewButton("Run 确认",
		func() {
			data.RunName = selectV.Selected
			data.CashData.Default = selectV.Selected
			data.SaveDefault()
			dialog.ShowInformation("Success", "切换服务成功", win)
		},
	)
	//button.Se
	return container.NewVBox(
		selectV,
		text,
		button,
	)
}
