import { createContext, useContext } from "react";
import { User } from "../interfaces";

export interface AuthData {
  isLoggedIn: boolean;
  setIsLoggedIn: React.Dispatch<React.SetStateAction<boolean>>;
  user?: User;
  setUser?:  React.Dispatch<React.SetStateAction<User | undefined>>;
}

export const AuthContext = createContext<AuthData>({ isLoggedIn: false, setIsLoggedIn: () => {} });

export function useAuth() {
  return useContext(AuthContext);
}
