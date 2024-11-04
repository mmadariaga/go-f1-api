export function UpcomingRaces () : JSX.Element {
    return (
        <section className="py-12 mx-auto bg-gray-50">
            <div className="container mx-auto px-4">
            <h3 className="text-3xl font-semibold mb-8">Upcoming Races</h3>
            <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
                <div className="bg-gray-700 p-6 rounded-lg">
                <h4 className="text-xl font-bold text-white">Japanese Grand Prix</h4>
                <p className="text-gray-400">October 22, 2024</p>
                </div>
                <div className="bg-gray-700 p-6 rounded-lg">
                <h4 className="text-xl font-bold text-white">United States Grand Prix</h4>
                <p className="text-gray-400">October 29, 2024</p>
                </div>
                <div className="bg-gray-700 p-6 rounded-lg">
                <h4 className="text-xl font-bold text-white">Mexican Grand Prix</h4>
                <p className="text-gray-400">November 5, 2024</p>
                </div>
            </div>
            </div>
        </section>
    );
}