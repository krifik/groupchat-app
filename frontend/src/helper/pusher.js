Pusher.logToConsole = true;

var pusher = new Pusher("711add371e9613586624", {
  cluster: "ap1",
});

var channel = pusher.subscribe("group-chat");
channel.bind("groupchat", function (data) {
  // app.messages.push(JSON.stringify(data));
});

export default channel;
