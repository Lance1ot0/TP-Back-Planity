import { useState } from "react";
import { getEmployee, addEmployee } from "../../api/employee/employeeApi";
import useSWR from "swr";
import style from './Employee.module.css';
import Availability from "../Availability/Availability";

interface EmployeeProps {
    employeeID?: number;
    firstname?: string;
    lastname?: string;
    hairSalonID: number;
}

export default function Employee(props: EmployeeProps) {
    const { data, mutate } = useSWR(`professional/employee/${props.hairSalonID}`, getEmployee)
    const [isAvailabilityOpen, setIsAvailabilityOpen] = useState<boolean>(false);
    const [employeeID, setEmployeeID] = useState<number | null>(null);

    const onHandleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        const firstname = (e.target as HTMLFormElement).firstname.value;
        const lastname = (e.target as HTMLFormElement).lastname.value;
        
        const res = await addEmployee(firstname, lastname, props.hairSalonID);
        if (res.success) {
            mutate();
        }
    }

    const onHandleAvailability = (id: number) => {
        setEmployeeID(id === employeeID ? null : id);
        setIsAvailabilityOpen(id !== employeeID);       
    }
    

    return (
        <div className={style.container}>
            <h1>Employe</h1>

            <form onSubmit={onHandleSubmit}>
                <label htmlFor="firstname">firstname</label>
                <input type="text" id="firstname" name="firstname"/>

                <label htmlFor="lastname">lastname</label>
                <input type="text" id="lastname" name="lastname"/>

                <button type="submit">Ajouter</button>
            </form>

            <div className={style.main}>
                <h2>Liste des employes</h2>
                <div className={style.employeeContainer}>
                    <table>
                        <thead>
                            <tr>
                                <th>Firstname</th>
                                <th>Lastname</th>
                                <th></th>
                            </tr>
                        </thead>
                        <tbody>
                            {data && data.map((emp: any) => (
                                <tr key={emp.employeeID}>
                                    <td>{emp.firstname}</td>
                                    <td>{emp.lastname}</td>
                                    <td>
                                        <button onClick={() => onHandleAvailability(emp.employeeID)}>Horaire</button>
                                    </td>
                                </tr>
                            ))}
                        </tbody>
                    </table>

                    {isAvailabilityOpen && (
                        <div>
                            <Availability EmployeeID={employeeID}/>
                        </div>
                    )}
                </div>
            </div>
        </div>
    );
}