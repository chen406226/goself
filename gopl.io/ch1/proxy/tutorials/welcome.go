package tutorials

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
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
