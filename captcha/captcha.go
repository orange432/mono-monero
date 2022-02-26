package captcha

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/orange432/mono-monero/cache"
	"github.com/orange432/mono-monero/enigma"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

var appCache *cache.AppCache

func AddLabel(img *image.RGBA, x, y int, label string) {
	col := color.RGBA{200, 100, 0, 255}
	point := fixed.Point26_6{fixed.Int26_6(x * 64), fixed.Int26_6(y * 64)}
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(label)
}

// Generate draws a captcha with a random string
func Generate() (filepath string, encryptedString string) {
	// Clean the captcha directory
	CleanCaptcha()
	// Generate a random string and hash
	randomString := enigma.RandomString(5)
	encrypted := enigma.Encrypt(string(time.Now().UnixMicro()) + "___" + randomString)
	// Draw image
	rectImage := image.NewRGBA(image.Rect(0, 0, 281, 100))
	green := color.RGBA{0, 100, 0, 255}

	draw.Draw(rectImage, rectImage.Bounds(), &image.Uniform{green}, image.ZP, draw.Src)
	AddLabel(rectImage, 32, 32, randomString)
	file, err := os.CreateTemp("./public/captcha", "*.png")
	if err != nil {
		fmt.Println(err.Error())
	}
	png.Encode(file, rectImage)
	// b64, err := os.ReadFile("./" + file.Name())
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// return "data:image/png;base64," + base64.StdEncoding.EncodeToString(b64), encrypte
	return strings.Split(file.Name(), "./public/captcha\\")[1], encrypted
}

func Validate(input string, encrypted string) bool {
	decrypted := enigma.Decrypt(input)
	split := strings.Split(decrypted, "___")
	return split[1] == input
}

// CleanCaptcha cleans the captcha directory
func CleanCaptcha() {
	if (appCache.LastCleanedCaptcha + 120) < time.Now().Unix() {
		directory := "./public/captcha/"
		files, err := ioutil.ReadDir(directory)
		if err != nil {
			fmt.Println(err.Error())
		}
		for _, file := range files {
			if file.ModTime().Add(time.Second*120).Unix() < time.Now().Unix() {
				os.Remove(directory + file.Name())
			}
		}
		appCache.LastCleanedCaptcha = time.Now().Unix()
	}
}

// LoadCache loads in the cache variable
func LoadCache(c *cache.AppCache) {
	appCache = c
}
