import { useEffect, useState } from 'react';
import useSWR from "swr";
import { useNavigate } from 'react-router-dom';
import style from './AdminPage.module.css';
import { FaCheck, FaTimes } from 'react-icons/fa';
import { getAllRequests, updateRequest } from '../../../api/request/requestApi';

export interface Salon {
    requestID: number;
    professionalID: number;
    salonName: string;
    address: string;
    city: string;
    postalCode: string;
    requestDate: string;
    requestStatus: string;
  }



export default function AdminPage() {
    const navigate = useNavigate();
    const [token, setToken] = useState<string | null>(null);
    const { data, mutate } = useSWR("admin/requests", getAllRequests);

    console.log(data);

    const logout = () => {
        localStorage.removeItem('token');
        setToken(null);
        navigate('/admin/login');
    };

    useEffect(() => {
        const storedToken = localStorage.getItem('token');
        if (storedToken) {
            setToken(storedToken);
        } else {
            navigate('/admin/login'); 
        }
    }, [navigate]);

    if (!token) {
        return null;
    }

    return (
        <div className={style.adminContainer}>
            <div className={style.adminMenu}>
                <h1 style={{ fontSize: '38px' }}>Admin Page</h1>
                <button
                    className={style.logoutButton}
                    onClick={() => { logout(); }}
                >Déconnexion</button>
            </div>

            <div className={style.adminMain}>
                <div>
                    <h2 style={{marginBottom: '20px', fontSize: '32px'}}>Demandes en cours</h2>
                    <table className={style.adminTable}>
                        <thead>
                            <tr>
                                <th>Nom de l'établissement</th>
                                <th>Adresse</th>
                                <th>Ville</th>
                                <th>Code postal</th>
                                <th>Date de demande</th>
                            </tr>
                        </thead>
                        <tbody>
                            {data && (
                                data.filter((salon: Salon) => salon.requestStatus === 'pending').map((salon: Salon) => (
                                    <tr key={salon.requestID}  className={style.requestList}>
                                        <td>{salon.salonName}</td>
                                        <td>{salon.address}</td>
                                        <td>{salon.city}</td>
                                        <td>{salon.postalCode}</td>
                                        <td>{salon.requestDate}</td>
                                        <div style={{display: 'flex', gap: '10px'}}>
                                            <button style={{backgroundColor: '#a7c957'}} onClick={() => {updateRequest(salon.requestID, {RequestStatus: "accepted"}, mutate)}}><FaCheck /></button>
                                            <button style={{backgroundColor: '#e63946'}} onClick={() => {updateRequest(salon.requestID, {RequestStatus: "rejected"}, mutate)}}><FaTimes /></button>
                                        </div>
                                    </tr>
                                    
                                ))
                            )}
                        </tbody>
                    </table>
                </div>

                <div style={{marginTop: '40px'}}>
                    <h2 style={{marginBottom: '20px', fontSize: '32px', color: '#a7c957'}}>Demandes accepté</h2>
                    <table className={style.adminTable}>
                        <thead>
                            <tr>
                                <th>Nom de l'établissement</th>
                                <th>Adresse</th>
                                <th>Ville</th>
                                <th>Code postal</th>
                                <th>Date de demande</th>
                            </tr>
                        </thead>
                        <tbody>
                            {data && (
                                data.filter((salon: Salon) => salon.requestStatus === 'accepted').map((salon: Salon) => (
                                    <tr key={salon.requestID}  className={style.requestList}>
                                        <td>{salon.salonName}</td>
                                        <td>{salon.address}</td>
                                        <td>{salon.city}</td>
                                        <td>{salon.postalCode}</td>
                                        <td>{salon.requestDate}</td>
                                    </tr>
                                    
                                ))
                            )}
                        </tbody>
                    </table>
                </div>

                <div style={{marginTop: '40px'}}>
                    <h2 style={{marginBottom: '20px', fontSize: '32px', color: '#e63946'}}>Demandes rejeté</h2>
                    <table className={style.adminTable}>
                        <thead>
                            <tr>
                                <th>Nom de l'établissement</th>
                                <th>Adresse</th>
                                <th>Ville</th>
                                <th>Code postal</th>
                                <th>Date de demande</th>
                            </tr>
                        </thead>
                        <tbody>
                            {data && (
                                data.filter((salon: Salon) => salon.requestStatus === 'rejected').map((salon: Salon) => (
                                    <tr key={salon.requestID}  className={style.requestList}>
                                        <td>{salon.salonName}</td>
                                        <td>{salon.address}</td>
                                        <td>{salon.city}</td>
                                        <td>{salon.postalCode}</td>
                                        <td>{salon.requestDate}</td>
                                    </tr>
                                    
                                ))
                            )}
                        </tbody>
                    </table>
                </div>
            </div>
        </div>
    );
}

