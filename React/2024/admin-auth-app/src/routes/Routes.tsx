import {
  BrowserRouter,
  Navigate,
  Route,
  Routes,
  createBrowserRouter,
} from "react-router-dom";
import { AppPage } from "../pages/AppPage";
import { App } from "../App";

import { Suspense, lazy } from "react";

const AdminPageLazy = lazy(() => import("../pages/AdminPage"));

const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
    children: [
      {
        index: true,
        element: <AppPage />,
      },
      {
        path: "/admin",
        element: <AdminPageLazy />,
      },
    ],
  },
]);

export function AppRoutes() {
  const isAdmin = true;
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<App />}>
          <Route index element={<AppPage />} />

          <Route
            path="admin"
            element={
              isAdmin ? (
                <Suspense fallback={<div>Loading...</div>}>
                  <AdminPageLazy />{" "}
                </Suspense>
              ) : (
                <Navigate to={"/"} />
              )
            }
          />
        </Route>
      </Routes>
    </BrowserRouter>
  );
}
