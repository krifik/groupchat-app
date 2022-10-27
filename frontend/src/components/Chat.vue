<template>
  <div>
    <div class="flex-1 p:2 sm:p-6 justify-between flex flex-col h-screen">
      <div class="flex sm:items-center justify-between py-3 border-b-2 border-gray-200">
        <div class="relative flex items-center space-x-4">
          <div class="relative">
            <span class="absolute text-green-500 right-0 bottom-0">
              <svg width="20"
                   height="20">
                <circle cx="8"
                        cy="8"
                        r="8"
                        fill="currentColor"></circle>
              </svg>
            </span>
            <img src="/src/assets/pusher.png"
                 alt=""
                 class="w-10 sm:w-16 h-10 sm:h-16 rounded-full" />
          </div>
          <div class="flex flex-col leading-tight">
            <div class="text-2xl mt-1 flex items-center">
              <span class="text-gray-700 mr-3">Pusher Developer Team</span>
            </div>
            <span class="text-lg text-gray-600">{{ online }} Online</span>
          </div>
        </div>
        <div class="flex items-center space-x-2">
          <button @click="logout"
                  class="inline-flex items-center justify-center rounded-lg border h-10 w-10 transition duration-500 ease-in-out text-gray-500 hover:bg-gray-300 focus:outline-none">
            <svg version="1.1"
                 class="h-6 w-6"
                 id="Layer_1"
                 xmlns="http://www.w3.org/2000/svg"
                 xmlns:xlink="http://www.w3.org/1999/xlink"
                 x="0px"
                 y="0px"
                 viewBox="0 0 64 64"
                 enable-background="new 0 0 64 64"
                 xml:space="preserve">
              <g id="Exit_1_">
                <path d="M52.4501991,28.7678509l-5-4.9990005c-0.3768997-0.3770008-0.9902-0.3770008-1.3671989,0
		c-0.3778992,0.3778992-0.3778992,0.9902,0,1.3671989l3.3171997,3.3164005H35.2666016v2h14.1320992l-3.3157005,3.3163986
		c-0.3778992,0.377903-0.3778992,0.9902,0,1.3672028c0.1884995,0.1884995,0.4365997,0.2831993,0.6835976,0.2831993
		c0.2471008,0,0.4951019-0.0946999,0.6836014-0.2831993l5-5.0010014c0.1817017-0.1816006,0.2831993-0.4277,0.2831993-0.6835995
		C52.7333984,29.1946507,52.6319008,28.9495506,52.4501991,28.7678509z" />
                <path d="M40.2666016,39.4524498c-0.5527,0-1,0.4473-1,1v10.7900009c0,1.0429993-0.8310013,2.2099991-1.9433022,2.2099991
		h-6.0566998V11.2394505V9.8677502L30.0191994,9.33395L14.0765009,2.56445l-0.2606955-0.112h23.507494
		c1.2168007,0,1.9433022,0.9921999,1.9433022,1.9511998v15.0487995c0,0.5527,0.4473,1,1,1c0.5527992,0,1-0.4473,1-1V4.4036498
		c0-2.1786997-1.7685013-3.9511998-3.9433022-3.9511998H12.2666006c-0.5215998,0-0.9358997,0.4029-0.9822998,0.9124
		L11.2666006,1.35725V1.45245V55.03405l17.1855011,7.3064003l2.8144989,1.2070999v-3.0951004v-5h6.0566998
		c2.3584023,0,3.9433022-2.1767998,3.9433022-4.2099991V40.4524498
		C41.2666016,39.8997498,40.8194008,39.4524498,40.2666016,39.4524498z M29.2665997,11.2394505v49.2129974l-15.999999-6.7766991
		V4.4524498l15.9906988,6.7728004l0.0093002,0.0038996V11.2394505z" />
              </g>
              <g>
              </g>
              <g>
              </g>
              <g>
              </g>
              <g>
              </g>
              <g>
              </g>
              <g>
              </g>
              <g>
              </g>
              <g>
              </g>
              <g>
              </g>
              <g>
              </g>
              <g>
              </g>
              <g>
              </g>
              <g>
              </g>
              <g>
              </g>
              <g>
              </g>
            </svg>
          </button>
        </div>
      </div>
      <div id="messages"
           class="flex flex-col space-y-4 p-3 overflow-y-auto scrollbar-thumb-blue scrollbar-thumb-rounded scrollbar-track-blue-lighter scrollbar-w-2 scrolling-touch">
        <ChatMessageVue />
      </div>
      <div class="border-t-2 border-gray-200 px-4 pt-4 mb-2 sm:mb-0">
        <form @submit.prevent="sendMessage(form)">

          <div class="
              relative
              flex">
            <span class="absolute inset-y-0 flex items-center">
              <button type="button"
                      class="inline-flex items-center justify-center rounded-full h-12 w-12 transition duration-500 ease-in-out text-gray-500 hover:bg-gray-300 focus:outline-none">
                <svg xmlns="http://www.w3.org/2000/svg"
                     fill="none"
                     viewBox="0 0 24 24"
                     stroke="currentColor"
                     class="h-6 w-6 text-gray-600">
                  <path stroke-linecap="round"
                        stroke-linejoin="round"
                        stroke-width="2"
                        d="M19 11a7 7 0 01-7 7m0 0a7 7 0 01-7-7m7 7v4m0 0H8m4 0h4m-4-8a3 3 0 01-3-3V5a3 3 0 116 0v6a3 3 0 01-3 3z">
                  </path>
                </svg>
              </button>
            </span>
            <input type="text"
                   placeholder="Write your message!"
                   class="w-full focus:outline-none focus:placeholder-gray-400 text-gray-600 placeholder-gray-600 pl-12 bg-gray-200 rounded-md py-3"
                   v-model="form.content" />
            <div class="absolute right-0 items-center inset-y-0 hidden sm:flex">
              <button type="submit"
                      class="inline-flex items-center justify-center rounded-lg px-4 py-3 transition duration-500 ease-in-out text-white bg-blue-500 hover:bg-blue-400 focus:outline-none">
                <span class="font-bold">Send</span>
                <svg xmlns="http://www.w3.org/2000/svg"
                     viewBox="0 0 20 20"
                     fill="currentColor"
                     class="h-6 w-6 ml-2 transform rotate-90">
                  <path
                        d="M10.894 2.553a1 1 0 00-1.788 0l-7 14a1 1 0 001.169 1.409l5-1.429A1 1 0 009 15.571V11a1 1 0 112 0v4.571a1 1 0 00.725.962l5 1.428a1 1 0 001.17-1.408l-7-14z">
                  </path>
                </svg>
              </button>
            </div>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>
