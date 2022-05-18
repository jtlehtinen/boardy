import { writable } from 'svelte/store';

const storedIsLoggedIn = localStorage.getItem('isLoggedIn');

export const isLoggedIn = writable(storedIsLoggedIn || false);

isLoggedIn.subscribe(value => localStorage.setItem('isLoggedIn', value.toString()));

