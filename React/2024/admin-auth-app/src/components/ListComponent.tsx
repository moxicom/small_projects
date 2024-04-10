type Props = {
  listName: string;
  listId: number;
  onDelete: (listId: number) => void;
};

export default function List({ listName, listId, onDelete }: Props) {
  const deleteClicked = () => {
    console.log("Delete clicked - " + listId);
  };
  return (
    <>
      <button
        className="w-full min-h-16 bg-zinc-700 text-left text-pretty break-words"
        onClick={() => console.log(listId + " opened")}
      >
        <p>{listName}</p>
      </button>
      <div className="mb-3">
        <button
          className="rounded-l-lg rounded-r-none"
          onClick={() => onDelete(listId)}
        >
          Delete
        </button>
        <button className="rounded-r-lg rounded-l-none">Modify</button>
      </div>
    </>
  );
}
