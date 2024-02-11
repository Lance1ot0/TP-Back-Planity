export async function getAvailability(url: string): Promise<any> {
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
          return { success: false, error: 'AVAILABILITY_DOES_NOT_EXIST' };
      }
  }

export async function addAvailability(id: number, item: object): Promise<any> {
    const response = await fetch(`http://127.0.0.1:8081/api/professional/employee/availability/${id}`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "Authorization": `Bearer ${localStorage.getItem('token')}`,
      },
      body: JSON.stringify(item)
    })

    if (response.ok) {
        return { success: true};
    } else {
        return { success: false, error: 'ERRO_WHILE_ADDING_AVAILABILITY'};
    }

  }