package pp

import (
	"fmt"
	"log"

	"github.com/vault-thirteen/SDLW/development/pp/models"
)

// GetFuncDataFromFile gets information about C functions and converts it into
// Go functions' data. This function is an entry point into this package.
func GetFuncDataFromFile(filePath string, fps *FuncProcessorSettings) (fds []*models.FuncData, err error) {
	log.Println(fmt.Sprintf("Parsing a file: %s.", filePath))

	var fp *FuncParser
	fp, err = NewFuncParser(fps)
	if err != nil {
		return nil, err
	}

	fds, err = fp.ParseCFunctions(filePath)
	if err != nil {
		return nil, err
	}

	var fproc *FuncProcessor
	fproc, err = NewFuncProcessor(fps)
	if err != nil {
		return nil, err
	}

	err = fproc.CToGo(fds)
	if err != nil {
		return nil, err
	}

	return fds, nil
}
