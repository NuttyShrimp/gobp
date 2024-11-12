import { usePermission } from "@/lib/hooks/usePermission"
import { NavLinkProps, NavLink } from "react-router-dom"

export const ProtectedLink = ({ permission, ...props }: NavLinkProps & { permission: string }) => {
  const { isAllowed, isLoading } = usePermission(permission);

  if (isLoading) {
    return <p>Checking permissions...</p>
  }

  return isAllowed ? (
    <NavLink {...props} />
  ) : null
}
