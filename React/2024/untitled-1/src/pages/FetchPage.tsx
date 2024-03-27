import { useEffect, useState } from "react";
import { getData } from "../apiHandlers/getData";
import { GetDataForm } from "./GetDataForm";

export default function FetchPage() {
  const [isLoading, setLoading] = useState(false);
  const [cityInfo, setCity] = useState("");

  const buttonClicked = async (city: string) => {
    if (city === "") {
      return;
    }
    console.log("button clicked");
    setLoading(true);

    try {
      const data = await getData(city);
      setCity(data);
    } catch (error) {
      console.error("error fetching data", error);
      setCity("City fetch error");
    } finally {
      setLoading(false);
    }
  };

  if (isLoading) {
    return <p className="text-center pt-5">Loading...</p>;
  } else {
    return (
      <div className="flex justify-center pt-8 content-center">
        <div className="w-2/3">
          <h1 className="font-sans font-bold text-xl text-center">
            Welcome to fetch page
          </h1>
          <GetDataForm onSubmit={buttonClicked} />
          <p className="text-center pt-5">{cityInfo}</p>
        </div>
      </div>
    );
  }
}
