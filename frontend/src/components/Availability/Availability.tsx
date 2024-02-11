import useSWR from 'swr';
import style from './Availability.module.css';
import { getAvailability, addAvailability } from '../../api/AvailabilityApi/AvailabilityApi';
import React from 'react';

function Availability(props: any) {
    const { data, mutate } = useSWR(`professional/employee/availability/${props.EmployeeID}`, getAvailability)

    function formatDate(time: string): string {
        const [hours, minutes, _] = time.split(':');
        return `${hours}:${minutes}`;
    }

    const onHandleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        const dayOfWeek = (e.target as HTMLFormElement).Day.value;
        const startTime = (e.target as HTMLFormElement).startTime.value;
        const endTime = (e.target as HTMLFormElement).endTime.value;

        const item = {
            dayOfWeek,
            startTime,
            endTime,
            intervalTime: 10
        }

        const res = await addAvailability(props.EmployeeID, item);
        if (res.success) {
            mutate();
        }
    }

    return (
        <div className={style.container}>
            <table>
                <thead>
                    <tr>
                        <th>Monday</th>
                        <th>Tuesday</th>
                        <th>Wednesday</th>
                        <th>Thursday</th>
                        <th>Friday</th>
                        <th>Saturday</th>
                        <th>Sunday</th>
                    </tr>
                </thead>
                <tbody>
                    <tr>
                        {data && data.map((day: any) => (
                            <React.Fragment key={day.availabilityID}>
                                <td>
                                    <span>{formatDate(day.startTime)} </span>
                                    <span>{formatDate(day.endTime)}</span>
                                </td>
                            </React.Fragment>
                        ))}
                    </tr>
                </tbody>
            </table>

            <form onSubmit={onHandleSubmit}>
                <div>
                    <label htmlFor="Day">Jour</label>
                    <select id="Day" name="Day">
                        <option value="Monday">Monday</option>
                        <option value="Tuesday">Tuesday</option>
                        <option value="Wednesday">Wednesday</option>
                        <option value="Thursday">Thursday</option>
                        <option value="Friday">Friday</option>
                        <option value="Saturday">Saturday</option>
                        <option value="Sunday">Sunday</option>
                    </select>
                </div>
                
                <div className={style.timeContainer}>
                    <label htmlFor="startTime">Début</label>
                    <input className={style.inputTime} type="time" id="startTime" name="startTime"/>

                    <label htmlFor="endTime">Fin</label>
                    <input className={style.inputTime} type="time" id="endTime" name="endTime"/>
                </div>

                <button type="submit">Ajouter</button>
            </form>
        </div>
    )
}

export default Availability
