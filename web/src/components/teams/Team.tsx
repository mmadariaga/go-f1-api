interface TeamProps {
    name: string;
    img: string;
    teamPrincipal: string;
    engine: string;
}

export function Team (props: TeamProps): JSX.Element {

    const { name, img, teamPrincipal, engine } = props;

    return (
        <div className="p-6 rounded-lg shadow-2xl">
            <div className="bg-white flex items-center">
                <img src={`/images/team/${img}`} alt={name} className="w-16 h-16 rounded-full mr-4" />
                <div>
                    <h4 className="text-xl font-bold mb-2 text-gray-800">{name}</h4>
                    <p className="text-lg text-gray-600">Team Principal: {teamPrincipal}</p>
                    <p className="text-lg text-gray-600">Engine: {engine}</p>
                </div>
            </div>
        </div>
    );


}