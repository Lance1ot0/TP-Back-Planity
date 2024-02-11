import Service from '../Service/Service';
import Employee from '../employee/Employee';
import style from './HairSalon.module.css';
import { useState } from 'react';

interface HairSalonProps {
    name: string;
    address: string;
    city: string;
    postalCode: string;
    hairSalonID: number;
}

export default function HairSalon(props: HairSalonProps) {
    const [isEmployeeOpen, setIsEmployeeOpen] = useState<boolean>(false);
    const [isServiceOpen, setIsServiceOpen] = useState<boolean>(false);

    return (
        <div className={style.container}>
            <div className={style.top}>
                <div className={style.topLeft}>
                    <h2>{props.name}</h2>
                    <p>{props.address}, {props.postalCode} {props.city}</p>
                    <p>id: {props.hairSalonID}</p>
                </div>
                <div className={style.topRight}>
                    <button onClick={() => {setIsEmployeeOpen(!isEmployeeOpen); setIsServiceOpen(false)}}>Employe</button>
                    <button onClick={() => {setIsServiceOpen(!isServiceOpen); setIsEmployeeOpen(false)}}>Service</button>
                </div>
            </div>

            <div className={style.main}>
                {isEmployeeOpen && (
                    <Employee hairSalonID={props.hairSalonID} />
                )}

                {isServiceOpen && (
                    <Service hairSalonID={props.hairSalonID} />
                )}
            </div>
        </div>
    );
}