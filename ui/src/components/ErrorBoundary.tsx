import { useRouteError } from "react-router-dom";
import * as Sentry from "@sentry/react";
import { useEffect } from "react";

export function SentryErrorBoundary() {
  const error = useRouteError() as Error;

  useEffect(() => {
    Sentry.captureException(error);
  }, [error]);

  return (
    <div>
      <h1>Oh no, it seems like something went terribly wrong!</h1>
      <p>But team software has been notified of this error</p>
    </div>
  );
}
