import React from "react";
import { useState } from "react";

export const API_URL = "http://localhost:8081/api";

export interface Client {
  email: string;
  password: string;
}

async function login(values: { email: string; password: string }) {
  const login = await fetch(`${API_URL}/client/login`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(values),
  }).then((r) => r.json());

  if (login) {
    //
  }
}

function LoginClient() {
  const [inputs, setInputs] = useState({});

  console.log(inputs);

  const handleChange = (event: any) => {
    const name = event.target.name;
    const value = event.target.value;
    setInputs((values) => ({ ...values, [name]: value }));
  };

  const handleSubmit = (event: any) => {
    event.preventDefault();

    login({
      email: "client@planity.com",
      password: "password",
    });
  };

  return (
    <>
      <h1>Login Client</h1>
      <form onSubmit={handleSubmit}>
        <label>
          Enter your email:
          <input
            type="email"
            name="email"
            value={inputs.email || ""}
            onChange={handleChange}
          />
        </label>
        <label>
          Enter your passsword:
          <input
            type="password"
            name="password"
            value={inputs.password || ""}
            onChange={handleChange}
          />
        </label>
        <input type="submit" />
      </form>
    </>
  );
}

export default LoginClient;
