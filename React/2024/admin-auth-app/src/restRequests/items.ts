import axios, { AxiosError } from "axios";
import { urlProvider } from "../utils/urlProvider";

export type ItemData = {
    description: string;
    done: boolean;
    id: number;
    title: string;
  };

export async function getItems(token: string, listId: number) {
    const backUrl = urlProvider.getBackUrl();
    try {
      const response = await axios.get(`${backUrl}/api/lists/${listId}/items/`, {
        timeout: 5000,
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });
      console.log(`list ${listId} items response = ` + response);
      return response.data.data as ItemData[];
    } catch (e) {
      console.log(e);
      const error = e as AxiosError;
      if (error.response?.status === 401) {
        throw error.response.status;
      }
      throw error;
    }
  }