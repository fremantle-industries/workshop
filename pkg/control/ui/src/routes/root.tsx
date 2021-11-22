// Import { useEffect } from 'react';
import {
	useLoaderData,
	useLocation,
	useNavigation,
	useSubmit,
	redirect,
	Outlet,
	NavLink,
	Form,
	type LoaderFunctionArgs,
} from 'react-router-dom';
import {getContacts, createContact, type ContactInfo} from '../contacts';

type RootLoaderData = {
	contacts: ContactInfo[];
	q?: string;
};

export async function loader({request}: LoaderFunctionArgs): Promise<RootLoaderData> {
	const url = new URL(request.url);
	const q = url.searchParams.get('q') || undefined;
	const contacts = await getContacts(q);
	return {contacts, q};
}

export async function action() {
	const contact = await createContact();
	return redirect(`/contacts/${contact.id}/edit`);
}

export default function Root() {
	const {contacts} = useLoaderData() as RootLoaderData;
	const navigation = useNavigation();

	// UseEffect(() => {
	//   // TODO: use a ref here!
	//   document.getElementById("q").value = q;
	// }, [q])

	return (
		<>
			<MainNav />

			<div className='grid grid-cols-2'>
				<div id='sidebar'>
					<div>
						<Form method='post'>
							<button type='submit'>New</button>
						</Form>
					</div>
					<nav>
						{contacts.length ? (
							<ul>
								{contacts.map(contact => (
									<li key={contact.id}>
										<NavLink
											to={`contacts/${contact.id}`}
											className={({isActive, isPending}) =>
												isActive
													? 'active'
													: isPending
														? 'pending'
														: ''
											}
										>
											{contact.first || contact.last ? (
												<>
													{contact.first} {contact.last}
												</>
											) : (
												<i>No Name</i>
											)}{' '}
											{contact.favorite && <span>★</span>}
										</NavLink>
									</li>
								))}
							</ul>
						) : (
							<p>
								<i>No contacts</i>
							</p>
						)}
					</nav>
				</div>

				<div
					id='detail'
					className={
						navigation.state === 'loading' ? 'loading' : ''
					}
				>
					<Outlet />
				</div>
			</div>
		</>
	);
}

