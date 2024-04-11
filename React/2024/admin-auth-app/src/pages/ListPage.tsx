import { useEffect, useState } from "react";
import { useSelector } from "react-redux";
import { Params, useParams } from "react-router-dom";
import { ItemData, getItems } from "../restRequests/items";

export default function ListPage() {
  const { token } = useSelector((state: any) => state.user);
  const listId = parseInt(useParams<Params>().listId!);  

  const [isLoading, setLoading] = useState(false);
  const [items, setItems] = useState<ItemData[]>([]);
  const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        setLoading(true);
        isNaN(listId) ? setError("Bad id") : setError(null);
        getItems(token, listId)
        .then((response: ItemData[]) => {
          console.log(response);
          if (Array.isArray(response)) {
            setItems(response);
          } else {
            console.error("Response is not an array");
          }
        })
        .catch(()=>{})
        .finally(() => setLoading(false))



        // getItems(token, listId)
        
      }, []);

    if (isLoading) {
    return <>Loading...</>;
    }

    return (<>
    <h1>List page {listId}</h1>
    {error && (
        <div className="justify-center flex">
          <div className="max-w-96 text-center font-extrabold bg-red-950 border-2
          border-red-500 p-3 mt-5 rounded-lg">
            {error}
          </div>
        </div>
      )}
    </>);
}