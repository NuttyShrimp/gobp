import { useAuth } from "@/lib/context/auth";

export const IndexPage = () => {
  const { name } = useAuth();

  return (
    <div>
      <p>Welcome {name}</p>
    </div>
  )
}
