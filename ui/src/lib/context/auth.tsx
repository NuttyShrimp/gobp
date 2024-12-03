import { useQuery } from "@tanstack/react-query";
import { PropsWithChildren, createContext, useContext, useEffect, useState } from "react";
import { User } from "../types";
import { notifications } from "@mantine/notifications";
import { useNavigate } from "react-router-dom";

declare type AuthContext = {
  name: string;
  logout: () => void;
}

const authContext = createContext<AuthContext | undefined>(undefined);

export const AuthProvider = ({ children }: PropsWithChildren) => {
  const [name, setName] = useState('');
  const navigate = useNavigate();
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
      setName(data.name)
    }
  }, [data])

  const logout = async () => {
    const resp = await fetch("/api/auth/logout")
    if (!resp.ok) {
      console.error("Failed to log out");
    }
    notifications.show({
      variant: "success",
      message: "Logged Out!"
    });
    setName("");
    navigate("/login")
  }

  return (
    <authContext.Provider value={{
      name,
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
