import React from "react";
import { useLocation } from "react-router-dom";
import useSWR from "swr";
import { useToasts } from "react-toast-notifications";
import Button from "../../components/button";
import styles from "./consent.module.css";
import useSWRPost from "../../hooks/useSWRPost";
import logoImg from "../../assets/logo.png";
import { ReactComponent as LoadingIcon } from "../../assets/three-dots.svg";

const useConsentChallenge = (cc: string) => {
  const { data, error } = useSWR(`/oauth/consent/${cc}`, (input: RequestInfo, init?: RequestInit) =>
    fetch(input, init).then((res) => res.json())
  );
  return {
    consent: data,
    error,
    loading: !(data || error),
  };
};

export const Consent: React.FC = () => {
  const location = useLocation();
  const { addToast } = useToasts();

  const query = new URLSearchParams(location.search);
  const cc = query.get("consent_challenge") || "";
  const [giveConsent, { isValidating }] = useSWRPost("/oauth/consent", {
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

  const { consent, error, loading } = useConsentChallenge(cc);
  if (loading)
    return (
      <div className={styles.container}>
        <LoadingIcon style={{ height: "3em" }} />
      </div>
    );
  if (error || consent?.status !== "success")
    return <div className={styles.container}>Some error occurred.</div>;

  const handleSubmit = (allow: boolean) => {
    giveConsent(
      JSON.stringify({
        allow: allow,
        challenge: cc,
      })
    );
  };

  const { user, client } = consent;

  return (
    <div className={styles.container}>
      <div className={styles.logoDiv}>
        <img src={logoImg} alt="logo" className={styles.mainLogo} />
        <h1 className={styles.title}>Account Login</h1>
      </div>
      <main className={styles.main}>
        <hr className={styles.hr} />
        {client.logo_uri && <img src={client.logo_uri} alt="client logo" className={styles.logo} />}
        <p className={styles.center}>
          Hi {user}, <b>{client.client_name || client.client_id}</b> wants to access your IIITM
          account.
        </p>
        <p>
          This will allow <b>{client.client_name || client.client_id}</b> to:
        </p>
        <ul className={styles.list}>
          <li>View your name and photo</li>
          <li>Associate you with your personal info on IIITM Accounts</li>
          <li>See your personal info, including details like your roll no. and course</li>
          <li>View your personal and contact details</li>
        </ul>
        {/* <hr className={styles.hr} /> */}
        <div className={styles.btnGroup}>
          <Button
            type="button"
            variant="link"
            disabled={isValidating}
            onClick={() => handleSubmit(false)}
          >
            Deny
          </Button>
          <Button
            type="button"
            disabled={isValidating}
            onClick={() => handleSubmit(true)}
            style={{ background: "var(--color-brand)" }}
          >
            {isValidating ? <LoadingIcon style={{ height: "1em" }} /> : `Allow`}
          </Button>
        </div>
      </main>
    </div>
  );
};

export default Consent;
