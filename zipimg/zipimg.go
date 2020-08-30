package zipimg

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	"math"
	"os"

	resize "github.com/nfnt/resize"
)

// calculateRatioFit 计算图片缩放后的尺寸
func calculateRatioFit(srcWidth, srcHeight int) (int, int) {
	const DefaultMaxWidth float64 = 750
	const DefaultMaxHeight float64 = 1334
	ratio := math.Min(DefaultMaxWidth/float64(srcWidth), DefaultMaxHeight/float64(srcHeight))
	return int(math.Ceil(float64(srcWidth) * ratio)), int(math.Ceil(float64(srcHeight) * ratio))
}

// MakeThumbnail 生成缩略图
func MakeThumbnail(imagePath, savePath string) bool {

	file, _ := os.Open(imagePath)
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return false
	}

	b := img.Bounds()
	width := b.Max.X
	height := b.Max.Y

	w, h := calculateRatioFit(width, height)

	// fmt.Println("width = ", width, " height = ", height)
	// fmt.Println("w = ", w, " h = ", h)

	// 调用resize库进行图片缩放
	m := resize.Resize(uint(w), uint(h), img, resize.Lanczos3)

	// 需要保存的文件
	imgfile, _ := os.Create(savePath)
	defer imgfile.Close()

	// 以PNG格式保存文件
	err = png.Encode(imgfile, m)
	if err != nil {
		return false
	}

	return true
}
