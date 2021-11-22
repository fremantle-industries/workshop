import React from 'react';
import ReactDOM from 'react-dom/client';
import {
	createBrowserRouter,
	createRoutesFromElements,
	RouterProvider,
	Route,
} from 'react-router-dom';
import Root, {
	loader as rootLoader,
	action as rootAction,
} from './routes/root';
import Index from './routes/index';
import Strategies, {
	loader as strategiesLoader,
} from './routes/strategies';
import Products, {
	loader as productsLoader,
} from './routes/products';
import Spot, {
	loader as spotLoader,
} from './routes/products/spot';
import Futures, {
	loader as futuresLoader,
} from './routes/products/futures';
import Perps, {
	loader as perpsLoader,
} from './routes/products/perps';
import Options, {
	loader as optionsLoader,
} from './routes/products/options';
import Nft, {
	loader as nftLoader,
} from './routes/products/nft';
import Topology, {
	loader as topologyLoader,
} from './routes/topology';
import Api, {
	loader as apiLoader,
} from './routes/api';
import Contact, {
	loader as contactLoader,
	action as contactAction,
} from './routes/contact';
import EditContact, {
	action as editAction,
} from './routes/edit';
import {action as destroyAction} from './routes/destroy';
import ErrorPage from './error-page';
import './index.css';

const router = createBrowserRouter(
	createRoutesFromElements(
		<Route
			path='/'
			element={<Root />}
			loader={rootLoader}
			action={rootAction}
			errorElement={<ErrorPage />}
		>
			<Route errorElement={<ErrorPage />}>
				<Route index element={<Index />} />
				<Route
					path='strategies'
					element={<Strategies />}
					loader={strategiesLoader}
				/>
				<Route
					path='products'
					element={<Products />}
					loader={productsLoader}
				/>
				<Route
					path='products/spot'
					element={<Spot />}
					loader={spotLoader}
				/>
				<Route
					path='products/futures'
					element={<Futures />}
					loader={futuresLoader}
				/>
				<Route
					path='products/perps'
					element={<Perps />}
					loader={perpsLoader}
				/>
				<Route
					path='products/options'
					element={<Options />}
					loader={optionsLoader}
				/>
				<Route
					path='products/nft'
					element={<Nft />}
					loader={nftLoader}
				/>
				<Route
					path='topology'
					element={<Topology />}
					loader={topologyLoader}
				/>
				<Route
					path='api'
					element={<Api />}
					loader={apiLoader}
				/>
				<Route
					path='contacts/:contactId'
					element={<Contact />}
					loader={contactLoader}
					action={contactAction}
				/>
				<Route
					path='contacts/:contactId/edit'
					element={<EditContact />}
					loader={contactLoader}
					action={editAction}
				/>
				<Route
					path='contacts/:contactId/destroy'
					action={destroyAction}
				/>
			</Route>
		</Route>,
	),
);

ReactDOM.createRoot(document.getElementById('root')!).render(
	<React.StrictMode>
		<RouterProvider router={router} />
	</React.StrictMode>,
);
