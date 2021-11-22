import {
	useLoaderData,
	useFetcher,
	Form,
	type ActionFunctionArgs,
	type LoaderFunctionArgs,
} from 'react-router-dom';
import {getContact, updateContact, type ContactInfo} from '../contacts';

export type ContactLoaderData = ContactInfo;

export async function loader({params}: LoaderFunctionArgs): Promise<ContactLoaderData> {
	// @ts-expect-error
	const contact = await getContact(params.contactId);
	if (!contact) {
		throw new Response('', {
			status: 404,
			statusText: 'Not Found',
		});
	}

	return contact;
}

export async function action({request, params}: ActionFunctionArgs) {
	const formData = await request.formData();
	// @ts-expect-error
	return updateContact(params.contactId, {
		favorite: formData.get('favorite') === 'true',
	});
}

export default function Contact() {
	const contact = useLoaderData() as ContactLoaderData;

	return (
		<div id='contact'>
			<div>
				<img key={contact.avatar} src={contact.avatar} />
			</div>

			<div>
				<h1>
					{contact.first || contact.last ? (
						<>
							{contact.first} {contact.last}
						</>
					) : (
						<i>No Name</i>
					)}{' '}
					<Favorite contact={contact} />
				</h1>

				{contact.twitter && (
					<p>
						<a
							target='_blank'
							href={`https://twitter.com/${contact.twitter}`} rel='noreferrer'
						>
							{contact.twitter}
						</a>
					</p>
				)}

				{contact.notes && <p>{contact.notes}</p>}

				<div>
					<Form action='edit'>
						<button type='submit'>Edit</button>
					</Form>
					<Form
						method='post'
						action='destroy'
						onSubmit={event => {
							if (
								!confirm(
									'Please confirm you want to delete this record.',
								)
							) {
								event.preventDefault();
							}
						}}
					>
						<button type='submit'>Delete</button>
					</Form>
				</div>
			</div>
		</div>
	);
}

type FavoriteProps = {
	contact: ContactInfo;
};

function Favorite({contact}: FavoriteProps) {
	const fetcher = useFetcher();
	// Yes, this is a `let` for later
	let {favorite} = contact;
	if (fetcher.formData) {
		favorite = fetcher.formData.get('favorite') === 'true';
	}

	return (
		<fetcher.Form method='post'>
			<button
				name='favorite'
				value={favorite ? 'false' : 'true'}
				aria-label={
					favorite
						? 'Remove from favorites'
						: 'Add to favorites'
				}
			>
				{favorite ? '★' : '☆'}
			</button>
		</fetcher.Form>
	);
}
