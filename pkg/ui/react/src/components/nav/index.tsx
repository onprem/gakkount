import React from "react";
import { useAuth } from "../../contexts/auth";
import logoImg from "../../assets/logo.png";
import { Link } from "react-router-dom";
import Button from "../button";

import styles from "./nav.module.css";

export const Nav: React.FC = () => {
  const { isLoggedIn, setIsLoggedIn, user } = useAuth();
  const handleLogout = () => {
    fetch("/api/logout")
      .then((res) => res.json())
      .then((res) => {
        if (res.status === "success") setIsLoggedIn(false);
      });
  };
  return (
    <nav className={styles.nav}>
      <div className={styles.left}>
        <img className={styles.logo} src={logoImg} alt="logo" />
        <Link className={styles.title} to="/">
          Accounts
        </Link>
      </div>
      <div>
        {user?.role === "admin" && (
          <Link className={styles.link} to="/admin">
            Manage
          </Link>
        )}
        {isLoggedIn ? (
          <>
            <Link className={styles.link} to="/profile">
              Profile
            </Link>
            <Button className={styles.btn} onClick={handleLogout}>
              LOGOUT
            </Button>
          </>
        ) : (
          <Button className={styles.btn} variant="pill">
            <Link to="/login">LOGIN</Link>
          </Button>
        )}
      </div>
    </nav>
  );
};

export default Nav;
