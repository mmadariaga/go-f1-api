interface PodiumProps {
    title: string;
    img: string;
    driver: string;
    team: string;
    time: string;
}

export function Podium (props: PodiumProps): JSX.Element {

    const { title, img, driver, team, time } = props;

    return (
        <div className="p-6 rounded-lg shadow-2xl">
            <h4 className="text-xl font-bold mb-4 text-gray-700">{title}</h4>
            <div className="bg-white flex items-center">
            <img src={img} alt={driver} className="w-16 h-16 rounded-full mr-4" />
            <div>

                <p className="text-lg">Driver: {driver}</p>
                <p className="text-lg">Team: {team}</p>
                <p className="text-lg">Time: {time}</p>
            </div>
            </div>
        </div>
    );


}