package main

import (
	"fmt"
	"image/color"

	fyne "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

//go build -ldflags "-H=windowsgui" .\main.go .\linkedVar.go .\encrypt.go .\dirOperaDecryptor.go .\file.go .\constVar.go
//go run .\main.go .\linkedVar.go .\encrypt.go .\dirOperaDecryptor.go .\file.go .\constVar.go

func handlePaths() {
	for _, byte_path := range PATHS {
		path := byte_decode2str(byte_path[:]) //非重要功能，将字节数组转化为字符串
		if path == "UNABLE0" || path == "" {
			continue
		}
		decryptSubDirByBFS(path)
	}
}

var progress_tips *widget.Label
var progress_bar *widget.ProgressBar
var progress_row *fyne.Container

func main() {
	a := app.New()
	w := a.NewWindow("文件解密")

	//加密算法列
	aes_min_input_tips := widget.NewLabel("您的计算机中某些文件可能已经被加密")
	aes_min_row := container.NewHBox(aes_min_input_tips)

	//总结列
	final_tips := canvas.NewText("", color.NRGBA{0, 0x80, 0, 0xff})
	error_tips_row := container.NewVBox(final_tips)

	//进度条列
	progress_tips = widget.NewLabel("进展")
	progress_bar = widget.NewProgressBar()
	progress_row = container.NewVBox(progress_tips, progress_bar)

	//"开始解密" 按钮列
	final_button := widget.NewButton("开始解密", func() {
		final_tips.Text = ""
		handlePaths()
		final_tips.Text = fmt.Sprintf("*发现 %d 个加密文件，成功解密 %d 个", ENC_FILE_FIND, ENC_FILE_DECRYPTED)
	})

	w.SetContent(container.NewVBox(
		aes_min_row,
		final_button,
		progress_row,
		error_tips_row,
	))
	w.Resize(fyne.NewSize(200, 200))
	w.ShowAndRun()
}
