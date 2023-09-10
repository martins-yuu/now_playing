import 'dart:async';
import 'dart:convert';
import 'dart:io';

import 'package:http/http.dart' as http;

void main() {
  final channel = File('channel.txt').readAsStringSync().trim();
  final url = Uri.parse(
      'https://api.streamelements.com/kappa/v2/songrequest/$channel/playing');
  final file = File('now_playing.txt');
  final delay = Duration(seconds: 5);
  final spacer = ' ' * 10;

  Timer.periodic(delay, (_) async {
    try {
      final response = await http.get(url);

      if (response.statusCode == HttpStatus.ok) {
        final data = json.decode(response.body);
        final title = data['title'];

        await file.writeAsString('$title$spacer');
        print('Now playing: $title');
      } else {
        await file.writeAsString('');
        print('Error: ${response.reasonPhrase}');
      }
    } catch (e) {
      print(e);
    }
  });
}
