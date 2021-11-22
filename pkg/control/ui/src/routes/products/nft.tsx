export type NftLoaderData = {
	connectors: Array<Record<string, unknown>>;
};

export async function loader(): Promise<NftLoaderData> {
	// Const spot = await getRegistry();
	const nft: NftLoaderData = {
		connectors: [],
	};
	return nft;
}

export default function Nft() {
	// Const registry = useLoaderData() as RegistryLoaderData;

	return (
		<div id='nft'>
      TODO NFT
		</div>
	);
}
