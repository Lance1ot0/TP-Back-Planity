import "./RegisterSelection.css";

import React from "react";
import { Link, Outlet } from "react-router-dom";

function RegisterSelection() {
  return (
    <section className="register-selection-container">
      <h1>Sign in as a Pro or a Client ?</h1>
      <nav className="account-selection-type-container">
        <Link className="link" to="/professional/register">
          Pro
        </Link>
        <Link className="link" to="/client/register">
          Client
        </Link>
      </nav>
      <Outlet />
    </section>
  );
}

export default RegisterSelection;
