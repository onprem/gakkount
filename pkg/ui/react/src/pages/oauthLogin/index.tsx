import React from "react";
import { useLocation } from "react-router-dom";
import { useForm } from "react-hook-form";
import { useToasts } from "react-toast-notifications";

import styles from "../login/login.module.css";
import { Text, Label } from "../../components/form";
import Button from "../../components/button";
import useSWRPost from "../../hooks/useSWRPost";
import useFormErrors from "../../hooks/useFormErrors";
import { ReactComponent as LoadingIcon } from "../../assets/three-dots.svg";
import logoImg from "../../assets/logo.png";
import { useAuth } from "../../contexts/auth";

interface Payload {
  email: string;
  password: string;
  challenge: string;
}

export const OAuthLogin: React.FC = () => {
  const { register, handleSubmit, errors } = useForm<Payload>();
  const location = useLocation();
  const { addToast } = useToasts();
  useFormErrors(errors);

  const query = new URLSearchParams(location.search);
  const lc = query.get("login_challenge");
  const { isLoggedIn, user } = useAuth();

  const [acceptChallenge, { isValidating }] = useSWRPost<string>("/oauth/challenge", {
    onSuccess: (res) => {
      if (res.status === "success") {
        window.location = res.redirectTo;
      } else {
        addToast(res?.error || "Something went wrong", { appearance: "error" });
      }
    },
    onError: (err) => {
      addToast(err?.error || "Something went wrong.", { appearance: "error" });
    },
  });

  const onSubmit = (values: Payload) => {
    acceptChallenge(JSON.stringify(values));
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
            placeholder="Enter e-mail address"
            inpRef={register({ required: "Email is required" })}
            defaultValue={isLoggedIn ? user?.email : ""}
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
        {lc && (
          <input
            type="hidden"
            name="challenge"
            value={lc}
            ref={register({ required: "Login challenge is required" })}
          />
        )}
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

export default OAuthLogin;
