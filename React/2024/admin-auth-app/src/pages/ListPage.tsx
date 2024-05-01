import { useEffect, useState } from "react";
import { useSelector } from "react-redux";
import { Navigate, Params, useNavigate, useParams } from "react-router-dom";
import { deleteList } from "../restRequests/lists";
import { ItemData, getItems, switchItem } from "../restRequests/items";
import Item from "../components/ItemComponent";

export default function ListPage() {
  const navigate = useNavigate();

  const { token } = useSelector((state: any) => state.user);
  const listId = parseInt(useParams<Params>().listId!);

  const [isLoading, setLoading] = useState(false);
  const [items, setItems] = useState<ItemData[]>([]);
  const [error, setError] = useState<string | null>(null);

  // Load items on page load
  useEffect(() => {
    setLoading(true);
    isNaN(listId) ? setError("Bad id") : setError(null);
    getItems(token, listId)
      .then((response: ItemData[]) => {
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
      switchItem(
        token,
        listId,
        itemID,
        !clickedItem.done,
        title,
        clickedItem.description
      ).catch(() => setError("Something went wrong"));
    }
  }

  if (isLoading) {
    return <>Loading...</>;
  }

  return (
    <>
      <div className="mx-auto flex max-w-7xl items-center justify-between p-6 lg:px-8">
        <h1 className="w-full">List page {listId}</h1>
        {!error && (
          <button
            className="float-end bg-red-700 hover:border-red-500"
            onClick={() => deleteListClicked(listId)}
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
        <div className="flex justify-center w-full">
          <div className="mt-6 max-w-7xl w-full">
            {items.map((item) => (
              <Item
                key={item.id}
                objectId={item.id}
                objectName={item.title}
                done={item.done}
                onDelete={() => {
                  console.log("Delete item " + item.id);
                }}
                onSwitch={() => switchItemClicked(item.id)}
              />
            ))}
          </div>
        </div>
      )}
    </>
  );
}
