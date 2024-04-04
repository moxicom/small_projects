import React from "react";
import ReactDOM from "react-dom/client";
import "./index.css";
import { Routes } from "./routes/Routes";
import { Provider } from "react-redux";
import { store } from "./store/store";

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <Provider store={store}>
      <Routes />
    </Provider>
  </React.StrictMode>
);
