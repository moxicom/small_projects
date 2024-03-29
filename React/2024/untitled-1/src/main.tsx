import React from "react";
import ReactDOM from "react-dom/client";
import "./index.css";
import { Routes } from "./Routes.tsx";
import { Provider } from "react-redux";
import { persistor, store } from "./Redux/store.ts";
import "./Redux/firebase.ts";
import { PersistGate } from "redux-persist/integration/react";

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <Provider store={store}>
      <PersistGate persistor={persistor} loading={null}>
        <Routes></Routes>
      </PersistGate>
    </Provider>
  </React.StrictMode>
);
