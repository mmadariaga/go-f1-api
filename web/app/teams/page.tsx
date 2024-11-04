import { Team } from "@/src/components/teams/Team";

export default function Teams() {
  return (
    <main>
        <section className="py-16 bg-gray-50">
            <div className="container mx-auto px-4 text-center">
                <h2 className="text-5xl font-bold text-gray-800 mb-4">F1 Teams</h2>
                <p className="text-lg text-gray-500 mb-8">Discover all the teams competing in the Formula 1 season, with detailed information about each one.</p>
            </div>
        </section>

        <div className="mx-auto container border-b-2 border-gray-200"></div>

        <section className="container mx-auto py-12 px-4">
            <h3 className="text-3xl font-semibold mb-6">2024 Teams</h3>
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
                <Team
                    name="Red Bull Racing"
                    img="redbull.avif"
                    teamPrincipal="Christian Horner"
                    engine="Honda RBPT"
                />

                <Team
                    name="Mercedes AMG Petronas"
                    img="mercedes.avif"
                    teamPrincipal="Toto Wolff"
                    engine="Mercedes"
                />

                <Team
                    name="Scuderia Ferrari"
                    img="ferrari.avif"
                    teamPrincipal="Fred Vasseur"
                    engine="Ferrari"
                />

                <Team
                    name="Aston Martin"
                    img="astonmartin.avif"
                    teamPrincipal="Mike Krack"
                    engine="Mercedes"
                />
            </div>
        </section>
    </main>
  );
}
