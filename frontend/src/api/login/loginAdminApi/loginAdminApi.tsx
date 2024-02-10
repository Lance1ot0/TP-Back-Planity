import { LoginResponse } from "./loginAdminType";

export const loginAdminApi = async (email: string, password: string): Promise<LoginResponse> => {
    const response = await fetch('http://127.0.0.1:8081/api/admin/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ email, password }),
    });

    if (response.ok) {
      const data = await response.json();
      return { success: true, token: data.token };
    } else {
      return { success: false, error: 'AUTHENTICATION_FAILED' };
    }
};

