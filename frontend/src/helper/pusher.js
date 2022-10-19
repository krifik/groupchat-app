import Pusher from "pusher-js";

Pusher.logToConsole = true;

var pusher = new Pusher("711add371e9613586624", {
  cluster: "ap1",
});

export default pusher;
