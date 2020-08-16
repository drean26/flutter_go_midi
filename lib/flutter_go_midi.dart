// You have generated a new plugin project without
// specifying the `--platforms` flag. A plugin project supports no platforms is generated.
// To add platforms, run `flutter create -t plugin --platforms <platforms> .` under the same
// directory. You can also find a detailed instruction on how to add platforms in the `pubspec.yaml` at https://flutter.dev/docs/development/packages-and-plugins/developing-packages#plugin-platforms.

import 'dart:async';

import 'package:flutter/services.dart';

class FlutterGoMidi {
  static const MethodChannel _channel = const MethodChannel('flutter_go_midi');

  static const EventChannel eventChannel =
      EventChannel('flutter_go_midi_event');

  static Future<String> get platformVersion async {
    final String version = await _channel.invokeMethod('getPlatformVersion');
    return version;
  }

  static Future<bool> init() async {
    final bool version = await _channel.invokeMethod('init');
    return version;
  }

  static Future<List<dynamic>> ins() async {
    final List<dynamic> version = await _channel.invokeMethod('ins');
    return version;
  }

  static Future<bool> listen(int port) async {
    final bool version = await _channel.invokeMethod('listen', {
      'port': port,
    });
    return version;
  }

  static Future<List<dynamic>> outs() async {
    final List<dynamic> version = await _channel.invokeMethod('outs');
    return version;
  }

  static Future<bool> openInPort(int port) async {
    final bool version = await _channel.invokeMethod('openInPort', {'port': port});
    return version;
  }

  static Future<bool> openOutPort(int port) async {
    final bool version = await _channel.invokeMethod('openOutPort', {'port': port});
    return version;
  }

  static Future<bool> closeInPort(int port) async {
    final bool version = await _channel.invokeMethod('closeInPort', {'port': port});
    return version;
  }
  static Future<bool> closeOutPort(int port) async {
    final bool version = await _channel.invokeMethod('closeOutPort', {'port': port});
    return version;
  }

  static Future<bool> closeAllDevice() async {
    final bool version = await _channel.invokeMethod('closeAllDevice');
    return version;
  }

  static Future<bool> noteOn(int number, int velocity) async {
    final bool version = await _channel.invokeMethod('noteOn', {
      'number': number,
      'velocity': velocity,
    });
    return version;
  }

  static Future<bool> noteOff(int number) async {
    final bool version = await _channel.invokeMethod('noteOff', {
      'number': number,
    });
    return version;
  }

  static Future<bool> lightOn(int number) async {
    final bool version = await _channel.invokeMethod('lightOn', {
      'number': number,
    });
    return version;
  }

  static Future<bool> lightOff(int number) async {
    final bool version = await _channel.invokeMethod('lightOff', {
      'number': number,
    });
    return version;
  }

  static Future<bool> lightOffAll() async {
    final bool version = await _channel.invokeMethod('lightOffAll');
    return version;
  }

  static Future<bool> initPopPiano() async {
    final bool version = await _channel.invokeMethod('initPopPiano');
    return version;
  }
}
