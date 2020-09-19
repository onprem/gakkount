import React from "react";
import { useForm } from "react-hook-form";

import styles from "./login.module.css";
import { Text, Label } from "../../components/form";
import Button from "../../components/button";
import { useAuth } from "../../contexts/auth";
import { Redirect, useLocation } from "react-router-dom";

export const Login: React.FC = () => {
  const { register, handleSubmit } = useForm();
  const { isLoggedIn, setIsLoggedIn } = useAuth();
  const location = useLocation<{ referrer?: Location }>();

  if (isLoggedIn) return <Redirect to={location.state?.referrer || "/dashboard"} />

  const onSubmit = (values: Record<string, any>) => {
    fetch("/api/login", {
      method: "POST",
      body: JSON.stringify(values),
      headers: {
        "Content-Type": "application/json",
      },
    })
      .then((res) => res.json())
      .then((res) => {
        if (res.status === "success") {
          setIsLoggedIn(true);
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
        <Button type="submit">Log In</Button>
      </form>
    </div>
  );
};

export default Login;
