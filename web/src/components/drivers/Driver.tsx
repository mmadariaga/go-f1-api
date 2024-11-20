interface DriverProps {
    name: string;
    img: string;
    team: string;
    number: number
}

export function Driver(props: DriverProps): JSX.Element {

    const { name, img, team, number } = props;

    return (
        <div className="p-6 rounded-lg shadow-2xl">
            <div className="bg-white flex items-center">
                <img src={img} width={64} height={64} alt={name} className="w-16 h-16 rounded-full mr-4" />
                <div>
                    <h4 className="text-xl font-bold mb-2 text-gray-800">{name}</h4>
                    <p className="text-lg text-gray-600">Team: {team}</p>
                    <p className="text-lg text-gray-600">Number: {number}</p>
                </div>
            </div>
        </div>
    );


}