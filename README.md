# Ausblick
Ausblick is an open-source email and calendar client for multiple platforms, developed by a second-year trainee.

Ausblick runs with MacOS, Windows and Linux

Exchange Online: Full support for shared mailboxes and calendars. IMAP/SMTP and CalDAV for other providers planned as a later milestone.

Multi-account: Multiple private and business accounts with calendar overlay (merged view).

Performance: approx. 5 mailboxes / 100,000 messages, locally indexed, responsive UI

Security: No web technologies (no Electron, no embedded browser, no JavaScript).

Sovereignty: Prefer EU-sovereign solutions. System SDKs from US providers are acceptable.



# Technologie stack:

Language: Go is used as the implementation language, chosen for its small surface area, extensive standard library (HTTP, TLS, JSON), and simple concurrency model.

GUI: Fyne is used for the user interface, providing a pure Go API with GPU-based rendering and no dependency on any web engine.

Exchange: Microsoft Graph is accessed via msgraph-sdk-go and azidentity, using delegated permissions only. The Entra app registration is performed by the trainee.

Local storage: SQLite is used via modernc.org/sqlite (cgo-free) for local data persistence, with FTS5 enabling full-text search.

Secrets: Tokens and credentials are stored in the OS keychain (e.g. via zalando/go-keyring); no tokens are written to files or logs.

IMAP/SMTP (later): emersion/go-imapv2 and emersion/go-message will be used to support IMAP and SMTP for non-Exchange providers.

CalDAV (later): emersion/go-webdav will be used for CalDAV support.
