package tutorials

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"gopl.io/ch1/proxy/data"
	"net/url"
)

func parseURL(urlStr string) *url.URL {
	link, err := url.Parse(urlStr)
	if err != nil {
		fyne.LogError("Could not parse URL", err)
	}

	return link
}

func welcomeScreen(_ fyne.Window) fyne.CanvasObject {
	logo := canvas.NewImageFromResource(data.FyneScene)
	logo.FillMode = canvas.ImageFillContain
	if fyne.CurrentDevice().IsMobile() {
		logo.SetMinSize(fyne.NewSize(540, 480))
	} else {
		logo.SetMinSize(fyne.NewSize(540, 480))
	}
	//data.BundleFile("fynescenedark", "dark.jpg")
	//data.BundleFile("fynescenelight", "light.jpg")
	runName := "Server Provider: " + data.RunName

	return container.NewCenter(container.NewVBox(
		widget.NewLabelWithStyle(runName, fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		logo,
	))
}

func EnpytT(i int) *widget.TextGrid {
	title := &widget.TextGrid{}
	title.SetRowStyle(i,nil)
	return title
}

var TestT string

func webBuild(win fyne.Window) fyne.CanvasObject {

	title := &widget.TextGrid{}
	title.SetText("If First Use This App Or Server Computer Was Restarted")

	title2 := &widget.TextGrid{}
	title2.SetText("Your project Dir:" + data.SourceFolder)

	var selectList []string
	selectList = append(selectList, "10.10.0.123")
	//for _, v := range data.CashData.ProviderList {
	//	selectList = append(selectList, v.Name)
	//}
	selectV := widget.NewSelect(selectList,
		func(s string) {

		},
	)
	selectV.Selected = "10.10.0.123"
	return container.NewVBox(
		selectV,
		widget.NewButton("Folder Open", func() {
			dialog.ShowFolderOpen(func(list fyne.ListableURI, err error) {
				if err != nil {
					dialog.ShowError(err, win)
					return
				}
				if list == nil {
					return
				}

				//children, err := list.List()
				_, err = list.List()
				if err != nil {
					dialog.ShowError(err, win)
					return
				}
				data.SourceFolder = list.String()[7:]
				title2.SetText("Your project Dir:" + data.SourceFolder)
				//out := fmt.Sprintf("Folder %s (%d children):\n%s", list.Name(), len(children), list.String())
				//title2.SetText()
				//dialog.ShowInformation("Folder Open", out, win)
			}, win)
		}),
		title2,
		widget.NewButton("Run Publish",
			func() {
				data.CashData.SourceFolder = data.SourceFolder
				//data.SourceFolder = title2.Text()
				data.SaveDefault()
				data.MoveFile(win)
			},
		),
		EnpytT(1),
		title,
		widget.NewButton("Connect Teleservice",
			func() {
				data.FirstConnection()
			},
		),
	)
}
