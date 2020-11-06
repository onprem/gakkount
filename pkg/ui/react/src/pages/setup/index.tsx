import React from "react";
import { useForm } from "react-hook-form";
import { Redirect, useHistory } from "react-router-dom";
import { useToasts } from "react-toast-notifications";

import { Text, Label } from "../../components/form";
import Button from "../../components/button";
import { useAuth } from "../../contexts/auth";
import useSWRPost from "../../hooks/useSWRPost";
import useFormErrors from "../../hooks/useFormErrors";
import { ReactComponent as LoadingIcon } from "../../assets/three-dots.svg";
import logoImg from "../../assets/logo.png";

import styles from "../login/login.module.css";

export const Setup: React.FC = () => {
  const { addToast } = useToasts();
  const { register, handleSubmit, errors } = useForm();
  useFormErrors(errors);
  const { isLoggedIn } = useAuth();
  const history = useHistory();
  const [runSetup, { isValidating }] = useSWRPost<string>("/api/setup", {
    onSuccess: (res) => {
      if (res.status === "success") {
        history.replace("/login");
      } else addToast(res?.error || "Something went wrong", { appearance: "error" });
    },
    onError: (err) => {
      addToast(err?.error || "Something went wrong.", { appearance: "error" });
    },
  });

  if (isLoggedIn) return <Redirect to="/profile" />;

  const onSubmit = (values: Record<string, any>) => {
    runSetup(JSON.stringify(values));
  };

  return (
    <div className={styles.container}>
      <div className={styles.logoDiv}>
        <img src={logoImg} alt="logo" className={styles.logo} />
        <h1 className={styles.title}>Admin Account Setup</h1>
      </div>
      <form className={styles.form} onSubmit={handleSubmit(onSubmit)}>
        <p>Enter name, admin email and new password to continue</p>
        <Label value="Name">
          <Text
            type="text"
            name="name"
            inpRef={register({ required: "name is required" })}
            placeholder="Enter your name"
          />
        </Label>
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
            autoComplete="new-password"
            placeholder="Enter password"
            inpRef={register({ required: "Password is required" })}
          />
        </Label>
        <Button type="submit" disabled={isValidating} style={{ background: "var(--color-brand)" }}>
          {isValidating ? <LoadingIcon style={{ height: "1em" }} /> : `Create`}
        </Button>
        <hr />
        <p>This account will have admin access.</p>
      </form>
    </div>
  );
};

export default Setup;
