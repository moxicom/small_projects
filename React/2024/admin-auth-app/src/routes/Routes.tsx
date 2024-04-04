import { RouterProvider, createBrowserRouter } from "react-router-dom";
import ProtectedRoute from "./ProtectedRoute";
import { HomePage } from "../pages/HomePage";
import { AppContent } from "../pages/AppContent";

const router = createBrowserRouter([
  {
    path: "/",
    element: <AppContent />,
    children: [
      {
        path: "account",
        element: (
          <ProtectedRoute>
            <div>HELLO ACCOUNT</div>
          </ProtectedRoute>
        ),
      },
    ],
  },
  {
    index: true,
    element: <HomePage />,
  },
]);

export function Routes() {
  return <RouterProvider router={router} />;
}
