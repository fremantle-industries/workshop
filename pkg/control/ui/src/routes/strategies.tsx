export type StrategiesLoaderData = {
	connectors: Array<Record<string, unknown>>;
};

export async function loader(): Promise<StrategiesLoaderData> {
	const strategies: StrategiesLoaderData = {
		connectors: [],
	};
	return strategies;
}

export default function Topology() {
	// Const strategies = useLoaderData() as StrategiesLoaderData;

	return (
		<div id='strategies'>
      TODO STRATEGIES
		</div>
	);
}
