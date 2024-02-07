import "./App.css";
import useSWR from "swr";
import { useState } from "react";

export const ENDPOINT = "http://localhost:8081/api";

export interface Salon {
  requestID: number;
  professionalID: number;
  salonName: string;
  address: string;
  city: string;
  postalCode: string;
  requestDate: string;
  requestStatus: string;
}

const getAllRequests = (url: string) =>
  fetch(`${ENDPOINT}/${url}`, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  }).then((res) => res.json());

async function createSalonRequest(
  values: {
    professionalID: number;
    salonName: string;
    address: string;
    city: string;
    postalCode: string;
  },
  mutate: () => void
) {
  const updated = await fetch(`${ENDPOINT}/professional/request`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(values),
  }).then((r) => r.json());

  if (updated) {
    await mutate();
  }
}

function App() {
  const { data, mutate } = useSWR("admin/requests", getAllRequests);
  const [newSalonName, setNewSalonName] = useState("");

  return (
    <>
      <h1>TP PLANITY PROJECT</h1>

      {data && (
        <div>
          <h2>Request for admin:</h2>
          <ul>
            {data.map((salon: Salon) => (
              <li key={salon.requestID}>
                {salon.salonName} : {salon.requestStatus}
              </li>
            ))}
          </ul>
        </div>
      )}

      <input
        type="text"
        value={newSalonName}
        onChange={(e) => setNewSalonName(e.target.value)}
      />

      <button
        onClick={() => {
          createSalonRequest(
            {
              professionalID: 1,
              salonName: newSalonName,
              address: "1 Rue du zboub",
              city: "Paris",
              postalCode: "75635",
            },
            mutate
          );
        }}
      >
        Create Request
      </button>
    </>
  );
}

export default App;
