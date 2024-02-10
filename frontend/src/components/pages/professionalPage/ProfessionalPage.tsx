import { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import style from './ProfessionalPage.module.css';

export default function ProfessionalPage() {
    const navigate = useNavigate();
    const [role, setRole] = useState<string | null>(null);
    const [token, setToken] = useState<string | null>(null);

    const logout = () => {
        localStorage.removeItem('role');
        localStorage.removeItem('token');
        setRole(null);
        setToken(null);
        navigate('/professional/login');
    };

    useEffect(() => {
        const role = localStorage.getItem('role');
        const storedToken = localStorage.getItem('token');
        if (role && storedToken) {
            setRole(role);
            setToken(storedToken);
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
        </div>
    );
}