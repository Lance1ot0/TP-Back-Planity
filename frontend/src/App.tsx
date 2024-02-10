import "./App.css";
import useSWR from "swr";
import { useState } from "react";

export const API_URL = "http://localhost:8081/api";

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
  fetch(`${API_URL}/${url}`, {
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
  const created = await fetch(`${API_URL}/professional/request`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(values),
  }).then((r) => r.json());

  if (created) {
    await mutate();
  }
}

async function updateRequest(
  id: number,
  RequestStatus: { RequestStatus: string },
  mutate: () => void
) {
  const updated = await fetch(`${API_URL}/admin/request/${id}`, {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(RequestStatus),
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
              <li key={salon.requestID} className="li">
                {salon.salonName} : {salon.requestStatus}
                <div className="button-wrapper">
                  {" "}
                  <button
                    onClick={() => {
                      updateRequest(
                        salon.requestID,
                        {
                          RequestStatus: "accepted",
                        },
                        mutate
                      );
                    }}
                  >
                    Accept
                  </button>
                  <button
                    onClick={() => {
                      updateRequest(
                        salon.requestID,
                        {
                          RequestStatus: "rejected",
                        },
                        mutate
                      );
                    }}
                  >
                    Decline
                  </button>
                </div>
              </li>
            ))}{" "}
          </ul>
        </div>
      )}

      <h2>Request for Pro:</h2>

      <span>
        Salon Name:
        <input
          type="text"
          value={newSalonName}
          onChange={(e) => setNewSalonName(e.target.value)}
        />
      </span>

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
