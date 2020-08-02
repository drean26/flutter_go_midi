// You have generated a new plugin project without
// specifying the `--platforms` flag. A plugin project supports no platforms is generated.
// To add platforms, run `flutter create -t plugin --platforms <platforms> .` under the same
// directory. You can also find a detailed instruction on how to add platforms in the `pubspec.yaml` at https://flutter.dev/docs/development/packages-and-plugins/developing-packages#plugin-platforms.

import 'dart:async';

import 'package:flutter/services.dart';

class FlutterGoMidi {
  static const MethodChannel _channel = const MethodChannel('flutter_go_midi');

  static const EventChannel _eventChannel =
      EventChannel('flutter_go_midi_event');

  static Future<String> get platformVersion async {
    final String version = await _channel.invokeMethod('getPlatformVersion');
    return version;
  }

  static Future<String> noteOn(int number) async {
    final String version = await _channel.invokeMethod('noteOn', {
      'number': number,
    });
    return version;
  }

  static Future<String> noteOff(int number) async {
    final String version = await _channel.invokeMethod('noteOff', {
      'number': number,
    });
    return version;
  }
}
