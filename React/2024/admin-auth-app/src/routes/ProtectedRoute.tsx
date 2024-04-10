import React from "react";
import { Navigate, useLocation } from "react-router-dom";
import { useAuth } from "../hooks/useAuth";
import { useDispatch } from "react-redux";

type Props = { children: React.ReactNode };

const ProtectedRoute = ({ children }: Props) => {
  const location = useLocation();
  const dispatch = useDispatch()
  const { isAuth, token } = useAuth(dispatch);
  return isAuth ? (
    <>{children}</>
  ) : (
    <Navigate to={"/login"} state={{ from: location }} replace />
  );
};

export default ProtectedRoute;
