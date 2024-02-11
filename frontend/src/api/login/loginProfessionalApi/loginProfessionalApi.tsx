import { LoginResponse } from "./loginProfessionalType";

export const loginProfessionalApi = async (email: string, password: string): Promise<LoginResponse> => {
    const response = await fetch('http://127.0.0.1:8081/api/professional/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ email, password }),
    });

    if (response.ok) {
      const data = await response.json();
      return { success: true, role: data.role, token: data.token };
    } else {
      return { success: false, error: 'AUTHENTICATION_FAILED' };
    }
};

