import "./LoginSelection.css";

import React from "react";
import { Link, Outlet } from "react-router-dom";

function LoginSelection() {
  return (
    <>
      <h1>Login Selection</h1>
      <nav className="account-selection-type-container">
        <Link className="link" to="/login/pro">
          Pro Account
        </Link>
        <Link className="link" to="/login/client">
          Client Account
        </Link>
      </nav>
      <Outlet />
    </>
  );
}

export default LoginSelection;
