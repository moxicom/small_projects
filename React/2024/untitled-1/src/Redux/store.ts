import { configureStore } from "@reduxjs/toolkit";
import userReducer from "./userSlice";
import { persistReducer } from "redux-persist";
import { combineReducers } from "@reduxjs/toolkit";
import storage from "redux-persist/lib/storage";
import persistStore from "redux-persist/es/persistStore";

const redusers = combineReducers({ user: userReducer });

const persistConfig = {
  key: "root",
  storage,
};

const persistedReduser = persistReducer(persistConfig, redusers);

export const store = configureStore({
  reducer: persistedReduser,
});

export const persistor = persistStore(store);
