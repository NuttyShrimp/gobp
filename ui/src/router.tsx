import { createBrowserRouter } from "react-router-dom";
import { IndexPage } from "./pages";
import { LoginPage } from "./pages/auth/login";
import { AuthenticatedLayout } from "./components/layout/Authenticated";
import { wrapCreateBrowserRouter } from "@sentry/react";
import { SentryErrorBoundary } from "./components/ErrorBoundary";
import { AuthProvider } from "./lib/context/auth";

const sentryCreateBrowserRouter = wrapCreateBrowserRouter(createBrowserRouter);

export const router = sentryCreateBrowserRouter([
  {
    path: "/",
    errorElement: import.meta.env.PROD && <SentryErrorBoundary />,
    children: [
      {
        path: "",
        element: <AuthProvider><AuthenticatedLayout /></AuthProvider>,
        children: [
          {
            element: <IndexPage />,
            index: true,
          }
        ]
      },
      {
        path: "login",
        element: <LoginPage />,
      },
    ]
  },
]);
