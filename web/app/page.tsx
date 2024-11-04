import { UpcomingRaces } from "@/src/components/home/UpcomingRaces";
import { Podium } from "@/src/components/results/Podium";

export default function Results() {
  return (
    <main>
      <section className="py-12 bg-gray-50">
        <div className="container mx-auto px-4 text-center">
          <h2 className="text-5xl font-bold text-gray-800 mb-4">Latest F1 Race Results</h2>
          <p className="text-lg text-gray-500 mb-8">Get the latest race results, standings, and upcoming schedules of the Formula 1 season.</p>
          <a href="#" className="bg-red-600 text-white py-3 px-6 rounded-lg text-lg hover:bg-red-500">See Full Results</a>
        </div>
      </section>
      <div className="mx-auto container border-b-2 border-gray-200" />

      { /* Latest Race Results */ }
      <section className="container mx-auto py-12 px-4">
        <h3 className="text-3xl font-semibold mb-6">Last Race: Italian Grand Prix</h3>
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
          <Podium
            title={'1st Place'}
            img={'/images/driver/alonso.jpeg'}
            driver={'Fernando Alonso'}
            team={'Aston Martin'}
            time={'1:35:15'}
          />
          <Podium
            title={'2nd Place'}
            img={'/images/driver/verstappen.jpeg'}
            driver={'Max Verstappen'}
            team={'Red Bull Racing'}
            time={'1:35:22'}
          />
          <Podium
            title={'3rd Place'}
            img={'/images/driver/hamilton.jpeg'}
            driver={'Lewis Hamilton'}
            team={'Mercedes'}
            time={'1:35:55'}
          />
        </div>
      </section>
      <div className="mx-auto container border-b-2 border-gray-200" />

      <UpcomingRaces />
    </main>
  );
}
