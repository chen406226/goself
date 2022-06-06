package tutorials

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"gopl.io/ch1/proxy/data"
	"image/color"
	"net/url"
)

func parseURL(urlStr string) *url.URL {
	link, err := url.Parse(urlStr)
	if err != nil {
		fyne.LogError("Could not parse URL", err)
	}

	return link
}

func welcomeScreen(win fyne.Window) fyne.CanvasObject {
	text := canvas.NewText("切换代理服务地址", nil)
	text.Alignment = fyne.TextAlignLeading
	text.TextStyle = fyne.TextStyle{Italic: true}

	//logo := canvas.NewImageFromResource(data.FyneScene)
	//logo.FillMode = canvas.ImageFillContain
	//if fyne.CurrentDevice().IsMobile() {
	//	logo.SetMinSize(fyne.NewSize(360, 320))
	//} else {
	//	logo.SetMinSize(fyne.NewSize(360, 320))
	//}
	//data.BundleFile("fynescenedark", "dark.jpg")
	//data.BundleFile("fynescenelight", "light.jpg")
	runName := "当前代理服务: " + data.RunName

	return container.NewVBox(container.NewVBox(
		widget.NewLabelWithStyle(runName, fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		text,
		windowScreen(win),
		//logo,
	))
}

func EnpytT(i int) *widget.TextGrid {
	title := &widget.TextGrid{}
	title.SetRowStyle(i,nil)
	return title
}

var TestT string

func webBuild(win fyne.Window) fyne.CanvasObject {

	title := canvas.NewText("服务连接测试", nil)

	title2 := canvas.NewText("当前项目地址:" + data.SourceFolder, nil)

	var selectList []string
	for _, v := range data.CashData.PublishList {
		selectList = append(selectList, v.Name)
	}
	selectV := widget.NewSelect(selectList,
		func(s string) {

		},
	)
	selectV.Selected = data.CashData.PubName
	infProgress := widget.NewProgressBarInfinite()
	//lb := widget.NewLabel("Please Wait ,Building...")
	lb := canvas.NewText("Please Wait Building...",color.RGBA{
		R: 255,
		G: 0,
		B: 0,
		A: 255,
	})
	lb.Hide()
	infProgress.Hide()
	return container.NewVBox(
		lb,
		infProgress,
		selectV,
		widget.NewButton("更改项目文件夹", func() {
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
				title2.Text = "当前项目地址:" + data.SourceFolder
				title2.Refresh()
				//title2.SetText("Your project Dir:" + data.SourceFolder)
				//out := fmt.Sprintf("Folder %s (%d children):\n%s", list.Name(), len(children), list.String())
				//title2.SetText()
				//dialog.ShowInformation("Folder Open", out, win)
			}, win)
		}),
		title2,
		widget.NewButton("开始发布",
			func() {
				data.CashData.SourceFolder = data.SourceFolder
				data.CashData.PubName = selectV.Selected
				data.SaveDefault()
				//data.SourceFolder = title2.Text()
				data.SaveDefault()
				lb.Text="Please Wait ,Building..."
				lb.Show()
				infProgress.Show()
				data.MoveFile(win, lb)
				lb.Hide()
				infProgress.Hide()
			},
		),
		EnpytT(1),
		title,
		widget.NewButton("连接服务测试",
			func() {
				lb.Text="请等待，发布中..."
				lb.Show()
				infProgress.Show()
				data.FirstConnection(win, lb)
				lb.Hide()
				infProgress.Hide()
			},
		),
	)
}
