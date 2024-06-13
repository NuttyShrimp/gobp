import { useQuery } from "@tanstack/react-query"
import { Navigate, Outlet } from "react-router-dom"
import { useAuth } from "../../lib/context/auth";
import { toast } from "sonner";

export const AuthenticatedLayout = () => {
  const { isLoading, isError } = useQuery({
    queryKey: ["session"],
    queryFn: () => fetch("/api/auth/session"),
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

  if (isError) {
    logout();
    toast.warning("Session has expired...")
    return (<Navigate to={"/login"} />)
  }

  return (
    <Outlet />
  )
}
