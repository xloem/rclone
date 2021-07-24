// Package lotus provides an interface to the Filecoin lotus client
package lotus

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"path"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"
	"github.com/rclone/rclone/fs"
	"github.com/rclone/rclone/fs/config"
	"github.com/rclone/rclone/fs/config/configmap"
	"github.com/rclone/rclone/fs/config/configstruct"
	"github.com/rclone/rclone/fs/config/obscure"
	"github.com/rclone/rclone/fs/fserrors"
	"github.com/rclone/rclone/fs/fshttp"
	"github.com/rclone/rclone/fs/hash"
	"github.com/rclone/rclone/fs/operations"
	"github.com/rclone/rclone/fs/walk"
	"github.com/rclone/rclone/lib/atexit"
	"github.com/rclone/rclone/lib/bucket"
	"github.com/rclone/rclone/lib/encoder"
	"github.com/rclone/rclone/lib/oauthutil"
	"github.com/rclone/rclone/lib/pacer"
	"github.com/rclone/rclone/lib/readers"
	"github.com/rclone/rclone/lib/rest"
	"github.com/rclone/rclone/lib/structs"
)

const (
)

// Register with Fs
func init() {
	fs.Register(&fs.RegInfo{
		Name:        "lotus",
		Description: "Lotus long description",
		NewFs:       NewFs,
		Options: []fs.Option{{
			Name: "miner",
			Help: "Miner address",
			Required: true,
		}, {
			Name: "api-url",
			Help: "API url",
			Default: nil,
		}/*, {
			Name: "option_name",
			Help: fmt.Sprintf(`Option help. %v

Option help.`, option_help_value),
E			Examples: []fs.OptionExmaple{{
				Value: "OptionalExamplesValue",
				Help: "OptionalExamplesDescription"
			}}
			Default: optionalDefault,
			Advanced: optionalTrue
		}*/},
		CommandHelp: commandHelp,
		//Config: func(ctx context.Context, name string, m configmap.Mapper, config fs.ConfigIn) (*fs.ConfigOut, error) {
		//},
	})
}

// Options defines the configuration for this backend
type Options struct {
	Miner		string			`config:"miner"`
	ApiUrl		string			`config:"api-url"`
	//OptionName	optiontype		`config::"option_name"`
}

// Fs represents a remote lotus provider
type Fs struct {
	name           string         // name of this provider
	root           string         // path we are working on if any
	opt            Options        // parsed options
	ci             *fs.ConfigInfo // global config
	features        *fs.Features  // optional features
	//rootCID		string	      // cid for root directory
	// add more stuff
	//ctx            context.Context // global context for reading config
	// some server connection
}

// Name of the remote (as passed into NewFs)
func (f *Fs) Name() string {
	return f.name
}

// Root of the remote (as passed into NewFs)
func (f *Fs) Root() string {
	return f.root
}

// String converts this Fs to a string
func (f *Fs) String() string {
	return fmt.Sprintf("Lotus miner %s", f.opt.Miner)
}

// Features returns the optional features of this Fs
func (f *Fs) Features() *fs.Features {
	return f.features
}

// NewFs constructs an Fs from the path, bucket:path
func NewFs(ctx context.Context, name, root string, m configmap.Mapper) (fs.Fs, error) {
	// Parse config into Options struct
	opt := new(Options)
	err := configstruct.Set(m, opt)
	if err != nil {
		return nil, err
	}

	if opt.IsSet(
	// FULLNODE_API_INFO
	// MINER_API_INFO
	// WORKER_API_INFO


	ci := fs.GetConfig(ctx)
	f := &fs{
		name:		name,
		opt:		*opt,
		ci:		ci
	}
}
