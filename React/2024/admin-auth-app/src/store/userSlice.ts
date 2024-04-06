import { createSlice } from "@reduxjs/toolkit";

const initialState = {
  username: null,
  token: null,
};

const userSlice = createSlice({
  name: "user",
  initialState,
  reducers: {
    setUser(state, action) {
      console.log(state);
      console.log(action);

      state.username = action.payload.username;
      state.token = action.payload.token;
      console.log("New user in store: " + state.username + " " + state.token);
    },
    removeUser(state) {
      state.username = null;
      state.token = null;
    },
  },
});

export const { setUser, removeUser } = userSlice.actions;
export default userSlice.reducer;
