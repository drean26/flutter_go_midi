package flutter_go_midi

import (
	"fmt"

	flutter "github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"

	// "reflect"
	"gitlab.com/gomidi/midi"
	"gitlab.com/gomidi/midi/reader"
	"gitlab.com/gomidi/midi/writer"
	driver "gitlab.com/gomidi/rtmididrv"
)

const channelName = "flutter_go_midi"
const channelEventName = "flutter_go_midi_event"

var wr *writer.Writer
var eventSink *plugin.EventSink

var drv *driver.Driver
var rd *reader.Reader
var err error

var out midi.Out
var in midi.In

// FlutterGoMidiPlugin implements flutter.Plugin and handles method.
type FlutterGoMidiPlugin struct {
}

var _ flutter.Plugin = &FlutterGoMidiPlugin{} // compile-time type check

// InitPlugin initializes the plugin.
func (p *FlutterGoMidiPlugin) InitPlugin(messenger plugin.BinaryMessenger) error {
	drv, err = driver.New()

	channel := plugin.NewMethodChannel(messenger, channelName, plugin.StandardMethodCodec{})

	eventChannel := plugin.NewEventChannel(messenger, channelEventName, plugin.StandardMethodCodec{})
	eventChannel.Handle(p)

	channel.HandleFunc("getPlatformVersion", p.handlePlatformVersion)
	channel.HandleFunc("outs", p.outs)
	channel.HandleFunc("ins", p.ins)
	channel.HandleFunc("openInPort", p.openInPort)
	channel.HandleFunc("openOutPort", p.openOutPort)
	channel.HandleFunc("closeInPort", p.closeInPort)
	channel.HandleFunc("closeOutPort", p.closeOutPort)
	channel.HandleFunc("listen", p.listen)
	channel.HandleFunc("noteOn", p.noteOn)
	channel.HandleFunc("noteOff", p.noteOff)
	channel.HandleFunc("lightOn", p.lightOn)
	channel.HandleFunc("lightOff", p.lightOff)
	channel.HandleFunc("lightOffAll", p.lightOffAll)
	channel.HandleFunc("initPopPiano", p.initPopPiano)

	rd = reader.New(
		// reader.Each(func(_ *reader.Position, msg midi.Message) {
		// 	fmt.Println("Each", msg.String())
		// }),

		reader.NoteOn(func(p *reader.Position, channel, key, velocity uint8) {
			// fmt.Println("NoteOn", channel, key, velocity)
			event := map[interface{}]interface{}{
				"event":    "NoteOn",
				"channel":  int32(channel),
				"key":      int32(key),
				"velocity": int32(velocity),
			}
			eventSink.Success(event)
		}),
		reader.NoteOff(func(p *reader.Position, channel, key, velocity uint8) {
			// fmt.Println("NoteOn", channel, key, velocity)
			event := map[interface{}]interface{}{
				"event":    "NoteOff",
				"channel":  int32(channel),
				"key":      int32(key),
				"velocity": int32(velocity),
			}
			eventSink.Success(event)
		}),
	)

	return nil
}

func (p *FlutterGoMidiPlugin) OnListen(arguments interface{}, sink *plugin.EventSink) {
	eventSink = sink
}

func (p *FlutterGoMidiPlugin) OnCancel(arguments interface{}) {

}

func (p *FlutterGoMidiPlugin) handlePlatformVersion(arguments interface{}) (reply interface{}, err error) {
	return "go-flutter " + flutter.PlatformVersion, nil
}

///获取输出端口
func (p *FlutterGoMidiPlugin) outs(arguments interface{}) (reply interface{}, err error) {
	// drv, err := driver.New()
	outs, err := drv.Outs()
	var size = len(outs)
	list := make([]interface{}, size)
	for i := 0; i < len(outs); i++ {
		list[i] = outs[i].String()
	}
	return list, nil
}

///获取输入端口
func (p *FlutterGoMidiPlugin) ins(arguments interface{}) (reply interface{}, err error) {
	// drv, err := driver.New()
	ins, err := drv.Ins()
	var size = len(ins)
	list := make([]interface{}, size)
	for i := 0; i < len(ins); i++ {
		list[i] = ins[i].String()
	}
	return list, nil
}

///打开输出端口
func (p *FlutterGoMidiPlugin) openOutPort(arguments interface{}) (reply interface{}, err error) {
	argsMap := arguments.(map[interface{}]interface{})
	port := argsMap["port"].(int32)
	// drv, err := driver.New()
	outs, err := drv.Outs()
	out = outs[port]
	err = out.Open()
	if err != nil {
		fmt.Println("打开端口失败")
		return false, nil
	}
	wr = writer.New(out)
	return true, nil
}

