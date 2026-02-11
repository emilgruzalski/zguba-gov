# Frontend - Zguba.gov

Angular 19 frontend for the lost & found reporting application.

## Setup

```bash
npm install
```

## Development

```bash
npm start
```

Open http://localhost:4200. The app reloads automatically on file changes.

The development server proxies API requests to `http://localhost:8000` (configured in `src/environments/environment.development.ts`).

## Build

```bash
ng build
```

Build artifacts are output to `dist/`. For production:

```bash
ng build --configuration production
```

Or with Docker:

```bash
docker build -t zguba-frontend .
docker run -p 80:80 zguba-frontend
```

The Docker image uses nginx to serve the SPA and proxy `/api/`, `/odata/`, and `/metadata/` routes to the backend.

## Testing

```bash
npm test       # Unit tests (Karma)
ng e2e         # End-to-end tests
```

## Project Structure

```
src/
├── app/
│   ├── app.component.ts       # Main component (form, search, export)
│   ├── app.component.html     # Template
│   ├── app.component.css      # Styles
│   └── services/
│       ├── found-items.service.ts        # API client for found items
│       └── territorial-units.service.ts  # Territorial unit data (from JSON)
├── assets/
│   └── territorial-units.json  # Polish administrative units dataset
├── environments/
│   ├── environment.ts              # Production config
│   └── environment.development.ts  # Development config
└── styles.css                  # Global styles (GOV.PL theme)
```

## Key Features

- Multi-step form for reporting found items
- Autocomplete for territorial units (voivodeships, counties, municipalities)
- Search and filter found items
- Export to JSON and CSV
- Responsive design following GOV.PL guidelines
- WCAG/ARIA accessibility attributes
