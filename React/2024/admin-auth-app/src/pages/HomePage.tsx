import reactLogo from "../assets/react.svg";
import "./App.css";
import { Link, Outlet } from "react-router-dom";

export function HomePage() {
  return (
    <div className="items-center">
      <div>
        <Link to={"/lists"}>
          <img src={reactLogo} className="logo react" alt="React logo" />
        </Link>
      </div>
      <h1>WELCOME</h1>
      <p className="read-the-docs">Click React icon to go to admin page</p>
      <Outlet />
    </div>
  );
}
