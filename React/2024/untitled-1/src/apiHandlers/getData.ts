export async function getData(cityName: string) {
  const apiURL = import.meta.env.VITE_ENDPOINT;
  const params = new URLSearchParams();
  params.append("format", "4");
  const URLWithParams = `${apiURL}/${cityName}?${params.toString()}`;

  const response = await fetch(URLWithParams);
  const body = await response.text();
  console.log(body);
  return body as string;
}
