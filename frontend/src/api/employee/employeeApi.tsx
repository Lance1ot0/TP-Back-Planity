export async function getEmployee(url: string): Promise<any> {
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
        return { success: false, error: 'EMPLOYEE_DOES_NOT_EXIST' };
    }
}

export async function addEmployee(firstname: string, lastname: string, hairSalonID: number): Promise<any> {
    const response = await fetch(`http://127.0.0.1:8081/api/professional/employee`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${localStorage.getItem('token')}`,
        },
        body: JSON.stringify({ firstname, lastname, hairSalonID }),
    })

    if (response.ok) {
        return { success: true }
    } else {
        return { success: false, error: 'EMPLOYEE_ALREADY_EXIST' }
    }
}