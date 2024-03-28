import logo from "/vite.svg";
import { NavLink, Link } from "react-router-dom";

export function Header() {
  return (
    <header>
      <div className="w-screen h-32 flex flex-row justify-center items-center">
        <div className="flex items-center">
          <Link to="">
            <img src={logo} className="w-24 p-4" alt="logo"></img>
          </Link>
          <div className="flex items-center">
            <h1 className="font-bold font-sans mr-10 ">Mospoly lab</h1>
          </div>
        </div>
      </div>
      <div className="bg-black h-10 flex flex-row justify-center items-center text-white">
        <NavLink to={"fetch"} className="ml-2 mr-2">
          К запросам
        </NavLink>
        <NavLink to={"json"} className="ml-2 mr-2">
          К Json
        </NavLink>
        <NavLink to={"account"} className="ml-2 mr-2">
          Аккаунт
        </NavLink>
      </div>
    </header>
  );
}
