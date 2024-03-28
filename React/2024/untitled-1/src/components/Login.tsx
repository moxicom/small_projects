import { UseDispatch, useDispatch } from "react-redux";
import { setUser } from "../Redux/userSlice";
import { getAuth, signInWithEmailAndPassword } from "firebase/auth";
import { AuthForm } from "./AuthForm";
import { useNavigate } from "react-router-dom";

export function Login() {
  const dispatch = useDispatch();
  const navigate = useNavigate();

  const handleLogin = (email: string, password: string) => {
    const auth = getAuth();
    signInWithEmailAndPassword(auth, email, password)
      .then(({ user }) => {
        console.log(user.email);
        console.log(user.uid);
        dispatch(
          setUser({
            email: user.email,
            id: user.uid,
          })
        );
        navigate("/account");
        return null;
      })
      .catch((error) => {
        const errorCode = error.code;
        const errorMessage = error.message;
        console.log(errorCode + errorMessage);
      });
  };

  return <AuthForm handleClick={handleLogin} title="sign in" />;
}
