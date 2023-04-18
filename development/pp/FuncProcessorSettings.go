package pp

import "github.com/vault-thirteen/SDLW/development/pp/models"

// FuncProcessorSettings is settings of the parser and processor.
type FuncProcessorSettings struct {
	// List of words to be removed from text during parsing of C functions.
	ignoredWords []models.Word

	// Function name prefix which will be removed from C function names during
	// the conversion to Go function names.
	funcNamePrefix string

	// Type mappings.
	// Key = C type name, Value = Go type name.
	cToGoTypeMapping map[string]string
}

// NewFuncProcessorSettings is a constructor of FuncProcessorSettings.
func NewFuncProcessorSettings(ignoredWords []models.Word, funcNamePrefix string, cToGoTypeMapping map[string]string) (fps *FuncProcessorSettings, err error) {
	fps = &FuncProcessorSettings{
		ignoredWords:     ignoredWords,
		funcNamePrefix:   funcNamePrefix,
		cToGoTypeMapping: cToGoTypeMapping,
	}

	err = fps.check()
	if err != nil {
		return nil, err
	}

	return fps, nil
}

// check checks the settings for correctness.
func (fps *FuncProcessorSettings) check() (err error) {
	for _, iw := range fps.ignoredWords {
		err = iw.Check()
		if err != nil {
			return err
		}
	}

	w := models.Word(fps.funcNamePrefix)
	err = w.Check()
	if err != nil {
		return err
	}

	return nil
}
