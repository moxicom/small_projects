import { useEffect, useState } from "react";
import { useSelector } from "react-redux";
import { Navigate, Params, useNavigate, useParams } from "react-router-dom";
import { deleteList } from "../restRequests/lists";
import * as requests from "../restRequests/items";
import Item from "../components/ItemComponent";

export default function ListPage() {
  const navigate = useNavigate();

  const { token } = useSelector((state: any) => state.user);
  const listID = parseInt(useParams<Params>().listId!);

  const [newTitle, setNewTitle] = useState("");
  const [isLoading, setLoading] = useState(false);
  const [items, setItems] = useState<requests.ItemData[]>([]);
  const [error, setError] = useState<string | null>(null);

  // Load items on page load
  useEffect(() => {
    setLoading(true);
    isNaN(listID) ? setError("Bad id") : setError(null);
    requests
      .getItems(token, listID)
      .then((response: requests.ItemData[]) => {
        if (Array.isArray(response)) {
          setItems(response);
        } else {
          console.error("Response is not an array");
        }
      })
      .catch(() => {
        setError("Something went wrong");
      })
      .finally(() => setLoading(false));
  }, []);

  // Delete list
  function deleteListClicked(listID: number) {
    setLoading(true);
    deleteList(token, listID)
      .then(() => {
        navigate("/lists");
      })
      .catch(() => {
        setError("Something went wrong");
      })
      .finally(() => {
        setLoading(false);
      });
  }

  // Switch item
  function switchItemClicked(itemID: number) {
    let title = "";

    const newItems = items.map((item) => {
      if (item.id === itemID) {
        title = item.title;
        return { ...item, done: !item.done };
      }
      return item;
    });

    setItems(newItems);

    const clickedItem = items.find((item) => item.id === itemID);
    if (clickedItem) {
      requests
        .switchItem(
          token,
          listID,
          itemID,
          !clickedItem.done,
          title,
          clickedItem.description
        )
        .catch(() => setError("Something went wrong"));
    }
  }

  function deleteItemClicked(itemID: number) {
    const clickedItem = items.find((item) => item.id === itemID);
    let newItems = items.filter((item) => item.id !== itemID);
    setItems(newItems);

    requests
      .deleteItem(token, listID, itemID)
      .catch(() => setError("Failed to delete an object"))
      .finally();
  }

  function createItemClicked() {
    if (newTitle.trim() == "") {
      return;
    }
    setLoading(true);
    requests
      .createItem(token, listID, newTitle)
      .catch()
      .then((id) => {
        items.push({
          id: id,
          description: "empty",
          done: false,
          title: newTitle,
        });
        setItems(items);
      })
      .finally(() => setLoading(false));
  }

  if (isLoading) {
    return <>Loading...</>;
  }

  return (
    <>
      <div className="mx-auto flex max-w-7xl items-center justify-between p-6 lg:px-8">
        <h1 className="w-full">List page {listID}</h1>
        {!error && (
          <button
            className="float-end bg-red-700 hover:border-red-500"
            onClick={() => deleteListClicked(listID)}
          >
            Delete
          </button>
        )}
      </div>
      {error && (
        <div className="justify-center flex">
          <div
            className="max-w-96 text-center font-extrabold bg-red-950 border-2
          border-red-500 p-3 mt-5 rounded-lg"
          >
            {error}
          </div>
        </div>
      )}
      {!error && (
        <>
          {/* add new field */}
          <div className="w-full h-auto flex justify-center">
            <div className="flex w-full max-w-7xl rounded-lg overflow-hidden">
              <input
                className="w-full flex-1 rounded-l-lg p-3 bg-zinc-700"
                placeholder="Note title"
                value={newTitle}
                onChange={(e) => setNewTitle(e.target.value)}
              ></input>
              <button
                className="rounded-l-none"
                onClick={() => createItemClicked()}
              >
                New
              </button>
            </div>
          </div>

          <div className="flex justify-center w-full">
            <div className="w-full max-w-7xl mt-6">
              {items.map((item) => (
                <Item
                  key={item.id}
                  objectId={item.id}
                  objectName={item.title}
                  done={item.done}
                  onDelete={() => deleteItemClicked(item.id)}
                  onSwitch={() => switchItemClicked(item.id)}
                />
              ))}
            </div>
          </div>
        </>
      )}
    </>
  );
}
