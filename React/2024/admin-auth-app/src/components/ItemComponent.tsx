type Props = {
  objectName: string;
  objectId: number;
  done: boolean;
  onDelete: (itemId: number) => void;
  onSwitch: (itemId: number) => void;
};

export default function Item(props: Props) {
  const border_green = "border-green-500";
  const border_red = "border-red-500";

  let border = "";

  if (props.done) {
    border = border_green;
  } else {
    border = border_red;
  }
  return (
    <>

      <div className="flex align-middle mb-3 max-w-7xl">
        <button
          className={`
            text-wrap
            flex-1
            border-t-4
            border-0
            rounded-r-none
            min-h-16
            text-left
            break-all
            bg-zinc-700
            ${border}
            hover:border-t-slate-50
           `}
          onClick={() => props.onSwitch(props.objectId)}
        >
          {props.objectName}
        </button>
        <button
          className="rounded-l-none mr-0"
          onClick={() => props.onDelete(props.objectId)}
        >
          del
        </button>
      </div>
    </>
  );
}
