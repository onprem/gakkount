import React, { useState } from "react";
import useSWR, { mutate } from "swr";
import { useForm } from "react-hook-form";
import { useToasts } from "react-toast-notifications";

import Layout from "../../components/layout";
import Button from "../../components/button";
import { Label, Text } from "../../components/form";
import Modal from "../../components/modal";

import { ReactComponent as LoadingIcon } from "../../assets/three-dots.svg";
import { ReactComponent as ClientIcon } from "../../assets/client.svg";
import { ExtClient } from "../../interfaces";
import useSWRPost from "../../hooks/useSWRPost";
import useFormErrors from "../../hooks/useFormErrors";

import styles from "./style.module.css";

interface Payload {
  name: string;
  grantTypes: string;
  responseTypes: string;
  scope: string;
  callbacks: string;
  origins: string;
  logoURI: string;
  privacyURI: string;
  tosURI: string;
}

const AddClient = () => {
  const [isOpen, setIsOpen] = useState<boolean>(false);
  const { addToast } = useToasts();
  const { register, handleSubmit, errors } = useForm<Payload>();
  useFormErrors(errors);

  const [createClient, { isValidating }] = useSWRPost<string>("/api/client", {
    onSuccess: (res) => {
      if (res.status === "success") {
        mutate("/api/clients");
        addToast(res?.message || "New client created successfully", { appearance: "success" });
        setIsOpen(false);
      } else addToast(res?.error || "Something went wrong", { appearance: "error" });
    },
    onError: (err) => {
      addToast(err?.error || "Something went wrong.", { appearance: "error" });
    },
  });

  const onSubmit = (values: Payload) => {
    const {
      name,
      grantTypes,
      responseTypes,
      scope,
      callbacks,
      origins,
      logoURI,
      tosURI,
      privacyURI,
    } = values;
    const p = {
      name,
      logoURI,
      privacyURI,
      tosURI,
      grantTypes: grantTypes.split(","),
      responseTypes: responseTypes.split(","),
      scope: scope.split(","),
      callbacks: callbacks.split(","),
      origins: origins.split(","),
    };
    createClient(JSON.stringify(p));
  };
  return (
    <>
      <Button
        type="button"
        style={{ background: "var(--color-brand)", width: "max-content" }}
        onClick={() => setIsOpen(true)}
      >
        Add New Client
      </Button>
      <Modal isOpen={isOpen} toggleModal={setIsOpen}>
        <form onSubmit={handleSubmit(onSubmit)} className={styles.addForm}>
          <h2>Create New OAuth Client</h2>
          <Label value="Name *">
            <Text
              name="name"
              placeholder="Name of the Application"
              inpRef={register({
                required: true,
              })}
            />
          </Label>
          <Label value="Grant Types *">
            <Text
              name="grantTypes"
              placeholder="Grant types you need"
              defaultValue="authorization_code,refresh_token"
              inpRef={register({
                required: true,
                pattern: {
                  value: /^[a-z_,]+$/,
                  message: "invalid grant types",
                },
              })}
            />
          </Label>
          <Label value="Response Types *">
            <Text
              name="responseTypes"
              placeholder="Response types you need"
              defaultValue="code,id_token"
              inpRef={register({
                required: true,
                pattern: {
                  value: /^[a-z_,]+$/,
                  message: "invalid response types",
                },
              })}
            />
          </Label>
          <Label value="Scope *">
            <Text
              name="scope"
              placeholder="Enter scope"
              defaultValue="openid,offline"
              inpRef={register({
                required: true,
                pattern: {
                  value: /^[a-z_,]+$/,
                  message: "invalid scope",
                },
              })}
            />
          </Label>
          <Label value="Callback URLs *">
            <Text
              name="callbacks"
              placeholder="https://example.org/callback,https://example.net/back"
              inpRef={register({
                required: true,
              })}
            />
          </Label>
          <Label value="Origins *">
            <Text
              name="origins"
              placeholder="https://example.org,http://localhost:8080"
              inpRef={register({
                required: true,
              })}
            />
          </Label>
          <Label value="Logo URI">
            <Text name="logoURI" placeholder="https://example.org/logo.png" inpRef={register()} />
          </Label>
          <Label value="Privacy policy URI">
            <Text name="privacyURI" placeholder="https://example.org/privacy" inpRef={register()} />
          </Label>
          <Label value="Terms of Service URI">
            <Text name="tosURI" placeholder="https://example.org/tos" inpRef={register()} />
          </Label>
          <Button
            type="button"
            onClick={() => setIsOpen(false)}
            style={{ background: "var(--color-danger)" }}
            disabled={isValidating}
          >
            Cancel
          </Button>
          <Button style={{ background: "var(--color-success)" }} disabled={isValidating}>
            {isValidating ? <LoadingIcon style={{ height: "1rem" }} /> : `Submit`}
          </Button>
        </form>
      </Modal>
    </>
  );
};

interface Response {
  status: string;
  message?: string;
  error?: string;
  clients?: ExtClient[];
}

export interface ClientsProps {
  all?: boolean;
}

const Clients: React.FC<ClientsProps> = ({ all = false }) => {
  const [client, setClient] = useState<ExtClient>();
  const { data, isValidating } = useSWR<Response>(`/api/clients?all=${all ? "true" : "false"}`);

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
          <h1>{all ? `All OAuth Clients` : `OAuth Clients`}</h1>
          <AddClient />
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
                <Text value={client.client.clientID} readOnly />
              </Label>
              <Label value="Client Secret">
                <Text value={client.client.secret} readOnly />
              </Label>
              {client.client.edges?.User && (
                <>
                  <Label value="User">
                    <Text value={client.client.edges.User.name} readOnly />
                  </Label>
                </>
              )}
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

export default Clients;
