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

export async function switchItem(
  token: string,
  listID: number,
  itemID: number,
  done: boolean,
  title: string,
  description: string
) {
  const backUrl = urlProvider.getBackUrl();
  description = "xasdasd";
  const data = JSON.stringify({
    title: title,
    description: description,
    done: done,
  });

  console.log("data = " + data);

  try {
    const response = await axios.put(
      `${backUrl}/api/lists/${listID}/items/${itemID}`,
      data,
      {
        timeout: 5000,
        headers: {
          Authorization: `Bearer ${token}`,
        },
      }
    );
    console.log(`itemID ${itemID} PUT response = ` + response);
    return response;
  } catch (e: any) {
    const error = e as AxiosError;
    console.log("Error occurred:", e.message); // Вывести сообщение об ошибке
    console.log("Error status:", e.response?.status); // Вывести статус ошибки, если есть
    console.log("Error data:", e.response?.data); // Вывести данные ошибки, если есть
    throw error;
  }
}
