import { combineReducers, configureStore } from "@reduxjs/toolkit";
import userReducer from "./userSlice";
import { persistStore, persistReducer } from "redux-persist";
import storage from "redux-persist/lib/storage";

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
