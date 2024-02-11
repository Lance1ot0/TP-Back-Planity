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
    const [isEmployeeOpen, setIsEmployeeOpen] = useState<any>(false);
    console.log("isEmployeeOpen: ", isEmployeeOpen);
    return (
        <div className={style.container}>
            <div className={style.top}>
                <div className={style.topLeft}>
                    <h2>{props.name}</h2>
                    <p>{props.address}, {props.postalCode} {props.city}</p>
                    <p>id: {props.hairSalonID}</p>
                </div>
                <div className={style.topRight}>
                    <button onClick={() => setIsEmployeeOpen(!isEmployeeOpen)}>Employe</button>
                    <button>Service</button>
                </div>
            </div>

            <div className={style.main}>
                {isEmployeeOpen && (
                    <Employee hairSalonID={props.hairSalonID} />
                )}
            </div>
        </div>
    );
}