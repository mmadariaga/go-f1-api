interface DriverProps {
    name: string;
    img: string;
    team: string;
    championships: number
}

export function Driver (props: DriverProps): JSX.Element {

    const { name, img, team, championships } = props;

    return (
        <div className="p-6 rounded-lg shadow-2xl">
            <div className="bg-white flex items-center">
                <img src={`/images/driver/${img}`} alt={name} className="w-16 h-16 rounded-full mr-4" />
                <div>
                    <h4 className="text-xl font-bold mb-2 text-gray-800">{name}</h4>
                    <p className="text-lg text-gray-600">Team: {team}</p>
                    <p className="text-lg text-gray-600">Championships: {championships}</p>
                </div>
            </div>
        </div>
    );


}