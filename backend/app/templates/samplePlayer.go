package templates

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"html/template"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"os/exec"
	"sync"
)

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

func CreateSamplePlayer(manifest string, native bool) *big.Int {
	var path = "backend/cli/app" + randomInt().String()
	createCopyOfTemplate(path)
	var wg sync.WaitGroup
	wg.Add(3)
	go generateTemplate(manifest, path, native, &wg)
	go generateAppJson(path, &wg)
	go generatePng(path, &wg)
	wg.Wait()
	random := randomInt()
	packageIpk(path, random)
	os.RemoveAll(path)
	return random
}

func packageIpk(path string, rand *big.Int) {
	cmd := exec.Command("ares-package", path, "-o", "frontend/static/ipk"+rand.String())
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

func randomInt() *big.Int {
	v, err := rand.Int(rand.Reader, big.NewInt(1000))
	if err != nil {
		panic(err)
	}

	return v
}
