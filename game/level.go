package game

import (
	"io/ioutil"
	"log"
	"encoding/json"
	"fmt"
	"strconv"
)

type LevelInfo struct {
	Name string `json:"name"`
	Author string `json:"author"`
	LevelIndex int `json:"level_index"`
	BeatsPerMinute int `json:"bpm"`
	Divisions int `json:"divisions"`
	AudioFormat string `json:"format"`
}

type SectionData struct {
	Speed int `json:"speed"`
	Divisions float64 `json:"divisions"`
	Pause bool `json:"pause"`
	SpawnsEnemies bool `json:"spawns_enemies"`
	DefinedRythm bool `json:"defined_rythm"`
	ShowTutorial string `json:"show_tutorial"`
}

type LevelData struct {
	Sections []int `json:"sections"`
	SectionData []SectionData `json:"section_data"`
}

type Level struct {
	Info *LevelInfo
	Data *LevelData
}

func GetLevels(gsm *GameStateManager) []Level {
	levels := make([]Level, 0)
	files, err := ioutil.ReadDir(gsm.game.Content.GetRoot()+"/sound/")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	for _, f := range files {
		if f.IsDir() {
			if f.Name() != "sfx" {
				fmt.Println(gsm.game.Content.GetRoot()+"/sound/"+f.Name())
				lvlinf := GetLevelInfo(gsm, f.Name())
				if lvlinf.Info != nil && lvlinf.Data != nil {
					levels = append(levels, lvlinf)
				}
			}
		}
	}
	return levels
}

func GetLevelInfo(gsm *GameStateManager, name string) Level {
	var lvlInfo LevelInfo
	var lvlData LevelData
	levelinfostr, err := gsm.game.Content.LoadResourceStr("sound/"+name+"/TRACK_INF.json")
	if err != nil {
		fmt.Println(name+"/TRACK_INF.json", err)
	} else {
		err = json.Unmarshal([]byte(levelinfostr), &lvlInfo)
		if err != nil {
			fmt.Println(name+"/TRACK_INF.json", err)
		}
	}
	leveldatastr, err := gsm.game.Content.LoadResourceStr("sound/"+name+"/LEVEL.json")
	if err != nil {
		fmt.Println(name+"/LEVEL.json", err)
		return Level{}
	} else {
		err = json.Unmarshal([]byte(leveldatastr), &lvlData)
		if err != nil {
			fmt.Println(name+"/LEVEL.json", err)
		}
	}
	return Level{&lvlInfo, &lvlData}
}

func LoadLevel(gsm *GameStateManager, level int) {
	gw := new(GameWorld)
	lvls := GetLevels(gsm)
	for _, lvl := range lvls {
		if lvl.Info.LevelIndex == level {
			if gsm.state != nil {
				gsm.state.(*GameWorld).music.Stop()
			}
			gsm.PushState(gw).SendInit(lvl)
			if gsm.state != nil {
				gsm.state.(*GameWorld).LID = func () map[int]*SectionData {
					var m map[int]*SectionData = make(map[int]*SectionData)
					for i, d := range lvl.Data.Sections {
						fmt.Println("Adding section " + strconv.Itoa(i) + "@" + strconv.Itoa(d)+"\n	-", lvl.Data.SectionData[i])
						m[d] = &lvl.Data.SectionData[i]
					}
					return m
				}()
			}
			return
		}
	}
	LoadLevel(gsm, 0)
}