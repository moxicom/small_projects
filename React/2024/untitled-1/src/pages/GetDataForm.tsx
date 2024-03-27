import { Form } from "react-router-dom";

type Props = {
  onSubmit: (city: string) => void;
};

export function GetDataForm({ onSubmit }: Props) {
  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    const formData = new FormData(event.currentTarget);
    const city = formData.get("city") as string;
    onSubmit(city);
  };

  return (
    <Form method="get" onSubmit={handleSubmit}>
      <div className="flex mb-2 justify-center pt-8 ">
        <input
          type="text"
          name="city"
          placeholder="Enter city name"
          className="bg-gray-950 text-whiterounded-lg text-white rounded-l p-2"
        />
        <button
          type="submit"
          className="border-black border-2 p-2 rounded-r pl-5 pr-5"
        >
          Submit
        </button>
      </div>
    </Form>
  );
}
