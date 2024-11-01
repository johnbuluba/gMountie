import {skeleton} from '@skeletonlabs/tw-plugin';
import forms from '@tailwindcss/forms';
import typography from '@tailwindcss/typography';
import {join} from 'path';
import type {Config} from 'tailwindcss';


export default {
  // 2. Opt for dark mode to be handled via the class method
  darkMode: 'class',
  content: [
    './src/**/*.{html,js,svelte,ts}',
    join(require.resolve('@skeletonlabs/skeleton'), '../**/*.{html,js,svelte,ts}')
  ],
  theme: {
    extend: {},
  },
  plugins: [
    typography,
    forms,
    skeleton({
      themes: {
        preset: ["modern", "skeleton", "modern", "crimson"]
      }
    })
  ],
} satisfies Config;
