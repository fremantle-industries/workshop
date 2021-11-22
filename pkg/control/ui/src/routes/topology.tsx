export type TopologyLoaderData = {
	connectors: Array<Record<string, unknown>>;
};

export async function loader(): Promise<TopologyLoaderData> {
	// Const topology = await getTopology();
	const topology: TopologyLoaderData = {
		connectors: [],
	};
	return topology;
}

export default function Topology() {
	// Const topology = useLoaderData() as TopologyLoaderData;

	return (
		<div id='topology'>
      TODO TOPOLOGY
		</div>
	);
}
