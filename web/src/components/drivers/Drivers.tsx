import { fetchDrivers } from "@/src/lib/api";
import { Driver } from "./Driver";

export async function Drivers(): Promise<JSX.Element> {

    // sleep 1 sec just to show the skeleton
    await new Promise((resolve) => setTimeout(resolve, 10000));

    const drivers = await fetchDrivers();

    return (
        <>
        {drivers.map((driver) => {
            return (
                <Driver
                    key={driver.driver_number}
                    name={driver.full_name}
                    img={driver.avatar}
                    team={driver.team_name}
                    number={driver.driver_number}
                />
            );
        })}
        </>
    );
}