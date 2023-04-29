import { useState } from "react";
import { createContext } from "react";

export const AdminFlagContext = createContext({});

export const AdminFlagProvider = props => {
  const { children } = props;
  const [isAdmin, setIsAdmin] = useState(false);
  return (
    <AdminFlagContext.Provider value={{isAdmin, setIsAdmin}}>
      {children}
    </AdminFlagContext.Provider>
  )
}
