import React from "react";
import { BrowserRouter, Route, Switch, Redirect } from "react-router-dom";

import { isAuthenticated } from "./services/auth";

import Main from "./components/layout/Main";
import SignIn from "./pages/SignIn";
import HomePage from "./pages/Home";
import SimplesN from "./pages/SimplesN";
import Faturamento from "./pages/Faturamento";

const PrivateRoute = ({ component: Component, ...rest }) => (
  <Route
    {...rest}
    render={props =>
      isAuthenticated() ? (
        <Component {...props} />
      ) : (
        <Redirect to={{ pathname: "/", state: { from: props.location } }} />
      )
    }
  />
);

const Routes = () => (
  <BrowserRouter>
    <Switch>
      <Route exact path="/" component={SignIn} />
      <PrivateRoute path="/app" component={() => <h1>App</h1>} />
      <Main>
        <PrivateRoute exact path="/home" component={HomePage} />
        <PrivateRoute exact path="/tabelaSN" component={SimplesN} />
        <PrivateRoute exact path="/faturamentos" component={Faturamento} />
      </Main>
      <Route path="*" component={() => <h1>Page not found</h1>} />
    </Switch>
  </BrowserRouter>
);

export default Routes;