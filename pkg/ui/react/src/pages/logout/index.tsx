import React from "react";
import { Redirect } from "react-router-dom";
import useSWR from "swr";
import { useToasts } from "react-toast-notifications";
import { useAuth } from "../../contexts/auth";
import { ReactComponent as LoadingIcon } from "../../assets/three-dots.svg";

export const Logout: React.FC = () => {
  const { isLoggedIn, setIsLoggedIn } = useAuth();
  const { addToast } = useToasts();

  useSWR("/api/logout", {
    onSuccess: (res) => {
      if (res.status === "success") setIsLoggedIn(false);
      else {
        addToast(res?.error || "Something went wrong", { appearance: "error" });
      }
    },
    onError: (err) => {
      addToast(err?.error || "Something went wrong.", { appearance: "error" });
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
