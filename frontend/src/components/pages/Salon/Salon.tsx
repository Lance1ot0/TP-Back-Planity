import useSWR from "swr";
import { getSalonInfos } from "../../../api/client/getSalonInfos";
import "./Salon.css";
import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";

function Salon() {
  const { id } = useParams();

  const [loadedData, setLoadedData] = useState<any>();

  const { data, mutate } = useSWR(`client/hairSalon/${id}`, getSalonInfos);

  useEffect(() => {
    setLoadedData(data);
  }, [data]);

  console.log(loadedData);

  return (
    <div className="salon-infos-resa-container">
      <div className="salon-infos-resa-wrapper">
        <div className="salon-infos">
          {loadedData && (
            <>
              <div className="salon-title">{loadedData.salon.name}</div>
              <div className="salon-title">{loadedData.salon.address}</div>
            </>
          )}
        </div>
        <div className="salon-resas">
          <div className="wrapper">
            <h3>Services</h3>
            {loadedData &&
              loadedData.services.map((service: any) => (
                <div className="card">
                  <input type="radio" id={service.serviceID} name="service" />
                  <label htmlFor={service.serviceID}>
                    <div>{service.description}</div>
                    <div>{service.price}$</div>
                    <div>{service.duration}</div>
                  </label>
                </div>
              ))}
          </div>

          <div className="wrapper">
            <h3>Employés</h3>
            {loadedData &&
              loadedData.employees.map((employee: any) => (
                <div className="card">
                  <input type="radio" id={employee.employeeID} name="service" />
                  <label htmlFor={employee.serviceID}>
                    <div>{employee.firstname}</div>
                    <div>{employee.lastname}</div>
                  </label>
                </div>
              ))}
          </div>

          <div className="wrapper hours">
            <h3>Horaires Disponibles</h3>
            <table>
              <thead>
                <tr>
                  <th>Lundi</th>
                  <th>Mardi</th>
                  <th>Mercredi</th>
                  <th>Jeudi</th>
                  <th>Vendredi</th>
                  <th>Samedi</th>
                  <th>Dimanche</th>
                </tr>
              </thead>
              <tbody>
                <tr>
                  {loadedData &&
                    loadedData.employees[0].Availabilities.map((hours: any) => (
                      <div className="card">
                        <td>
                          <input
                            type="radio"
                            id={hours.employeeID}
                            name="hours"
                          />
                          {/* <label htmlFor={employee.serviceID}>
                            <div>{employee.firstname}</div>
                            <div>{employee.lastname}</div>
                          </label> */}
                        </td>
                      </div>
                    ))}
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </div>
  );
}

export default Salon;
