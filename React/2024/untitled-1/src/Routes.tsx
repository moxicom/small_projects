import { createBrowserRouter, RouterProvider } from "react-router-dom";
import App from "./App";
import FetchPage from "./pages/FetchPage";
import ErrorPage from "./pages/ErrorPage";
import { HomePage } from "./pages/HomePage";

const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
    errorElement: <ErrorPage />,
    children: [
      {
        index: true,
        element: <HomePage />,
      },
      {
        path: "fetch",
        element: <FetchPage />,
      },
    ],
  },
]);

export function Routes() {
  return <RouterProvider router={router} />;
}
