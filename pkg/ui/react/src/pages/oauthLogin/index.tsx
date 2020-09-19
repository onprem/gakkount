import React from "react";
import { useLocation } from "react-router-dom";
import { useForm } from "react-hook-form";

import styles from "./login.module.css";
import { Text, Label } from "../../components/form";
import Button from "../../components/button";
import useSWRPost from "../../hooks/useSWRPost";
import { ReactComponent as LoadingIcon } from "../../assets/three-dots.svg";
import { useAuth } from "../../contexts/auth";

export const OAuthLogin: React.FC = () => {
  const { register, handleSubmit } = useForm();
  const location = useLocation();
  const query = new URLSearchParams(location.search);
  const lc = query.get("login_challenge");
  const { isLoggedIn, user } = useAuth();

  const [acceptChallenge, { isValidating }] = useSWRPost<string>("/oauth/challenge", {
    onSuccess: (res) => {
      if (res.status === "success") {
        window.location = res.redirectTo;
      }
    },
  });

  console.log("lc", lc);

  const onSubmit = (values: Record<string, any>) => {
    acceptChallenge(JSON.stringify(values));
  };

  return (
    <div className={styles.container}>
      <form className={styles.form} onSubmit={handleSubmit(onSubmit)}>
        <h2>Login to continue</h2>
        <Label value="Email">
          <Text
            type="text"
            name="email"
            inpRef={register({ required: "Email is required" })}
            defaultValue={isLoggedIn ? user?.email : ""}
          />
        </Label>
        <Label value="Password">
          <Text
            type="password"
            name="password"
            inpRef={register({ required: "Password is required" })}
          />
        </Label>
        {lc && (
          <input
            type="hidden"
            name="challenge"
            value={lc}
            ref={register({ required: "Login challenge is required" })}
          />
        )}
        <Button type="submit" disabled={isValidating}>
          {isValidating ? <LoadingIcon style={{ height: "1em" }} /> : `Log In`}
        </Button>
      </form>
    </div>
  );
};

export default OAuthLogin;
