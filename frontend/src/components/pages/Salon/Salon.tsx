import useSWR from "swr";
import { getSalonInfos } from "../../../api/client/getSalonInfos";
import "./Salon.css";
import React from "react";
import { useParams } from "react-router-dom";

function Salon() {
  const { id } = useParams();

  const { data, mutate } = useSWR(`client/hairSalon/${id}`, getSalonInfos);

  console.log(data);

  return <div>Salon</div>;
}

export default Salon;
