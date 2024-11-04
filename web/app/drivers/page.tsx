import { Driver } from "@/src/components/drivers/Driver";

export default function Drivers() {
  return (
    <main>
        <section className="py-16 bg-gray-50">
            <div className="container mx-auto px-4 text-center">
                <h2 className="text-5xl font-bold text-gray-800 mb-4">F1 Drivers</h2>
                <p className="text-lg text-gray-500 mb-8">Explore the current drivers competing in the 2024 Formula 1 season, including their teams and key stats.</p>
            </div>
        </section>

        <div className="mx-auto container border-b-2 border-gray-200"></div>

        <section className="container mx-auto py-12 px-4">
            <h3 className="text-3xl font-semibold mb-6">2024 Drivers</h3>
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
                <Driver
                    name="Lewis Hamilton"
                    img="hamilton.jpeg"
                    team="Mercedes"
                    championships={7}
                />
                <Driver
                    name="Max Verstappen"
                    img="verstappen.jpeg"
                    team="Red Bull Racing"
                    championships={2}
                />
                <Driver
                    name="Fernando Alonso"
                    img="alonso.jpeg"
                    team="Aston Martin"
                    championships={2}
                />
            </div>
        </section>
    </main>
  );
}
