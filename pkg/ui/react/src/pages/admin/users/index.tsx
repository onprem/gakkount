import React, { useState } from "react";
import useSWR from "swr";

import { Select } from "../../../components/form";
import Layout from "../../../components/layout";
import AddUser from "./addUser";

import { User } from "../../../interfaces";

import styles from "./users.module.css";

const UserRow: React.FC<{ user: User }> = ({ user }) => {
  const {
    id,
    name,
    email,
    rollNo,
    edges: { Course },
  } = user;
  return (
    <div className={styles.user}>
      <span className={styles.idSpan}>{id}</span>
      <span className={styles.nameSpan}>{name}</span>
      <span className={styles.emailSpan}>{email}</span>
      <span className={styles.rollSpan} style={{ textTransform: "uppercase" }}>
        {rollNo}
      </span>
      <span className={styles.courseSpan}>{Course?.name}</span>
    </div>
  );
};

const Users = () => {
  const [role, setRole] = useState<User["role"]>("student");
  const { data, isValidating } = useSWR<{ status: string; users: User[] }>(
    `/api/users?role=${role}`
  );

  const list = data?.users.map((u) => <UserRow user={u} key={u.id} />);

  return (
    <Layout>
      <div className={styles.main}>
        <h1 className={styles.head}>All Users</h1>
        <div className={styles.options}>
          <AddUser />
          <Select onChange={(e) => setRole(e.target.value as User["role"])} value={role}>
            <option value="student">Student</option>
            <option value="faculty">Faculty</option>
            <option value="staff">Staff</option>
            <option value="admin">Admin</option>
            <option value="misc">Misc.</option>
          </Select>
        </div>
        {!isValidating && (
          <div className={styles.list}>
            <div className={`${styles.user} ${styles.topUser}`}>
              <span className={styles.idSpan}>ID</span>
              <span className={styles.nameSpan}>Name</span>
              <span className={styles.emailSpan}>Email</span>
              <span className={styles.rollSpan}>RollNo.</span>
              <span className={styles.courseSpan}>Course</span>
            </div>
            {list}
          </div>
        )}
      </div>
    </Layout>
  );
};

export default Users;
