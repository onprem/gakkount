import React from "react";
import { useLocation } from "react-router-dom";
import useSWR from "swr";
import Button from "../../components/button";
import styles from "./consent.module.css";
import useSWRPost from "../../hooks/useSWRPost";
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
  const query = new URLSearchParams(location.search);
  const cc = query.get("consent_challenge") || "";
  const [giveConsent, { isValidating }] = useSWRPost("/oauth/consent", {
    onSuccess: (res) => {
      console.log(res);
      if (res.status === "success") {
        window.location = res.redirectTo;
      }
    },
  });

  const { consent, error, loading } = useConsentChallenge(cc);
  if (loading)
    return (
      <div className={styles.container}>
        <LoadingIcon style={{ height: "3em" }} />
      </div>
    );
  if (error) return <div className={styles.container}>Some error occurred.</div>;

  const handleSubmit = (allow: boolean) => {
    giveConsent(
      JSON.stringify({
        allow: allow,
        challenge: cc,
      })
    );
  };

  const { user, client } = consent;
  console.log("concent", consent, client);
  return (
    <div className={styles.container}>
      <main className={styles.main}>
        {client.logo_uri && <img src={client.logo_uri} alt="client logo" className={styles.logo} />}
        <p>
          Hi {user}, <b>{client.client_name || client.client_id}</b> wants to access your account.
        </p>
        <p>
          This will allow <b>{client.client_name || client.client_id}</b> to:
        </p>
        <ul className={styles.list}>
          <li>Associate you with your personal info on IIITM Accounts</li>
          <li>See your personal info, including details like your roll no. and course</li>
          <li>View your email address</li>
        </ul>
        <hr />
        <div className={styles.btnGroup}>
          <Button
            type="button"
            variant="link"
            disabled={isValidating}
            onClick={() => handleSubmit(false)}
          >
            Deny
          </Button>
          <Button type="button" disabled={isValidating} onClick={() => handleSubmit(true)}>
            {isValidating ? <LoadingIcon style={{ height: "1em" }} /> : `Allow`}
          </Button>
        </div>
      </main>
    </div>
  );
};

export default Consent;
