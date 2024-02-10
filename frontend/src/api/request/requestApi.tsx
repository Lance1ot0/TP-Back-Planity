export const getAllRequests = (url: string) =>
fetch(`http://127.0.0.1:8081/api/${url}`, {
  method: "GET",
  headers: {
    "Content-Type": "application/json",
    "Authorization": `Bearer ${localStorage.getItem('token')}`,
  },
}).then((res) => res.json());

export async function updateRequest(id: number, RequestStatus: { RequestStatus: string }, mutate: () => void) {
    const updated = await fetch(`http://127.0.0.1:8081/api/admin/request/${id}`, {
        method: "PUT",
        headers: {
        "Content-Type": "application/json",
        "Authorization": `Bearer ${localStorage.getItem('token')}`,
        },
        body: JSON.stringify(RequestStatus),
    }).then((r) => r.json());

    if (updated) {
        await mutate();
    }
}