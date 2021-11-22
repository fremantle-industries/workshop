export type SpotLoaderData = {
	connectors: Array<Record<string, unknown>>;
};

export async function loader(): Promise<SpotLoaderData> {
	// Const spot = await getRegistry();
	const spot: SpotLoaderData = {
		connectors: [],
	};
	return spot;
}

export default function Spot() {
	// Const registry = useLoaderData() as RegistryLoaderData;

	return (
		<div id='spot'>
      TODO SPOT
		</div>
	);
}
