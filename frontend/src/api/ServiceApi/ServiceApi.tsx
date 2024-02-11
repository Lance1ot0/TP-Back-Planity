export async function getServices(url: string): Promise<any> {
    const response = await fetch(`http://127.0.0.1:8081/api/${url}`, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          "Authorization": `Bearer ${localStorage.getItem('token')}`,
        },
      })
        if (response.ok) {
            const res = await response.json();
            return res
        } else {
            return { success: false, error: 'SERVICE_DOES_NOT_EXIST' };
        }
    }

export async function addService(item: any): Promise<any> {
    const response = await fetch(`http://127.0.0.1:8081/api/professional/service`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${localStorage.getItem('token')}`,
        },
        body: JSON.stringify(item),
    })  
    if (response.ok) {
        return { success: true }
    } else {
        return { success: false, error: 'SERVICE_ALREADY_EXIST' }
    }
}