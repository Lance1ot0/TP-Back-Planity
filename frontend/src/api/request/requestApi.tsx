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

export async function sendRequest(item: object): Promise<any> {
  const response = await fetch(`http://127.0.0.1:8081/api/professional/request`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      "Authorization": `Bearer ${localStorage.getItem('token')}`,
    },
    body: JSON.stringify(item),
  })

  if (response.ok) {
    const res = await response.json();
    return { success: true, data: res.data };
  } else {
    return { success: false, error: 'REQUEST_FAILED' };
  }
}

export async function getHairSalon(): Promise<any> {
  const response = await fetch(`http://127.0.0.1:8081/api/professional/hairSalon`, {
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
    return { success: false, error: 'HAIRSALON_DOES_NOT_EXIST' };
  }
}

export async function getRequest(): Promise<any> {
  const response = await fetch(`http://127.0.0.1:8081/api/professional/request`, {
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
    return { success: false, error: 'REQUEST_DOES_NOT_EXIST' };
  }
}