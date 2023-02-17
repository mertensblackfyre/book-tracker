// ignore_for_file: prefer_const_constructors
import 'package:flutter/material.dart';
import 'package:english_words/english_words.dart';

class RandomWords extends StatefulWidget {
  @override
  RandomWordState createState() => RandomWordState();
}

class RandomWordState extends State<RandomWords> {
  final __randomWord = <WordPair>[];
  final __savedWordPairs = <WordPair>{};

  Widget __buildList() {
    return ListView.builder(
        padding: const EdgeInsets.all(16),
        itemBuilder: (context, item) {
          if (item.isOdd) return Divider();
          final index = item ~/ 2;
          if (index >= __randomWord.length) {
            __randomWord.addAll(generateWordPairs().take(10));
          }
          return __buildRow(__randomWord[index]);
        });
  }

  Widget __buildRow(WordPair pair) {
    final alreadySaved = __savedWordPairs.contains(pair);
    return ListTile(
      title: Text(
        pair.asPascalCase,
        style: TextStyle(fontSize: 18),
      ),
      trailing: Icon(alreadySaved ? Icons.favorite : Icons.favorite_border,
          color: alreadySaved ? Colors.red : null),
      onTap: () {
        setState(() {
          if (alreadySaved) {
            __savedWordPairs.remove(pair);
          } else {
            __savedWordPairs.add(pair);
          }
        });
      },
    );
  }

  void __pushSaved() {
    Navigator.of(context).push(MaterialPageRoute(builder: (BuildContext) {
      final Iterable<ListTile> tiles = __savedWordPairs.map((WordPair pair) {
        return ListTile(
            title: Text(pair.asPascalCase, style: TextStyle(fontSize: 16.0)));
      });
      final List<Widget> divided =
          ListTile.divideTiles(context: context, tiles: tiles).toList();

      return Scaffold(
          appBar: AppBar(title: Text('Saved WordPairs')),
          body: ListView(children: divided));
    }));
  }

  Widget build(BuildContext context) {
    return Scaffold(
        appBar: AppBar(
          title: Text('WordPair Generator'),
          actions: <Widget>[
            IconButton(icon: Icon(Icons.list), onPressed: __pushSaved)
          ],
        ),
        body: __buildList());
  }
}
