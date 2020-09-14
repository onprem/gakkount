import React from "react";
import { useLocation } from "react-router-dom";
import useSWR from "swr";
import Button from "../../components/button";
import styles from "./consent.module.css";

const useConsentChallenge = (cc: string) => {
  const { data, error } = useSWR(`/consent/${cc}`, (input: RequestInfo, init?: RequestInit) =>
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

  const { consent, error, loading } = useConsentChallenge(cc);
  if (loading) return <div>Loading...</div>;
  if (error) return <div>Some error occurred.</div>;

  const handleSubmit = (allow: boolean) => {
    fetch("/consent", {
      method: "POST",
      body: JSON.stringify({
        allow: allow,
        challenge: cc,
      }),
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

  const { user, client } = consent;
  console.log("concent", consent);
  return (
    <div className={styles.container}>
      <main className={styles.main}>
        <p>
          Hi {user}, {client.client_name || client.client_id} wants to access:
        </p>
        <ul>
          <li>Your profile</li>
        </ul>
        <hr />
        <div className={styles.btnGroup}>
          <Button type="button" variant="link" onClick={() => handleSubmit(false)}>
            Deny
          </Button>
          <Button type="button" onClick={() => handleSubmit(true)}>
            Allow
          </Button>
        </div>
      </main>
    </div>
  );
};

export default Consent;
