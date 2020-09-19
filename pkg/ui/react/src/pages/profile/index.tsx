import React from "react";

import styles from "./profile.module.css";
import { useAuth } from "../../contexts/auth";
import { User } from "../../interfaces";

export interface UserDetailsProps {
  user: User;
}
const UserDetails: React.FC<UserDetailsProps> = ({ user }) => {
  return (
    <div className={styles.udMain}>
      <div className={styles.photo}>
        <img
          src={
            user?.photo ||
            `https://ui-avatars.com/api/?name=${user.name}&size=196&background=465062&color=f1f1f1` ||
            `https://api.adorable.io/avatars/196/${user.email}.png`
          }
          alt="user"
        />
      </div>
      <div className={styles.details}>
        <h1 className={styles.name}>{user.name}</h1>
        <span className={styles.role}>{user.role}</span>
        <div className={styles.info}>
          <span>
            <b>Email: </b>
            <span>{user.email}</span>
          </span>
          {user.altEmail && (
            <span>
              <b>Alt Email: </b>
              <span>{user.altEmail}</span>
            </span>
          )}
          {user.role === "student" && (
            <>
              <span>
                <b>Roll No.: </b>
                <span style={{ textTransform: "uppercase" }}>{user.rollNo}</span>
              </span>
            </>
          )}
        </div>
      </div>
    </div>
  );
};

const Profile = () => {
  const { user } = useAuth();

  if (!user) return <h2>Loading...</h2>;

  return (
    <div className={styles.main}>
      <UserDetails user={user} />
    </div>
  );
};

export default Profile;
