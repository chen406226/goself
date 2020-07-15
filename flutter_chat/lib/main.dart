//import 'package:flutter/material.dart';
//
//import 'api/chat_service.dart';
//
//import 'blocs/application_bloc.dart';
//import 'blocs/bloc_provider.dart';
//import 'blocs/message_events.dart';
//
//import 'pages/home.dart';
//
//import 'theme.dart';
//
import 'dart:async';
import 'dart:convert';
import 'dart:io';
import 'dart:typed_data';

///// main is entry point of Flutter application
//void main() {
//  return runApp(BlocProvider<ApplicationBloc>(
//    bloc: ApplicationBloc(),
//    child: App(),
//  ));
//}
//
//// Stateful application widget
//class App extends StatefulWidget {
//  @override
//  State<StatefulWidget> createState() => _AppState();
//}
//
//// State for application widget
//class _AppState extends State<App> {
//  // BLoc for application
//  ApplicationBloc _appBloc;
//
//  /// Chat client service
//  ChatService _service;
//
//  bool _isInit = false;
//
//  @override
//  void didChangeDependencies() {
//    super.didChangeDependencies();
//
//    // As the context of not yet available at initState() level,
//    // if not yet initialized, we get application BLoc to start
//    // gRPC isolates
//    if (_isInit == false) {
//      _appBloc = BlocProvider.of<ApplicationBloc>(context);
//
//      // initialize Chat client service
//      _service = ChatService(
//          onMessageSent: _onMessageSent,
//          onMessageSendFailed: _onMessageSendFailed,
//          onMessageReceived: _onMessageReceived,
//          onMessageReceiveFailed: _onMessageReceiveFailed);
//      _service.start();
//
//      _listenMessagesToSend();
//
//      if (mounted) {
//        setState(() {
//          _isInit = true;
//        });
//      }
//    }
//  }
//
//  void _listenMessagesToSend() async {
//    await for (var event in _appBloc.outMessageSend) {
//      _service.send(event.message);
//    }
//  }
//
//  @override
//  Widget build(BuildContext context) {
//    return MaterialApp(
//      title: 'Friendlychat',
//      theme: isIOS(context) ? kIOSTheme : kDefaultTheme,
//      home: HomePage(),
//    );
//  }
//
//  @override
//  void dispose() {
//    // close Chat client service
//    _service.shutdown();
//    _service = null;
//
//    super.dispose();
//  }
//
//  /// 'outgoing message sent to the server' event
//  void _onMessageSent(MessageSentEvent event) {
//    debugPrint('Message "${event.id}" sent to the server');
//    _appBloc.inMessageSent.add(event);
//  }
//
//  /// 'failed to send message' event
//  void _onMessageSendFailed(MessageSendFailedEvent event) {
//    debugPrint(
//        'Failed to send message "${event.id}" to the server: ${event.error}');
//    _appBloc.inMessageSendFailed.add(event);
//  }
//
//  /// 'new incoming message received from the server' event
//  void _onMessageReceived(MessageReceivedEvent event) {
//    debugPrint('Message received from the server: ${event.text}');
//    _appBloc.inMessageReceived.add(event);
//  }
//
//  /// 'failed to receive messages' event
//  void _onMessageReceiveFailed(MessageReceiveFailedEvent event) {
//    debugPrint('Failed to receive messages from the server: ${event.error}');
//  }
//}
//
//
//
//
//





import 'package:flutter/material.dart';
void main() => runApp(App());
class App extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter Demo',
      theme: ThemeData(
        primarySwatch: Colors.blue,
      ),
      home: MyHomePage(title: 'Flutter Demo Home Page'),
    );
  }
}

class MyHomePage extends StatefulWidget {
  MyHomePage({Key key, this.title}) : super(key: key);

  final String title;

  @override
  _MyHomePageState createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  int _counter = 0;
  void _incrementCounter() {
//      socket.write("kslfjlklsaf;lkj;ljfas");
//    socket.writeln('Sendasfdasfdfasdafsdafsd.');

//      socket.flush();
    socket.writeln('Send.');

//    socket.close();
    setState(() {
//      socket.destroy();
//      _counter++;
    });
    print("insert");
//    socket.writeCharCode(123);

  }

  Socket socket;

  @override
  void dispose() {
    // TODO: implement dispose
//    socket.close();
    super.dispose();
  }

  @override
  void initState() {
    // TODO: implement initState
    contcp();
    super.initState();
  }

  handleMsg(msg) {
    print('Message received: $msg');
  }

  contcp () async{
    print("start");
//    Socket.connect('192.168.127.103', 8000).then((value) => {
//      value.write("Hello imsocket");
//      print(value);
//      value.listen
//    });
    socket = await Socket.connect('192.168.127.103', 8000);
    print('connected');

    // listen to the received data event stream
    socket.listen((List<int> event) {
//      print(utf8.decode(event));
      setState(() {
        msglist.add(utf8.decode(event));
      });
//      socket.add(utf8.encode('hello'));
//      socket.write("writesocket");
    });
//    var timer = new Timer.periodic(
//        new Duration(seconds: 1),
//            (_)  {socket.writeln('Send.'); timer.cancel();});

    // send hello
//    var message = Uint8List(4);
//    var bytedata = ByteData.view(message.buffer);
//
//    bytedata.setUint8(0, 0x01);
//    bytedata.setUint8(1, 0x07);
//    bytedata.setUint8(2, 0xFF);
//    bytedata.setUint8(3, 0x88);

//    socket.add(message);
//    socket.addStream(Stream<List<int>> stream);
//      var server = await HttpServer.bind('192.168.127.103', 8000);
//    print('connected');
//  server.listen((event) { })
//
//      await for (var req in server) {
//        print(req.uri.path);
//        if (req.uri.path == '/ws') {
//          // Upgrade a HttpRequest to a WebSocket connection.
//          var socket = await WebSocketTransformer.upgrade(req);
//          socket.listen(handleMsg);
//        };
//      }
//    print('end');



//    socket.listen((List<int> event) {
//      print(utf8.decode(event));
//    });
//
//    socket.add(utf8.encode('hello'));

//    await Future.delayed(Duration(seconds: 5));

//    socket.close();
  }
  List msglist=[];

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(widget.title),
      ),
      body: Center(
//        child: Column(
//          mainAxisAlignment: MainAxisAlignment.center,
//          children: <Widget>[
//            Text(
//              'You have pushed the button this many times:',
//            ),
//            Text(
//              '$_counter',
//              style: Theme.of(context).textTheme.display1,
//            ),
//          ],
//        ),
      child: ListView.builder(
          itemCount: msglist.length,
          itemBuilder: (context, index){
            return Container(
              child: Text(msglist[index]),
            );
          }
      ),
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: _incrementCounter,
        tooltip: 'Increment',
        child: Icon(Icons.add),
      ), // This trailing comma makes auto-formatting nicer for build methods.
    );
  }
}














