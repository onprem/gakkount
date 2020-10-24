import React from "react";
import { Redirect } from "react-router-dom";
import useSWR from "swr";
import { useAuth } from "../../contexts/auth";
import { ReactComponent as LoadingIcon } from "../../assets/three-dots.svg";

export const Logout: React.FC = () => {
  const { isLoggedIn, setIsLoggedIn } = useAuth();
  useSWR("/api/logout", {
    onSuccess: (res) => {
      if (res.status === "success") setIsLoggedIn(false);
    },
  });

  if (!isLoggedIn) return <Redirect to="/login" />;

  return (
    <div
      style={{
        display: "flex",
        alignItems: "center",
        justifyContent: "center",
        width: "100%",
        height: "100vh",
      }}
    >
      <LoadingIcon style={{ height: "3rem" }} />
    </div>
  );
};

export default Logout;
