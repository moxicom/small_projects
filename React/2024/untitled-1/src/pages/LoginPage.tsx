import { redirect, useNavigate } from "react-router-dom";
import { Login } from "../components/Login";
import { useAuth } from "../hooks/use-auth";

export function LoginPage() {
  const { isAuth } = useAuth();
  const navigate = useNavigate();

  if (isAuth) {
    navigate("/account");
    return null;
  }

  return (
    <div className="w-full mt-5">
      <Login></Login>
    </div>
  );
}
