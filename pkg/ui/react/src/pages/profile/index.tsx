import React, { useState } from "react";

import Layout from "../../components/layout";
import styles from "./profile.module.css";
import { useAuth } from "../../contexts/auth";
import { Label, Text } from "../../components/form";
import Button from "../../components/button";
import { useForm } from "react-hook-form";
import useSWRPost from "../../hooks/useSWRPost";
import { User } from "../../interfaces";
import { mutate } from "swr";

interface Response {
  status: string;
  error?: string;
  message?: string;
  user?: User;
}

interface Payload {
  phone: string;
  altEmail: string;
  linkedin: string;
  github: string;
  twitter: string;
}

const Profile = () => {
  const [editing, setEditing] = useState<boolean>(false);
  const { user } = useAuth();
  const { register, handleSubmit } = useForm();
  const [runUpdate] = useSWRPost<string>(
    "/api/user",
    {
      onSuccess: (data: Response) => {
        if (data.status === "success") {
          mutate("/api/user", data.user);
          console.log(data.message);
          setEditing(false);
        } else {
          console.log(data.error);
        }
      },
      onError: console.log,
    },
    "PATCH"
  );

  if (!user) return <h2>Loading...</h2>;

  const onSubmit = (values: Payload) => {
    runUpdate(JSON.stringify(values));
  };

  return (
    <Layout>
      <form className={styles.main} onSubmit={handleSubmit<Payload>(onSubmit)}>
        <section className={styles.content}>
          <h1>User Profile</h1>
          <Label value="Name">
            <Text defaultValue={user.name} disabled={editing} readOnly />
          </Label>
          <div className={styles.group}>
            <Label value="Email">
              <Text defaultValue={user.email} disabled={editing} readOnly />
            </Label>
            <Label value="Role">
              <Text
                defaultValue={user.role}
                style={{ textTransform: "capitalize" }}
                disabled={editing}
                readOnly
              />
            </Label>
          </div>

          {user.role === "student" && (
            <>
              <div className={styles.group}>
                <Label value="Roll No.">
                  <Text
                    style={{ textTransform: "uppercase" }}
                    defaultValue={user.rollNo}
                    disabled={editing}
                    readOnly
                  />
                </Label>
                {user?.admissionTime && (
                  <Label value="Admission Year">
                    <Text
                      defaultValue={new Date(user.admissionTime).getFullYear()}
                      disabled={editing}
                      readOnly
                    />
                  </Label>
                )}
              </div>
              <Label value="Course">
                <Text defaultValue={user.edges.Course?.name} disabled={editing} readOnly />
              </Label>
            </>
          )}
          {user.role === "faculty" && (
            <Label value="Department">
              <Text defaultValue={user.edges.Department?.name} disabled={editing} readOnly />
            </Label>
          )}
          <h2>Contact Details</h2>
          <div className={styles.group}>
            <Label value="Phone">
              <Text
                name="phone"
                inpRef={register()}
                defaultValue={user?.phone}
                readOnly={!editing}
              />
            </Label>
            <Label value="Alt Email">
              <Text
                name="altEmail"
                inpRef={register()}
                defaultValue={user?.altEmail}
                readOnly={!editing}
              />
            </Label>
          </div>
          <h2>Social Media Profile</h2>
          <Label value="LinkedIn">
            <Text
              name="linkedin"
              inpRef={register()}
              defaultValue={user?.linkedin}
              readOnly={!editing}
            />
          </Label>
          <Label value="GitHub">
            <Text
              name="github"
              inpRef={register()}
              defaultValue={user?.github}
              readOnly={!editing}
            />
          </Label>
          <Label value="Twitter">
            <Text
              name="twitter"
              inpRef={register()}
              defaultValue={user?.twitter}
              readOnly={!editing}
            />
          </Label>
        </section>
        <section className={styles.aside}>
          {editing ? (
            <>
              <Button
                style={{ background: "var(--color-danger)" }}
                type="button"
                onClick={() => setEditing(false)}
              >
                Cancel
              </Button>
              <Button style={{ background: "var(--color-success)" }}>Submit</Button>
            </>
          ) : (
            <Button
              style={{ background: "var(--color-hint)" }}
              type="button"
              onClick={() => setEditing(true)}
            >
              Modify Details
            </Button>
          )}
        </section>
      </form>
    </Layout>
  );
};

export default Profile;
