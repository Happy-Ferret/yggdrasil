package main

import (
	"flag"

	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilectron-bootstrap"
	"github.com/asticode/go-astilog"
	"github.com/pkg/errors"
	i18n "github.com/middangeard-fiction/rosetta-go"
)

var (
	AppName string
	BuiltAt string
	debug   = flag.Bool("d", false, "enables the debug mode")
	w       *astilectron.Window
	a       *astilectron.Astilectron
)

func main() {
	// Shortcuts
	newAccel := astilectron.NewAccelerator("ctrl+n")
	openAccel := astilectron.NewAccelerator("ctrl+o")
	saveAccel := astilectron.NewAccelerator("ctrl+s")
	quitAccel := astilectron.NewAccelerator("ctrl+q")
	importAccel := astilectron.NewAccelerator("ctrl+shift+o")

	// Init
	flag.Parse()
	astilog.FlagInit()

	// Run bootstrap
	astilog.Debugf("Running app built at %s", BuiltAt)
	if err := bootstrap.Run(bootstrap.Options{
		Asset:    Asset,
		AssetDir: AssetDir,
		AstilectronOptions: astilectron.Options{
			AppName:            AppName,
			AppIconDarwinPath:  "resources/icon.icns",
			AppIconDefaultPath: "resources/icon.png",
		},
		Debug: *debug,
		MenuOptionsFunc: func(app *astilectron.Astilectron) []*astilectron.MenuItemOptions {
			a = app
			return []*astilectron.MenuItemOptions{{
				Label: astilectron.PtrStr("File"),
				SubMenu: []*astilectron.MenuItemOptions{
					{
						Label:       astilectron.PtrStr("New Map"),
						Accelerator: newAccel,
						OnClick: func(e astilectron.Event) (deleteListener bool) {
							if err := bootstrap.SendMessage(w, "new_map", nil); err != nil {
								astilog.Error(errors.Wrap(err, "sending new_map event failed"))
							}
							return false
						},
					},
					{Type: astilectron.MenuItemTypeSeparator},
					{
						Label:       astilectron.PtrStr("Open Map"),
						Accelerator: openAccel,
						OnClick: func(e astilectron.Event) (deleteListener bool) {
							if err := bootstrap.SendMessage(w, "open_map", "yggdrasil"); err != nil {
								astilog.Error(errors.Wrap(err, "sending open_map event failed"))
							}
							return false
						},
					},
					{
						Label:       astilectron.PtrStr("Open Trizbort Map"),
						Accelerator: importAccel,
						OnClick: func(e astilectron.Event) (deleteListener bool) {
							if err := bootstrap.SendMessage(w, "open_map", "trizbort"); err != nil {
								astilog.Error(errors.Wrap(err, "sending open_map event failed"))
							}
							return false
						},
					},
					{Type: astilectron.MenuItemTypeSeparator},
					{
						Label:       astilectron.PtrStr("Save As"),
						Accelerator: saveAccel,
						OnClick: func(e astilectron.Event) (deleteListener bool) {
							if err := bootstrap.SendMessage(w, "save_map", nil); err != nil {
								astilog.Error(errors.Wrap(err, "sending save_map event failed"))
							}
							return false
						},
					},
					{
						Label: astilectron.PtrStr("Save As Image"),
						OnClick: func(e astilectron.Event) (deleteListener bool) {
							if err := bootstrap.SendMessage(w, "export_map", "image"); err != nil {
								astilog.Error(errors.Wrap(err, "sending export_map event failed"))
							}
							return false
						},
					},
					{
						Label: astilectron.PtrStr("Export"),
						SubMenu: []*astilectron.MenuItemOptions{
							{
								Label: astilectron.PtrStr("Middangeard 1"),
								OnClick: func(e astilectron.Event) (deleteListener bool) {
									if err := bootstrap.SendMessage(w, "export_map", "middangeard1"); err != nil {
										astilog.Error(errors.Wrap(err, "sending export_map event failed"))
									}
									return false
								},
							},
							{Type: astilectron.MenuItemTypeSeparator},
							{
								Label: astilectron.PtrStr("TADS"),
								OnClick: func(e astilectron.Event) (deleteListener bool) {
									if err := bootstrap.SendMessage(w, "export_map", "tads"); err != nil {
										astilog.Error(errors.Wrap(err, "sending export_map event failed"))
									}
									return false
								},
							},
							{
								Label: astilectron.PtrStr("Inform 7"),
								OnClick: func(e astilectron.Event) (deleteListener bool) {
									if err := bootstrap.SendMessage(w, "export_map", "i7"); err != nil {
										astilog.Error(errors.Wrap(err, "sending export_map event failed"))
									}
									return false
								},
							},
							{
								Label: astilectron.PtrStr("Alan 2"),
								OnClick: func(e astilectron.Event) (deleteListener bool) {
									if err := bootstrap.SendMessage(w, "export_map", "a2"); err != nil {
										astilog.Error(errors.Wrap(err, "sending export_map event failed"))
									}
									return false
								},
							},
							{
								Label: astilectron.PtrStr("Alan 3"),
								OnClick: func(e astilectron.Event) (deleteListener bool) {
									if err := bootstrap.SendMessage(w, "export_map", "a3"); err != nil {
										astilog.Error(errors.Wrap(err, "sending export_map event failed"))
									}
									return false
								},
							},
							{
								Label: astilectron.PtrStr("Quest"),
								OnClick: func(e astilectron.Event) (deleteListener bool) {
									if err := bootstrap.SendMessage(w, "export_map", "quest"); err != nil {
										astilog.Error(errors.Wrap(err, "sending export_map event failed"))
									}
									return false
								},
							},
						},
					},
					{Type: astilectron.MenuItemTypeSeparator},
					{Type: astilectron.MenuItemTypeSeparator},
					{
						Label:       astilectron.PtrStr("Exit"),
						Accelerator: quitAccel,
						OnClick: func(e astilectron.Event) (deleteListener bool) {
							if err := bootstrap.SendMessage(w, "prompt_quit", nil); err != nil {
								astilog.Error(errors.Wrap(err, "sending prompt_quit event failed"))
							}
							return false
						},
					},
				},
			},
				{
					Label: astilectron.PtrStr("Edit"),
					SubMenu: []*astilectron.MenuItemOptions{
						{Role: astilectron.MenuItemRoleCopy},
						{Role: astilectron.MenuItemRolePaste},
					},
				},
				{
					Label: astilectron.PtrStr("Maps"),
					SubMenu: []*astilectron.MenuItemOptions{
						{
							Label: astilectron.PtrStr("Castle of Doom"),
							OnClick: func(e astilectron.Event) (deleteListener bool) {
								if err := bootstrap.SendMessage(w, "load_map", "castleofdoom"); err != nil {
									astilog.Error(errors.Wrap(err, "sending load_map event failed"))
								}
								return false
							},
						},
						{
							Label: astilectron.PtrStr("Colossal Cave Adventure"),
							OnClick: func(e astilectron.Event) (deleteListener bool) {
								if err := bootstrap.SendMessage(w, "load_map", "adventure"); err != nil {
									astilog.Error(errors.Wrap(err, "sending load_map event failed"))
								}
								return false
							},
						},
						{
							Label: astilectron.PtrStr("The Hitchhiker’s Guide to the Galaxy"),
							OnClick: func(e astilectron.Event) (deleteListener bool) {
								if err := bootstrap.SendMessage(w, "load_map", "hhg"); err != nil {
									astilog.Error(errors.Wrap(err, "sending load_map event failed"))
								}
								return false
							},
						},
						{
							Label: astilectron.PtrStr("The Hobbit"),
							OnClick: func(e astilectron.Event) (deleteListener bool) {
								if err := bootstrap.SendMessage(w, "load_map", "hobbit"); err != nil {
									astilog.Error(errors.Wrap(err, "sending load_map event failed"))
								}
								return false
							},
						},
						{
							Label: astilectron.PtrStr("Zork"),
							OnClick: func(e astilectron.Event) (deleteListener bool) {
								if err := bootstrap.SendMessage(w, "load_map", "zork"); err != nil {
									astilog.Error(errors.Wrap(err, "sending load_map event failed"))
								}
								return false
							},
						},
					},
				},
				{
					Label: astilectron.PtrStr("Settings"),
					SubMenu: []*astilectron.MenuItemOptions{
						{
							Label: astilectron.PtrStr("Map Settings"),
							OnClick: func(e astilectron.Event) (deleteListener bool) {
								if err := bootstrap.SendMessage(w, "open_settings", "map"); err != nil {
									astilog.Error(errors.Wrap(err, "sending open_settings event failed"))
								}
								return false
							},
						},
						{
							Label: astilectron.PtrStr("Render Settings"),
							OnClick: func(e astilectron.Event) (deleteListener bool) {
								if err := bootstrap.SendMessage(w, "open_settings", "render"); err != nil {
									astilog.Error(errors.Wrap(err, "sending open_settings event failed"))
								}
								return false
							},
						},
					},
				},
				{
					Label: astilectron.PtrStr("Help"),
					SubMenu: []*astilectron.MenuItemOptions{
						{
							Label: astilectron.PtrStr("About"),
							OnClick: func(e astilectron.Event) (deleteListener bool) {
								if err := bootstrap.SendMessage(w, "open_window", "about"); err != nil {
									astilog.Error(errors.Wrap(err, "sending load_map event failed"))
								}
								return false
							},
						},
						{
							Label: astilectron.PtrStr("Changelog"),
							OnClick: func(e astilectron.Event) (deleteListener bool) {
								if err := bootstrap.SendMessage(w, "open_window", "changelog"); err != nil {
									astilog.Error(errors.Wrap(err, "sending load_map event failed"))
								}
								return false
							},
						},
					},
				}}
		},
		OnWait: func(app *astilectron.Astilectron, ws []*astilectron.Window, _ *astilectron.Menu, _ *astilectron.Tray, _ *astilectron.Menu) error {
			w = ws[0]
			return nil
		},
		RestoreAssets: RestoreAssets,
		Windows: []*bootstrap.Window{{
			Homepage:       "index.html?lang=" + i18n.GetUILanguage(),
			MessageHandler: handleMessages,
			Options: &astilectron.WindowOptions{
				BackgroundColor: astilectron.PtrStr("#5cc070"),
				Center:          astilectron.PtrBool(true),
				Height:          astilectron.PtrInt(700),
				Width:           astilectron.PtrInt(700),
			},
		}},
	}); err != nil {
		astilog.Fatal(errors.Wrap(err, "running bootstrap failed"))
	}
}
