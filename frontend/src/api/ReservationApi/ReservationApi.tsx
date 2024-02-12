export async function getReservations(id: number): Promise<any> {
    const response = await fetch(`http://127.0.0.1:8081/api/professional/reservation/${id}`, {
        method: "GET",
        headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${localStorage.getItem('token')}`,
        },
        })

        if (response.ok) {
            const res = await response.json();
            return { success: true, res };
        } else {
            return { success: false, error: 'RESERVATION_DOES_NOT_EXIST' };
        }
}

export async function getClientReservations(): Promise<any> {
    const response = await fetch(`http://127.0.0.1:8081/api/client/reservations`, {
        method: "GET",
        headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${localStorage.getItem('token')}`,
        },
        })

        if (response.ok) {
            const res = await response.json();
            return { success: true, res };
        } else {
            return { success: false, error: 'RESERVATION_DOES_NOT_EXIST' };
        }
}