import { createBrowserRouter } from "react-router-dom";
import { IndexPage } from "./pages";
import { LoginPage } from "./pages/auth/login";
import { AuthenticatedLayout } from "./components/layout/Authenticated";
import { wrapCreateBrowserRouter } from "@sentry/react";
import { SentryErrorBoundary } from "./components/ErrorBoundary";

const sentryCreateBrowserRouter = wrapCreateBrowserRouter(createBrowserRouter);

export const router = sentryCreateBrowserRouter([
  {
    path: "/",
    errorElement: import.meta.env.PROD && <SentryErrorBoundary />,
    children: [
      {
        path: "",
        element: <AuthenticatedLayout />,
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
