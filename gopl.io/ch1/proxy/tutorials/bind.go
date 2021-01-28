package tutorials

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func bindingScreen(_ fyne.Window) fyne.CanvasObject {

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
	host.SetPlaceHolder("http://10.10.0.123:8088")
	//host.Validator = validation.NewRegexp(`\w{1,}@\w{1,}\.\w{1,4}`, "not a valid email")

	password := widget.NewPasswordEntry()
	password.SetPlaceHolder("Password")

	largeText := widget.NewMultiLineEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Name", Widget: name, HintText: "server provider name"},
			{Text: "Host", Widget: host, HintText: "server host:port"},
		},
		//OnCancel: func() {
		//	fmt.Println("Cancelled")
		//},
		OnSubmit: func() {
			fmt.Println("Form submitted")
			fmt.Println(name.Text)
			fmt.Println(host.Text)
			fyne.CurrentApp().SendNotification(&fyne.Notification{
				Title:   "Form for: " + name.Text,
				Content: largeText.Text,
			})
		},
	}
	return container.NewVBox(form)
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
