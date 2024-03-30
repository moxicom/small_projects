import { useState } from "react";
import reactLogo from "../assets/react.svg";
import viteLogo from "/vite.svg";
import "./App.css";
import { Link } from "react-router-dom";

export function AppPage() {
  return (
    <div>
      <div>
        <Link to={"/admin"}>
          <img src={reactLogo} className="logo react" alt="React logo" />
        </Link>
      </div>
      <h1>Home page</h1>
      <p className="read-the-docs">Click React icon to go to admin page</p>
    </div>
  );
}
