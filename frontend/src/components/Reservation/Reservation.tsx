import React, { useEffect, useState } from 'react'
import { getClientReservations } from '../../api/ReservationApi/ReservationApi';

function Reservation() {
    const [reservations, setReservations] = useState<any[]>([]);
    const fetchReservations = async () => {
        const res = await getClientReservations();

        if (res.success) {
            setReservations(res.res.data);
        }
    }

    useEffect(() => {
        fetchReservations();
    }, []);
  return (
    <div>
        <h1>Reservation</h1>
        <div>
            <ul>
                {reservations.map((reservation, index) => (
                    <li key={index}>
                        <p>ClientId : {reservation.clientID}</p>
                        <p>EmployeeID : {reservation.employeeID}</p>
                        <p>hairSalonID : {reservation.hairSalonID}</p>
                        <p>reservationDate : {reservation.reservationDate}</p>
                        <p>reservationID : {reservation.reservationID}</p>
                        <p>reservationStatus : {reservation.reservationStatus}</p>
                        <p>serviceID : {reservation.serviceID}</p>
                    </li>
                ))}
            </ul>
        </div>
    </div>
  )
}

export default Reservation
