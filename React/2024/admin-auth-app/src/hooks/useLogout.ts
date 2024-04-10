import { Dispatch } from "@reduxjs/toolkit";
import { removeUser } from "../store/userSlice";

export default function useLogOut(dispatch: Dispatch) {
    dispatch(removeUser())
}