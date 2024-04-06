import React from "react";
import { Navigate, useLocation } from "react-router-dom";
import { useAuth } from "../hooks/useAuth";

type Props = { children: React.ReactNode };

const OnlyUnregisteredRoute = ({ children }: Props) => {
  const location = useLocation();
  const { isAuth, token } = useAuth();
  return isAuth ? (
    <Navigate to={"/"} state={{ from: location }} replace />
  ) : (
    <>{children}</>
  );
};

export default OnlyUnregisteredRoute;
