import "./Homepage.css";
import React, { useState } from "react";
import { researchSalonByName } from "../../../api/client/searchSalonByName";

export interface Salon {
  hairSalonID: number;
  professionalID: number;
  name: string;
  address: string;
  city: string;
  postalCode: string;
  requestDate: string;
  requestStatus: string;
}

function Homepage() {
  const [errorSearchMessage, setErrorSearchMessage] = useState("");
  const [allSearchedSalons, setAllSearchedSalons] = useState();

  const submitSearchSalon = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const searchSalon = (e.target as HTMLFormElement).searchSalon.value;

    if (!searchSalon) {
      return;
    }

    const res = await researchSalonByName(searchSalon);
    console.log(res);
    if (res) {
      console.log(res);
      setAllSearchedSalons(res);
    }
    if (res == null) setErrorSearchMessage("No salon found");
  };

  return (
    <div className="search">
      <form onSubmit={submitSearchSalon}>
        <input type="text" id="searchSalon" name="search-salon" />
        <button type="submit">Search</button>
        <>{errorSearchMessage && <div>{errorSearchMessage}</div>}</>
      </form>

      <div className="search-result-container">
        {allSearchedSalons &&
          allSearchedSalons.map((salon: Salon) => (
            <a
              className="salon-card"
              href={"/client/hairsalon/" + salon.hairSalonID}
            >
              <div className="salon-title">{salon.name}</div>
              <div className="salon-address">
                {salon.address} {salon.city} {salon.postalCode}
              </div>
            </a>
          ))}
      </div>
    </div>
  );
}

export default Homepage;
