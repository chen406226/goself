import 'package:flutter/material.dart';
import 'package:uuid/uuid.dart';

/// Message is class defining message data (id and text)
class Message {
  /// id is unique ID of message
  final String id;

  static var _uuid = Uuid();
  /// text is content of message
  final String text;

  /// Class constructor
  Message({String id, @required String text})
//      : this.id = id ?? UniqueKey().toString(),
      : this.id = id ?? _uuid.v4(),
        this.text = text;
}
