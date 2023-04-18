package parser

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	c "github.com/vault-thirteen/SDLW/development/common"
	"github.com/vault-thirteen/auxie/reader"
	"github.com/vault-thirteen/errorz"
)

const (
	ErrNoHeader                = "no header"
	ErrFmtNoHeader             = ErrNoHeader + ": %s"
	ErrBadHeader               = "bad header: %s"
	ErrNoArgs                  = "no args"
	ErrFmtNoArgs               = ErrNoArgs + ": %s"
	ErrUnsupportedArgumentText = "unsupported argument text: %s"
)

type FuncParser struct {
	ignoredWords []c.Word
}

// NewFuncParser is a FuncParser's constructor.
func NewFuncParser(ignoredWords []string) (fp *FuncParser, err error) {
	fp = &FuncParser{
		ignoredWords: make([]c.Word, 0),
	}

	for _, iw := range ignoredWords {
		w := c.Word(iw)
		err = w.Check()
		if err != nil {
			return nil, err
		}
		fp.ignoredWords = append(fp.ignoredWords, w)
	}

	return fp, nil
}

// ParseCLine parses C function data from a text line.
// On error, a null pointer is returned.
// On success, a set pointer is returned.
func (fp *FuncParser) ParseCLine(line string) (fd *c.FuncData, err error) {
	fd = &c.FuncData{
		C: &c.FuncDatum{
			SourceCode: []string{line},
		},
		Go: &c.FuncDatum{},
	}

	// Remove words which are ignored, detach stars, remove double spaces.
	clearLine := c.StringWithoutWords(line, fp.ignoredWords)
	clearLine = c.DetachPointerMarks(clearLine)
	clearLine = c.RemoveDoubleSpaces(strings.TrimSpace(clearLine))

	err = fp.parseCHead(clearLine, fd)
	if err != nil {
		return nil, err
	}

	err = fp.parseCArgs(clearLine, fd)
	if err != nil {
		return nil, err
	}

	return fd, nil
}

// parseCHead parses C function head data and saves it into dst.
// Function header consists of an output argument and function name.
func (fp *FuncParser) parseCHead(line string, dst *c.FuncData) (err error) {
	parts := strings.Split(line, "(")
	if len(parts) != 2 {
		return fmt.Errorf(ErrFmtNoHeader, line)
	}
	header := parts[0]

	headerParts := strings.Split(header, " ")
	if len(headerParts) < 2 {
		return fmt.Errorf(ErrBadHeader, header)
	}
	funcName := headerParts[len(headerParts)-1]
	funcRetTypeParts := headerParts[:len(headerParts)-1]
	funcRetType := c.GlueParts(funcRetTypeParts)

	dst.C.Name.Original = funcName
	dst.C.ReturnedValues = []*c.Argument{
		{
			Type: c.Parameter{
				Original: funcRetType,
			},
		},
	}

	return nil
}

// parseCArgs parses C function's input arguments and saves them into dst.
func (fp *FuncParser) parseCArgs(line string, dst *c.FuncData) (err error) {
	dst.C.Arguments = make([]*c.Argument, 0)

	parts := strings.Split(line, "(")
	if len(parts) != 2 {
		dst.C.Arguments = nil
		return fmt.Errorf(ErrFmtNoHeader, line)
	}
	nonHead := parts[1]

	nonHeadParts := strings.Split(nonHead, ")")
	if len(nonHeadParts) < 1 {
		dst.C.Arguments = nil
		return fmt.Errorf(ErrFmtNoArgs, line)
	}
	argsPart := nonHeadParts[0]

	argSections := strings.Split(argsPart, ",")

	var arg *c.Argument
	for _, argSection := range argSections {
		arg, err = parseCArgSection(argSection)
		if err != nil {
			dst.C.Arguments = nil
			return err
		}

		dst.C.Arguments = append(dst.C.Arguments, arg)
	}

	return nil
}

// parseCArgSection parses C argument text.
// Examples of argument texts: void, char * s, int size.
func parseCArgSection(s string) (arg *c.Argument, err error) {
	s = strings.TrimSpace(s)
	parts := strings.Split(s, " ")
	partsCount := len(parts)
	switch partsCount {
	case 1:
		// Only type is set. Example: void.
		arg = &c.Argument{
			Type: c.Parameter{
				Original: parts[0],
			},
		}
		return arg, nil

	case 2:
		// Name & type are set. Example: int size.
		arg = &c.Argument{
			Type: c.Parameter{
				Original: parts[0],
			},
			Name: c.Parameter{
				Original: parts[1],
			},
		}
		return arg, nil

	case 3:
		// Name & pointer type ? Example: char * name.
		if parts[1] == "*" {
			arg = &c.Argument{
				Type: c.Parameter{
					Original: c.GlueParts([]string{parts[0], parts[1]}),
				},
				Name: c.Parameter{
					Original: parts[2],
				},
			}
			return arg, nil
		} else {
			return nil, fmt.Errorf(ErrUnsupportedArgumentText, s)
		}

	default:
		return nil, fmt.Errorf(ErrUnsupportedArgumentText, s)
	}
}

// ParseCFunctions reads a file containing definitions of C functions and
// parses its lines into function data. Lines which do not have a simple
// function header are ignored. Function definition must be on a single line.
//
// N.B.
// Golang currently does not support callbacks and C function pointers, so
// implementation of such things in this parser would be a waste of time.
func ParseCFunctions(filePath string) (funcData []*c.FuncData, err error) {
	// Parser.
	var ignoredWords = []string{
		"extern",
		"DECLSPEC",
		"SDLCALL",
		"const",
	}
	var fp *FuncParser
	fp, err = NewFuncParser(ignoredWords)
	if err != nil {
		return nil, err
	}

	// File.
	var f *os.File
	f, err = os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func() {
		derr := f.Close()
		if derr != nil {
			err = errorz.Combine(err, derr)
		}
	}()

	// Result Holder.
	funcData = make([]*c.FuncData, 0)
	var fd *c.FuncData

	// Line Reader.
	lr := reader.New(f)
	var line []byte
	line, err = lr.ReadLineEndingWithCRLF()
	for {
		if err != nil {
			if err == io.EOF {
				err = nil
				break
			} else {
				return nil, err
			}
		}

		fd, err = fp.ParseCLine(strings.TrimSpace(string(line)))
		// Clear error for special cases.
		if err != nil {
			if strings.HasPrefix(err.Error(), ErrNoHeader) {
				err = nil
				log.Print(fmt.Sprintf("ignoring line without header: %s", string(line)))
			} else if strings.HasPrefix(err.Error(), ErrNoArgs) {
				err = nil
				log.Print(fmt.Sprintf("ignoring line without arguments: %s", string(line)))
			}
		}
		if err != nil {
			return nil, err
		}

		if fd != nil {
			funcData = append(funcData, fd)
		}

		// Next Line.
		line, err = lr.ReadLineEndingWithCRLF()
	}

	return funcData, nil
}
