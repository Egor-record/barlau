package templates

import (
	"encoding/json"
	"fmt"
	"html/template"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"sync"
)

const nativePlayer = `<!DOCTYPE html>
<html>
<head>
<title>Example Web App</title>
<style>
    body {
        width: 100%;
        height: 100%;
        background-color:#202020;
    }
</style>

</head>
<body>
    <div>
        <video src="{{ . }}" autoplay>
        </video>
    </div>
</body>
</html>`

const shaka = `<!DOCTYPE html>
<html>
<head>
<title>Example Web App</title>
<style>
    body {
        width: 100%;
        height: 100%;
        background-color:#202020;
    }
</style>

</head>
<body>
    <div>
        <video src="{{ . }}" autoplay>
        </video>
    </div>
</body>
</html>`

func generateTemplate(manifest, path string, native bool, wg *sync.WaitGroup) {
	defer wg.Done()
	var html = nativePlayer
	if !native {
		html = shaka
	}

	t, err := template.New("webpage").Parse(html)

	if err != nil {
		log.Println("error generation html: ", err)
	}

	f, err := os.Create(path + "/index.html")
	if err != nil {
		log.Println("create file: ", err)
		return
	}

	err = t.Execute(f, manifest)
	if err != nil {
		log.Println("error writing to file: ", err)
		return
	}

	f.Close()
}

func CreateSamplePlayer(manifest string, native bool) int {
	var path = "backend/cli/app" + strconv.Itoa(rand.Int())
	createCopyOfTemplate(path)
	var wg sync.WaitGroup
	wg.Add(3)
	go generateTemplate(manifest, path, native, &wg)
	go generateAppJson(path, &wg)
	go generatePng(path, &wg)
	wg.Wait()
	v := rand.Int()
	packageIpk(path, v)
	os.RemoveAll(path)
	return v
}

func packageIpk(path string, rand int) {
	cmd := exec.Command("ares-package", path, "-o", "frontend/static/ipk"+strconv.Itoa(rand))
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Print the output
	fmt.Println(string(stdout))
}

func generatePng(path string, wg *sync.WaitGroup) {
	defer wg.Done()
	img := image.NewRGBA(image.Rect(0, 0, 80, 80))

	img.Set(2, 3, color.RGBA{255, 0, 0, 255})

	f, _ := os.OpenFile(path+"/icon.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)
}

func generateAppJson(path string, wg *sync.WaitGroup) {
	defer wg.Done()
	type Appinfo struct {
		ID                  string   `json:"id"`
		Version             string   `json:"version"`
		Vendor              string   `json:"vendor"`
		Type                string   `json:"type"`
		Main                string   `json:"main"`
		Title               string   `json:"title"`
		Icon                string   `json:"icon"`
		RequiredPermissions []string `json:"requiredPermissions"`
	}

	data := Appinfo{
		ID:      "com.sample.player",
		Version: "1.0.0",
		Vendor:  "My Company",
		Type:    "web",
		Main:    "index.html",
		Title:   "Sample Player",
		Icon:    "icon.png",
	}

	file, _ := json.MarshalIndent(data, "", " ")

	_ = ioutil.WriteFile(path+"/appinfo.json", file, 0644)
}

func createCopyOfTemplate(path string) {
	err := os.Mkdir(path, 0755)
	if err != nil {
		log.Fatal(err)
	}
}
