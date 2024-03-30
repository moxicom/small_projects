import { Link } from "react-router-dom";

export default function AdminPage() {
  console.log("admin page loaded");
  return (
    <div>
      <h1>Admin page</h1>
      <Link to="/">
        <button>Go to home page</button>
      </Link>
      <p> SECURE INFO</p>
      <p>My name is alexander markov</p>
    </div>
  );
}
