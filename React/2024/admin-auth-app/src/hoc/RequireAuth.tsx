import React, { Children, Suspense, lazy } from "react";
import { Navigate, Outlet } from "react-router-dom";

export function PrivateRoute() {
  const isAuth = false;
  return isAuth ? <Outlet /> : <Navigate to={"/"} />;
}
