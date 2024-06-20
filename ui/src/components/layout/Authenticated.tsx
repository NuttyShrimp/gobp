import { useQuery } from "@tanstack/react-query"
import { Outlet } from "react-router-dom"
import { useAuth } from "../../lib/context/auth";
import { toast } from "sonner";

export const AuthenticatedLayout = () => {
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
    logout();
    toast.warning("Session has expired...")
    return (
      <div className="h-svh w-svw flex flex-col items-center justify-center">
        <p>Logging out...</p>
      </div>
    )
  }

  return (
    <Outlet />
  )
}