<style>
.scrollbar-w-2::-webkit-scrollbar {
  width: 0.25rem;
  height: 0.25rem;
}

.scrollbar-track-blue-lighter::-webkit-scrollbar-track {
  --bg-opacity: 1;
  background-color: #f7fafc;
  background-color: rgba(247, 250, 252, var(--bg-opacity));
}

.scrollbar-thumb-blue::-webkit-scrollbar-thumb {
  --bg-opacity: 1;
  background-color: #edf2f7;
  background-color: rgba(237, 242, 247, var(--bg-opacity));
}

.scrollbar-thumb-rounded::-webkit-scrollbar-thumb {
  border-radius: 0.25rem;
}
</style>

<script setup>
import ChatMessageVue from "./ChatMessage.vue";
import { sendMessage } from "../helper/message";
import { checkToken } from "../helper/auth";
import { ref, reactive, onMounted, onUpdated, watchEffect } from "vue"
import { useChatStore } from "../stores/chat";
import { generateAvatar } from "../helper/avatar"
import { useRouter } from "vue-router";
import pusher from "../helper/pusher";
const router = useRouter()
const form = ref({
  user_id: "",
  content: "",
  user: {
    email: "",
    name: "",
    image: ""
  }
})
const decodedToken = ref(null);
checkToken().then((res) => {
  decodedToken.value = res.data.data;
  form.value.user_id = decodedToken.value.id
  form.value.user.email = decodedToken.value.email
  form.value.user.name = decodedToken.value.name
  form.value.user.image = decodedToken.value.image
});

const online = ref(null)
var channel = pusher.subscribe("groupchat-channel");
const logout = () => {
  localStorage.removeItem("token")
  channel.unsubscribe("groupchat-channel")
  router.push({ name: "login" })
}
onMounted(() => {
  channel.bind("pusher:subscription_count", function (data) {
    online.value = data.subscription_count;
  });
})
</script>
