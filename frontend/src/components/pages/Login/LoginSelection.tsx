import "./LoginSelection.css";

import React from "react";
import { Link, Outlet } from "react-router-dom";

function LoginSelection() {
  return (
    <section className="login-selection-container">
      <h1>What account do you have ?</h1>
      <nav className="account-selection-type-container">
        <Link className="link" to="/professional/login">
          Pro
        </Link>
        <Link className="link" to="/client/login">
          Client
        </Link>
        <Link className="link" to="/admin/login">
          Admin
        </Link>
      </nav>
      <Outlet />
    </section>
  );
}

export default LoginSelection;
