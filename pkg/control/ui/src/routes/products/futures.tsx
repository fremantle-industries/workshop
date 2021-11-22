export type FuturesLoaderData = {
	connectors: Array<Record<string, unknown>>;
};

export async function loader(): Promise<FuturesLoaderData> {
	// Const spot = await getRegistry();
	const spot: FuturesLoaderData = {
		connectors: [],
	};
	return spot;
}

export default function Futures() {
	// Const registry = useLoaderData() as RegistryLoaderData;

	return (
		<div id='futures'>
      TODO FUTURES
		</div>
	);
}
