import { Podium } from "@/src/components/results/Podium";

export default function Results() {
  return (
    <main>
      <section className="py-12 bg-gray-50">
        <div className="container mx-auto px-4 text-center">
          <h2 className="text-5xl font-bold text-gray-800 mb-4">TODO</h2>
          <p className="text-lg text-gray-500 mb-8">Get the latest race results, standings, and upcoming schedules of the Formula 1 season.</p>
          <a href="#" className="bg-red-600 text-white py-3 px-6 rounded-lg text-lg hover:bg-red-500">See Full Results</a>
        </div>
      </section>
      <div className="mx-auto container border-b-2 border-gray-200" />
    </main>
  );
}
