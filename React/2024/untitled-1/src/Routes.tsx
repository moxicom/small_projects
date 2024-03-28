import { createBrowserRouter, RouterProvider } from "react-router-dom";
import App from "./App";
import FetchPage from "./pages/FetchPage";
import ErrorPage from "./pages/ErrorPage";
import { HomePage } from "./pages/HomePage";
import { JsonParsePage } from "./pages/JsonParsePage";
import { LoginPage } from "./pages/LoginPage";
import { AccountPage } from "./pages/AccountPage";

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
      {
        path: "json",
        element: <JsonParsePage />,
      },
      {
        path: "login",
        element: <LoginPage />,
      },
      {
        path: "account",
        element: <AccountPage />,
      },
    ],
  },
]);

export function Routes() {
  return <RouterProvider router={router} />;
}
