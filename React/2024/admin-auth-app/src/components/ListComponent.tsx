type Props = {
  listName: string;
  listId: number;
};

export default function List({ listName, listId }: Props) {
  return (
    <button
      className="w-full mb-3 min-h-16 bg-zinc-700 text-left text-pretty break-words"
      onClick={() => console.log(listId + " opened")}
    >
      <p>{listName}</p>
    </button>
  );
}
