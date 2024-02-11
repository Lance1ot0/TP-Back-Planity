import Service from '../Service/Service';
import Employee from '../employee/Employee';
import style from './HairSalon.module.css';
import { useEffect, useState } from 'react';
import { getReservations } from '../../api/ReservationApi/ReservationApi';

interface HairSalonProps {
    name: string;
    address: string;
    city: string;
    postalCode: string;
    hairSalonID: number;
}

export default function HairSalon(props: HairSalonProps) {
    const [isRDVOpen, setIsRDVOpen] = useState<boolean>(true);
    const [isEmployeeOpen, setIsEmployeeOpen] = useState<boolean>(false);
    const [isServiceOpen, setIsServiceOpen] = useState<boolean>(false);
    const [reservations, setReservations] = useState<any>([]);

    const fetchReservations = async () => {
        const resReservation = await getReservations(props.hairSalonID);
        if (resReservation.success) {
            const filteredReservations = resReservation.res.map((reservation: any) => {
                const {
                    employeeFirstname,
                    employeeLastname,
                    clientFirstname,
                    clientLastname,
                    serviceName,
                    reservationDate,
                    reservationStatus,
                } = reservation;
                return {
                    employeeFirstname,
                    employeeLastname,
                    clientFirstname,
                    clientLastname,
                    serviceName,
                    reservationDate,
                    reservationStatus,
                };
            });
       
            setReservations(filteredReservations);
        }
    };

    useEffect(() => {
        fetchReservations();
    }, []);

    const handleRDVClick = () => {
        setIsRDVOpen(true);
        setIsEmployeeOpen(false);
        setIsServiceOpen(false);
    };

    const handleEmployeeClick = () => {
        setIsRDVOpen(false);
        setIsEmployeeOpen(true);
        setIsServiceOpen(false);
    };

    const handleServiceClick = () => {
        setIsRDVOpen(false);
        setIsEmployeeOpen(false);
        setIsServiceOpen(true);
    };

    return (
        <div className={style.container}>
            <div className={style.top}>
                <div className={style.topLeft}>
                    <h2>{props.name}</h2>
                    <p>{props.address}, {props.postalCode} {props.city}</p>
                </div>
                <div className={style.topRight}>
                    <button onClick={handleRDVClick}>RDV</button>
                    <button onClick={handleEmployeeClick}>Employe</button>
                    <button onClick={handleServiceClick}>Service</button>
                </div>
            </div>

            <div className={style.main}>
                {isRDVOpen && (
                    <div className={style.rdv}>
                        <h1>RDV</h1>
                        <table>
                            <thead>
                                <tr>
                                    <th>Employe</th>
                                    <th>Client</th>
                                    <th>Service</th>
                                    <th>Date</th>
                                    <th>Status</th>
                                </tr>
                            </thead>
                            <tbody>
                                {reservations.map((reservation: any, index: number) => (
                                    <tr key={index}>
                                        <td><span>{reservation.employeeFirstname} {reservation.employeeLastname}</span></td>
                                        <td><span>{reservation.clientFirstname} {reservation.clientLastname}</span></td>
                                        <td>{reservation.serviceName}</td>
                                        <td>{reservation.reservationDate}</td>
                                        <td>{reservation.reservationStatus}</td>
                                    </tr>
                                ))}
                            </tbody>
                        </table>
                    </div>
                )}

                {isEmployeeOpen && (
                    <div className={style.employee}>
                        <Employee hairSalonID={props.hairSalonID} />
                    </div>
                )}

                {isServiceOpen && (
                    <div className={style.service}>
                        <Service hairSalonID={props.hairSalonID} />
                    </div>
                )}
            </div>
        </div>
    );
}