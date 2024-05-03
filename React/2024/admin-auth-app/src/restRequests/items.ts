import axios, { AxiosError } from "axios";
import { urlProvider } from "../utils/urlProvider";

export type ItemData = {
  description: string;
  done: boolean;
  id: number;
  title: string;
};

export async function getItems(token: string, listID: number) {
  const backUrl = urlProvider.getBackUrl();
  try {
    const response = await axios.get(`${backUrl}/api/lists/${listID}/items/`, {
      timeout: 5000,
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
    console.log(`list ${listID} items response = ` + response);
    return response.data.data as ItemData[];
  } catch (e) {
    console.log(e);
    const error = e as AxiosError;
    console.log("Error occurred:", error.message); // Вывести сообщение об ошибке
    console.log("Error status:", error.response?.status); // Вывести статус ошибки, если есть
    console.log("Error data:", error.response?.data); // Вывести данные ошибки, если есть
    throw error;
  }
}

export async function switchItem(
  token: string,
  listID: number,
  itemID: number,
  done: boolean,
  title: string,
  description: string = "empty"
) {
  const backUrl = urlProvider.getBackUrl();
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
    console.log("Error occurred:", e.message);
    console.log("Error status:", e.response?.status);
    console.log("Error data:", e.response?.data);
    throw error;
  }
}

export async function deleteItem(
  token: string,
  listID: number,
  itemID: number
) {
  const backUrl = urlProvider.getBackUrl();
  try {
    const response = await axios.delete(
      `${backUrl}/api/lists/${listID}/items/${itemID}`,
      {
        timeout: 5000,
        headers: {
          Authorization: `Bearer ${token}`,
        },
      }
    );
    console.log(`list ${listID} item ${itemID} delete response = ` + response);
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

export async function createItem(
  token: string,
  listID: number,
  title: string,
  description: string = "empty"
) {
  const backUrl = urlProvider.getBackUrl();
  const data = JSON.stringify({
    title: title,
    description: description,
    done: false,
  });

  console.log("data = " + data);

  try {
    const response = await axios.post(
      `${backUrl}/api/lists/${listID}/items/`,
      data,
      {
        timeout: 5000,
        headers: {
          Authorization: `Bearer ${token}`,
        },
      }
    );
    console.log(`listID ${listID} POST response = ` + response);
    return response.data.id as number;
  } catch (e: any) {
    const error = e as AxiosError;
    console.log("Error occurred:", e.message);
    console.log("Error status:", e.response?.status);
    console.log("Error data:", e.response?.data);
    throw error;
  }
}