const MainNav = () => {
	const {q} = useLoaderData() as RootLoaderData;
	const navigation = useNavigation();
	const location = useLocation();
	const submit = useSubmit();
	const searching = navigation.location && new URLSearchParams(navigation.location.search).has('q');
	const isApisActive = location.pathname.startsWith('/api');
	const isProductsActive = location.pathname.startsWith('/products');

	return (
		<>
			<div className='flex items-center justify-between border-b-2 pt-2 pb-2 h-20 bg-base-200'>
				<div className='ml-8'>
					<NavLink to='/' className='float-left'>
						<svg xmlns='http://www.w3.org/2000/svg' fill='none' viewBox='0 0 24 24' strokeWidth={1.5} stroke='currentColor' className='w-10 h-10'>
							<path strokeLinecap='round' strokeLinejoin='round' d='M11.42 15.17L17.25 21A2.652 2.652 0 0021 17.25l-5.877-5.877M11.42 15.17l2.496-3.03c.317-.384.74-.626 1.208-.766M11.42 15.17l-4.655 5.653a2.548 2.548 0 11-3.586-3.586l6.837-5.63m5.108-.233c.55-.164 1.163-.188 1.743-.14a4.5 4.5 0 004.486-6.336l-3.276 3.277a3.004 3.004 0 01-2.25-2.25l3.276-3.276a4.5 4.5 0 00-6.336 4.486c.091 1.076-.071 2.264-.904 2.95l-.102.085m-1.745 1.437L5.909 7.5H4.5L2.25 3.75l1.5-1.5L7.5 4.5v1.409l4.26 4.26m-1.745 1.437l1.745-1.437m6.615 8.206L15.75 15.75M4.867 19.125h.008v.008h-.008v-.008z' />
						</svg>
					</NavLink>
					<form id='search-form' role='search' className='inline-block ml-4'>
						<div className='form-control'>
							<div className='input-group'>
								<input
									placeholder='Search…'
									id='q'
									className={(searching ? 'loading' : '') + ' input input-bordered w-96'}
									aria-label='Search contacts'
									type='search'
									name='q'
									defaultValue={q}
									onChange={event => {
										const isFirstSearch = (q === null);
										submit(event.currentTarget.form, {
											replace: !isFirstSearch,
										});
									}}
								/>
								<button className='btn btn-square'>
									<svg xmlns='http://www.w3.org/2000/svg' className='h-6 w-6' fill='none' viewBox='0 0 24 24' stroke='currentColor'><path strokeLinecap='round' strokeLinejoin='round' strokeWidth='2' d='M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z' /></svg>
								</button>
							</div>
						</div>
					</form>
				</div>

				<div className='mr-8'>
					<nav className='grid'>
						<ul>
							<li className='inline-block'>
								<NavLink to='/strategies'>Strategies</NavLink>
							</li>
							<li className='inline-block'>
                &nbsp;|&nbsp;
								<NavLink to='/products'>Products</NavLink>
							</li>
							<li className='inline-block'>
                &nbsp;|&nbsp;
								<NavLink to='/topology'>Topology</NavLink>
							</li>
							<li className='inline-block'>
                &nbsp;|&nbsp;
								<NavLink to='/api'>API</NavLink>
							</li>
							<li className='inline-block'>
                &nbsp;|&nbsp;
								<a href='http://redpanda.localhost'>Redpanda</a>
							</li>
							<li className='inline-block'>
                &nbsp;|&nbsp;
								<a href='http://grafana.localhost'>Grafana</a>
							</li>
							<li className='inline-block'>
                &nbsp;|&nbsp;
								<a href='http://prometheus.localhost'>Prometheus</a>
							</li>
						</ul>
					</nav>
				</div>
			</div>

			{(isProductsActive || isApisActive) && (
				<div className='flex items-center justify-between border-b-2 pt-2 pb-2 pl-8 pr-8'>
					{isProductsActive
            && <ul>
            	<li className='inline-block'>
            		<NavLink to='/products/spot'>Spot</NavLink>
            	</li>
            	<li className='inline-block'>
                &nbsp;|&nbsp;
            		<NavLink to='/products/futures'>Futures</NavLink>
            	</li>
            	<li className='inline-block'>
                &nbsp;|&nbsp;
            		<NavLink to='/products/perps'>Perps</NavLink>
            	</li>
            	<li className='inline-block'>
                &nbsp;|&nbsp;
            		<NavLink to='/products/options'>Options</NavLink>
            	</li>
            	<li className='inline-block'>
                &nbsp;|&nbsp;
            		<NavLink to='/products/nft'>NFT</NavLink>
            	</li>
            </ul>}

					{isApisActive
            && <ul>
            	<li className='inline-block'>
            		<a href='http://vgraph.api.localhost'>Virtual Graph</a>
            	</li>
            	<li className='inline-block'>
                &nbsp;|&nbsp;
            		<a href='http://search.api.localhost'>Search</a>
            	</li>
            	<li className='inline-block'>
                &nbsp;|&nbsp;
            		<a href='http://optionchain.api.localhost'>Option Chain</a>
            	</li>
            	<li className='inline-block'>
                &nbsp;|&nbsp;
            		<a href='http://orderbook.api.localhost'>Order Book</a>
            	</li>
            	<li className='inline-block'>
                &nbsp;|&nbsp;
            		<a href='http://products.api.localhost'>Registry</a>
            	</li>
            	<li className='inline-block'>
                &nbsp;|&nbsp;
            		<a href='http://rollup.api.localhost'>Rollup</a>
            	</li>
            	<li className='inline-block'>
                &nbsp;|&nbsp;
            		<a href='http://nftactivity.api.localhost'>NFT Activity</a>
            	</li>
            </ul>
					}
				</div>
			)}
		</>
	);
};
