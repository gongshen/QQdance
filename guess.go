package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"gocv.io/x/gocv"
	"image"
	"log"
	"time"
)

func guess() {
	// 加载预训练的 ONNX 模型
	modelPath := "arrow_classifier.onnx" // 替换为你的模型路径
	net := gocv.ReadNet(modelPath)
	if net.Empty() {
		log.Fatalf("无法加载模型: %s\n", modelPath)
	}
	defer net.Close()

	fmt.Println("程序已启动，开始识别屏幕中心箭头...")

	for {
		// 截取屏幕中心区域（假设箭头大小为 100x100）
		centerX, centerY := robotgo.GetScreenSize()
		centerX /= 2
		centerY /= 2
		img := robotgo.CaptureScreen(centerX-50, centerY-50, 100, 100) // 截取 100x100 区域
		defer robotgo.FreeBitmap(img)

		// 将截图转换为 gocv.Mat
		imgBytes := robotgo.ToBitmapBytes(img)
		mat, err := gocv.ImageToMatRGB(imgBytes)
		if err != nil {
			log.Fatalf("无法转换图像: %v\n", err)
		}
		defer mat.Close()

		// 预处理图像（调整大小、归一化等）
		blob := gocv.BlobFromImage(mat, 1.0, image.Pt(100, 100), gocv.NewScalar(0, 0, 0, 0), true, false)
		defer blob.Close()

		// 输入模型进行推理
		net.SetInput(blob, "")
		prob := net.Forward("")
		defer prob.Close()

		// 获取分类结果
		classID := getPredictedClass(prob)
		switch classID {
		case 0:
			fmt.Println("识别结果: 1 (向下)")
		case 1:
			fmt.Println("识别结果: 2 (向上)")
		case 2:
			fmt.Println("识别结果: 3 (向左)")
		case 3:
			fmt.Println("识别结果: 4 (向右)")
		default:
			fmt.Println("未识别到箭头")
		}

		// 等待一段时间后继续
		time.Sleep(1 * time.Second)
	}
}

// 获取预测的类别
func getPredictedClass(prob gocv.Mat) int {
	// 获取概率最大的类别
	_, maxVal, _, maxLoc := gocv.MinMaxLoc(prob)
	if maxVal < 0.5 { // 置信度阈值
		return -1
	}
	return maxLoc.X
}
