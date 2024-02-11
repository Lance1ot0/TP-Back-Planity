export async function researchSalonByName(name: string): Promise<any> {
  const response = await fetch(`http://127.0.0.1:8081/api/client/hairsalon`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ name }),
  });

  if (response.ok) {
    return response.json();
  } else {
    return { success: false, error: "NO SALON FOUND" };
  }
}
