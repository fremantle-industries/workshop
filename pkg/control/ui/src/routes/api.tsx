export type ApiLoaderData = {
	connectors: Array<Record<string, unknown>>;
};

export async function loader(): Promise<ApiLoaderData> {
	const api: ApiLoaderData = {
		connectors: [],
	};
	return api;
}

export default function Topology() {
	// Const api = useLoaderData() as ApiLoaderData;

	return (
		<div id='api'>
      TODO API
		</div>
	);
}
