import { useSelector } from "react-redux";
import List from "../components/ListComponent";
import { useEffect, useState } from "react";
import { createList, getLists, ListData } from "../restRequests/lists";
import { AxiosError, AxiosResponse } from "axios";
import { useNavigate } from "react-router-dom";

export default function ListsPage() {
  const { token } = useSelector((state: any) => state.user);
  const navigate = useNavigate();

  const [isLoading, setLoading] = useState(true);
  const [lists, setLists] = useState<ListData[]>([]);
  const [error, setError] = useState<string | null>(null);

  const openListClicked = (listId: number) => {
    navigate(`${listId}`)
  }

  const addListClicked = () => {
    console.log("Clicked - Add new list");
    setLoading(true);
    setError(null);
    createList(token)
      .then((response: number) => {
        console.log("created list id = " + response);
        const newArray = [...lists];
        newArray.push({
          id: response,
          title: "MyNewTitle",
          description: "MyNewDescription",
        });
        setLists(newArray);
      })
      .catch((error: AxiosError) => {
        console.log(error);
        setError(error.message);
      })
      .finally(() => {
        setLoading(false);
      });
  };

  useEffect(() => {
    setLoading(true);
    setError(null);
    getLists(token)
      .then((response: ListData[]) => {
        console.log("response data = " + response);
        if (Array.isArray(response)) {
          setLists(response);
        } else {
          console.error("Response is not an array");
        }
      })
      .catch((error: AxiosError) => {
        console.log(error);
        setError(error.message);
      })
      .finally(() => {
        setLoading(false);
      });
  }, []);

  if (isLoading) {
    return <>Loading...</>;
  }
  return (
    <>
      {error && (
        <div className="justify-center flex">
          <div className="max-w-96 text-center font-extrabold bg-red-950 border-2 border-red-500 p-5 rounded-lg">
            {error}
          </div>
        </div>
      )}
      {!error && (
        <>
          <div className="flex justify-center">
            <div className="mt-6 max-w-7xl">
              {lists.map((list) => (
                <List
                  key={list.id}
                  listId={list.id}
                  listName={list.title}
                  onOpen={() => openListClicked(list.id) }
                />
              ))}
            </div>
          </div>
          <button
            className="text-stone-800 bg-white hover:bg-transparent hover:text-white hover:border-white"
            onClick={addListClicked}
          >
            Add new
          </button>
        </>
      )}
    </>
  );
}
