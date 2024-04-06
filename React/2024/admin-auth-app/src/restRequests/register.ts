import axios, { Axios, AxiosError } from "axios";
import { urlProvider } from "../utils/urlProvider";

export type Data = {
  name: string;
  username: string;
  password: string;
};

export async function registerRequest(data: Data) {
  const backUrl = urlProvider.getBackUrl();
  const jsonData = JSON.stringify(data);
  try {
    const response = await axios.post(backUrl + "/auth/sign-up", jsonData, {
      timeout: 5000,
    });
    console.log(response);
    return response;
  } catch (e) {
    const error = e as AxiosError;
    if (error.response?.status === 409) {
      throw error.response.status;
    }
  }
}
