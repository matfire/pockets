// @ts-check
import { defineConfig } from 'astro/config';
import starlight from '@astrojs/starlight';

// https://astro.build/config
export default defineConfig({
  site: "https://pockets.matteogassend.com",

  integrations: [
      starlight({
          title: 'Pockets',
          social: {
              github: 'https://github.com/matfire/pockets',
              blueSky: 'https://bsky.app/profile/matteogassend.com'
          },
          sidebar: [
              {
                  label: 'Guides',
                  items: [
                      // Each item here is one entry in the navigation menu.
                      { label: 'Get Started', slug: 'guides/get-started' },
                  ],
              },
              {
                label: "Tutorials",
                autogenerate: { directory: 'tutorials'}
              },
              {
                  label: 'References',
                  autogenerate: { directory: 'references' },
              },
          ],
      }),
	],
});