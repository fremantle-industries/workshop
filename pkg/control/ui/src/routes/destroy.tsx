import {type ActionFunctionArgs, redirect} from 'react-router-dom';
import {deleteContact} from '../contacts';

export async function action({params}: ActionFunctionArgs) {
	// @ts-expect-error
	await deleteContact(params.contactId);
	return redirect('/');
}
