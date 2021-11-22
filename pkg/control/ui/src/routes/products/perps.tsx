export type PerpsLoaderData = {
	connectors: Array<Record<string, unknown>>;
};

export async function loader(): Promise<PerpsLoaderData> {
	// Const spot = await getRegistry();
	const perps: PerpsLoaderData = {
		connectors: [],
	};
	return perps;
}

export default function Perps() {
	// Const registry = useLoaderData() as RegistryLoaderData;

	return (
		<div id='perps'>
      TODO PERPS
		</div>
	);
}
