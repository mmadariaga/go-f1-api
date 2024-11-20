import { Suspense } from "react";

import { Drivers as DriversComponent } from "@/src/components/drivers/Drivers";
import { DriversSkeleton } from "@/src/components/drivers/DriversSkeleton";

export default function Drivers() {

  const year = (new Date()).getFullYear();

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
            <h3 className="text-3xl font-semibold mb-6">{year} Drivers</h3>
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
                <Suspense fallback={<DriversSkeleton />}>
                  <DriversComponent />
                </Suspense>
            </div>
        </section>
    </main>
  );
}
