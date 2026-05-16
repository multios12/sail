export async function fetchJson<T>(url: string, init?: RequestInit): Promise<T | null> {
  const response = await fetch(url, init);
  const text = await response.text();

  if (!response.ok) {
    console.error(`Request failed: ${response.status} ${response.statusText}`, url, text);
    return null;
  }

  if (text.trim() === "") {
    return null;
  }

  return JSON.parse(text) as T;
}
