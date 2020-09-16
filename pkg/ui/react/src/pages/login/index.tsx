import React from "react";
import { useLocation } from "react-router-dom";
import { useForm } from "react-hook-form";

import styles from "./login.module.css";
import { Text, Label } from "../../components/form";
import Button from "../../components/button";

export const Login: React.FC = () => {
  const { register, handleSubmit } = useForm();
  const location = useLocation();
  const query = new URLSearchParams(location.search);
  const lc = query.get("login_challenge");

  console.log("lc", lc);

  const onSubmit = (values: Record<string, any>) => {
    fetch("/oauth/challenge", {
      method: "POST",
      body: JSON.stringify(values),
      headers: {
        "Content-Type": "application/json",
      },
    })
      .then((res) => res.json())
      .then((res) => {
        console.log(res);
        if (res.status === "success") {
          window.location = res.redirectTo;
        }
      });
  };

  return (
    <div className={styles.container}>
      <form className={styles.form} onSubmit={handleSubmit(onSubmit)}>
        <h2>Login to continue</h2>
        <Label value="Email">
          <Text type="text" name="email" inpRef={register({ required: "Email is required" })} />
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
        <Button type="submit">Log In</Button>
      </form>
    </div>
  );
};

export default Login;
