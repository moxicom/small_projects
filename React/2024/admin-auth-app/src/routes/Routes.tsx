import { RouterProvider, createBrowserRouter } from "react-router-dom";
import ProtectedRoute from "./ProtectedRoute";
import { HomePage } from "../pages/HomePage";
import { AppContent } from "../pages/AppContent";
import ListsPage from "../pages/ListsPage";
import RegisterPage from "../pages/RegisterPage";
import LoginPage from "../pages/LoginPage";
import OnlyUnregisteredRoute from "./OnlyUnRegisteredRoute";

const router = createBrowserRouter([
  {
    path: "/",
    element: <AppContent />,
    children: [
      {
        path: "lists",
        element: (
          <ProtectedRoute>
            <ListsPage />
          </ProtectedRoute>
        ),
      },
    ],
  },
  {
    index: true,
    element: <HomePage />,
  },
  {
    path: "/register",
    element: (
      <OnlyUnregisteredRoute>
        <RegisterPage />
      </OnlyUnregisteredRoute>
    ),
  },
  {
    path: "/login",
    element: (
      <OnlyUnregisteredRoute>
        <LoginPage />
      </OnlyUnregisteredRoute>
    ),
  },
]);

export function Routes() {
  return <RouterProvider router={router} />;
}
