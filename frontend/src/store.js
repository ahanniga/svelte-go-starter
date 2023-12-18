import { readable, writable } from 'svelte/store';

export let sidebarVisible = writable(true);
export let showComponent = writable("comp-1");

