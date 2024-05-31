import { useSelector } from "react-redux";
import { Dispatch } from "@reduxjs/toolkit";
import { isJWTValid } from "../utils/jwt";
import { removeUser } from "../store/userSlice";

export function useAuth(dispatch: Dispatch) {

  const { username, token } = useSelector((state: any) => state.user);
  if (token && !isJWTValid(new Date(Date.now()), token)) {
    dispatch(removeUser())
  }
  return {
    isAuth: !!username,
    // isAuth: true,
    token,
  };
}
