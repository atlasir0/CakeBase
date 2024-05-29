package config

import (
	"errors"
	"flag"
	"os"
)

type Flags struct {
	ShowSymlinks bool
	ShowDirs     bool
	ShowFiles    bool
	SpecificExt  bool
}

type Parameters struct {
	Path      string
	Extension string
}

var ErrInvalidFlagsCombination = errors.New("flag -ext works only with flag -f")
var ErrDirectoryNotFound = errors.New("no such directory")
var ErrMissingDirectory = errors.New("no directory passed")

func ParseArgs() (params Parameters, flags Flags, err error) {
	flags = Flags{}
	params = Parameters{}

	flag.BoolVar(&flags.ShowSymlinks, "sl", false, "Print only symlinks")
	flag.BoolVar(&flags.ShowDirs, "d", false, "Print only directories")
	flag.BoolVar(&flags.ShowFiles, "f", false, "Print only files")
	flag.StringVar(&params.Extension, "ext", "", "Print only specific extension. Works only with -f")
	flag.Parse()

	if params.Extension != "" {
		flags.SpecificExt = true
	}

	if flags.SpecificExt && !flags.ShowFiles {
		err = ErrInvalidFlagsCombination
		return
	}

	if len(flag.Args()) > 0 {
		params.Path = flag.Arg(0)
	}

	if params.Path == "" {
		err = ErrMissingDirectory
		return
	}

	if _, err = os.Stat(params.Path); os.IsNotExist(err) {
		err = ErrDirectoryNotFound
		return
	}

	return
}
