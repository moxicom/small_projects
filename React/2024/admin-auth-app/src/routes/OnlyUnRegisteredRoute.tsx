import React from "react";
import { Navigate, useLocation } from "react-router-dom";
import { useAuth } from "../hooks/useAuth";
import { useDispatch } from "react-redux";

type Props = { children: React.ReactNode };

const OnlyUnregisteredRoute = ({ children }: Props) => {
  const location = useLocation();
  const dispatch = useDispatch()
  const { isAuth, token } = useAuth(dispatch);
  return isAuth ? (
    <Navigate to={"/"} state={{ from: location }} replace />
  ) : (
    <>{children}</>
  );
};

export default OnlyUnregisteredRoute;
