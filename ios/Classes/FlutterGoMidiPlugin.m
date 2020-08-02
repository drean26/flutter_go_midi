#import "FlutterGoMidiPlugin.h"
#if __has_include(<flutter_go_midi/flutter_go_midi-Swift.h>)
#import <flutter_go_midi/flutter_go_midi-Swift.h>
#else
// Support project import fallback if the generated compatibility header
// is not copied when this plugin is created as a library.
// https://forums.swift.org/t/swift-static-libraries-dont-copy-generated-objective-c-header/19816
#import "flutter_go_midi-Swift.h"
#endif

@implementation FlutterGoMidiPlugin
+ (void)registerWithRegistrar:(NSObject<FlutterPluginRegistrar>*)registrar {
  [SwiftFlutterGoMidiPlugin registerWithRegistrar:registrar];
}
@end
