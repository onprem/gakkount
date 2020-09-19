import React from "react";
import { Route, RouteProps, Redirect, useLocation } from "react-router-dom";
import { useAuth } from "../../contexts/auth";

export const ProtectedRoute: React.FC<RouteProps> = ({ children, ...rest }) => {
  const { isLoggedIn } = useAuth();
  const loc = useLocation();
  return (
    <Route {...rest}>
      {isLoggedIn ? children : <Redirect to={{ pathname: "/login", state: { referrer: loc } }} />}
    </Route>
  );
};

export default ProtectedRoute;