///打开输入端口
func (p *FlutterGoMidiPlugin) openInPort(arguments interface{}) (reply interface{}, err error) {
	argsMap := arguments.(map[interface{}]interface{})
	port := argsMap["port"].(int32)
	// drv, err := driver.New()
	ins, err := drv.Ins()
	in = ins[port]
	err = in.Open()
	if err != nil {
		fmt.Println("打开端口失败")
		return false, nil
	}
	rd.ListenTo(in)
	return true, nil
}

///关闭输出端口
func (p *FlutterGoMidiPlugin) closeOutPort(arguments interface{}) (reply interface{}, err error) {
	if out == nil {
		return true, nil
	}
	err = out.Close()
	if err != nil {
		fmt.Println("关闭端口失败")
		return false, nil
	}
	return true, nil
}

///关闭输入端口
func (p *FlutterGoMidiPlugin) closeInPort(arguments interface{}) (reply interface{}, err error) {
	if in == nil {
		return true, nil
	}
	err = in.Close()
	if err != nil {
		fmt.Println("关闭端口失败")
		return false, nil
	}
	return true, nil
}

///监听输入端口
func (p *FlutterGoMidiPlugin) listen(arguments interface{}) (reply interface{}, err error) {
	if in == nil {
		fmt.Println("端口未打开")
		return false, nil
	}
	rd.ListenTo(in)
	return true, nil
}

func (p *FlutterGoMidiPlugin) noteOn(arguments interface{}) (reply interface{}, err error) {
	argsMap := arguments.(map[interface{}]interface{})
	number := argsMap["number"].(int32)
	velocity := argsMap["velocity"].(int32)
	// fmt.Println("noteOn", number, velocity)
	writer.NoteOn(wr, uint8(number), uint8(velocity))
	return true, nil
}

func (p *FlutterGoMidiPlugin) noteOff(arguments interface{}) (reply interface{}, err error) {
	argsMap := arguments.(map[interface{}]interface{})
	number := argsMap["number"].(int32)
	// fmt.Println("%d off", number)
	writer.NoteOff(wr, uint8(number))
	return true, nil
}

func (p *FlutterGoMidiPlugin) initPopPiano(arguments interface{}) (reply interface{}, err error) {
	initPiano := []midi.Message{NewDefauleTurnOn1(), NewDefauleTurnOn2()}
	writer.WriteMessages(wr, initPiano)
	return true, nil
}

func (p *FlutterGoMidiPlugin) lightOn(arguments interface{}) (reply interface{}, err error) {
	argsMap := arguments.(map[interface{}]interface{})
	number := argsMap["number"].(int32)
	// fmt.Println("%d off", number)
	msg := []midi.Message{LightControl{key: uint8(number), turnOn: 0x11}}
	writer.WriteMessages(wr, msg)
	return true, nil
}

func (p *FlutterGoMidiPlugin) lightOff(arguments interface{}) (reply interface{}, err error) {
	argsMap := arguments.(map[interface{}]interface{})
	number := argsMap["number"].(int32)
	// fmt.Println("%d off", number)
	msg := []midi.Message{LightControl{key: uint8(number), turnOn: 0x00}}
	writer.WriteMessages(wr, msg)
	return true, nil
}

func (p *FlutterGoMidiPlugin) lightOffAll(arguments interface{}) (reply interface{}, err error) {
	lightOff := []midi.Message{NewDefauleLightOff()}
	writer.WriteMessages(wr, lightOff)
	return true, nil
}

type LightControl struct {
	key    uint8
	turnOn uint8
}

type TurnOn struct {
	data []byte
}

func NewDefauleTurnOn1() TurnOn {
	return TurnOn{data: []byte{0xf0, 0x0f, 0x2e, 0x63, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf7}}
}

func NewDefauleTurnOn2() TurnOn {
	return TurnOn{data: []byte{0xf0, 0x07, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf7}}
}

func NewDefauleLightOff() TurnOn {
	return TurnOn{data: []byte{0xf0, 0x4d, 0x4c, 0x4e, 0x45, 0x30, 0x03, 0x00, 0xf7}}
}

// Raw returns the bytes for the noteon message.
func (n TurnOn) Raw() []byte {
	return n.data
}

// String returns human readable information about the note-on message.
func (n TurnOn) String() string {
	return "连接成功初始化命令"
}

// Raw returns the bytes for the noteon message.
func (n LightControl) Raw() []byte {
	return []uint8{0xf0, 0x4d, 0x4c, 0x4e, 0x45, n.key, n.turnOn, 0x00, 0xf7}
}

// String returns human readable information about the note-on message.
func (n LightControl) String() string {
	return "灯光"
}
