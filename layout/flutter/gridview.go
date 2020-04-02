package flutter

import (
	"image/jpeg"
	"os"
)

const (
	vertical = iota
	landscape
	square
	transparent
)

type direction int

type Size struct {
	Width  int
	Height int
}

type SizeRatio struct {
	Width  int
	Height int
}

//normalized 归一化
//对输入的宽高进行`近似`化的比值处理
func normalized(width, height int) direction {

	if width > height {
		return landscape
	}

	if width == height {
		return square
	}

	return vertical
}

//loadPic 加载图片，返回图片的宽和高
func LoadPic(path string) (int, int) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	img, err := jpeg.Decode(f)
	if err != nil {
		panic(err)
	}

	size := img.Bounds().Size()
	return size.X, size.Y
}

//Calc 计算Flutter使用的布局数据
func Calc(s []Size) []SizeRatio {
	var ratio []direction
	var addon bool = false
	var sr []SizeRatio

	for _, _s := range s {
		ratio = append(ratio, normalized(_s.Width, _s.Height))
	}

	if len(ratio)%2 == 1 {
		addon = true
		ratio = append(ratio, transparent)
	}

	for i := 0; i+1 < len(ratio); i += 2 {
		if ratio[i] == vertical {
			sr = append(sr, SizeRatio{
				Width:  2,
				Height: 4,
			})
		} else if ratio[1] == landscape {
			sr = append(sr, SizeRatio{
				Width:  4,
				Height: 2,
			})
		} else if ratio[i] == square {
			sr = append(sr, SizeRatio{
				Width:  2,
				Height: 2,
			})
		} else {
			sr = append(sr, SizeRatio{
				Width:  4,
				Height: 4,
			})
		}

		if ratio[i+1] == vertical {
			sr = append(sr, SizeRatio{
				Width:  2,
				Height: 4,
			})
		} else if ratio[i+1] == landscape {
			sr = append(sr, SizeRatio{
				Width:  4,
				Height: 2,
			})
		} else if ratio[i+1] == square {
			sr = append(sr, SizeRatio{
				Width:  2,
				Height: 2,
			})
		} else {
			sr = append(sr, SizeRatio{
				Width:  4,
				Height: 4,
			})
		}
	}

	if addon {
		return sr[:len(sr)-1]
	}

	return sr
}
