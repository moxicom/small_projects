type Props = {
  objectName: string;
  objectId: number;
  done: boolean;
  onDelete: (itemId: number) => void;
  onSwitch: (itemId: number) => void;
};

export default function Item(props: Props) {
  let border = "";
  if (props.done) {
    border = "border-green-500";
  } else {
    border = "border-red-500";
  }
  return (
    <>
      <div className="w-full flex align-middle mb-3">
        <button
          className={`w-auto h-full flex-1 border-t-4 border-0 rounded-r-none min-h-16 text-left text-pretty break-words bg-zinc-700 ${border}`}
          onClick={() => props.onSwitch(props.objectId)}
        >
          <span>{props.objectName}</span>
        </button>
        <button
          className="rounded-l-none"
          onClick={() => props.onDelete(props.objectId)}
        >
          del
        </button>
      </div>
    </>
  );
}
