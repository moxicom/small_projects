import { UseDispatch, useDispatch } from "react-redux";
import { setUser } from "../Redux/userSlice";
import { getAuth, signInWithEmailAndPassword, sendPasswordResetEmail } from "firebase/auth";
import { AuthForm } from "./AuthForm";
import { useNavigate } from "react-router-dom";
import { useState } from "react";

export function Login() {
  const dispatch = useDispatch();
  const navigate = useNavigate();
  const [isLoading, setLoading] = useState(false);

  const handleLogin = (email: string, password: string) => {
    const auth = getAuth();
    setLoading(true)
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
        alert(error)
      }).finally(() => setLoading(false));
  };

  const handleReset = (email: string) => {
    const auth = getAuth();
    setLoading(true)
    sendPasswordResetEmail(auth, email)
      .then(() => {
        alert("Message has been send")
      })
      .catch((error) => {
        alert(error)
      }).finally(() => setLoading(false));
  }
  if (isLoading) {
    return <div className="text-center">Loading...</div>
  } else {
    return <AuthForm handleClick={handleLogin} handleReset={handleReset} title="sign in" />;
  }
}
