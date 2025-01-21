import { useQuery } from "@tanstack/react-query"
import { Outlet, useNavigate } from "react-router-dom"
import { useAuth } from "../../lib/context/auth";
import { notifications } from "@mantine/notifications";
import { AppShell } from "../appshell";

export const AuthenticatedLayout = () => {
  const navigate = useNavigate();
  const { isLoading, isError, data } = useQuery({
    queryKey: ["session"],
    queryFn: () => fetch("/api/auth/session", {
      redirect: "error"
    }),
    retry: false,
    refetchOnWindowFocus: false,
  })
  const { logout } = useAuth();
  if (isLoading) {
    return (
      <div className="h-svh w-svw flex flex-col items-center justify-center">
        <p>Loading Session</p>
      </div>
    )
  }

  if (isError || data?.status !== 200) {
    Promise.resolve(logout()).then(() => navigate("/login"))
    notifications.show({
      variant: "warning",
      message: "Session has expired..."
    })
    return (
      <div className="h-svh w-svw flex flex-col items-center justify-center">
        <p>Logging out...</p>
      </div>
    )
  }

  return (
    <>
      <AppShell>
        <Outlet />
      </AppShell>
    </>
  )
}
