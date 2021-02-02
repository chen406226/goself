// Package main provides various examples of Fyne API capabilities.
package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"gopl.io/ch1/proxy/data"
	"gopl.io/ch1/proxy/tutorials"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

const preferenceCurrentTutorial = "currentTutorial"

var topWindow fyne.Window

func shortcutFocused(s fyne.Shortcut, w fyne.Window) {
	if focused, ok := w.Canvas().Focused().(fyne.Shortcutable); ok {
		focused.TypedShortcut(s)
	}
}

// $ fyne package -os windows -icon icon.jpg
// go tool arguments:  -ldflags -H=windowsgui
func main() {
	a := app.NewWithID("io.fyne.demo")
	//a.SetIcon(theme.FyneLogo())
	iconContent, _ := base64.StdEncoding.DecodeString(data.IconBase64)
	var windowIcon = data.HardentoolsWindowIconStruct{
		NameInt:    "HardenToolsWindowIcon",
		ContentInt: iconContent,
	}
	a.SetIcon(windowIcon)
	w := a.NewWindow("发布换源工具")
	topWindow = w
	mainMenu := fyne.NewMainMenu(
		// a quit item will be appended to our first menu
		fyne.NewMenu("Set", fyne.NewMenuItemSeparator()),
	)
	w.SetMainMenu(mainMenu)
	w.SetMaster()
	w.SetIcon(windowIcon)

	content := container.NewMax()
	title := widget.NewLabel("Component name")
	intro := widget.NewLabel("An introduction would probably go\nhere, as well as a")
	intro.Wrapping = fyne.TextWrapWord
	setTutorial := func(t tutorials.Tutorial) {
		if fyne.CurrentDevice().IsMobile() {
			child := a.NewWindow(t.Title)
			topWindow = child
			child.SetContent(t.View(topWindow))
			child.Show()
			child.SetOnClosed(func() {
				topWindow = w
			})
			return
		}

		title.SetText(t.Title)
		intro.SetText(t.Intro)

		content.Objects = []fyne.CanvasObject{t.View(w)}
		content.Refresh()
	}

	tutorial := container.NewBorder(
		container.NewVBox(title, widget.NewSeparator(), intro), nil, nil, nil, content)
	if fyne.CurrentDevice().IsMobile() {
		w.SetContent(makeNav(setTutorial, false))
	} else {
		split := container.NewHSplit(makeNav(setTutorial, true), tutorial)
		split.Offset = 0.2
		w.SetContent(split)
	}
	go runProxy(w)
	w.Resize(fyne.NewSize(640, 460))
	w.ShowAndRun()
}

func Succ(w fyne.Window, e chan <- error)  {
	time.Sleep(time.Second * 5)
	dialog.ShowInformation("Server Status", "Proxy Service Started Successfully: localhost:9876", w)
}

func runProxy(w fyne.Window)  {
	fmt.Println("%%%%%%%%%%")
	var err error
	http.HandleFunc("/", handler)
	time.AfterFunc(time.Duration(time.Second * 3), func() {
		if err == nil {
			dialog.ShowInformation("Server Status", "Proxy Service Started Successfully: localhost:9876", w)
		}
	})

	//cErr := make(chan error)
	//go Succ(w, cErr)
	err = http.ListenAndServe("0.0.0.0:9876", nil)
	if err != nil {
		//cErr <- err
		dialog.ShowError(errors.New("ListenAndServe Error Check The 9876 Port Occupied"), w)
	}
	//log.Fatal(http.ListenAndServe("0.0.0.0:9876", nil))
}
func handler(w http.ResponseWriter, r *http.Request) {
	var proxyHost = "http://localhost:8088"
	for _, v := range data.CashData.ProviderList {
		if  v.Name == data.RunName {
			proxyHost = v.Host
		}
	}
	//fmt.Println(proxyHost)
	remote, err := url.Parse(proxyHost)
	if err != nil {
		panic(err)
	}
	r.Host = "localhost:8080"
	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.ServeHTTP(w, r)
}

func makeNav(setTutorial func(tutorial tutorials.Tutorial), loadPrevious bool) fyne.CanvasObject {
	a := fyne.CurrentApp()

	tree := &widget.Tree{
		ChildUIDs: func(uid string) []string {
			return tutorials.TutorialIndex[uid]
		},
		IsBranch: func(uid string) bool {
			children, ok := tutorials.TutorialIndex[uid]

			return ok && len(children) > 0
		},
		CreateNode: func(branch bool) fyne.CanvasObject {
			return widget.NewLabel("Collection Widgets")
		},
		UpdateNode: func(uid string, branch bool, obj fyne.CanvasObject) {
			t, ok := tutorials.Tutorials[uid]
			if !ok {
				fyne.LogError("Missing tutorial panel: "+uid, nil)
				return
			}
			obj.(*widget.Label).SetText(t.Title)
		},
		OnSelected: func(uid string) {
			if t, ok := tutorials.Tutorials[uid]; ok {
				a.Preferences().SetString(preferenceCurrentTutorial, uid)
				setTutorial(t)
			}
		},
	}

	if loadPrevious {
		currentPref := a.Preferences().StringWithFallback(preferenceCurrentTutorial, "welcome")
		tree.Select(currentPref)
	}

	themes := fyne.NewContainerWithLayout(layout.NewGridLayout(2),
		widget.NewButton("Dark", func() {
			a.Settings().SetTheme(theme.DarkTheme())
		}),
		widget.NewButton("Light", func() {
			a.Settings().SetTheme(theme.LightTheme())
		}),
	)

	return container.NewBorder(nil, themes, nil, nil, tree)
}
