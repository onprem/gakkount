import React from "react";
import { Link, NavLink } from "react-router-dom";
import { useAuth } from "../../contexts/auth";
import { User } from "../../interfaces";
import logoImg from "../../assets/logo.png";
import styles from "./style.module.css";

const UserInfo = ({ user }: { user: User }) => {
  return (
    <>
      <img
        className={styles.navPhoto}
        src={
          user?.photo ||
          `https://ui-avatars.com/api/?name=${user.name}&size=196&background=465062&color=f1f1f1`
        }
        alt="user"
      />
      <h4 className={styles.navName}>{user.name}</h4>
    </>
  );
};

const Nav: React.FC = () => {
  const { user } = useAuth();
  return (
    <nav className={styles.nav}>
      <div className={styles.logoDiv}>
        <img alt="logo" src={logoImg} className={styles.logo} />
        <h1>ABV-IIITM</h1>
      </div>
      <div className={styles.accountDiv}>
        {user && (
          <>
            <UserInfo user={user} />
            <div className={styles.dropdown}>
              <Link to="/logout">Logout</Link>
            </div>
          </>
        )}
      </div>
    </nav>
  );
};

export const Layout: React.FC = ({ children }) => {
  const { user } = useAuth();
  return (
    <div className={styles.container}>
      <Nav />
      <section className={styles.main}>
        <input type="checkbox" className={styles.exInput} id="expand" />
        <label className={styles.expand} htmlFor="expand">
          <span />
          <span />
          <span />
        </label>
        <nav className={styles.sidebar}>
          <div className={styles.smolAccountDiv}>{user && <UserInfo user={user} />}</div>
          <NavLink exact activeClassName={styles.active} to="/profile">
            Profile
          </NavLink>
          <NavLink activeClassName={styles.active} to="/profile/clients">
            OAuth Clients
          </NavLink>
          {user?.role === "admin" && (
            <>
              <hr />
              <NavLink activeClassName={styles.active} to="/admin/users">
                All Users
              </NavLink>
              <NavLink activeClassName={styles.active} to="/admin/clients">
                All OAuth Clients
              </NavLink>
              <NavLink activeClassName={styles.active} to="/admin/courses">
                Courses
              </NavLink>
            </>
          )}
          <div className={styles.mobileUser}>
            <hr />
            <Link to="/logout">Logout</Link>
          </div>
        </nav>
        <div className={styles.content}>{children}</div>
      </section>
    </div>
  );
};

export default Layout;
