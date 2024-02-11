import { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import style from './ProfessionalPage.module.css';
import { getHairSalon } from '../../../api/request/requestApi';
import Request from '../../request/Request';
import HairSalon from '../../hairSalon/HairSalon';

export default function ProfessionalPage() {
    const navigate = useNavigate();
    const [role, setRole] = useState<string | null>(null);
    const [token, setToken] = useState<string | null>(null);
    const [hairSalon, setHairSalon] = useState<any>(null);

    const fetchHairSalon = async () => {
        const response = await getHairSalon();
        if (response.success) {
            setHairSalon(response.res);
        } 
    };

    const logout = () => {
        localStorage.removeItem('role');
        localStorage.removeItem('token');
        setRole(null);
        setToken(null);
        navigate('/professional/login');
    };

    useEffect(() => {
        const storedRole = localStorage.getItem('role');
        const storedToken = localStorage.getItem('token');
        if (storedRole && storedToken) {
            setRole(storedRole);
            setToken(storedToken);
            fetchHairSalon();
        } else {
            navigate('/professional/login'); 
        }
    }, [navigate]);

    if (!role || role !== 'professional' || !token) {
        navigate('/professional/login');
        return null;
    }

    return (
        <div className={style.container}>
            <div className={style.menu}>
                <h1 style={{ fontSize: '38px' }}>Professional Page</h1>
                <button
                    className={style.logoutButton}
                    onClick={() => { logout(); }}
                >Déconnexion</button>
            </div>
            <div className={style.main}>
                <div className={style.hairSalon}>
                    {!hairSalon ? (
                        <div>
                            <Request/>
                        </div>
                    ) : (
                        <div>
                            <HairSalon {...hairSalon}/>
                        </div>
                    )}  
                </div>
            </div>
        </div>
    );
}