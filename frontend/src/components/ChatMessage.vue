<template>
  <!-- <pre> {{ item.user.id }} </pre> -->
  <div>
    <template v-for="(item, index) in messages">
      <div class="chat-message"
           ref="messagesRef">
        <div class="flex items-end"
             :class="[item.user_id === decodedToken.id ? 'justify-end' : '']">
          <div class="flex space-y-2 flex-col text-xs max-w-xs mx-2 "
               :class="[item.user_id === decodedToken.id ? 'order-1 items-start'
                 : 'order-2 items-end'
               ]">
            <div>
              <span v-if="item.user && index === 0">{{ item.user.email }}</span>
              <span v-if="index > 0">
                {{
                    item.user.email === messages[index - 1].email
                      ? ""
                      : item.user.email
                }}
              </span>
            </div>
            <div class="">
              <span class="px-4 py-2 rounded-lg inline-block "
                    :class="[item.user_id === decodedToken.id ? 'rounded-br-none bg-blue-600 text-white' : 'rounded-bl-none bg-gray-300 text-gray-600']">{{
                        item.content
                    }}</span>
            </div>
          </div>
          <img :src="publicUrl + item.user.image"
               alt="My profile"
               class="w-6 h-6 rounded-full order-1" />
        </div>
      </div>
    </template>
  </div>
</template>

<script setup>
import { onMounted, reactive, ref, nextTick, watchEffect, onUpdated } from "vue";
import pusher from "../helper/pusher";
import axios from "axios";
import { checkToken } from "../helper/auth";
import { sendMessage } from "../helper/message";
import { useChatStore } from "../stores/chat";
const messages = ref([]);
const messagesRef = ref(null)
const chatStore = useChatStore()
axios
  .get(`${import.meta.env.VITE_API_URL}/messages`)
  .then((res) => {
    messages.value = res.data.data;
  })
  .catch((err) => {
    messages.value = err.response.data.data;
  });
const decodedToken = ref(null);
checkToken().then((res) => {
  decodedToken.value = res.data.data;
});
let publicUrl = import.meta.env.VITE_PUBLIC_URL
let messageList
onUpdated(() => {
  messageList = document.getElementsByClassName("chat-message")
  if (messageList.length > 0) {

    messageList[messageList.length - 1].setAttribute("id", "last-message")
    messageList[messageList.length - 1].scrollIntoView(false, { behavior: "smooth" })
    chatStore.$patch({
      lastDiv: messageList[messageList.length - 1]
    })
  }
})
onMounted(() => {
  var channel = pusher.subscribe("groupchat-channel");
  channel.bind("message", function (data) {
    messages.value.push(data);
  });
});


</script>

<style lang="scss" scoped>

</style>
