import { useSelector } from "react-redux";
import List from "../components/ListComponent";
import { useEffect, useState } from "react";
import { createList, getLists, ListData } from "../restRequests/lists";
import { AxiosError, AxiosResponse } from "axios";
import { useNavigate } from "react-router-dom";

export default function ListsPage() {
  const { token } = useSelector((state: any) => state.user);
  const navigate = useNavigate();

  const [newTitle, setNewTitle] = useState("")
  const [isLoading, setLoading] = useState(true);
  const [lists, setLists] = useState<ListData[]>([]);
  const [error, setError] = useState<string | null>(null);

  const openListClicked = (listId: number) => {
    navigate(`${listId}`);
  };

  const addListClicked = () => {
    if (newTitle.trim() == "") {
      return
    }
    setLoading(true);
    setError(null);
    createList(token, newTitle)
      .then((createdID: number) => {
        lists.push({
          id: createdID,
          description: "",
          title: newTitle,
        })
        setLists(lists);
      })
      .catch((error: AxiosError) => {
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
          <div className="w-full h-auto flex justify-center">
            <div className="flex w-full max-w-7xl rounded-lg overflow-hidden">
              <input
                className="w-full flex-1 rounded-l-lg p-3 bg-zinc-700"
                placeholder="List title"
                value={newTitle}
                onChange={(e) => setNewTitle(e.target.value)}
              ></input>
              <button
                className="rounded-l-none"
                onClick={() => addListClicked()}
              >
                New
              </button>
            </div>
          </div>

          <div className="flex justify-center w-full">
            <div className="mt-6 max-w-7xl w-full">
              {lists.map((list) => (
                <List
                  key={list.id}
                  objectId={list.id}
                  objectName={list.title}
                  onOpen={() => openListClicked(list.id)}
                />
              ))}
            </div>
          </div>
        </>
      )}
    </>
  );
}
