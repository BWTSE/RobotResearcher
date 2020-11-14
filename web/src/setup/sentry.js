import Vue from 'vue'
import * as Sentry from '@sentry/browser'
import { Vue as VueIntegration } from '@sentry/integrations'
import { Integrations } from '@sentry/tracing'

Sentry.init({
  dsn: 'https://c3812d572f7f47418eb2bc2dd5e32179@o408198.ingest.sentry.io/5516819',
  integrations: [
    new VueIntegration({
      Vue,
      tracing: true,
      logErrors: true,
      attachProps: true,
    }),
    new Integrations.BrowserTracing(),
  ],
  environment: process.env.NODE_ENV,
})
