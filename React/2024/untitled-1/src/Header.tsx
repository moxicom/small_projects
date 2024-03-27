import logo from "/vite.svg";
import { NavLink, Link } from "react-router-dom";

export function Header() {
  return (
    <header>
      <div className="w-screen h-32 flex justify-center">
        <div className="flex items-center">
          <Link to="">
            <img src={logo} className="w-24 p-4" alt="logo"></img>
          </Link>
        </div>
        <div className="flex items-center">
          <h1 className="font-bold font-sans ">Mospoly lab 2</h1>
        </div>
      </div>
      <div className="bg-black h-10 flex justify-center items-center">
        <NavLink to={"fetch"} className="text-white">
          К запросам
        </NavLink>
      </div>
    </header>
  );
}
