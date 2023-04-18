package main

import (
	"fmt"
	"log"
	"os"

	"github.com/vault-thirteen/SDLW/development/pp"
	"github.com/vault-thirteen/SDLW/development/pp/models"
)

func mustBeNoError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// Command line arguments.
	if len(os.Args) != 2 {
		log.Fatal("os args")
	}
	filePath := os.Args[1]

	fps, err := pp.NewFuncProcessorSettings(
		[]models.Word{
			"extern",
			"const",
			"DECLSPEC",
			"SDLCALL",
		},
		"Mix_",
		map[string]string{
			"void":    "",
			"void*":   "*m.Void",
			"char":    "m.Char",
			"char*":   "string",
			"int":     "m.Int",
			"int*":    "*m.Int",
			"uint":    "m.Uint",
			"Int8":    "m.Int8",
			"Uint8":   "m.Uint8",
			"Sint8":   "m.Int8",
			"Uint8*":  "*m.Uint8",
			"Int16":   "m.Int16",
			"Uint16":  "m.Uint16",
			"Sint16":  "m.Int16",
			"Uint16*": "*m.Uint16",
			"Int32":   "m.Int32",
			"Uint32":  "m.Uint32",
			"Sint32":  "m.Int32",
			"float":   "m.Float",
			"double":  "m.Double",
			"size_t":  "m.SizeT",

			// SDL Types.
			"SDL_bool":     "*m.Bool",
			"SDL_RWops*":   "*m.RWops",
			"SDL_version*": "*m.Version",

			// SDL Mixer Types.
			"Mix_Chunk*":       "*mm.Chunk",
			"Mix_Fading":       "mm.Fading",
			"Mix_EffectFunc_t": "mm.EffectFunc",
			"Mix_EffectDone_t": "mm.EffectDone",
			"Mix_Music*":       "*mm.Music",
			"Mix_MusicType":    "mm.MusicType",
		})
	mustBeNoError(err)

	err = showFuncData(filePath, fps)
	mustBeNoError(err)
}

func showFuncData(filePath string, fps *pp.FuncProcessorSettings) (err error) {
	var fds []*models.FuncData
	fds, err = pp.GetFuncDataFromFile(filePath, fps)
	if err != nil {
		return err
	}

	fmt.Println(fds) //TODO

	return nil
}
