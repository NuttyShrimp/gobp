import { useQuery } from "@tanstack/react-query";
import { PropsWithChildren, createContext, useContext, useEffect, useState } from "react";
import { User } from "../types";
import { toast } from "sonner";

declare type AuthContext = {
  loggedIn: boolean;
  name: string;
  logout: () => void;
}

const authContext = createContext<AuthContext | undefined>(undefined);

export const AuthProvider = ({ children }: PropsWithChildren) => {
  const [loggedIn, setLoggedIn] = useState(false);
  const [name, setName] = useState('');
  const { data } = useQuery<User>({
    queryKey: ["user"],
    queryFn: async () => {
      const resp = await fetch("/api/user/me")
      return resp.json();
    },
    refetchOnWindowFocus: false,
  })

  useEffect(() => {
    if (data) {
      setLoggedIn(true);
      setName(data.name)
    }
  }, [data])

  const logout = async () => {
    const resp = await fetch("/auth/logout")
    if (!resp.ok) {
      console.error("Failed to log out");
    }
    toast.success("Logged Out!");
    setLoggedIn(false);
    setName("");
    location.href = "/login"
  }

  return (
    <authContext.Provider value={{
      name,
      loggedIn,
      logout,
    }}>
      {children}
    </authContext.Provider>
  )
}

export const useAuth = () => {
  const context = useContext(authContext);
  if (!context) {
    throw Error('useAuth must be used within an AuthProvider');

  }
  return context;
}
