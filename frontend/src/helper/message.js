import axios from "axios";
export const sendMessage = (form) => {
  if (form.content === "") {
    return;
  }
  return new Promise((resolve, reject) => {
    axios
      .post(`${import.meta.env.VITE_API_URL}/message`, {
        user_id: form.user_id,
        content: form.content,
        user: form.user,
      })
      .then((res) => {
        form.content = "";
        resolve(res);
      })
      .catch((err) => {
        reject(err);
      });
  });
};
