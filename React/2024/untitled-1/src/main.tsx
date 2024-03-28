import React from "react";
import ReactDOM from "react-dom/client";
import "./index.css";
import { Routes } from "./Routes.tsx";
import { Provider } from "react-redux";
import { store } from "./Redux/store.ts";
import "./Redux/firebase.ts";

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <Provider store={store}>
      <Routes></Routes>
    </Provider>
  </React.StrictMode>
);
