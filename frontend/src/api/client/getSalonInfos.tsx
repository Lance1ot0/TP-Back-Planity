export async function getSalonInfos(url: string): Promise<any> {
  const response = await fetch(`http://127.0.0.1:8081/api/${url}`, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  });
  if (response.ok) {
    const res = await response.json();
    return res;
  } else {
    return { success: false, error: "Salon not found" };
  }
}
