type Props = {
  listName: string;
  listId: number;
  onOpen: (listId: number) => void;
};

export default function List(props: Props) {
  return (
    <>
      <button
        className="w-full min-h-16 bg-zinc-700 text-left text-pretty break-words mb-3"
        onClick={() => props.onOpen(props.listId)}
      >
        <p>{props.listName}</p>
      </button>
    </>
  );
}
