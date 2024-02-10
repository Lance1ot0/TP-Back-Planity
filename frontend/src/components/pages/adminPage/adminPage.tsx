import { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';

export default function AdminPage() {
    const navigate = useNavigate();
    const [token, setToken] = useState<string | null>(null);

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
        <div>
            <h1>Admin Page</h1>

        </div>
    );
}
