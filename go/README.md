# flutter_go_midi

This Go package implements the host-side of the Flutter [flutter_go_midi](https://github.com/drean26/flutter_go_midi) plugin.

## Usage

Import as:

```go
import flutter_go_midi "github.com/drean26/flutter_go_midi/go"
```

Then add the following option to your go-flutter [application options](https://github.com/go-flutter-desktop/go-flutter/wiki/Plugin-info):

```go
flutter.AddPlugin(&flutter_go_midi.FlutterGoMidiPlugin{}),
```
