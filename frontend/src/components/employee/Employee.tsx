// import { useEffect, useState } from "react";
import { getEmployee, addEmployee } from "../../api/employee/employeeApi";
import useSWR from "swr";
import style from './Employee.module.css';

interface EmployeeProps {
    employeeID?: number;
    firstname?: string;
    lastname?: string;
    hairSalonID: number;
}

export default function Employee(props: EmployeeProps) {
    const { data, mutate } = useSWR(`professional/employee/${props.hairSalonID}`, getEmployee)

    const onHandleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        const firstname = (e.target as HTMLFormElement).firstname.value;
        const lastname = (e.target as HTMLFormElement).lastname.value;
        
        const res = await addEmployee(firstname, lastname, props.hairSalonID);
        if (res.success) {
            mutate();
        }
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
                                    <button>Horaire</button>
                                </td>
                            </tr>
                        ))}
                    </tbody>
                </table>
            </div>
        </div>
    );
}