import React from 'react'
import useSWR from 'swr'
import { getServices, addService } from '../../api/ServiceApi/ServiceApi'
import style from './Service.module.css';

function Service(props: any) {
    const { data, mutate } = useSWR(`professional/service/${props.hairSalonID}`, getServices)

    const formatDuration = (duration: number) => {
        if (duration > 59) {
            const hours = Math.floor(duration / 60);
            const minutes = duration % 60;
            return `${hours}h ${minutes}min`;
        } else {
            return `${duration}min`;
        }
    }

    const onHandleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        const name = (e.target as HTMLFormElement).serviceName.value;
        const description = (e.target as HTMLFormElement).description.value;
        const price = parseFloat((e.target as HTMLFormElement).price.value);
        const duration = parseInt((e.target as HTMLFormElement).duration.value);

        const item = {
            name,
            description,
            price,
            duration,
            hairSalonID: props.hairSalonID
        }

        const res = await addService(item);
        if (res.success) {
            mutate();
        }
    }

    return (
        <div className={style.container}>
            <h1>Service</h1>

            <form onSubmit={onHandleSubmit}>
                <label htmlFor="serviceName">Nom du service</label>
                <input type="text" id="serviceName" name="serviceName"/>

                <label htmlFor="description">Description</label>
                <input type="text" id="description" name="description"/>

                <label htmlFor="price">Prix</label>
                <input type="number" id="price" name="price"/>

                <label htmlFor="duration">Durée</label>
                <input type="number" id="duration" name="duration"/>

                <button type="submit">Ajouter</button>
            </form>

            <h2>Liste des services</h2>
            {data && (
                <table>
                    <thead>
                        <tr>
                            <th>Nom du service</th>
                            <th>Description</th>
                            <th>Prix (euros)</th>
                            <th>Durée</th>
                        </tr>
                    </thead>
                    <tbody>
                        {data.map((service: any) => (
                            <tr key={service.id}>
                                <td>{service.name}</td>
                                <td>{service.description}</td>
                                <td>{service.price}</td>
                                <td>{formatDuration(service.duration)}</td>
                            </tr>
                        ))}
                    </tbody>
                </table>
            )}
        </div>
    )
}

export default Service
