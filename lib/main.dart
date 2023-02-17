// ignore_for_file: prefer_const_constructors
import 'package:flutter/material.dart';
import './random.dart';

void main() => runApp(App());

class App extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
        theme: ThemeData(primaryColor: Colors.cyan), home: RandomWords());
  }
}
