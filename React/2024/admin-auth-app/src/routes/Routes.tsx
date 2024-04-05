import { RouterProvider, createBrowserRouter } from "react-router-dom";
import ProtectedRoute from "./ProtectedRoute";
import { HomePage } from "../pages/HomePage";
import { AppContent } from "../pages/AppContent";
import ListsPage from "../pages/ListsPage";
import RegisterPage from "../pages/RegisterPage";

const router = createBrowserRouter([
  {
    path: "/",
    element: <AppContent />,
    children: [
      {
        path: "lists",
        element: (
          <ProtectedRoute>
            {" "}
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
    element: <RegisterPage />,
  },
]);

export function Routes() {
  return <RouterProvider router={router} />;
}
