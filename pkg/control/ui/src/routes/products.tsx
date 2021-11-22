export type RegistryLoaderData = {
	connectors: Array<Record<string, unknown>>;
};

export async function loader(): Promise<RegistryLoaderData> {
	// Const registry = await getRegistry();
	const registry: RegistryLoaderData = {
		connectors: [],
	};
	return registry;
}

export default function Registry() {
	// Const registry = useLoaderData() as RegistryLoaderData;

	return (
		<div id='registry'>
      TODO REGISTRY
		</div>
	);
}
