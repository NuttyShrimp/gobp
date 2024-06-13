import { createBrowserRouter } from "react-router-dom";
import { IndexPage } from "./pages";
import { LoginPage } from "./pages/auth/login";
import { AuthenticatedLayout } from "./components/layout/Authenticated";

export const router = createBrowserRouter([
  {
    path: "/",
    element: <AuthenticatedLayout />,
    children: [
      {
        element: <IndexPage />,
        index: true,
      }
    ]
  },
  {
    path: "/login",
    element: <LoginPage />,
  },
]);
