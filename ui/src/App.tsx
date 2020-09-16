import React from "react";
import { Login, Consent } from "./pages";
import "./App.css";
import { Switch, Route } from "react-router-dom";

function App() {
  return (
    <div className="App">
      <Switch>
        <Route exact path="/">
          <header className="App-header">
            <code>IIITM Accounts</code>
          </header>
        </Route>
        <Route exact path="/oauth/login">
          <Login />
        </Route>
        <Route exact path="/oauth/consent">
          <Consent />
        </Route>
      </Switch>
    </div>
  );
}

export default App;
