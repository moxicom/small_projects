import axios from "axios";
import { urlProvider } from "../utils/urlProvider";

export type FormData = {
  name: string;
  username: string;
  password: string;
};

export function registerRequest(data: FormData) {
  const backUrl = urlProvider.getBackUrl();
  console.log(backUrl);
  axios
    .post(backUrl + "/auth/sign-up", data)
    .then((response) => {
      console.log(response);
    })
    .catch((error) => {
      console.log(error);
    });
}
