// +build ignore

package main

import (
	"fmt"
	"os"
	"path"
	fileP "path/filepath"
	"runtime"

	"fyne.io/fyne/v2"
)

func bundleFile(name string, filepath string, f *os.File) {
    res, err := fyne.LoadResourceFromPath(filepath)
    fmt.Println(res)
    fmt.Println("sdflkfkdfs")
    if err != nil {
        fyne.LogError("Unable to load file "+filepath, err)
        return
    }

    _, err = f.WriteString(fmt.Sprintf("var %s = %#v\n", name, res))
    if err != nil {
        fyne.LogError("Unable to write to bundled file", err)
    }
}

func BundleFile(name string, filepath string) {
	f := openFile("bundled-scene.go")
	_, dirname, _, _ := runtime.Caller(0)
	m := fileP.Clean(path.Join(path.Dir(dirname), filepath))
	fmt.Println(m)
	res, err := fyne.LoadResourceFromPath(m)
	fmt.Println(res)
	fmt.Println("sdflkfkdfs")
	if err != nil {
		fyne.LogError("Unable to load file "+filepath, err)
		return
	}

	_, err = f.WriteString(fmt.Sprintf("var %s = %#v\n", name, res))
	if err != nil {
		fyne.LogError("Unable to write to bundled file", err)
	}
}
func openFile(filename string) *os.File {
    os.Remove(filename)
    _, dirname, _, _ := runtime.Caller(0)
    f, err := os.Create(path.Join(path.Dir(dirname), filename))
    if err != nil {
        fyne.LogError("Unable to open file "+filename, err)
        return nil
    }

    _, err = f.WriteString("// **** THIS FILE IS AUTO-GENERATED, PLEASE DO NOT EDIT IT **** //\n\npackage data\n\nimport \"fyne.io/fyne/v2\"\n\n")
    if err != nil {
        fyne.LogError("Unable to write file "+filename, err)
        return nil
    }

    return f
}

func iconDir() string {
    _, dirname, _, _ := runtime.Caller(0)
    return path.Join(path.Dir(dirname), "icons")
}
func init() {
    fmt.Println("sdflkjksdf")

}
func main() {
    fmt.Println("sdflkjksdf")
    f := openFile("bundled-scene.go")

    bundleFile("fynescenedark", "sf.jpeg", f)
    bundleFile("fynescenelight", "light.jpg", f)

    f.Close()
}
