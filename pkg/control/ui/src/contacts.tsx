import localforage from 'localforage';
import {matchSorter} from 'match-sorter';
// @ts-expect-error
import sortBy from 'sort-by';

export type ContactInfo = {
	id: string;
	createdAt: number;
	favorite: boolean;
	first?: string;
	last?: string;
	avatar?: string;
	twitter?: string;
	notes?: string;
};

export async function getContacts(query?: string): Promise<ContactInfo[]> {
	await fakeNetwork(`getContacts:${query}`);
	let contacts: ContactInfo[] = await localforage.getItem('contacts') || [];
	if (!contacts) {
		contacts = [];
	}

	if (query) {
		contacts = matchSorter(contacts, query, {keys: ['first', 'last']});
	}

	return contacts.sort(sortBy('last', 'createdAt'));
}

export async function createContact() {
	await fakeNetwork();
	const id = Math.random().toString(36).substring(2, 9);
	const contact: ContactInfo = {id, createdAt: Date.now(), favorite: false};
	const contacts = await getContacts();
	contacts.unshift(contact);
	await set(contacts);
	return contact;
}

export async function getContact(id: ContactInfo['id']) {
	await fakeNetwork(`contact:${id}`);
	const contacts: ContactInfo[] = await localforage.getItem('contacts') || [];
	const contact = contacts.find(contact => contact.id === id);
	return contact ?? null;
}

export async function updateContact(id: ContactInfo['id'], updates: Omit<ContactInfo, 'none'>) {
	await fakeNetwork();
	const contacts: ContactInfo[] = await localforage.getItem('contacts') || [];
	const contact = contacts.find(contact => contact.id === id);
	if (!contact) {
		throw new Error(`No contact found for ${id}`);
	}

	Object.assign(contact, updates);
	await set(contacts);
	return contact;
}

export async function deleteContact(id: ContactInfo['id']) {
	const contacts: ContactInfo[] = await localforage.getItem('contacts') || [];
	const index = contacts.findIndex(contact => contact.id === id);
	if (index > -1) {
		contacts.splice(index, 1);
		await set(contacts);
		return true;
	}

	return false;
}

async function set(contacts: ContactInfo[]): Promise<ContactInfo[]> {
	return localforage.setItem('contacts', contacts);
}

// Fake a cache so we don't slow down stuff we've already seen
const fakeCache = new Map<string, boolean>();

async function fakeNetwork(key?: string): Promise<void> {
	if (key === undefined || fakeCache.get(key)) {
		return;
	}

	fakeCache.set(key, true);
	return new Promise(res => {
		setTimeout(res, Math.random() * 800);
	});
}
