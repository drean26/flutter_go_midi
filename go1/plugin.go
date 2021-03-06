package flutte_go_midi

import (
	"github.com/gen2brain/dlgs"
	"github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"
	"github.com/pkg/errors"
	
	"gitlab.com/gomidi/midi"
	"gitlab.com/gomidi/midi/writer"
	driver "gitlab.com/gomidi/rtmididrv"
	// when using portmidi, replace the line above with
	// driver gitlab.com/gomidi/portmididrv
	// "gitlab.com/gomidi/midi"
	// "gitlab.com/gomidi/midi/writer"
	// driver "gitlab.com/gomidi/rtmididrv"
	// // when using portmidi, replace the line above with
	// // driver gitlab.com/gomidi/portmididrv
)

const channelName = "flutter_go_midi"

type FilePickerPlugin struct{}

var _ flutter.Plugin = &FilePickerPlugin{} // compile-time type check

func (p *FilePickerPlugin) InitPlugin(messenger plugin.BinaryMessenger) error {
	channel := plugin.NewMethodChannel(messenger, channelName, plugin.StandardMethodCodec{})
	channel.CatchAllHandleFunc(p.handleFilePicker)
	return nil
}

func (p *FilePickerPlugin) handleFilePicker(methodCall interface{}) (reply interface{}, err error) {
	method := methodCall.(plugin.MethodCall).Method

	if "noteOn" == method {
		drv, err := driver.New()
		outs, err := drv.Outs()
		out := outs[0]
		out.Open()
		wr := writer.New(out)
		writer.NoteOn(wr, 60, 100)
	}

	if "dir" == method {
		dirPath, err := dirDialog("Select a directory")
		if err != nil {
			return nil, errors.Wrap(err, "failed to open dialog picker")
		}
		return dirPath, nil
	}

	arguments := methodCall.(plugin.MethodCall).Arguments.(map[interface{}]interface{})

	var allowedExtensions []string

	// Parse extensions
	if arguments != nil && arguments["allowedExtensions"] != nil {
		allowedExtensions = make([]string, len(arguments["allowedExtensions"].([]interface{})))
		for i := range arguments["allowedExtensions"].([]interface{}) {
			allowedExtensions[i] = arguments["allowedExtensions"].([]interface{})[i].(string)
		}
	}

	selectMultiple, ok := arguments["allowMultipleSelection"].(bool) //method.Arguments.(bool)
	if !ok {
		return nil, errors.Wrap(err, "invalid format for argument, not a bool")
	}

	filter, err := fileFilter(method, allowedExtensions, len(allowedExtensions), selectMultiple)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get filter")
	}

	if selectMultiple {
		filePaths, _, err := dlgs.FileMulti("Select one or more files", filter)
		if err != nil {
			return nil, errors.Wrap(err, "failed to open dialog picker")
		}

		// type []string is not supported by StandardMessageCodec
		sliceFilePaths := make([]interface{}, len(filePaths))
		for i, file := range filePaths {
			sliceFilePaths[i] = file
		}

		return sliceFilePaths, nil
	}

	filePath, err := fileDialog("Select a file", filter)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open dialog picker")
	}
	return filePath, nil
}
