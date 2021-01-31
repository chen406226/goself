package tutorials

import (
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"gopl.io/ch1/proxy/data"
)

func bindingScreen(win fyne.Window) fyne.CanvasObject {

	//formStruct := struct {
	//	Name, Email string
	//	Subscribe   bool
	//}{}
	//
	//formData := binding.BindStruct(&formStruct)
	//form := newFormWithData(formData)
	//form.OnSubmit = func() {
	//	fmt.Println("Struct:\n", formStruct)
	//}
	name := widget.NewEntry()
	name.SetPlaceHolder("EP: DEV")

	host := widget.NewEntry()
	host.Text = "http://10.10.0.123:8088"
	host.SetPlaceHolder("http://10.10.0.123:8088")
	host.Validator = validation.NewRegexp(`^http`, "not a valid website")

	//password := widget.NewPasswordEntry()
	//password.SetPlaceHolder("Password")

	//largeText := widget.NewMultiLineEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Name", Widget: name, HintText: "server provider name"},
			{Text: "Host", Widget: host, HintText: "server host:port"},
		},
		//OnCancel: func() {
		//	fmt.Println("Cancelled")
		//},
		OnSubmit: func() {
			if name.Text == "" {
				dialog.ShowError(errors.New("Please Input Name"), win)
				return
			}
			data.CashData.ProviderList = append(data.CashData.ProviderList, data.User{Name: name.Text, Host: host.Text})
			data.SaveDefault()
			dialog.ShowInformation("Information", "Save Server Provider Succeeded", win)
			//fyne.CurrentApp().SendNotification(&fyne.Notification{
			//	Title:   "Form for: " + name.Text,
			//	Content: largeText.Text,
			//})
		},
	}
	title := &widget.TextGrid{}
	title.SetText(" Data Save In C:/Users/Administrator/proxyConfig.json")
	title1 := &widget.TextGrid{}
	title1.SetText(" dev.env.js setting localhost(ip):9876")
	title.SetRowStyle(2,nil)
	return container.NewVBox(title1,title,form)
	//return container.NewBorder(nil, nil, container.NewGridWithColumns(2, listPanel, form), nil, )
}

func newFormWithData(data binding.DataMap) *widget.Form {
	keys := data.Keys()
	items := make([]*widget.FormItem, len(keys))
	for i, k := range keys {
		data, err := data.GetItem(k)
		if err != nil {
			items[i] = widget.NewFormItem(k, widget.NewLabel(err.Error()))
		}
		items[i] = widget.NewFormItem(k, createBoundItem(data))
	}


	return widget.NewForm(items...)
}

func createBoundItem(v binding.DataItem) fyne.CanvasObject {
	switch val := v.(type) {
	case binding.Bool:
		return widget.NewCheckWithData("", val)
	case binding.Float:
		s := widget.NewSliderWithData(0, 1, val)
		s.Step = 0.01
		return s
	case binding.Int:
		return widget.NewEntryWithData(binding.IntToString(val))
	case binding.String:
		return widget.NewEntryWithData(val)
	default:
		return widget.NewLabel("")
	}
}
