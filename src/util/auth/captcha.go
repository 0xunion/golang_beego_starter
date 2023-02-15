package auth

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"math/big"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

var font *truetype.Font
var colors [10]color.RGBA

func init() {
	font_data, err := ioutil.ReadFile("assets/Consolas.ttf")
	if err != nil {
		panic(err)
	}
	font, err = freetype.ParseFont(font_data)
	if err != nil {
		panic(err)
	}
	colors[0] = color.RGBA{255, 0, 0, 255}
	colors[1] = color.RGBA{25, 25, 112, 255}
	colors[2] = color.RGBA{220, 20, 60, 255}
	colors[3] = color.RGBA{25, 25, 112, 255}
	colors[4] = color.RGBA{255, 165, 0, 255}
	colors[5] = color.RGBA{128, 0, 0, 255}
	colors[6] = color.RGBA{178, 34, 34, 255}
	colors[7] = color.RGBA{255, 140, 0, 255}
	colors[8] = color.RGBA{0, 100, 0, 255}
	colors[9] = color.RGBA{100, 149, 237, 255}
}

func random(min int, max int) int {
	res, _ := rand.Int(rand.Reader, big.NewInt(int64(max-min+1)))
	return min + int(res.Int64())
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func drawLine(x1 int, y1 int, x2 int, y2 int, callback func(x int, y int)) {
	dx := abs(x1 - x2)
	dy := abs(y1 - y2)
	sx, sy := 1, 1
	if x1 >= x2 {
		sx = -1
	}
	if y1 >= y2 {
		sy = -1
	}

	angle := dx - dy

	for {
		callback(x1, y1)
		if x1 == x2 && y1 == y2 {
			return
		}
		a2 := angle * 2
		if a2 > -dy {
			angle -= dy
			x1 += sx
		}
		if a2 < dx {
			angle += dx
			y1 += sy
		}
	}
}

/*
由给出的exp表达式和w与h绘制图片
*/
func GenerateImageFromText(exp string, char_size int) (string, bool) {
	padding := 10
	word_size := char_size + char_size/3
	w := padding*2 + word_size*len(exp) + 30
	h := padding*2 + word_size + 30

	img := image.NewNRGBA(image.Rect(0, 0, w, h))

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			img.Set(x, y, color.RGBA{255, 255, 255, 255})
		}
	}

	ctx := freetype.NewContext()
	ctx.SetFont(font)
	ctx.SetDPI(72)
	ctx.SetFontSize(float64(char_size))
	ctx.SetClip(img.Bounds())
	ctx.SetDst(img)
	ctx.SetSrc(&image.Uniform{color.RGBA{255, 0, 0, 255}})

	//随便画2~4条线
	lines := random(2, 4)
	for i := 0; i < lines; i++ {

		color := colors[random(0, 9)]

		mode := random(0, 1)
		switch mode {
		case 0:
			//随机出两个高度
			h1, h2 := random(0, h-1), random(0, h-1)
			drawLine(0, h1, w-1, h2, func(x, y int) {
				img.Set(x, y, color)
			})
			break
		case 1:
			//随机出两个宽度
			w1, w2 := random(0, w-1), random(0, w-1)
			drawLine(w1, 0, w2, h-1, func(x, y int) {
				img.Set(x, y, color)
			})
			break
		}
	}

	//一个一个画字符，间距不能一样，要骚一点，先随机出一个高度区间
	rect_left := random(padding, w-word_size*len(exp)-padding)
	rect_top := random(padding, h-word_size-padding)

	for i := range exp {
		pos := freetype.Pt(rect_left+i*word_size+random(0, 10), rect_top+random(0, 10)+word_size)
		ctx.SetSrc(&image.Uniform{colors[random(0, 9)]})
		_, err := ctx.DrawString(string(exp[i]), pos)
		if err != nil {
			fmt.Println(err)
			return "", false
		}
	}

	buff := bytes.NewBuffer(nil)
	err := png.Encode(buff, img)
	if err != nil {
		return "", false
	}

	//对比了一下向下取整再做if的效率和使用float向上取整的效率，好像用if效率更高，任何数学库的底层也都还是if
	buf_size := buff.Len()
	if buf_size%3 != 0 {
		padding = 4
	} else {
		padding = 0
	}
	b64_buf_size := (buf_size/3.0)*4 + padding
	b64_buf := make([]byte, b64_buf_size)
	base64.StdEncoding.Encode(b64_buf, buff.Bytes())

	return "data:image/png;base64," + string(b64_buf), true
}
