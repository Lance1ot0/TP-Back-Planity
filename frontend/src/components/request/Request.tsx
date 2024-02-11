import { useState } from 'react';
import style from './Request.module.css';
import { sendRequest, getRequest } from '../../api/request/requestApi';
import { useEffect } from 'react';

export default function Request() {
    const [request, setRequest] = useState<any>(null);

    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        const data = new FormData(e.currentTarget);
        const salonName = data.get('salonName');
        const address = data.get('address');
        const city = data.get('city');
        const postalCode = data.get('postalCode');

       const request = {
            salonName,
            address,
            city,
            postalCode
        };
        
        const res = await sendRequest(request);
        console.log("res: ", res);

        if (res.success) {
            fetchRequest();
        }
    };

    const fetchRequest = async () => {
        const response = await getRequest();
        if (response.success) {
            console.log("response req: ", response);
            setRequest(response.res);
        }
    };
    useEffect(() => {
        fetchRequest();
    }, []);
    return (
        <div className={style.container}>
            {!request ? (
                <div>
                    <h2>Vous n'avez pas encore de salon de coiffure</h2>
                    <form onSubmit={handleSubmit}>
                        <label htmlFor="salonName">Nom du salon</label>
                        <input type="text" id="salonName" name="salonName"/>

                        <label htmlFor="address">Adresse</label>
                        <input type="text" id="address" name="address"/>

                        <label htmlFor="city">Ville</label>
                        <input type="text" id="city" name="city"/>

                        <label htmlFor="postalCode">Code postal</label>
                        <input type="text" id="postalCode" name="postalCode"/>

                        <button type="submit">Envoyer la demande</button>
                    </form>
                </div>
            ) : (
                <div>
                    <h2>Demande en cours</h2>
                    <p><span>Nom du salon :</span> {request.salonName}</p>
                    <p><span>Adresse :</span> {request.address}, {request.postalCode} {request.city}</p>
                    <p><span>Date d'émission:</span> {request.requestDate}</p>
                    <p><span>Status:</span> {request.requestStatus}</p>
                </div>
            )}  
            
        </div>
    );
}