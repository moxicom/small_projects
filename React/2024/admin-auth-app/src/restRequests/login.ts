import axios, { AxiosError } from "axios";
import { urlProvider } from "../utils/urlProvider";

export type Data = {
  username: string;
  password: string;
};

export async function loginRequest(data: Data) {
  const backUrl = urlProvider.getBackUrl();
  const jsonData = JSON.stringify(data);
  try {
    const response = await axios.post(backUrl + "/auth/sign-in", jsonData, {
      timeout: 5000,
    });
    console.log("response.data = ", response.data);
    return response.data;
  } catch (e) {
    const error = e as AxiosError;
    if (error.response?.status === 401) {
      throw error.response.status;
    }
    throw e;
  }
}
