export type OptionsLoaderData = {
	connectors: Array<Record<string, unknown>>;
};

export async function loader(): Promise<OptionsLoaderData> {
	// Const spot = await getRegistry();
	const options: OptionsLoaderData = {
		connectors: [],
	};
	return options;
}

export default function Options() {
	// Const registry = useLoaderData() as RegistryLoaderData;

	return (
		<div id='options'>
      TODO OPTIONS
		</div>
	);
}
