import axios, { AxiosError } from "axios";
import { urlProvider } from "../utils/urlProvider";

export type ListData = {
  description: string;
  id: number;
  title: string;
};

export async function getLists(token: string) {
  const backUrl = urlProvider.getBackUrl();
  try {
    const response = await axios.get(backUrl + "/api/lists/", {
      timeout: 5000,
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
    console.log(response);
    return response.data.data as ListData[];
  } catch (e) {
    console.log(e);
    const error = e as AxiosError;
    if (error.response?.status === 401) {
      throw error.response.status;
    }
    throw error;
  }
}

// Returns id of the created list
export async function createList(token: string) {
  const backUrl = urlProvider.getBackUrl();
  const data = JSON.stringify({
    description: "MyNewDescription",
    id: 0,
    title: "MyNewTitle",
  });

  try {
    const response = await axios.post(backUrl + "/api/lists/", data, {
      timeout: 5000,
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
    return response.data.id;
  } catch (e) {
    console.log(e);
    const error = e as AxiosError;
    if (error.response?.status === 401) {
      throw error.response.status;
    }
    throw error;
  }
}
