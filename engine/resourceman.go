package engine

import (
	"os"
	"github.com/faiface/pixel"
	"image"

	_ "image/png"
	_ "image/jpeg"
	_ "image/gif"
	"fmt"
	"bitbucket.org/Member1221/music-gj-engine/engine/audio"
	"io/ioutil"
	"github.com/golang/freetype/truetype"
	"github.com/faiface/pixel/text"
)

type rescon struct {
	typestr string
	item interface{}
}

type ResourceManager struct {
	root string
	//content map[string]*rescon
}

func newresman(root string) *ResourceManager {
	res := new(ResourceManager)
	res.root = root
	//res.content = make(map[string]*rescon)
	return res
}

func (rm *ResourceManager) SetRoot(path string) {
	rm.root = path
}

func (rm *ResourceManager) GetRoot() string {
	return rm.root
}

func (rm *ResourceManager) CreateSpriteFont(name string, size float64, runesets ...rune) (*SpriteFont, error) {
	if rm.root == "" {
		rm.root = "dat"
	}
	file, err := os.Open(rm.root + "/" + name)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	font, err := truetype.Parse(bytes)
	if err != nil {
		return nil, err
	}
	ttf := truetype.NewFace(font, &truetype.Options{
		Size: size,
		GlyphCacheEntries: 1,
	})

	atlas := text.NewAtlas(ttf, runesets)

	return &SpriteFont{atlas, nil}, nil
}

func (rm *ResourceManager) LoadSpriteSheet(name string, selects [][]pixel.Rect) (*Spritesheet, error) {
	if rm.root == "" {
		rm.root = "dat"
	}
	/*if val, ok := rm.content[name]; ok {
		if val.typestr == "sprite" {
			return val.item.(*Spritesheet), nil
		}
		return nil, errors.New("tried to load type "+val.typestr+" as sprite!")
	}
	*/
	file, err := os.Open(rm.root + "/" + name)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	pic := pixel.PictureDataFromImage(img)
	spr := pixel.NewSprite(pic, pic.Bounds())
	var rects [][]pixel.Rect = make([][]pixel.Rect, 0)
	for xy, y := range selects {
		rects = append(rects, make([]pixel.Rect, 0))
		for xi, _ := range y {
			rects[xy] = append(rects[xy], selects[xy][xi])
		}
	}
	fmt.Println("Loaded spritesheet", name, "with size", spr.Picture().Bounds().String(), "with rect", len(rects))
	/*rm.content[name] = &rescon{
		"sprite",
		&Spritesheet { spr: spr, rect: rects },
	}*/
	return &Spritesheet { spr: spr, rect: rects }, nil //rm.LoadSpriteSheet(name, selects)
}

func (rm *ResourceManager) LoadAudio(name string) (*audio.AudioClip, error) {
	if rm.root == "" {
		rm.root = "dat"
	}
	/*
	if val, ok := rm.content[name]; ok {
		if val.typestr == "audioclip" {
			return val.item.(*audio.AudioClip), nil
		}
		return nil, errors.New("tried to load type "+val.typestr+" as audioclip!")
	}
	*/
	file, err := os.Open(rm.root + "/" + name)
	if err != nil {
		return nil, err
	}
	au, format, err := audio.Decode(file)
	if err != nil {
		return nil, err
	}
	fmt.Println("Loaded audioclip", name, "with size", au.Len())
	/*rm.content[name] = &rescon{
		"audioclip",
		&audio.AudioClip{au, &format },
	}*/
	return &audio.AudioClip{Source:au, Format:&format }, nil // rm.LoadAudio(name)
}

func (rm *ResourceManager) LoadResourceStr(name string) (string, error) {
	if rm.root == "" {
		rm.root = "dat"
	}
	file, err := os.Open(rm.root + "/" + name)
	if err != nil {
		return "", err
	}
	defer file.Close()
	f, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	return string(f), nil
}

func (rm *ResourceManager) LoadSprite(name string) (*Sprite, error) {
	if rm.root == "" {
		rm.root = "dat"
	}
	/*
	if val, ok := rm.content[name]; ok {
		if val.typestr == "sprite" {
			return val.item.(*Sprite), nil
		}
		return nil, errors.New("tried to load type "+val.typestr+" as sprite!")
	}
	*/
	file, err := os.Open(rm.root + "/" + name)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	pic := pixel.PictureDataFromImage(img)
	spr := pixel.NewSprite(pic, pic.Bounds())
	fmt.Println("Loaded sprite", name, "with size", spr.Picture().Bounds().String())
	/*rm.content[name] = &rescon{
		"sprite",
		&Sprite {spr },
	}*/
	return &Sprite {spr }, nil //rm.LoadImage(name)
}