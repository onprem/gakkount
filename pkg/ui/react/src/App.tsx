import React, { useState, useEffect } from "react";
import { Switch, Route } from "react-router-dom";
import cookie from "js-cookie";
import { AuthContext } from "./contexts/auth";
import { OAuthLogin, Consent } from "./pages";
import "./App.css";
import { User } from "./interfaces";

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
    }
  }, [isLoggedIn, setIsLoggedIn, user, setUser]);
  return (
    <AuthContext.Provider value={{ isLoggedIn, setIsLoggedIn, user, setUser }}>
      <div className="App">
        <Switch>
          <Route exact path="/">
            <header className="App-header">
              <code>IIITM Accounts</code>
            </header>
          </Route>
          <Route exact path="/oauth/login">
            <OAuthLogin />
          </Route>
          <Route exact path="/oauth/consent">
            <Consent />
          </Route>
        </Switch>
      </div>
    </AuthContext.Provider>
  );
}

export default App;
