export const registerClientApi = async (item: object): Promise<any> => {
    const response = await fetch('http://127.0.0.1:8081/api/client/register', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(item),
    });

    if (response.ok) {
        return { success: true };
    } else {
        return { success: false, error: 'AUTHENTICATION_FAILED' };
    }
}