import { redirect, useNavigate } from "react-router-dom";
import { useAuth } from "../hooks/use-auth";
import { useDispatch } from "react-redux";
import { removeUser } from "../Redux/userSlice";
import { useEffect } from "react";

export function AccountPage() {
  const dispatch = useDispatch();
  const navigate = useNavigate();
  const { isAuth, email } = useAuth();
  console.log("account page - " + isAuth);

  useEffect(() => {
    if (!isAuth) {
      navigate("/login");
    }
  }, [isAuth, navigate]);

  return (
    <div className="w-full h-auto">
      <h1>Account info...{email}</h1>
      <button onClick={() => dispatch(removeUser())}>LogOut</button>
    </div>
  );
}
