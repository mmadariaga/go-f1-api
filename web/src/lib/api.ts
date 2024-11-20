import { DriverInterface } from "../types/driver";

export async function fetchJson<T>(
    url: string,
    reqDetails: RequestInit = {},
) : Promise<T> {

    const defaultHeaders = {
      Accept: "application/json",
      Authorization: "Basic Z286YXBp",
    };

    reqDetails.headers = {
      ... defaultHeaders,
      ...(reqDetails.headers ?? {}),
    }

    const response = await fetch(
      url,
      reqDetails
    );

    return await response.json();
}

export async function fetchDrivers(): Promise<DriverInterface[]> {

  try {

    return await fetchJson<DriverInterface[]>(
      `${process.env.NEXT_PUBLIC_BASE_URL}/api/drivers`
    );

  } catch (error) {
    console.error('Unable to fetch drivers', error);
  }

  return [];
}