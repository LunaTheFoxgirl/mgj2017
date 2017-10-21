package game

import (
	"bitbucket.org/Member1221/music-gj-engine/engine"
	"time"
	"bitbucket.org/Member1221/music-gj-engine/engine/audio"
	"fmt"
	"github.com/faiface/pixel/text"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"image/color"
	"strconv"
	"math/rand"
)

//The value that all audio pulses follow
var pulse float64 = 0
//The last beat
var last = -1
//If the pulse should be flipped
var flip = false

type GameWorld struct {
	//Inherits
	engine.Drawable

	GameState *GameStateManager

	//Level Information and data
	Info LevelInfo
	Data LevelData

	//The player
	player *Player

	//The background music in the level
	music *audio.AudioClip

	//The font for splash text, etc.
	igfont *engine.SpriteFont

	//The scrolling background
	igbgsprite *engine.Sprite
	igbgdarkensprite *engine.Sprite

	//The area the player may move in.
	GameBounds *pixel.Rect

	LID map[int]*SectionData

	CurrentSection SectionData

	Score int
	Combo int
	ComboFalloff int

	//The position of the scrolling backgrounds
	bgp pixel.Vec

	//Entities in the game.
	entities []engine.Drawable
}

func (gw *GameWorld) Init(game *engine.Game, createInfo interface{}) {
	var err error
	gw.player = new(Player)
	gw.player.Init(game, createInfo)
	gw.player.AssignWorld(gw)

	cinf := createInfo.(Level)
	gw.Info = *cinf.Info
	gw.Data = *cinf.Data

	gw.CurrentSection = SectionData{
		gw.Info.BeatsPerMinute,
		1,
		true,
		false,
		false,
		"none",
	}

	game.Window.SetTitle("Rektangel > " + gw.Info.Name)
	gw.music, err = game.Content.LoadAudio("sound/"+cinf.Info.Name+"/TRACK."+cinf.Info.AudioFormat)
	if err != nil {
		fmt.Println(err)
		return
	}

	gw.igfont, err = game.Content.CreateSpriteFont("fonts/Bedstead.ttf", 18, text.ASCII...)
	if err != nil {
		fmt.Println(err)
		return
	}

	gw.igbgsprite, err = game.Content.LoadSprite("images/PLAY_FIELD.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	gw.igbgdarkensprite, err = game.Content.LoadSprite("images/PLAY_DARKEN.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	minc, maxc := game.Window.GetBounds().Center(), game.Window.GetBounds().Center()
	minc.Y += (320/2)-32
	maxc.Y -= (320/2)-32
	gw.ComboFalloff = 1
	gw.GameBounds = &pixel.Rect{minc, maxc}
	/*for i := 0; i < int((float64(gw.music.Source.Len()))/gw.music.Period(float64(gw.Info.BeatsPerMinute), 1)); i++ {

		n := new(Note)
		n.noteheight = float64((-128)+(64*(rand.Int()%5)))
		n.Init(game, float64((gw.music.Position()+1)/period))
		n.AssignWorld(gw)
		gw.entities = append(gw.entities, n)
	}*/
	go gw.music.Play(nil)
}

var fwd = 0
func (gw *GameWorld) Update(game *engine.Game, delta time.Duration) {
	if game.Input.KeyDown(pixelgl.KeyD) {
		gw.music.Seek(gw.music.Position()+((gw.music.Source.Len()/audio.SAMPLER_RATE))*fwd)
		fwd++
	}
	for _, e := range gw.entities {
		e.Update(game, delta)
	}
	var period int = int(gw.music.Period(float64(gw.CurrentSection.Speed), gw.CurrentSection.Divisions))
	game.Window.SetTitle(fmt.Sprint(gw.LID[int(gw.music.Position()/period)]))
	if gw.LID[int(gw.music.Position()/period)] != nil {
		if gw.CurrentSection != *gw.LID[int(gw.music.Position()/period)] {
			fmt.Println("Next section...")
			gw.CurrentSection = *gw.LID[int(gw.music.Position()/period)]
			if gw.CurrentSection.Divisions == 0 {
				gw.CurrentSection.Divisions = 1
			}
		}
	}
	gw.player.Update(game, delta)
}

func (gw *GameWorld) Draw(game *engine.Game, delta time.Duration) {

	//Move the playfield around to the music, including the bopping.
	gw.bgp = game.Window.GetBounds().Center()
	gw.HandleTiming(game, delta)
	gw.HandleScoreDraw(game, delta)
}

func (gw *GameWorld) HandleScoreDraw(game *engine.Game, delta time.Duration) {
	game.SpriteBatch.Begin()
	otext := strconv.Itoa(gw.Score)
	combo := ""
	if gw.Combo > 300 {
		combo = "\nx" + strconv.Itoa(gw.Combo)
	}
	ce := pulse
	if !flip {
		ce = -pulse
	}
	game.SpriteBatch.DrawString(gw.igfont, otext + combo, 0, 2, pixel.V(float64(game.Window.GetBounds().W()-float64((20*(len(otext))+(2*len(otext)))))-16, ((game.Window.GetBounds().H()-32)-float64(ce))-16), nil)
	game.SpriteBatch.End()
}

func (gw *GameWorld) IncreasePoints(amnt int) {
	gw.Score += amnt*gw.Combo
	gw.Combo += ((amnt+1)/gw.ComboFalloff)
	gw.ComboFalloff += 1
}

func (gw *GameWorld) EndCombo() {
	gw.Combo = 0
	gw.ComboFalloff = 1
}

func (gw *GameWorld) HandleTiming(game *engine.Game, delta time.Duration) {
	var period int = int(gw.music.Period(float64(gw.CurrentSection.Speed), gw.CurrentSection.Divisions))
	//xfmt.Println(period, gw.CurrentSection.Divisions, gw.CurrentSection.Speed)
	if (gw.music.Position()+1)/period != last {
		pulse = 5
		last = (gw.music.Position()+1)/period
		flip = !flip
		if !gw.CurrentSection.Pause {
			if (gw.music.Position()+1)/period < (gw.music.Source.Len())/period {
				n := new(Note)
				n.noteheight = float64((-128) + (64 * (rand.Int() % 5)))
				n.AssignWorld(gw)
				n.Init(game, float64((gw.music.Position() / period)))
				gw.entities = append(gw.entities, n)
			}
		}
	}
	if pulse > 0 {
		pulse -= 0.1
	}
	if !gw.CurrentSection.DefinedRythm {
		pulse = 0
	}
	/*if flip {
		gw.bgp.Y += float64(pulse)
	} else {
		gw.bgp.Y -= float64(pulse)
	}*/


	game.SpriteBatch.Begin()
	game.SpriteBatch.Draw(gw.igbgsprite, gw.bgp.Add(pixel.V(0, 0)), 0, 1, color.RGBA{255, 255, 255, 255}, nil)
	game.SpriteBatch.End()
	for _, e := range gw.entities {
		e.Draw(game, delta)
	}
	game.SpriteBatch.Begin()
	game.SpriteBatch.DrawEx(gw.igbgdarkensprite, gw.bgp.Add(pixel.V(0, 0)), 0, pixel.Vec{1, 5}, color.RGBA{255, 255, 255, 255}, nil)
	game.SpriteBatch.End()
	gw.player.Draw(game, delta)
}

