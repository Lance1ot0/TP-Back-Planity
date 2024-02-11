import "./Navbar.css";

import React from "react";
import { Link } from "react-router-dom";

function Navbar() {
  return (
    <nav className="navbar">
      <Link to="/">Acceuil</Link>
      <div className="right-wrapper">
        <Link to="/register">Sign In</Link>
        <Link to="/login">Login</Link>
      </div>
    </nav>
  );
}

export default Navbar;
