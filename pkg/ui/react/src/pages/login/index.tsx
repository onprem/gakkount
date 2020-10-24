import React from "react";
import { useForm } from "react-hook-form";

import styles from "./login.module.css";
import { Text, Label } from "../../components/form";
import Button from "../../components/button";
import { useAuth } from "../../contexts/auth";
import { Redirect, useLocation } from "react-router-dom";
import useSWRPost from "../../hooks/useSWRPost";
import { ReactComponent as LoadingIcon } from "../../assets/three-dots.svg";
import logoImg from "../../assets/logo.png";

export const Login: React.FC = () => {
  const { register, handleSubmit } = useForm();
  const { isLoggedIn, setIsLoggedIn } = useAuth();
  const location = useLocation<{ referrer?: Location }>();
  const [runLogin, { isValidating }] = useSWRPost<string>("/api/login", {
    onSuccess: (res) => {
      if (res.status === "success") {
        setIsLoggedIn(true);
      }
    },
  });

  if (isLoggedIn) return <Redirect to={location.state?.referrer || "/profile"} />;

  const onSubmit = (values: Record<string, any>) => {
    runLogin(JSON.stringify(values));
  };

  return (
    <div className={styles.container}>
      <div className={styles.logoDiv}>
        <img src={logoImg} alt="logo" className={styles.logo} />
        <h1 className={styles.title}>Account Login</h1>
      </div>
      <form className={styles.form} onSubmit={handleSubmit(onSubmit)}>
        <p>Enter email address and password to continue</p>
        <Label value="Email">
          <Text
            type="text"
            name="email"
            inpRef={register({ required: "Email is required" })}
            placeholder="Enter e-mail address"
          />
        </Label>
        <Label value="Password">
          <Text
            type="password"
            name="password"
            placeholder="Enter password"
            inpRef={register({ required: "Password is required" })}
          />
        </Label>
        <Button type="submit" disabled={isValidating} style={{ background: "var(--color-brand)" }}>
          {isValidating ? <LoadingIcon style={{ height: "1em" }} /> : `Log In`}
        </Button>
        <hr />
        <p>
          Don't have an account?{" "}
          <a href="mailto:admin@iiitm.ac.in" className={styles.link}>
            Contact Admin
          </a>
        </p>
      </form>
    </div>
  );
};

export default Login;
