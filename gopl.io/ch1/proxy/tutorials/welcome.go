package tutorials

import (
	"fmt"
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
	fmt.Println("sdlkfj")
	if fyne.CurrentDevice().IsMobile() {
		logo.SetMinSize(fyne.NewSize(171, 125))
	} else {
		logo.SetMinSize(fyne.NewSize(228, 167))
	}


	return container.NewCenter(container.NewVBox(
		widget.NewLabelWithStyle("", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		logo,
	))
}
