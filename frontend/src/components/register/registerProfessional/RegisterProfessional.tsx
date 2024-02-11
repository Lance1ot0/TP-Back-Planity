import { useState } from 'react';
import style from './RegisterProfessional.module.css';
import { useNavigate } from 'react-router-dom';
import { registerProfessionalApi } from '../../../api/register/registerProfessionalApi/registerProfessionalApi';

export default function registerProfessional() {
    const navigate = useNavigate();
    const  [errorMessage, setErrorMessage] = useState(false);

    const onHandleLogin = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        const firstname = (e.target as HTMLFormElement).firstname.value;
        const lastname = (e.target as HTMLFormElement).lastname.value;
        const email = (e.target as HTMLFormElement).email.value;
        const phone = (e.target as HTMLFormElement).phone.value;
        const address = (e.target as HTMLFormElement).address.value;
        const password = (e.target as HTMLFormElement).password.value;

        const item = {
            firstname,
            lastname,
            email,
            phone,
            address,
            password
        }

        const response = await registerProfessionalApi(item);
        if (response.success) {
            navigate('/professional/login');
        } else {
            setErrorMessage(true);
        }
    };

    return (
        <div className={style.loginWrapper}>
            <form className={style.loginForm} onSubmit={onHandleLogin}>
                <h2 className={style.loginTitle}>Register Professional</h2>
                <label className={style.loginLabel} htmlFor="firstname">
                    Firstname
                </label>
                <input
                    className={style.loginInput}
                    type="text"
                    id="firstname"
                    name="firstname"
                    required
                />
                <label className={style.loginLabel} htmlFor="lastname">
                    Lastname
                </label>
                <input
                    className={style.loginInput}
                    type="text"
                    id="lastname"
                    name="lastname"
                    required
                />
                <label className={style.loginLabel} htmlFor="email">
                    Email
                </label>
                <input
                    className={style.loginInput}
                    type="text"
                    id="email"
                    name="email"
                    required
                />
                <label className={style.loginLabel} htmlFor="phone">
                    Phone
                </label>
                <input
                    className={style.loginInput}
                    type="text"
                    id="phone"
                    name="phone"
                    required
                />
                <label className={style.loginLabel} htmlFor="address">
                    Address
                </label>
                <input
                    className={style.loginInput}
                    type="text"
                    id="address"
                    name="address"
                    required
                />
                <label className={style.loginLabel} htmlFor="password">
                    Password
                </label>
                <input
                    className={style.loginInput}
                    type="password"
                    id="password"
                    name="password"
                    required
                />


                <button className={style.loginButton} type="submit">
                    Register
                </button>
            </form>
            <a className={style.link} href="/professional/login">Se connecter</a>
            {errorMessage && (
                <p className={style.errorMessage}>Error during registration</p>
            )}
        </div>
    );
}