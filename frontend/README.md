# Security Code Scanner – Frontend

This is the frontend for the Security Code Scanner project, built with [Next.js](https://nextjs.org) and [Tailwind CSS](https://tailwindcss.com/).

## Features

- Modern UI for submitting and viewing code scan results
- Integration with the backend API for security scanning
- Responsive and accessible design
- Uses [shadcn/ui](https://ui.shadcn.com) for visual components
- Uses [zod](https://zod.dev/) for form validation
- Uses [React Query](https://tanstack.com/query/latest) for data fetching and caching

## Getting Started

1. Install dependencies:

   ```bash
   pnpm install
   ```

2. Run the development server:

   ```bash
   pnpm dev
   ```

3. Open http://localhost:3000 in your browser.

## Project Structure

- src/app/ – Main app and page components
- src/lib/dto/ – TypeScript DTOs for API communication
- src/app/componests/ – UI components for forms and results

## API

- This frontend expects the backend API to be running. See the main project or backend README for details.

## Libs

- UI: [shadcn/ui](https://ui.shadcn.com)
- Validation: [zod](https://zod.dev/)
- Data fetching: [React Query](https://tanstack.com/query/latest)
- Tailwind CSS is used for styling; edit globals.css or use Tailwind utility classes.
