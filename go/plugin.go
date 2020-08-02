package flutter_go_midi

import (
	flutter "github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"

	"gitlab.com/gomidi/midi"
	"gitlab.com/gomidi/midi/writer"
	driver "gitlab.com/gomidi/rtmididrv"
)

const channelName = "flutter_go_midi"
var wr *writer.Writer

// FlutterGoMidiPlugin implements flutter.Plugin and handles method.
type FlutterGoMidiPlugin struct{}

var _ flutter.Plugin = &FlutterGoMidiPlugin{} // compile-time type check

// InitPlugin initializes the plugin.
func (p *FlutterGoMidiPlugin) InitPlugin(messenger plugin.BinaryMessenger) error {
	fmt.
	channel := plugin.NewMethodChannel(messenger, channelName, plugin.StandardMethodCodec{})
	channel.HandleFunc("getPlatformVersion", p.handlePlatformVersion)
	channel.HandleFunc("init", p.handleInit)
	return nil
}

func (p *FlutterGoMidiPlugin) handlePlatformVersion(arguments interface{}) (reply interface{}, err error) {
	return "go-flutter " + flutter.PlatformVersion, nil
}

func (p *FlutterGoMidiPlugin) handleInit(arguments interface{}) (reply interface{}, err error) {
	drv, err := driver.New()
	outs, err := drv.Outs()
	out := outs[0]
	out.Open()
	wr = writer.New(out)
	return true, nil
}

func (p *FlutterGoMidiPlugin) handleNoteOn(arguments interface{}) (reply interface{}, err error) {
	argsMap := arguments.(map[interface{}]interface{})
	number := argsMap["number"].(int32)
	writer.NoteOn(wr, number, 100)
	return true, nil
}
