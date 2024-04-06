import axios, { AxiosError } from "axios";
import { urlProvider } from "../utils/urlProvider";

export type ListData = {
  description: string;
  id: number;
  title: string;
};

export async function getLists(token: string) {
  const backUrl = urlProvider.getBackUrl();
  const jsonData = JSON.stringify(token);
  try {
    const response = await axios.get(backUrl + "/api/lists/", {
      timeout: 5000,
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
    console.log(response);
    return response.data;
  } catch (e) {
    console.log(e);
    const error = e as AxiosError;
    if (error.response?.status === 401) {
      throw error.response.status;
    }
    throw error;
  }
}
