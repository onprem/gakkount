import React, { useState } from "react";
import useSWR from "swr";

import Layout from "../../components/layout";
import styles from "./style.module.css";
import { ExtClient } from "../../interfaces";
import { ReactComponent as LoadingIcon } from "../../assets/three-dots.svg";
import { ReactComponent as ClientIcon } from "../../assets/client.svg";
import Button from "../../components/button";
import { Label, Text } from "../../components/form";

interface Response {
  status: string;
  error?: string;
  clients?: ExtClient[];
}

const Profile = () => {
  const [client, setClient] = useState<ExtClient>();
  const { data, isValidating } = useSWR<Response>("/api/clients");

  if (isValidating || data?.error)
    return (
      <Layout>
        <main className={styles.center}>
          <LoadingIcon style={{ height: "2rem" }} />
        </main>
      </Layout>
    );

  return (
    <Layout>
      <main className={styles.main}>
        <section className={styles.left}>
          <h1>OAuth Clients</h1>
          <div className={styles.cards}>
            {data?.clients?.map((e) => (
              <div className={styles.client} key={e.client.clientID}>
                <div className={styles.nameDiv}>
                  <img
                    src={
                      e.payload.logo_uri ||
                      `https://ui-avatars.com/api/?name=${e.client.name}&size=196&background=465062&color=f1f1f1`
                    }
                    alt="client logo"
                  />
                  <h3>{e.client.name}</h3>
                </div>
                <hr />
                <div className={styles.dateDiv}>
                  <p>Added on {new Date(e.payload.created_at).toDateString()}</p>
                  <Button variant="link" onClick={() => setClient(e)}>
                    Manage
                  </Button>
                </div>
              </div>
            ))}
          </div>
        </section>
        <section className={styles.right}>
          {client ? (
            <>
              <h2>{client.client.name}</h2>
              <Label value="Client ID">
                <Text defaultValue={client.client.clientID} readOnly />
              </Label>
              <Label value="Client Secret">
                <Text defaultValue={client.client.secret} readOnly />
              </Label>
            </>
          ) : (
            <div className={styles.empty}>
              <ClientIcon className={styles.icon} />
              <h3>Select a client</h3>
            </div>
          )}
        </section>
      </main>
    </Layout>
  );
};

export default Profile;
