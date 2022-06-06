package tutorials

import (
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
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
	name.SetPlaceHolder("智管PC DEV")

	host := widget.NewEntry()
	host.Text = "http://10.10.0.123:8088"
	host.SetPlaceHolder("http://10.10.0.123:8088")
	host.Validator = validation.NewRegexp(`^http`, "not a valid website")

	//password := widget.NewPasswordEntry()
	//password.SetPlaceHolder("Password")

	//largeText := widget.NewMultiLineEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Name:", Widget: name, HintText: "服务名称 name"},
			{Text: "Host:", Widget: host, HintText: "服务地址 host:port"},
		},
		//OnCancel: func() {
		//	fmt.Println("Cancelled")
		//},
		OnSubmit: func() {
			if name.Text == "" {
				dialog.ShowError(errors.New("请输入名称"), win)
				return
			}
			data.CashData.ProviderList = append(data.CashData.ProviderList, data.User{Name: name.Text, Host: host.Text})
			data.SaveDefault()
			dialog.ShowInformation("Success", "保存成功", win)
			//fyne.CurrentApp().SendNotification(&fyne.Notification{
			//	Title:   "Form for: " + name.Text,
			//	Content: largeText.Text,
			//})
		},
	}
	form.SubmitText = "确认提交"
	//title := &widget.TextGrid{}
	//title.SetText("数据存储在 In C:/Users/Administrator/proxyConfig.json")
	title1 := canvas.NewText(" 使用请修改项目配置", nil)
	title2 := canvas.NewText(" EP：智管PC dev.env.js服务地址设为 localhost(ip):9876", nil)
	title1.TextStyle = fyne.TextStyle{Bold: true}
	//title1.SetText("使用请修改项目启动配置 EP：智管PC dev.env.js服务地址设为 localhost(ip):9876")
	//title.SetRowStyle(2,nil)
	return container.NewVBox(title1,title2,form)
	//return container.NewVBox(title1,title,form)
	//return container.NewBorder(nil, nil, container.NewGridWithColumns(2, listPanel, form), nil, )
}

func addPubScreen(win fyne.Window) fyne.CanvasObject {
	name := widget.NewEntry()
	name.SetPlaceHolder("PC测试")

	ip := widget.NewEntry()
	ip.Text = "10.10.0.123"
	ip.SetPlaceHolder("10.10.0.123")
	ip.Validator = validation.NewRegexp(`^\d`, "not a valid website")

	user := widget.NewEntry()
	user.Text = "Administrator"
	user.SetPlaceHolder("Administrator")

	password := widget.NewEntry()
	password.Text = "Nc@test"
	password.SetPlaceHolder("Nc@test")

	dir := widget.NewEntry()
	dir.Text = "D\\public\\NcManageUI"
	dir.SetPlaceHolder("D\\public\\NcManageUI")

	build := widget.NewEntry()
	build.Text = "npm run build:test"
	build.SetPlaceHolder("项目打包命令")
	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Name:", Widget: name, HintText: "服务名称"},
			{Text: "IP:", Widget: ip, HintText: "服务器IP地址"},
			{Text: "User:", Widget: user, HintText: "服务器登录名"},
			{Text: "PassWord:", Widget: password, HintText: "服务器登录密码"},
			{Text: "服务目录:", Widget: dir, HintText: "服务项目目录"},
			{Text: "打包:", Widget: build, HintText: "项目打包命令"},
		},
		//OnCancel: func() {
		//	fmt.Println("Cancelled")
		//},
		OnSubmit: func() {
			if name.Text == "" {
				dialog.ShowError(errors.New("请输入名称"), win)
				return
			}
			pub :=data.Pub{
				Name: name.Text,
				IP: ip.Text,
				User: user.Text,
				Pass: password.Text,
				ToDir: dir.Text,
				Build: build.Text,
			}
			data.CashData.PublishList = append(data.CashData.PublishList, pub)
			data.SaveDefault()
			dialog.ShowInformation("Success", "保存成功", win)

		},
	}
	form.SubmitText = "确认提交"
	return container.NewVBox(form)
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
