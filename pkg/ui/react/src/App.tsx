import React, { useState, useEffect } from "react";
import { Switch, Route, Redirect } from "react-router-dom";
import cookie from "js-cookie";
import { SWRConfig } from "swr";

import { AuthContext } from "./contexts/auth";
import fetcher from "./utils/fetcher";
import { User } from "./interfaces";
import ProtectedRoute from "./components/protectedRoute";

import {
  OAuthLogin,
  Consent,
  Profile,
  Login,
  Users,
  Logout,
  Clients,
  Courses,
  Setup,
  Departments,
} from "./pages";

import "./App.css";

function App() {
  const [isLoggedIn, setIsLoggedIn] = useState(cookie.get("signedin") === "true");
  const [user, setUser] = useState<User>();

  useEffect(() => {
    if (isLoggedIn && !user) {
      fetch("/api/user")
        .then<{ status: string; user: User }>((res) => res.json())
        .then((res) => {
          if (res.status === "success") setUser(res.user);
          else setIsLoggedIn(false);
        })
        .catch(() => setIsLoggedIn(false));
    } else if (!isLoggedIn && user) {
      setUser(undefined);
    }
  }, [isLoggedIn, setIsLoggedIn, user, setUser]);
  return (
    <AuthContext.Provider value={{ isLoggedIn, setIsLoggedIn, user, setUser }}>
      <SWRConfig
        value={{
          fetcher,
          revalidateOnFocus: false,
        }}
      >
        <div className="App">
          {/* <Nav /> */}
          <Switch>
            <Route exact path="/">
              <Redirect to="/login" />
            </Route>
            <Route exact path="/oauth/login">
              <OAuthLogin />
            </Route>
            <Route exact path="/oauth/consent">
              <Consent />
            </Route>
            <Route path="/login">
              <Login />
            </Route>
            <ProtectedRoute path="/dashboard">
              <header className="App-header">
                <code>Hello, {user?.name} </code>
              </header>
            </ProtectedRoute>
            <ProtectedRoute exact path="/profile">
              <Profile />
            </ProtectedRoute>
            <ProtectedRoute exact path="/profile/clients">
              <Clients />
            </ProtectedRoute>
            <ProtectedRoute path="/admin/users">
              <Users />
            </ProtectedRoute>
            <ProtectedRoute exact path="/admin/clients">
              <Clients all />
            </ProtectedRoute>
            <ProtectedRoute exact path="/admin/courses">
              <Courses />
            </ProtectedRoute>
            <ProtectedRoute exact path="/admin/departments">
              <Departments />
            </ProtectedRoute>
            <Route path="/setup">
              <Setup />
            </Route>
            <Route path="/logout">
              <Logout />
            </Route>
          </Switch>
        </div>
      </SWRConfig>
    </AuthContext.Provider>
  );
}

export default App;
