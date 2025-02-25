import { createBrowserRouter } from "react-router-dom";
import { IndexPage } from "./pages";
import { LoginPage } from "./pages/auth/login";
import { AuthenticatedLayout } from "./components/layout/Authenticated";
import { wrapCreateBrowserRouterV6 } from "@sentry/react";
import { SentryErrorBoundary } from "./components/ErrorBoundary";
import { AuthProvider } from "./lib/context/auth";

const sentryCreateBrowserRouter = wrapCreateBrowserRouterV6(createBrowserRouter);

export const router = sentryCreateBrowserRouter([
  {
    path: "/",
    errorElement: import.meta.env.PROD && <SentryErrorBoundary />,
    children: [
      {
        path: "login",
        element: <LoginPage />,
      },
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
    ]
  },
]);
