import axios from "axios";

export const checkToken = () => {
  const token = localStorage.getItem("token") || null;
  axios.defaults.headers.common.Authorization = `Bearer ${token}`;
  return new Promise((resolve, reject) => {
    axios
      .get(`${import.meta.env.VITE_API_URL}/user`)
      .then((res) => {
        resolve(res);
        return res;
      })
      .catch((err) => {
        localStorage.removeItem("token");
        reject(err);
        // return err;
      });
  });
};
