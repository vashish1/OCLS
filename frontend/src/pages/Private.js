import React from "react";
import { Route } from "react-router-dom";
import Dashboard from "./Dashboard";
export default function PrivateRoute({ component: Component, ...rest }) {
  const { isLoggedIn } = useLogin();

  return (
    <Route
      {...rest}
      render={(props) => {
        return isLoggedIn ? (
          <Component {...props} />
        ) : (
          <Route exact path="/" component={Dashboard} />
        );
      }}
    ></Route>
  );
}