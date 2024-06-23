import { useState } from "react";
import "./App.css";
import axios from "axios";

type Product = {
  id: number;
  category_id: number;
  name: string;
  price: number;
  properties: any; // Подставьте сюда нужный тип для properties
};

function App() {
  const [searchTerm, setSearchTerm] = useState("");
  const [error, setError] = useState("");

  const [products, setProducts] = useState<Product[]>([]);

  const handleChange = (e: { target: { value: string } }) => {
    setSearchTerm(e.target.value);
    setError("поиск...");
    if (e.target.value.trim() != "") {
      axios
        .get<Product[]>("http://localhost:8080/search/" + e.target.value)
        .then((response) => {
          setProducts(response.data);
          setError("");
          if (products.length == 0) {
            setError("Не найдено");
          }
        })
        .catch((error) => {
          console.error("Error fetching products:", error);
        });
    } else {
      setProducts([]);
      setError(" ");
    }
  };

  return (
    <>
      <input placeholder="Поиск" value={searchTerm} onChange={handleChange} />
      <div>
        <ul>
          {products.map((item) => (
            <li>
              {item.name} - {item.price}
            </li>
          ))}
        </ul>
        {error}
      </div>
    </>
  );
}

export default App;
