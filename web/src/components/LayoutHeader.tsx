'use client'

import Link from "next/link";
import { usePathname } from "next/navigation";

export function LayoutHeader() {

  const pathname = usePathname();

  type LinkItemType = {
    link: string,
    title: string
  }
  const links: Array<LinkItemType> = [
    { link: '/', title: 'Home' },
    { link: '/results', title: 'Results' },
    { link: '/drivers', title: 'Drivers' },
    { link: '/teams', title: 'Teams' },
  ];

  return (
    <header className="bg-red-600 py-6">
    <div className="container mx-auto flex justify-between items-center px-4">
      <h1 className="text-3xl font-bold text-white">F1 Results</h1>
      <nav>
        <ul className="flex space-x-6">
          {links.map(
            (linkItem) => {

              const { link, title } = linkItem;

              return (
                <li key={link}>
                  <Link href={link} className={`text-white ${link === pathname ? 'font-bold text-gray-300' : 'hover:underline'}`}>
                    {title}
                  </Link>
                </li>
              )
            }
          )}
        </ul>
      </nav>
    </div>
  </header>
  );
}
